package romannumerals

import "strings"

type RomanNumeral struct {
	Value uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}


func ConvertToArabic(romanNumeral string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(romanNumeral, numeral.Symbol) {
			arabic += numeral.Value
			romanNumeral = strings.TrimPrefix(romanNumeral, numeral.Symbol)
		}
	}

	return arabic
}