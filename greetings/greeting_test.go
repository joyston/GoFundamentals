package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Onan"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := HelloMessage(name)

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Onan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := HelloMessage("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
