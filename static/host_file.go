package static

import (
	"bufio"
	"io"
	"log"
	"regexp"
)

var (
	HostFileRecord = regexp.MustCompile(`^0\.0\.0\.0\s+(\S+)`)
)

type HostFile struct {
	logger *log.Logger
	reader io.Reader
}

func NewHostFile(logger *log.Logger, reader io.Reader) *HostFile {
	list := new(HostFile)
	list.logger = log.New(logger.Writer(), "[host file] ", logger.Flags())
	list.reader = reader
	return list
}

func (list *HostFile) Domains() ([]string, error) {
	list.logger.Println("Searching host file for block records...")
	return list.parseRecords(list.reader)
}

func (list *HostFile) parseRecord(record string) string {
	var domain string
	if matches := HostFileRecord.FindStringSubmatch(record); matches != nil {
		domain = matches[1]
	}
	return domain
}

func (list *HostFile) parseRecords(reader io.Reader) ([]string, error) {
	var domains []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		domain := list.parseRecord(scanner.Text())
		if domain != "" {
			domains = append(domains, domain)
		}
	}

	list.logger.Printf("Found %d domain records.\n", len(domains))
	return domains, scanner.Err()
}
