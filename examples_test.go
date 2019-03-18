package stringz_test

import (
	"fmt"

	"github.com/jzelinskie/stringz"
)

func ExampleTrimSurrounding() {
	unquoted := stringz.TrimSurrounding(`"quoted text"`, `"`)
	fmt.Println(unquoted)
	// Output:
	// quoted text
}
