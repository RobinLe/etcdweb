package delete

import "context"

// EtcdDelete delete a key from etcd
func EtcdDelete(key string, dir bool) error {
	kapi, err := EtcdHandler()
	if err != nil {
		return err
	}
	_, err = kapi.Delete(
		context.Background(),
		key,
		&client.DeleteOptions{Dir: dir, Recursive: true},
	)
	if err != nil {
		return err
	}
	return nil
}
