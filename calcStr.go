package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	var logKeyEnd, _ = utf8.DecodeRuneInString(str[len(str)-2:])
	var logKeyStart, _ = utf8.DecodeRuneInString(str[0:1])

	switch {
	case logKeyStart != '"':
		fmt.Println("Ошибка. Недопустимый формат ввода")
		return

	}

	if logKeyEnd == '"' {
		strings.TrimFunc(str, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		masStr := strings.Split(str, "\"")

		var str1, sign, str2 = masStr[1], masStr[2], masStr[3]
		sign = strings.TrimSpace(sign)

		switch {
		case len(str1) > 10:
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		case len(str2) > 10:
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		case sign == "*":
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		case sign == "/":
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		}

		if sign == "+" {
			fmt.Println("\"" + str1 + str2 + "\"")

		} else if sign == "-" {
			fmt.Println("\"" + strings.ReplaceAll(str1, str2, "") + "\"")
		}

	} else {
		str = strings.TrimLeft(str, "\"")
		masStr := strings.Split(str, "\"")

		var str1 = masStr[0]

		var tempStr = masStr[1]
		tempStr1 := tempStr[1:]

		tempStr2 := strings.Fields(tempStr1)
		var sign, str2 = tempStr2[0], tempStr2[1]

		str2Int, _ := strconv.Atoi(str2)

		switch {
		case str2Int < 1 || str2Int > 10:
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		case len(str1) > 10:
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		case sign == "+":
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		case sign == "-":
			fmt.Println("Ошибка. Недопустимый формат ввода")
			return
		}

		if sign == "*" {
			mul := strings.Repeat(str1, str2Int)
			if len(mul) > 40 {
				fmt.Println("\"" + mul[0:40] + "..." + "\"")
			} else {
				fmt.Println("\"" + mul + "\"")
			}

		} else if sign == "/" {

			inputFmt := str1[0:(len(str1) / 3)]
			fmt.Println("\"" + inputFmt + "\"")
		}

	}

}
