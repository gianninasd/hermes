package utils

import (
	"testing"
)

func TestStringEcho(t *testing.T) {
	if StringEcho("blue") != "echo::blue" {
		t.Error("Expected blue to equal echo::blue")
	}
}

func TestToStringStruct(t *testing.T) {
	type User struct {
		name     string
		email    string
		adminInt bool
	}

	david := User{name: "David James", adminInt: true}

	if x := ToStringStruct(david); x != "utils.User{name:\"David James\", email:\"\", adminInt:true}" {
		t.Error("Expected to be main.User{name:\"David James\", email:\"\", adminInt:true}", x)
	}
}
