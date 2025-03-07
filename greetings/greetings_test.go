package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Bilal"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)

	if !want.MatchString(message) || err != nil {
		t.Errorf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""

	msg, err := Hello(name)

	if msg != "" || err == nil {
		t.Errorf(`Hello("%v") = %q, %v, want "", error`, name, msg, err)
	}
}
