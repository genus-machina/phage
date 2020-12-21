package trackerradar

type Cname struct {
	Original string `json:"original"`
	Resolved string `json:"resolved"`
}

type DomainRecord struct {
	Categories     []string `json:"categories"`
	Cnames         []*Cname `json:"cnames"`
	Domain         string   `json:"domain"`
	Fingerprinting int      `json:"fingerprinting"`
	Subdomains     []string `json:"subdomains"`
}

func (domain *DomainRecord) Domains() []string {
	domains := []string{domain.Domain}
	return domains
}
