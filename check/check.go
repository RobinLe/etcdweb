package check

import "strings"

// EtcdCheck check if key or value exist in etcd
func EtcdCheck(key string, value string, dir bool) (bool, error) {
	res, err := EtcdGet(key)
	// check if key exist in etcd
	if err != nil {
		return false, err
	}
	// check if key exist in dir, dir's key is the same to value
	if dir {
		for _, i := range res.Node.Nodes {
			key := strings.Split(i.Key, "/")
			if key[len(key)-1] == value {
				return true, nil
			}
		}
	}
	// check if value exist in etcd
	if value == res.Node.Value {
		return true, nil
	}
	return false, nil
}

// EtcdKeyCheck check if key exist in etcd
func EtcdKeyCheck(key string) (bool, error) {
	_, err := EtcdGet(key)
	if err != nil {
		return false, err
	}
	return true, nil
}
