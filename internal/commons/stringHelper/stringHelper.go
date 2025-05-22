package stringHelper

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"
)

func GenerateRandString(lg int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, lg)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func RemoveDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func SortSlice(arr []string) []string {
	arr1 := arr
	sort.Slice(arr1, func(i1, i2 int) bool {
		numA := cast.ToInt(arr1[i1])
		numB := cast.ToInt(arr1[i2])
		return numA < numB
	})
	return arr1
}

var (
	LowerCaseRune   = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseRune   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NumericRune     = []rune("1234567890")
	SpecialCharRune = []rune("@$!%*#?&")
)

func Generate(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "ExcbsVQs"

	return str
}

func GenerateNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := NumericRune

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "1234567890"

	return str
}

func GenerateAlphabet(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := LowerCaseRune

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "123"

	return str
}

func GenerateCapitalAlphabet(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := UpperCaseRune

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "123"

	return str
}

func GenerateAlphanumeric(length, numOfNumeric int) (string, error) {
	if length < 3 {
		return "", errors.New("minimum len is 3")
	}

	if numOfNumeric > length-2 {
		return "", errors.New("number of numeric character exceed (should not be more than length-2)")
	}

	randomNumeric := GenerateNumber(numOfNumeric)

	remainingLength := length - numOfNumeric

	//determine random number of uppercase alphabet
	upperCaseLen := rand.Intn(remainingLength/2) + 1
	randomUpperCase := GenerateCapitalAlphabet(upperCaseLen)

	//determine random number of lowercase alphabet
	lowerCaseLen := remainingLength - upperCaseLen
	randomLowerCase := GenerateAlphabet(lowerCaseLen)

	str := randomNumeric + randomUpperCase + randomLowerCase

	//shuffle string
	inRune := []rune(str)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune), nil
}

func GenerateRandomPassword(length int) (password string) {
	rand.Seed(time.Now().UnixNano())
	allChars := append(LowerCaseRune, UpperCaseRune...)
	allChars = append(allChars, NumericRune...)
	allChars = append(allChars, SpecialCharRune...)

	b := make([]rune, length)
	// select 1 upper, 1 lower, 1 number and 1 special
	b[0] = LowerCaseRune[rand.Intn(len(LowerCaseRune))]
	b[1] = UpperCaseRune[rand.Intn(len(UpperCaseRune))]
	b[2] = NumericRune[rand.Intn(len(NumericRune))]
	b[3] = SpecialCharRune[rand.Intn(len(SpecialCharRune))]
	for i := 4; i < length; i++ {
		// randomly select 1 character from given charset
		b[i] = allChars[rand.Intn(len(allChars))]
	}

	//shuffle character
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
	return string(b)
}

func GenerateBrandCode(lastCode *string) string {
	var nextNumber int

	// Jika lastCode nil atau kosong, mulai dari 1
	if lastCode == nil || *lastCode == "" {
		nextNumber = 1
	} else {
		code := *lastCode // Dereference pointer

		// Pastikan kode memiliki format yang cukup panjang
		if len(code) < 3 || code[:2] != "B-" {
			return "" // Format tidak valid, return kosong
		}

		// Extract angka dari lastCode (format "B-XX")
		numStr := code[2:]
		parsedNum, err := strconv.Atoi(numStr)
		if err != nil {
			return "" // Return empty jika parsing gagal
		}
		nextNumber = parsedNum + 1
	}

	// Format kode baru
	return fmt.Sprintf("B-%02d", nextNumber)
}

func GenerateVehicleCode(lastCode *string) string {
	var nextNumber int

	// Jika lastCode nil atau kosong, mulai dari 1
	if lastCode == nil || *lastCode == "" {
		nextNumber = 1
	} else {
		code := *lastCode // Dereference pointer

		// Pastikan kode memiliki format yang cukup panjang dan prefix yang benar
		if len(code) < 6 || code[:2] != "V-" {
			return "" // Format tidak valid, return kosong
		}

		// Extract angka dari lastCode (format "V-XXXX")
		numStr := code[2:]
		parsedNum, err := strconv.Atoi(numStr)
		if err != nil {
			return "" // Return empty jika parsing gagal
		}
		nextNumber = parsedNum + 1
	}

	// Format kode baru dengan padding nol (4 digit)
	return fmt.Sprintf("V-%04d", nextNumber)
}

func GenerateAreaCode(lastCode *string) string {
	var nextNumber int

	// Jika lastCode nil atau kosong, mulai dari 0
	if lastCode == nil || *lastCode == "" {
		nextNumber = 0
	} else {
		code := *lastCode // Dereference pointer

		// Pastikan kode memiliki format yang benar (minimal "A-0")
		if len(code) < 3 || code[:2] != "A-" {
			return "" // Format tidak valid, return kosong
		}

		// Ambil angka setelah "A-"
		numStr := code[2:]
		parsedNum, err := strconv.Atoi(numStr)
		if err != nil {
			return "" // Return empty jika parsing gagal
		}
		nextNumber = parsedNum + 1
	}

	// Format kode baru tanpa padding nol
	return fmt.Sprintf("A-%d", nextNumber)
}

func GenerateChannelCode(lastCode *string) string {
	var nextNumber int

	// Jika lastCode nil atau kosong, mulai dari 1
	if lastCode == nil || *lastCode == "" {
		nextNumber = 1
	} else {
		code := *lastCode // Dereference pointer

		// Pastikan kode memiliki format yang cukup panjang dan prefix yang benar
		if len(code) < 6 || code[:2] != "C-" {
			return "" // Format tidak valid, return kosong
		}

		// Extract angka dari lastCode (format "V-XXXX")
		numStr := code[2:]
		parsedNum, err := strconv.Atoi(numStr)
		if err != nil {
			return "" // Return empty jika parsing gagal
		}
		nextNumber = parsedNum + 1
	}

	// Format kode baru dengan padding nol (4 digit)
	return fmt.Sprintf("V-%04d", nextNumber)
}
