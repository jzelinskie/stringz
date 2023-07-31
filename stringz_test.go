// Copyright 2019 Jimmy Zelinskie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stringz

import (
	"testing"
)

func TestSliceContains(t *testing.T) {
	table := []struct {
		description string
		ys          []string
		x           string
		expected    bool
	}{
		{"empty slice, empty str", []string{}, "", false},
		{"empty slice, non-empty str", []string{}, "wut", false},
		{"single item slice, empty str", []string{"wut"}, "", false},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := SliceContains(tt.ys, tt.x)
			if actual != tt.expected {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestDedup(t *testing.T) {
	table := []struct {
		description string
		xs          []string
		expected    []string
	}{
		{"empty", []string{}, []string{}},
		{"single item", []string{"x"}, []string{"x"}},
		{"single duplicate removed", []string{"x", "x"}, []string{"x"}},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := Dedup(tt.xs)
			if !SliceEqual(actual, tt.expected) {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestDefault(t *testing.T) {
	table := []struct {
		description string
		val         string
		fallback    string
		zeroValues  []string
		expected    string
	}{
		{"no zero values is val", "test", "", []string{}, "test"},
		{"basic case fallsback", "", "test", []string{""}, "test"},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := Default(tt.val, tt.fallback, tt.zeroValues...)
			if actual != tt.expected {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestSliceEqual(t *testing.T) {
	table := []struct {
		description string
		xs          []string
		ys          []string
		expected    bool
	}{
		{"empty slices are equal", []string{}, []string{}, true},
		{"single items are equal", []string{"x"}, []string{"x"}, true},
		{"duplicates aren't equal", []string{"x", "x"}, []string{"x"}, false},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := SliceEqual(tt.xs, tt.ys)
			if actual != tt.expected {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestMatrixEqual(t *testing.T) {
	table := []struct {
		description string
		xs          [][]string
		ys          [][]string
		expected    bool
	}{
		{"nil slices are equal", nil, nil, true},
		{"empty slices are equal", [][]string{}, [][]string{}, true},
		{"single items are equal", [][]string{{"x"}}, [][]string{{"x"}}, true},
		{"duplicates aren't equal", [][]string{{"x"}, {"x"}}, [][]string{{"x"}}, false},
		{"differing orders aren't equal", [][]string{{"x"}, {"y"}}, [][]string{{"y"}, {"x"}}, false},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := MatrixEqual(tt.xs, tt.ys)
			if actual != tt.expected {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestTrimPrefixIndex(t *testing.T) {
	table := []struct {
		description string
		input       string
		index       string
		expected    string
	}{
		{"no index match returns original", "test/this", "@", "test/this"},
		{"empty string index", "test/this", "", "test/this"},
		{"multichar index", "test/this/that", "/this/", "that"},
		{"simple", "test/this", "/", "this"},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := TrimPrefixIndex(tt.input, tt.index)
			if actual != tt.expected {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestTrimSurrounding(t *testing.T) {
	table := []struct {
		description string
		s           string
		surrounding string
		expected    string
	}{
		{"empty string remains empty", "", "", ""},
		{"empty surrounding trims nothing", "test", "", "test"},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := TrimSurrounding(tt.s, tt.surrounding)
			if actual != tt.expected {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestCopyStringMap(t *testing.T) {
	table := []struct {
		description string
		xs          map[string]string
		expected    map[string]string
	}{
		{"nil maps", nil, nil},
		{"empty maps", map[string]string{}, map[string]string{}},
		{"kv maps", map[string]string{"k": "v"}, map[string]string{"k": "v"}},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := CopyStringMap(tt.xs)
			switch {
			case actual == nil && tt.expected == nil:
				return
			case actual == nil && tt.expected != nil:
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			case actual != nil && tt.expected == nil:
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			default:
				for k := range tt.xs {
					if actual[k] != tt.expected[k] {
						t.Errorf("actual = %v; want = %v", actual, tt.expected)
					}
				}
			}
		})
	}
}

func TestSlicePermutationsR(t *testing.T) {
	table := []struct {
		description string
		xs          []string
		r           int
		expected    [][]string
	}{
		{"r is negative", []string{"a", "b"}, -1, nil},
		{"r > len(xs)", []string{"a", "b"}, 10, nil},
		{"nil slice", nil, 0, nil},
		{"empty slice", []string{}, 0, nil},
		{"two items r = 0", []string{"a", "b"}, 0, nil},
		{"two items r = 1", []string{"a", "b"}, 1, [][]string{{"a"}, {"b"}}},
		{"two items r = 2", []string{"a", "b"}, 2, [][]string{{"a", "b"}, {"b", "a"}}},
		{"duplicates", []string{"a", "a"}, 2, [][]string{{"a", "a"}, {"a", "a"}}},
		{"common case", []string{"a", "b", "c"}, 3, [][]string{
			{"a", "b", "c"},
			{"a", "c", "b"},
			{"b", "a", "c"},
			{"b", "c", "a"},
			{"c", "a", "b"},
			{"c", "b", "a"},
		}},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := SlicePermutationsR(tt.xs, tt.r)
			if !MatrixEqual(actual, tt.expected) {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestSliceCombinationsR(t *testing.T) {
	table := []struct {
		description string
		xs          []string
		r           int
		expected    [][]string
	}{
		{"r is negative", []string{"a", "b"}, -1, nil},
		{"r > len(xs)", []string{"a", "b"}, 10, nil},
		{"nil slice", nil, 0, nil},
		{"empty slice", []string{}, 0, nil},
		{"two items r = 0", []string{"a", "b"}, 0, nil},
		{"two items r = 1", []string{"a", "b"}, 1, [][]string{{"a"}, {"b"}}},
		{"two items r = 2", []string{"a", "b"}, 2, [][]string{{"a", "b"}}},
		{"duplicates", []string{"a", "a"}, 2, [][]string{{"a", "a"}}},
		{"common case", []string{"a", "b", "c"}, 3, [][]string{{"a", "b", "c"}}},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := SliceCombinationsR(tt.xs, tt.r)
			if !MatrixEqual(actual, tt.expected) {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestSliceCombinationsWithReplacement(t *testing.T) {
	table := []struct {
		description string
		xs          []string
		r           int
		expected    [][]string
	}{
		{"r is negative", []string{"a", "b"}, -1, nil},
		{"r > len(xs)", []string{"a", "b"}, 10, nil},
		{"nil slice", nil, 0, nil},
		{"empty slice", []string{}, 0, nil},
		{"two items r = 0", []string{"a", "b"}, 0, nil},
		{"two items r = 1", []string{"a", "b"}, 1, [][]string{{"a"}, {"b"}}},
		{"two items r = 2", []string{"a", "b"}, 2, [][]string{{"a", "a"}, {"a", "b"}, {"b", "b"}}},
		{"duplicates", []string{"a", "a"}, 2, [][]string{{"a", "a"}, {"a", "a"}, {"a", "a"}}},
		{"common case", []string{"a", "b", "c"}, 3, [][]string{
			{"a", "a", "a"},
			{"a", "a", "b"},
			{"a", "a", "c"},
			{"a", "b", "b"},
			{"a", "b", "c"},
			{"a", "c", "c"},
			{"b", "b", "b"},
			{"b", "b", "c"},
			{"b", "c", "c"},
			{"c", "c", "c"},
		}},
	}

	for _, tt := range table {
		t.Run(tt.description, func(t *testing.T) {
			actual := SliceCombinationsWithReplacement(tt.xs, tt.r)
			if !MatrixEqual(actual, tt.expected) {
				t.Errorf("actual = %v; want = %v", actual, tt.expected)
			}
		})
	}
}

func TestSplitExact(t *testing.T) {
	testCases := []struct {
		src         string
		expectedErr error
		parts       []string
	}{
		{"", nil, []string{""}},
		{"one/two", nil, []string{"one", "two"}},
		{"one/two/three/four/five", nil, []string{"one", "two", "three", "four", "five"}},
		{"one/two", ErrInconsistentUnpackLen, []string{}},
		{"one", ErrInconsistentUnpackLen, []string{"one", "two"}},
	}
	for _, tc := range testCases {
		t.Run(tc.src, func(t *testing.T) {
			dests := make([]string, len(tc.parts))
			destVars := make([]*string, len(tc.parts))

			for i := range dests {
				destVars[i] = &dests[i]
			}

			err := SplitExact(tc.src, "/", destVars...)
			if err != tc.expectedErr {
				t.Fatalf("actual = %s, want = %s", err, tc.expectedErr)
			}
			if err == nil {
				for i := range tc.parts {
					if tc.parts[i] != dests[i] {
						t.Fatalf("actual[%d] = %s, expected[%d] = %s", i, dests[i], i, tc.parts[i])
					}
				}
			}
		})
	}
}

func TestSplitInto(t *testing.T) {
	testCases := []struct {
		src         string
		expectedErr error
		parts       []string
	}{
		{"", nil, []string{""}},
		{"one/two", nil, []string{"one", "two"}},
		{"one/two/three/four/five", nil, []string{"one", "two/three/four/five"}},
		{"one/two", nil, []string{"one/two"}},
		{"one", ErrInconsistentUnpackLen, []string{"one", "two"}},
	}
	for _, tc := range testCases {
		t.Run(tc.src, func(t *testing.T) {
			dests := make([]string, len(tc.parts))
			destVars := make([]*string, len(tc.parts))

			for i := range dests {
				destVars[i] = &dests[i]
			}

			err := SplitInto(tc.src, "/", destVars...)
			if err != tc.expectedErr {
				t.Fatalf("actual = %s, want = %s", err, tc.expectedErr)
			}
			if err == nil {
				for i := range tc.parts {
					if tc.parts[i] != dests[i] {
						t.Fatalf("actual[%d] = %s, expected[%d] = %s", i, dests[i], i, tc.parts[i])
					}
				}
			}
		})
	}
}
