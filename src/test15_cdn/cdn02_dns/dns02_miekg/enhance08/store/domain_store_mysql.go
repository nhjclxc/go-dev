package store

import "github.com/miekg/dns"

type DomainStoreMySQL struct {
}

func (d *DomainStoreMySQL) GetRecords(qname string) ([]dns.RR, bool) {
	return nil, false
}
func (d *DomainStoreMySQL) Reload() error {

	return nil
}
