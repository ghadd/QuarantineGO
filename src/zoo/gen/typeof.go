package gen

import (
	"fmt"
	"strings"
)

const (
	nameStartsFrom = 5
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func GetKey(v interface{}) string {
	return strings.ToLower(typeof(v)[nameStartsFrom:])
}
