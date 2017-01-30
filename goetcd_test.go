package goetcd

import (
	"testing"
	"log"
)

func TestSetKey(t *testing.T) {
	cases := []string{"http://*:2379"}

	for _, c := range cases {
		_, err := SetConf(c)
		if (err != nil) {
			log.Fatal(err)
			t.Errorf("SetConf(%q) return error!", err)
		}
		err = SetKey("/foo", "bar")
		if (err != nil) {
			log.Fatal(err)
			t.Errorf("SetKey(%q) return error!", err)
		}
	}
}
