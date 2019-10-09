package deps

import (
	"bytes"
	"fmt"
)

// Greet says hello to the specified user
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
