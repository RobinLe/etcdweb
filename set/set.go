package set

import "context"

// EtcdSet save data in etcd
func EtcdSet(key string, value string, dir bool) error {
	kapi, err := EtcdHandler()
	if err != nil {
		return err
	}
	_, err = kapi.Set(context.Background(), key, value, &client.SetOptions{Dir: dir})
	if err != nil {
		return err
	}
	return nil
}
