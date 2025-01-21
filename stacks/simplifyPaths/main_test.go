package main

import "testing"

func Test_SimplifyPath(t *testing.T) {
	cases := []struct {
		Name     string
		Path     string
		Expected string
	}{
		{
			Name:     "Empty path",
			Path:     "",
			Expected: "",
		},
		{
			Name:     "test1",
			Path:     "/home/..",
			Expected: "/home/",
		},
	}

	for _, test := range cases {
		result := SimplifyPath(test.Path)
		t.Logf("Test: %s, result: %s", test.Name, result)
		if result != test.Expected {
			t.Errorf("Expected: %s, result: %s", test.Expected, result)
		}
	}
}
