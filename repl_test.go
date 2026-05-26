package main

import "testing"


func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
	{
		input: "   hello world   ",
		expected: []string{"hello", "world"},
	},
	{
		input: "My     name is   Moaz   ",
		expected: []string{"My", "name", "is", "Moaz"},
	},
}
	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("lengths aren't equal")
			continue
		}
		
		for i:= range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("expected %s recieved %s", expectedWord, word)
			}
		}
	}
}
