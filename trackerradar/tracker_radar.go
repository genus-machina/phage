package trackerradar

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/genus-machina/phage/utilities"
)

const (
	RepoPath = "tracker-radar"
	RepoUrl  = "https://github.com/duckduckgo/tracker-radar.git"
)

type DomainList struct {
	logger  *log.Logger
	records []*DomainRecord
	rule    SelectionRule
}

func NewDomainList(logger *log.Logger, rules ...SelectionRule) *DomainList {
	tracker := new(DomainList)
	tracker.logger = log.New(logger.Writer(), "[tracker radar] ", logger.Flags())
	tracker.rule = AnyOf(rules...)
	return tracker
}

func (tracker *DomainList) cloneDataset() error {
	tracker.logger.Printf("Cloning dataset from '%s' to '%s'...\n", RepoUrl, RepoPath)
	return utilities.CloneGitRepository(tracker.logger, RepoUrl, RepoPath)
}

func (tracker *DomainList) Domains() ([]string, error) {
	tracker.logger.Println("Initializing dataset...")
	if err := tracker.initializeDataset(); err != nil {
		return nil, err
	}
	tracker.logger.Println("Loading domain records...")
	if err := tracker.loadDataset(); err != nil {
		return nil, err
	}
	tracker.logger.Printf("Loaded %d domain records.\n", len(tracker.records))
	return tracker.filterDomains(), nil
}

func (tracker *DomainList) filterDomains() []string {
	var domains []string

	for _, record := range tracker.records {
		if tracker.rule(record) {
			domains = append(domains, record.Domains()...)
		}
	}

	tracker.logger.Printf("Selected %d domains for blocking.\n", len(domains))
	return domains
}

func (tracker *DomainList) handleDomainFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		var record *DomainRecord
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&record)
		if err != nil {
			return err
		}

		tracker.records = append(tracker.records, record)
	}

	return nil
}

func (tracker *DomainList) initializeDataset() error {
	var err error

	if tracker.shouldUpdateDataset(RepoPath) {
		err = tracker.updateDataset()
	} else {
		err = tracker.cloneDataset()
	}

	return err
}

func (tracker *DomainList) loadDataset() error {
	tracker.records = nil
	path := filepath.Join(RepoPath, "domains")
	return filepath.Walk(path, tracker.handleDomainFile)
}

func (tracker *DomainList) shouldUpdateDataset(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

func (tracker *DomainList) updateDataset() error {
	tracker.logger.Printf("Updating dataset at '%s'...\n", RepoPath)
	return utilities.UpdateGitRepository(tracker.logger, RepoPath)
}
