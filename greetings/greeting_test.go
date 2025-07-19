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

func TestHellosNames(t *testing.T) {
	names := []string{"Er", "Onan", "Shelah"}

	msgs, err := Hellos(names)

	//testnames := []string{"Perez", "Zerah"}
	for _, name := range names {
		want := regexp.MustCompile(`\b` + name + `\b`)
		if !want.MatchString(msgs[name]) || err != nil {
			t.Errorf(`Hellos(names) = %v, %v, want match for %#q, error`, msgs, err, want)
		}
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{"Tamar", ""}
	msgs, err := Hellos(names)
	if msgs != nil || err == nil {
		t.Errorf(`Hellos([]"Tamar", ""]) = %v, %v want <nil>, error`, msgs, err)
	}
}
