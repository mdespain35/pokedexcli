package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " Nidoking CrOaGuNk NinJAsK    ",
			expected: []string{"nidoking", "croagunk", "ninjask"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Invalid length of result")
			t.FailNow()
		}
		for i := range actual {
			word, expected := actual[i], c.expected[i]
			if word != expected {
				t.Errorf("Expected: %s: Actual: %s", expected, word)
				t.FailNow()
			}
		}
	}
}
