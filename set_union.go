package phage

import (
	"strings"
)

type SetUnion []DomainList

func Union(lists ...DomainList) SetUnion {
	return lists
}

func (list SetUnion) Domains() ([]string, error) {
	var domains []string
	set := make(map[string]bool)

	for _, component := range list {
		records, err := component.Domains()
		if err != nil {
			return domains, err
		}

		for _, name := range records {
			normalized := strings.ToLower(name)
			set[normalized] = true
		}
	}

	for domain, _ := range set {
		domains = append(domains, domain)
	}

	return domains, nil
}
