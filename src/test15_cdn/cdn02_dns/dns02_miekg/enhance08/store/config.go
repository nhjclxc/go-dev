package store

type Config struct {
	Records  []RrRecord
	Upstream Upstream
}
type RrRecord struct {
	Domain  string
	Ttl     uint32
	Answers []Answer
}
type Answer struct {
	Qtype string
	Value string
}
type Upstream struct {
	Enable  bool
	Servers []string
}
