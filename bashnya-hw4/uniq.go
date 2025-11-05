package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	cFlag = flag.Bool("c", false, "count of strings")
	dFlag = flag.Bool("d", false, "only repeat strings")
	uFlag = flag.Bool("u", false, "onle no repeat strings")
	iFlag = flag.Bool("i", false, "ignore register")
	fFlag = flag.Int("f", 0, "missing first num of fields")
	sFlag = flag.Int("s", 0, "missing first num of chars")
)

type Config struct {
	Count          bool
	Repeat         bool
	Uniq           bool
	IgnoreRegister bool
	SkipFields     int
	SkipChars      int
}

func readLines(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func createKey(s string, conf Config) string {
	key := s
	if conf.IgnoreRegister {
		key = strings.ToLower(key)
	}
	if conf.SkipFields > 0 {
		temp := strings.Fields(key)
		if len(temp) > conf.SkipFields {
			temp = temp[conf.SkipFields:]
		}
		key = strings.Join(temp, " ")
	}
	if conf.SkipChars > 0 {
		if len(key) > conf.SkipChars {
			key = key[conf.SkipChars:]
		}
	}
	return key
}

func processLines(lines []string, conf Config) []string {
	seen := make(map[string]bool)
	count := make(map[string]int)
	var uniquePairs []struct {
		original string
		key      string
	}

	for _, line := range lines {
		key := createKey(line, conf)
		count[key]++
		if !seen[key] {
			uniquePairs = append(uniquePairs,
				struct {
					original string
					key      string
				}{line, key})
			seen[key] = true
		}
	}

	var result []string
	for _, pair := range uniquePairs {
		switch {
		case conf.Count:
			result = append(result, fmt.Sprintf("%d %s\n", count[pair.key], pair.original))
		case conf.Repeat:
			if count[pair.key] > 1 {
				result = append(result, pair.original+"\n")
			}
		case conf.Uniq:
			if count[pair.key] == 1 {
				result = append(result, pair.original+"\n")
			}
		default:
			result = append(result, pair.original+"\n")
		}
	}

	return result
}

func main() {
	flag.Parse()

	if (*cFlag && *dFlag) || (*cFlag && *uFlag) || (*uFlag && *dFlag) {
		fmt.Fprintln(os.Stderr, "Error. You cannot use -c,-d,-u flags together")
		os.Exit(1)
	}

	remainingArgs := flag.Args()

	var reader io.Reader

	config := Config{
		Count:          *cFlag,
		Repeat:         *dFlag,
		Uniq:           *uFlag,
		IgnoreRegister: *iFlag,
		SkipFields:     *fFlag,
		SkipChars:      *sFlag,
	}

	if len(remainingArgs) > 0 {
		file, err := os.Open(remainingArgs[0])
		if err != nil {
			fmt.Printf("Error opening file : %s", err)
			return
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	s, err := readLines(reader)
	if err != nil {
		fmt.Printf("Error reading file : %s", err)
		return
	}

	res := processLines(s, config)
	if len(remainingArgs) == 2 {
		file, err := os.Create(remainingArgs[1])
		if err != nil {
			fmt.Printf("Error creating file : %v", err)
			return
		}
		defer file.Close()
		for _, line := range res {
			fmt.Fprint(file, line)
		}

	} else {
		fmt.Println(strings.Join(res, ""))
	}
}
