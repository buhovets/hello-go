package main

import (
	"hello-go/hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Hello()
	want := "Hello World!"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
