package db

import "testing"

func TestMongo(t *testing.T) {
	err := initDatabase(nil)
	if err != nil {
		t.Error(err)
		return
	}
}
