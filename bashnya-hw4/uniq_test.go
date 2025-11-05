package main

import (
	"testing"
)

func TestCreateKey(t *testing.T) {
	testTable := []struct {
		stringIn  string
		conf      Config
		stringOut string
	}{
		{
			stringIn:  "Here we go Again",
			conf:      Config{IgnoreRegister: true},
			stringOut: "here we go again",
		},
		{
			stringIn:  "Here we go Again",
			conf:      Config{IgnoreRegister: true, SkipFields: 1},
			stringOut: "we go again",
		},
		{
			stringIn:  "Here we go Again",
			conf:      Config{IgnoreRegister: true, SkipFields: 5},
			stringOut: "here we go again",
		},
		{
			stringIn:  "Here we go Again",
			conf:      Config{IgnoreRegister: true, SkipFields: 1, SkipChars: 1},
			stringOut: "e go again",
		},
		{
			stringIn:  "Here we go Again",
			conf:      Config{IgnoreRegister: true, SkipFields: 3, SkipChars: 3},
			stringOut: "in",
		},
		{
			stringIn:  "",
			conf:      Config{SkipFields: 1},
			stringOut: "",
		},
		{
			stringIn:  "   here           we   ",
			conf:      Config{SkipFields: 1},
			stringOut: "we",
		},
		{
			stringIn:  "   here       we   ",
			conf:      Config{SkipFields: 1, SkipChars: 12},
			stringOut: "we",
		},
		{
			stringIn:  "Hello WORLD",
			conf:      Config{SkipFields: 1},
			stringOut: "WORLD",
		},
	}
	for _, testCase := range testTable {
		result := createKey(testCase.stringIn, testCase.conf)

		t.Logf("Calling CreateKey(%s,%v), result %s", testCase.stringIn, testCase.conf, result)

		if result != testCase.stringOut {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.stringOut, result)
		}
	}
}

func TestProcessLines(t *testing.T) {
	testTable := []struct {
		stringsIn  []string
		conf       Config
		stringsOut []string
	}{
		{
			stringsIn: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{},
			stringsOut: []string{
				"apple\n",
				"banana\n",
				"cherry\n",
			},
		},
		{
			stringsIn: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{Count: true},
			stringsOut: []string{
				"2 apple\n",
				"1 banana\n",
				"1 cherry\n",
			},
		},
		{
			stringsIn: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{Repeat: true},
			stringsOut: []string{
				"apple\n",
			},
		},
		{
			stringsIn: []string{
				"apple",
				"banana",
				"apple",
				"cherry",
			},
			conf: Config{Uniq: true},
			stringsOut: []string{
				"banana\n",
				"cherry\n",
			},
		},
		{
			stringsIn: []string{
				"APPLE",
				"apple",
				"Banana",
				"BANANA",
			},
			conf: Config{IgnoreRegister: true},
			stringsOut: []string{
				"APPLE\n",
				"Banana\n",
			},
		},
		{
			stringsIn: []string{
				"We love music",
				"I love music",
				"They love music",
			},
			conf: Config{SkipFields: 1},
			stringsOut: []string{
				"We love music\n",
			},
		},
		{
			stringsIn: []string{
				"Apple",
				"Bpple",
				"Cpple",
			},
			conf: Config{SkipChars: 1},
			stringsOut: []string{
				"Apple\n",
			},
		},
		{
			stringsIn: []string{
				"We LOVE music",
				"I love MUSIC",
				"They Love music",
			},
			conf: Config{IgnoreRegister: true, SkipFields: 1},
			stringsOut: []string{
				"We LOVE music\n",
			},
		},
		{
			stringsIn:  []string{},
			conf:       Config{},
			stringsOut: []string{},
		},
		{
			stringsIn: []string{
				"same",
				"same",
				"same",
			},
			conf: Config{},
			stringsOut: []string{
				"same\n",
			},
		},
		{
			stringsIn: []string{
				"APPLE",
				"apple",
				"BANANA",
				"banana",
				"Apple",
			},
			conf: Config{Count: true, IgnoreRegister: true},
			stringsOut: []string{
				"3 APPLE\n",
				"2 BANANA\n",
			},
		},
		{
			stringsIn: []string{
				"short",
				"a b",
				"only one",
			},
			conf: Config{SkipFields: 5},
			stringsOut: []string{
				"short\n",
				"a b\n",
				"only one\n",
			},
		},
		{
			stringsIn: []string{
				"a",
				"bc",
				"def",
			},
			conf: Config{SkipChars: 10},
			stringsOut: []string{
				"a\n",
				"bc\n",
				"def\n",
			},
		},
		{
			stringsIn: []string{
				"third",
				"first",
				"second",
				"first",
				"third",
			},
			conf: Config{},
			stringsOut: []string{
				"third\n",
				"first\n",
				"second\n",
			},
		},
		{
			stringsIn: []string{
				"hello world",
				"test",
			},
			conf: Config{SkipFields: 3, SkipChars: 10},
			stringsOut: []string{
				"hello world\n",
				"test\n",
			},
		},
	}

	for _, tc := range testTable {

		result := processLines(tc.stringsIn, tc.conf)
		t.Logf("Calling CreateKey(%s,%v), result %s", tc.stringsIn, tc.conf, result)
		for i := range result {
			if result[i] != tc.stringsOut[i] {
				t.Errorf("Line %d: expected '%s', got '%s'", i, tc.stringsOut[i], result[i])
			}
		}
	}
}
