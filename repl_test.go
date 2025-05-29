package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected	[]string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	//loop over cases and run dem tests :0
	for _, c := range cases {
		actual := cleanInput(c.input)

		if actual == nil && c.expected != nil {
			t.Error("yo, output of cleanInput in nil when its not supposed to be")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("ERROR:Words no match \n actual: %v\n expected: %v\n", word, expectedWord)
			}
		}
	}
}

