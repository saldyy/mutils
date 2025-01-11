package main

import (
	"reflect"
	"strconv"
	"testing"
)

type MapTestCase[T, S any] struct {
	name       string
	input      []T
	transfomer Transfomer[T, S]
	expected   []S
}

func TestMapIntToInt(t *testing.T) {
	tests := []MapTestCase[int, int]{
		{
			name:     "Map integer in slice multiply by 2",
			input:    []int{1, 2, 3},
			expected: []int{2, 4, 6},
			transfomer: func(t int) int {
				return t * 2
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Map[int, int](testCase.input, testCase.transfomer)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Failed %s: expected %v, got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}

func TestMapIntToString(t *testing.T) {
	tests := []MapTestCase[int, string]{
		{
			name:     "Map integer in slice to string",
			input:    []int{1, 2, 3},
			expected: []string{"1a", "2a", "3a"},
			transfomer: func(t int) string {
				return strconv.Itoa(t) + "a"
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Map[int, string](testCase.input, testCase.transfomer)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Failed %s: expected %v, got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}

func TestMapStruct(t *testing.T) {
	type People struct {
		name string
		age  int
	}
	tests := []MapTestCase[People, People]{
		{
			name:     "Map struct People by masking name if person under 18",
			input:    []People{{name: "John", age: 19}, {name: "Alice", age: 16}},
			expected: []People{{name: "John", age: 19}, {name: "***", age: 16}},
			transfomer: func(t People) People {
				n := t.name
				if t.age < 18 {
					n = "***"
				}
				return People{
					name: n,
					age:  t.age,
				}
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Map[People, People](testCase.input, testCase.transfomer)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Failed %s: expected %v, got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}
