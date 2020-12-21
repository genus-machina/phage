package trackerradar

type SelectionRule func(record *DomainRecord) bool

func AllOf(rules ...SelectionRule) SelectionRule {
	return func(record *DomainRecord) bool {
		selected := true
		for _, rule := range rules {
			selected = selected && rule(record)
		}
		return selected
	}
}

func AnyOf(rules ...SelectionRule) SelectionRule {
	return func(record *DomainRecord) bool {
		for _, rule := range rules {
			if rule(record) {
				return true
			}
		}
		return false
	}
}

func FingerprintingScoreAtLeast(score int) SelectionRule {
	return func(record *DomainRecord) bool {
		return record.Fingerprinting >= score
	}
}

func HasCategoryIn(categories ...string) SelectionRule {
	return func(record *DomainRecord) bool {
		for _, member := range categories {
			for _, category := range record.Categories {
				if category == member {
					return true
				}
			}
		}
		return false
	}
}

func Not(rule SelectionRule) SelectionRule {
	return func(record *DomainRecord) bool {
		return !rule(record)
	}
}
