package static

import (
	"log"
	"os"
)

type LocalHostFile struct {
	*HostFile
	path string
}

func NewLocalHostFile(logger *log.Logger, path string) *LocalHostFile {
	list := new(LocalHostFile)
	list.HostFile = new(HostFile)
	list.logger = log.New(logger.Writer(), "[local host file] ", logger.Flags())
	list.path = path
	return list
}

func (list *LocalHostFile) Domains() ([]string, error) {
	list.logger.Printf("Opening host file at '%s'...\n", list.path)
	file, err := os.Open(list.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	list.reader = file
	return list.HostFile.Domains()
}
