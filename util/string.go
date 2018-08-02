package util

import (
	"strings"
	"regexp"
	"strconv"
)

func StringTrim(str string) string {
	return strings.TrimSpace(str)
}

func StringSplit(str string) string {
	if str == "None" {
		return "None"
	}

	var (
		stringList []string
		lenght     int
	)

	stringList = strings.Split(str, "/")
	lenght = len(stringList)

	return stringList[lenght-1]
}

func StringReplace(str string) string {
	var NewReplacer *strings.Replacer

	NewReplacer = strings.NewReplacer("\t", "")
	return NewReplacer.Replace(str)
}

func StringSplitByDot(str string) (int, int, int) {
	var stringList []string

	stringList = strings.Split(str, "•")

	var numberList []int
	for index, oneString := range stringList {
		if index > 0 && index < 4 {
			number := StringRegexp(oneString)
			numberList = append(numberList, number)
		}
	}

	return numberList[0], numberList[1], numberList[2]
}

func StringRegexp(str string) int {
	reg := regexp.MustCompile(`\d+`)
	stringContent := reg.FindString(str)
	number, _ := strconv.Atoi(stringContent)
	return number
}
