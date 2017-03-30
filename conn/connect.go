package conn

import "github.com/coreos/etcd/client"

// EtcdClient return etcd keysapi
func EtcdClient() (client.KeysAPI, error) {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	kapi := client.NewKeysAPI(c)
	return kapi, nil
}
