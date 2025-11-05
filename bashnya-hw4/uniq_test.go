package main

import (
	"testing"
)

func TestCreateKey(t *testing.T) {
	testTable := []struct {
		string_in  string
		conf       Config
		string_out string
	}{
		{
			string_in:  "Here we go Again",
			conf:       Config{IgnoreRegister: true},
			string_out: "here we go again",
		},
		{
			string_in:  "Here we go Again",
			conf:       Config{IgnoreRegister: true, SkipFields: 1},
			string_out: "we go again",
		},
		{
			string_in:  "Here we go Again",
			conf:       Config{IgnoreRegister: true, SkipFields: 5},
			string_out: "here we go again",
		},
		{
			string_in:  "Here we go Again",
			conf:       Config{IgnoreRegister: true, SkipFields: 1, SkipChars: 1},
			string_out: "e go again",
		},
		{
			string_in:  "Here we go Again",
			conf:       Config{IgnoreRegister: true, SkipFields: 3, SkipChars: 3},
			string_out: "in",
		},
		{
			string_in:  "",
			conf:       Config{SkipFields: 1},
			string_out: "",
		},
		{
			string_in:  "   here           we   ",
			conf:       Config{SkipFields: 1},
			string_out: "we",
		},
		{
			string_in:  "   here       we   ",
			conf:       Config{SkipFields: 1, SkipChars: 12},
			string_out: "we",
		},
		{
			string_in:  "Hello WORLD",
			conf:       Config{SkipFields: 1},
			string_out: "WORLD",
		},
	}
	for _, testCase := range testTable {
		result := createKey(testCase.string_in, testCase.conf)

		t.Logf("Calling CreateKey(%s,%v), result %s", testCase.string_in, testCase.conf, result)

		if result != testCase.string_out {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.string_out, result)
		}
	}
}

func TestProcessLines(t *testing.T) {
	testTable := []struct {
		strings_in  []string
		conf        Config
		strings_out []string
	}{
		{
			strings_in: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{},
			strings_out: []string{
				"apple\n",
				"banana\n",
				"cherry\n",
			},
		},
		{
			strings_in: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{Count: true},
			strings_out: []string{
				"2 apple\n",
				"1 banana\n",
				"1 cherry\n",
			},
		},
		{
			strings_in: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{Repeat: true},
			strings_out: []string{
				"apple\n",
			},
		},
		{
			strings_in: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{Uniq: true},
			strings_out: []string{
				"banana\n",
				"cherry\n",
			},
		},
		{
			strings_in: []string{
				"APPLE",
				"apple",
				"Banana",
				"BANANA",
			},
			conf: Config{IgnoreRegister: true},
			strings_out: []string{
				"APPLE\n",
				"Banana\n",
			},
		},
		{
			strings_in: []string{
				"We love music",
				"I love music",
				"They love music",
			},
			conf: Config{SkipFields: 1},
			strings_out: []string{
				"We love music\n",
			},
		},
		{
			strings_in: []string{
				"Apple",
				"Bpple",
				"Cpple",
			},
			conf: Config{SkipChars: 1},
			strings_out: []string{
				"Apple\n",
			},
		},
		{
			strings_in: []string{
				"We LOVE music",
				"I love MUSIC",
				"They Love music",
			},
			conf: Config{IgnoreRegister: true, SkipFields: 1},
			strings_out: []string{
				"We LOVE music\n",
			},
		},
		{
			strings_in:  []string{},
			conf:        Config{},
			strings_out: []string{},
		},
		{
			strings_in: []string{
				"same",
				"same",
				"same",
			},
			conf: Config{},
			strings_out: []string{
				"same\n",
			},
		},
		{
			strings_in: []string{
				"APPLE",
				"apple",
				"BANANA",
				"banana",
				"Apple",
			},
			conf: Config{Count: true, IgnoreRegister: true},
			strings_out: []string{
				"3 APPLE\n",
				"2 BANANA\n",
			},
		},
		{
			strings_in: []string{
				"short",
				"a b",
				"only one",
			},
			conf: Config{SkipFields: 5},
			strings_out: []string{
				"short\n",
				"a b\n",
				"only one\n",
			},
		},
		{
			strings_in: []string{
				"a",
				"bc",
				"def",
			},
			conf: Config{SkipChars: 10},
			strings_out: []string{
				"a\n",
				"bc\n",
				"def\n",
			},
		},
		{
			strings_in: []string{
				"third",
				"first",
				"second",
				"first",
				"third",
			},
			conf: Config{},
			strings_out: []string{
				"third\n",
				"first\n",
				"second\n",
			},
		},
		{
			strings_in: []string{
				"hello world",
				"test",
			},
			conf: Config{SkipFields: 3, SkipChars: 10},
			strings_out: []string{
				"hello world\n",
				"test\n",
			},
		},
	}

	for _, tc := range testTable {

		result := processLines(tc.strings_in, tc.conf)
		t.Logf("Calling CreateKey(%s,%v), result %s", tc.strings_in, tc.conf, result)
		for i := range result {
			if result[i] != tc.strings_out[i] {
				t.Errorf("Line %d: expected '%s', got '%s'", i, tc.strings_out[i], result[i])
			}
		}
	}
}
