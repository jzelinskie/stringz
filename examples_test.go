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

func ExampleDefaultEmpty() {
	myFunc := func() string {
		// Oh no! I failed!
		// Return the empty string!
		return ""
	}

	myFuncOutput := stringz.DefaultEmpty(myFunc(), "sane default")

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

func ExampleUnpack() {
	var first, second string
	if err := stringz.Unpack([]string{"hello", "world"}, &first, &second); err != nil {
		panic(err)
	}
	fmt.Println(first, second)

	// Output:
	// hello world
}

func ExampleTrimPrefixIndex() {
	fmt.Println(stringz.TrimPrefixIndex("this:that", ":"))
	fmt.Println(stringz.TrimPrefixIndex("this::that", "::"))

	// Output:
	// that
	// that
}
