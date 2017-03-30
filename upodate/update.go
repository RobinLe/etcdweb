package upodate

import "context"

// EtcdUpdate update specfic key's value
func EtcdUpdate(key string, value string) error {
	kapi, err := EtcdHandler()
	if err != nil {
		return err
	}
	_, err = kapi.Update(context.Background(), key, value)
	if err != nil {
		return err
	}
	return nil
}
