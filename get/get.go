package get

import (
	"context"
	"path/filepath"

	"github.com/coreos/etcd/client"
)

// GetValue get key from etcd
func GetValue(kapi client.KeysAPI, key string) (string, error) {
	value := ""
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		return value, err
	}
	value = resp.Node.Value
	return value, nil
}

// GetKeyValue get key value from etcd
func GetKeyValue(kapi client.KeysAPI, key string) (map[string]string, error) {
	keyValue := map[string]string{}
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		return keyValue, err
	}
	baseKey := filepath.Base(key)
	keyValue[baseKey] = resp.Node.Value
	return keyValue, nil
}
