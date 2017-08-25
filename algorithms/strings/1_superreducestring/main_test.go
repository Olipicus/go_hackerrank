package main

import "testing"

func TestSolution(t *testing.T) {

	tt := []struct {
		input  string
		result string
	}{
		{input: "aaabccddd", result: "abd"},
		{input: "baab", result: ""},
		{input: "kagoyzkgfjnyvjewazalxngpdcfahneqoqgiyjgpzobhaghmgzmnwcmeykqzgajlmuerhhsanpdtmrzibswswzjjbjqytgfewiuu", result: "kagoyzkgfjnyvjewazalxngpdcfahneqoqgiyjgpzobhaghmgzmnwcmeykqzgajlmuersanpdtmrzibswswzbjqytgfewi"},
	}

	for index, tc := range tt {
		if r := Sub(tc.input); r != tc.result {
			t.Errorf("Error on case[%d] input[%s] result[%s] expect[%s]", index, tc.input, r, tc.result)
		}
	}

}
