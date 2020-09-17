package utils

import "fmt"

// StringEcho test function
func StringEcho(s string) string {
	return "echo::" + s
}

// ToStringStruct test function
func ToStringStruct(s interface{}) string {
	return fmt.Sprintf("%#v", s) //"echo2::" + s
}
