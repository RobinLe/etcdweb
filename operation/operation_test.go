package operation

import "testing"

func TestGetValue(t *testing.T) {
	_, err := GetKeyValue("registry")
	if err != nil {
		t.Error("false")
	}
}

func TestGetDirKeys(t *testing.T) {
	_, err := GetDirKeys("registry")
	if err != nil {
		t.Error("false")
	}
}
