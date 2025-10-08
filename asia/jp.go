package asia

import "strings"

// IsJP determines whether the phone number belongs to Japan
func (AsiaProcessor) IsJP(phone string) bool {
	// Remove '+'
	phone = strings.TrimPrefix(phone, "+")

	if strings.HasPrefix(phone, "81") {
		// Check other parts
		numberPart := phone[2:]

		// Check length
		return len(numberPart) == 10 || len(numberPart) == 11
	}

	return false
}
