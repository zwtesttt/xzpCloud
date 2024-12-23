package jwt

import (
	"fmt"
	"testing"
)

func TestVaild(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzUwMjg1MzksImlhdCI6MTczNDk0MjEzOSwicm9sZV9pZCI6MSwidXNlcl9pZCI6IjY3NjkxOTFiNDgyOWM4M2Q3Zjk1ZjkyZSJ9.DnKvSE2wIt6EwtrSvl95tSbef8tErJ717jgzC2-yR4Q"

	data, err := ValidateJWT(token)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("data.user_id", data["user_id"])
}
