package store

import "github.com/miekg/dns"

type DomainStore interface {
	GetRecords(qname string) ([]dns.RR, bool)
	Reload() error
}
