package operation

// ETCD etcd base information
type ETCD struct {
	Protocol string
	IP       string
	Port     string
	Certs    string
}

// Client etcd client
// type Client struct {
// }

// Element object of a etcd value
type Element struct {
	Key         string
	Dir         bool
	Value       string
	ModifyIndex uint64
	CreateIndex uint64
}
