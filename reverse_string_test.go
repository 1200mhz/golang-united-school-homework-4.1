package reverse_string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "Hell low world!"

	actual := ReverseString(s)
	expected := "!dlrow wol lleH"

	if !reflect.DeepEqual(actual, expected) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}

func TestCyrillicEmoReverseString(t *testing.T) {
	s := "С ДР! 😸"

	actual := ReverseString(s)
	expected := "😸 !РД С"

	if !reflect.DeepEqual(actual, expected) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}
