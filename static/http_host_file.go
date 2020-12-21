package static

import (
	"log"
	"net/http"
)

type HttpHostFile struct {
	*HostFile
	url string
}

func NewHttpHostFile(logger *log.Logger, url string) *HttpHostFile {
	list := new(HttpHostFile)
	list.HostFile = new(HostFile)
	list.logger = log.New(logger.Writer(), "[http host file] ", logger.Flags())
	list.url = url
	return list
}

func (list *HttpHostFile) Domains() ([]string, error) {
	list.logger.Printf("Opening host file at '%s'...\n", list.url)
	response, err := http.Get(list.url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	list.reader = response.Body
	return list.HostFile.Domains()
}
