package main

import (
	"reflect"
	"testing"
)

type FilterTestCase[T any] struct {
	name      string
	input     []T
	predicate Predicate[T]
	expected  []T
}

func TestFilterWithInteger(t *testing.T) {
	tests := []FilterTestCase[int]{
		{
			name:  "Filter less than 3",
			input: []int{1, 2, 3},
			predicate: func(i int) bool {
				return i < 3
			},
			expected: []int{1, 2},
		},
		{
			name:  "No element match",
			input: []int{1, 2, 3},
			predicate: func(i int) bool {
				return i < 0
			},
			expected: []int{},
		},
		{
			name:  "No none-unique elements",
			input: []int{1, 2, 2, 3},
			predicate: func(i int) bool {
				return i < 3
			},
			expected: []int{1, 2, 2},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Filter[int](testCase.input, testCase.predicate)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Failed %s: expected %v, got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}

func TestFilterWithString(t *testing.T) {
	tests := []FilterTestCase[string]{
		{
			name:  "Filter less than 3",
			input: []string{"a", "bb", "ccc"},
			predicate: func(s string) bool {
				return len(s) > 2
			},
			expected: []string{"ccc"},
		},
		{
			name:  "No element match",
			input: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return len(s) < 0
			},
			expected: []string{},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Filter[string](testCase.input, testCase.predicate)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Failed %s: expected %v, got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}

func TestFilterWithStruct(t *testing.T) {
	type People struct {
		name string
		age  int
	}
	tests := []FilterTestCase[People]{
		{
			name: "Filter person with age > 18",
			input: []People{{
				name: "Bob",
				age:  18,
			}, {
				name: "Alice",
				age:  16,
			}},
			predicate: func(p People) bool {
				return p.age >= 18
			},
			expected: []People{{name: "Bob", age: 18}},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Filter[People](testCase.input, testCase.predicate)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Failed %s: expected %v, got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}
