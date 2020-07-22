package model

import (
	"fmt"
	"testing"
)

func TestApiAccessInsert(t *testing.T) {
	if err := ApiAccessInsertMinute("x", "[::1]:60539"); err != nil {
		t.Fatal(err)
	}
}

func TestApiAcessCountMinute(t *testing.T) {
	count, err := ApiAcessCountMinute("[::1]:63332")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(count)
}
