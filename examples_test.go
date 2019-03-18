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

func ExampleEmptyDefault() {
	myFunc := func() string {
		// Oh no! I failed!
		// Return the empty string!
		return ""
	}

	myFuncOutput := stringz.EmptyDefault(myFunc(), "sane default")

	fmt.Println(myFuncOutput)

	// Output:
	// sane default
}
