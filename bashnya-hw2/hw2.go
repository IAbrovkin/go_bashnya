package main

import (
	"fmt"
)

func newNumber(n int) int {
	for n < 12307 {
		if n < 0 {
			n *= -1
		} else if n%7 == 0 {
			n *= 39
		} else if n%9 == 0 {
			n *= 13
			n += 1
			continue
		} else {
			n += 2
			n *= 3
		}

		if n%9 == 0 && n%13 == 0 {
			break
		} else {
			n += 1
		}
	}
	return n
}

func testing() {
	maxX := -100000
	minX := 100000000000
	var a int
	for i := -100000; i < 100000; i++ {
		a = newNumber(i)
		if a < minX {
			minX = a
		}
		if a > maxX {
			maxX = a
		}
	}
	fmt.Printf("Min number of func is %d. Max number of func is %d", minX, maxX)
}

func numberToText(x int) string {
	var res string
	x_start := x
	a1 := x % 10
	switch a1 {
	case 1:
		res = "один " + res
	case 2:
		res = "два " + res
	case 3:
		res = "три " + res
	case 4:
		res = "четыре " + res
	case 5:
		res = "пять " + res
	case 6:
		res = "шесть " + res
	case 7:
		res = "семь " + res
	case 8:
		res = "восемь " + res
	case 9:
		res = "девять " + res
	}
	x /= 10

	a2 := x % 10
	switch a2 {
	case 1:
		switch a1 {
		case 0:
			res = "десять"
		case 1:
			res = "одиннадцать"
		case 2:
			res = "двенадцать"
		case 3:
			res = "тринадцать"
		case 4:
			res = "четырныдцать"
		case 5:
			res = "пятнадцать"
		case 6:
			res = "шестнадцать"
		case 7:
			res = "семнадцать"
		case 8:
			res = "восемнадцать"
		case 9:
			res = "девятнадцать"
		}
	case 2:
		res = "двадцать " + res
	case 3:
		res = "тридцать " + res
	case 4:
		res = "сорок " + res
	case 5:
		res = "пятьдесят " + res
	case 6:
		res = "шестьдясят " + res
	case 7:
		res = "семьдесят " + res
	case 8:
		res = "восемьдесят " + res
	case 9:
		res = "девяносто " + res
	}
	x /= 10

	a3 := x % 10
	switch a3 {
	case 1:
		res = "сто " + res
	case 2:
		res = "двести " + res
	case 3:
		res = "триста " + res
	case 4:
		res = "четыреста " + res
	case 5:
		res = "пятьсот " + res
	case 6:
		res = "шестьсот " + res
	case 7:
		res = "семьсот " + res
	case 8:
		res = "восемьсот " + res
	case 9:
		res = "девятьсот " + res
	}
	x /= 10

	a4 := x % 10
	if x_start < 10000 || x_start > 19999 {
		switch a4 {
		case 0:
			res = "тысяч " + res
		case 1:
			res = "одна тысяча " + res
		case 2:
			res = "две тысячи " + res
		case 3:
			res = "три тысячи " + res
		case 4:
			res = "четыре тысячи " + res
		case 5:
			res = "пять тысяч " + res
		case 6:
			res = "шесть тысяч " + res
		case 7:
			res = "семьтысяч " + res
		case 8:
			res = "восемь тысяч " + res
		case 9:
			res = "девять тысяч " + res
		}
	}
	x /= 10

	a5 := x % 10
	switch a5 {
	case 1:
		switch a4 {
		case 0:
			res = "десять тысяч "
		case 1:
			res = "одиннадцать тысяч " + res
		case 2:
			res = "двенадцать тысяч " + res
		case 3:
			res = "тринадцать тысяч " + res
		case 4:
			res = "четырныдцать тысяч " + res
		case 5:
			res = "пятнадцать тысяч " + res
		case 6:
			res = "шестнадцать тысяч " + res
		case 7:
			res = "семнадцать тысяч " + res
		case 8:
			res = "восемнадцать тысяч " + res
		case 9:
			res = "девятнадцать тысяч " + res
		}
	case 2:
		res = "двадцать " + res
	case 3:
		res = "тридцать " + res
	case 4:
		res = "сорок " + res
	case 5:
		res = "пятьдесят " + res
	case 6:
		res = "шестьдясят " + res
	case 7:
		res = "семьдесят " + res
	case 8:
		res = "восемьдесят " + res
	case 9:
		res = "девяносто " + res
	}
	x /= 10

	a6 := x % 10
	switch a6 {
	case 1:
		res = "сто " + res
	case 2:
		res = "двести " + res
	case 3:
		res = "триста " + res
	case 4:
		res = "четыреста " + res
	case 5:
		res = "пятьсот " + res
	case 6:
		res = "шестьсот " + res
	case 7:
		res = "семьсот " + res
	case 8:
		res = "восемьсот " + res
	case 9:
		res = "девятьсот " + res
	}
	x /= 10
	return res
}

func main() {
	var n int
	flag := false

	fmt.Printf("Input a number: ")
	fmt.Scan(&n)

	if n >= 12307 {
		flag = true
	}
	fmt.Printf("Your number is %d\n", n)

	n = newNumber(n)
	var s = numberToText(n)
	if n < 12307 {
		fmt.Printf("Service error\n")
	} else if flag {
		fmt.Println("Your number is greater than 12307, so it remains unchanged")
	} else {
		fmt.Printf("New number is %s (%d)\n ", s, n)
	}

}
