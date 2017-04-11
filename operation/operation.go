package operation

import (
	"context"
	"path/filepath"

	"github.com/coreos/etcd/client"
)

// Client return etcd keysapi
func Client() (client.KeysAPI, error) {
	cfg := client.Config{
		Endpoints: []string{"http://192.168.0.200:2379"},
		Transport: client.DefaultTransport,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	kapi := client.NewKeysAPI(c)
	return kapi, nil
}

// Set save data in etcd
func Set(key string, value string, dir bool) error {
	kapi, err := Client()
	if err != nil {
		return err
	}
	_, err = kapi.Set(context.Background(), key, value, &client.SetOptions{Dir: dir})
	if err != nil {
		return err
	}
	return nil
}

// GetValue get key from etcd
func GetValue(key string) (string, error) {
	value := ""
	kapi, err := Client()
	if err != nil {
		return value, err
	}
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		return value, err
	}
	value = resp.Node.Value
	return value, nil
}

// GetKeyValue get key value from etcd
func GetKeyValue(key string) (map[string]string, error) {
	keyValue := map[string]string{}
	kapi, err := Client()
	if err != nil {
		return keyValue, err
	}
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		return keyValue, err
	}
	baseKey := filepath.Base(key)
	keyValue[baseKey] = resp.Node.Value
	return keyValue, nil
}

// Update update specfic key's value
func Update(key string, value string) error {
	kapi, err := Client()
	if err != nil {
		return err
	}
	_, err = kapi.Update(context.Background(), key, value)
	if err != nil {
		return err
	}
	return nil
}

// Delete delete a key from etcd
func Delete(key string, dir bool) error {
	kapi, err := Client()
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

/*
// Check check if key or value exist in etcd
func Check(key string, value string, dir bool) (bool, error) {
	res, err := GetKeyValue(key)
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
*/

// CheckKey check if key exist in etcd
func CheckKey(key string) (bool, error) {
	_, err := GetKeyValue(key)
	if err != nil {
		return false, err
	}
	return true, nil
}
