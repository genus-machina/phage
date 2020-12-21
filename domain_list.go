package phage

type DomainList interface {
	Domains() ([]string, error)
}
