package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
)

func buildCombinedString(
	numDecimal, numOctal, numHexadecimal int,
	pi float64,
	name string,
	isActive bool,
	complexNum complex64,
) string {
	return strconv.Itoa(numDecimal) +
		strconv.Itoa(numOctal) +
		strconv.Itoa(numHexadecimal) +
		strconv.FormatFloat(pi, 'f', -1, 64) +
		name +
		strconv.FormatBool(isActive) +
		fmt.Sprintf("%g", complexNum)
}

func addSaltToRunes(runes []rune, salt []rune) []rune {
	middle := len(runes) / 2
	result := make([]rune, 0, len(runes)+len(salt))
	result = append(result, runes[:middle]...)
	result = append(result, salt...)
	result = append(result, runes[middle:]...)
	return result
}

func hashRunesSHA256(runes []rune) [32]byte {
	return sha256.Sum256([]byte(string(runes)))
}

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Println(reflect.TypeOf(numDecimal))
	fmt.Println(reflect.TypeOf(numOctal))
	fmt.Println(reflect.TypeOf(numHexadecimal))
	fmt.Println(reflect.TypeOf(pi))
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(isActive))
	fmt.Println(reflect.TypeOf(complexNum))

	combined := buildCombinedString(
		numDecimal, numOctal, numHexadecimal,
		pi, name, isActive, complexNum,
	)

	runes := []rune(combined)
	salt := []rune("go-2024")
	runesWithSalt := addSaltToRunes(runes, salt)

	hash := hashRunesSHA256(runesWithSalt)
	fmt.Printf("%x\n", hash)
}
