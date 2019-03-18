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

// Package stringz implements a collection of utility functions for
// manipulating strings and lists of strings.
package stringz

import "strings"

// Contains returns true if the provided string is in the provided string
// slice.
func Contains(ys []string, x string) bool {
	for _, y := range ys {
		if x == y {
			return true
		}
	}
	return false
}

// Dedup returns a new slice with any duplicates removed.
func Dedup(xs []string) []string {
	xsSet := make(map[string]struct{}, 0)
	for _, x := range xs {
		xsSet[x] = struct{}{}
	}

	ys := make([]string, 0, len(xsSet))
	for x := range xsSet {
		ys = append(ys, x)
	}

	return ys
}

// Default returns a fallback value when the provided value is equal to any
// of the provided zero values.
func Default(val, fallback string, zeroValues ...string) string {
	for _, zeroValue := range zeroValues {
		if val == zeroValue {
			return fallback
		}
	}

	return val
}

// TrimSurrounding returns a string with both a prefix and suffix trimmed from
// it.
func TrimSurrounding(s, surrounding string) string {
	s = strings.TrimPrefix(s, surrounding)
	return strings.TrimSuffix(s, surrounding)
}
