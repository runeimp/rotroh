package rotroh

import (
	"log"
	"testing"
)

func TestRot13String(t *testing.T) {
	got := Rot13String("Hi Ma!")
	// want := "asdf"
	log.Printf("TestRot13String() | got: %s\n", got)
}
