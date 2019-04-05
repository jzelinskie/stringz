package stringz_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

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

func ExampleSliceMap() {
	files := []string{"/home/my/file", "/home/your/file"}

	if err := stringz.SliceMap(files, func(path string) error {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close() // Defer runs for every iteration!

		_, err = io.Copy(ioutil.Discard, f)
		return err
	}); err != nil {
		fmt.Println(err)
	}
}
