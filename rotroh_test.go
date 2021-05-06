package rotroh

import (
	"testing"
)

func TestBase64String(t *testing.T) {
	want := "VXNlck5hbWU6UGFzc3dvcmQ="
	got, err := Base64String("UserName:Password")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
	if got != want {
		t.Fatalf("expected: %q, got: %q", want, got)
	}
}

func TestRot13String(t *testing.T) {
	want := "Uv Zn!"
	got := Rot13String("Hi Ma!")
	if got != want {
		t.Fatalf("expected: %q, got: %q", want, got)
	}
}

func TestRot47String(t *testing.T) {
	want := "(96C6 :D (2=5@n"
	got := Rot47String("Where is Waldo?")
	if got != want {
		t.Fatalf("expected: %q, got: %q", want, got)
	}
}

func TestRotRoh47String(t *testing.T) {
	want := "KDk2QzYgOkQgKDI9NUBu"
	got, err := RotRoh47String("Where is Waldo?")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
	if got != want {
		t.Fatalf("expected: %q, got: %q", want, got)
	}
}

func TestRotCustomString(t *testing.T) {
	want := "0879 Fast Ln."
	got, err := RotCustomString("1342 Fast Ln.", "1234567890")
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
	if got != want {
		t.Fatalf("expected: %q, got: %q", want, got)
	}
}
