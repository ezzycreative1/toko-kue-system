package generateCodeHelper

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func GenerateRomanDate() string {
	// Get the current time
	now := time.Now()

	// Map for converting month number to Roman numeral
	romanNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
	}

	// Extract the month and year
	month := now.Month()
	year := now.Year()

	// Convert the month to Roman numeral
	romanMonth := romanNumerals[int(month)]

	// Format the result as "IX/2024"
	return fmt.Sprintf("%s/%d", romanMonth, year)
}

func IncrementCode(code string) (string, error) {
	// Define the regular expression to extract the numeric prefix
	re := regexp.MustCompile(`^(\d+)/`)

	// Find the prefix match
	match := re.FindStringSubmatch(code)
	if match == nil || len(match) < 2 {
		return "", fmt.Errorf("no numeric prefix found in the code")
	}

	// Convert the numeric prefix to an integer
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return "", fmt.Errorf("failed to convert numeric prefix to integer: %v", err)
	}

	// Increment the numeric value
	num++

	// Format the new code by replacing the old numeric prefix with the new one
	newCode := fmt.Sprintf("%03d%s", num, code[len(match[1]):])

	return newCode, nil
}
