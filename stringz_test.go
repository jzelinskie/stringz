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

import "testing"

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
