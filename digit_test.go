package main

import (
	"reflect"
	"testing"
)

func TestDigit(t *testing.T) {
	//digit with 1234 return [1,2,3,4]
	//digit with 0 return [0]
	expected := []int{4, 3, 2, 1}
	actual := digit(1234)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(expected, actual)
	}
}
