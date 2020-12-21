package phage

import (
	"strings"
)

type SetDifference []DomainList

func Difference(lists ...DomainList) SetDifference {
	return lists
}

func (list SetDifference) Domains() ([]string, error) {
	var domains []string

	if len(list) == 0 {
		return domains, nil
	}

	set := make(map[string]bool)
	reference, removed := list[0], list[1:]

	records, err := reference.Domains()
	if err != nil {
		return domains, err
	}

	for _, name := range records {
		normalized := strings.ToLower(name)
		set[normalized] = true
	}

	for _, component := range removed {
		records, err := component.Domains()
		if err != nil {
			return domains, err
		}

		for _, name := range records {
			normalized := strings.ToLower(name)
			set[normalized] = false
		}
	}

	for domain, included := range set {
		if included {
			domains = append(domains, domain)
		}
	}

	return domains, nil
}
