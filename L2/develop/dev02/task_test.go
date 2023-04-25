package main

//import "unicode"
import "testing"

func TestСulc(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""}, // (некорректная строка)
		{"", ""},   // (некорректная строка)
		{"as222", ""},
		{"a1s2d3f4g5h6j7j8k9l1", "assdddffffggggghhhhhhjjjjjjjjjjjjjjjkkkkkkkkkl"},
	}
	for _, test := range tests {
		if res, _ := Culc(test.input); res != test.output {
			t.Errorf("Decode(%s) = %s", test.input, test.output)
		}
	}
}
