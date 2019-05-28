package main

import (
	"testing"
	"os/exec"
	"fmt"
	"bytes"
)

var testCases = []struct {
	input    string
	key		 string
	expected string
}{
{
		input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		key:	 		"cryptii",
		expected: "XYZABCDEFGHIJKLMNOPQRSTUVW",
	},
	{
		input:    "ATTACKATONCE",
		key:	  "secretkey",
		expected: "SXVRGDKXMFGG",
	},
	{
		input:    "cipher",
		key:	  "supersecret",
		expected: "ucelvj",
	},
	{
		input:    "caesar",
		key:	  "test",
		expected: "vewltv",
	},
	{
		input:    "TESTlolKek",
		key:	  "KrOp",
		expected: "DVGIvfzZob",
	},
}

func TestVigenere(t *testing.T) {
	var actual bytes.Buffer
	for _, tt := range testCases {
		command := fmt.Sprintf("vigenere -e --key %s %s", tt.key, tt.input)
		vigenere := exec.Command(command)
		vigenere.Stdout = &actual
		err := vigenere.Run()

		if err != nil {
			t.Fatalf("Vigenere(%q) returned error %q.  Error not expected.", tt.input, err)
		}
		if actual.String() != tt.expected {
			t.Fatalf("Vigenere(%q) was expected to return %v but returned %v.",
				tt.input, tt.expected, actual)
		}
	}
}
