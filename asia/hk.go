package asia

import "strings"

// IsHK determines whether the phone number belongs to Hongkong
func (AsiaProcessor) IsHK(phone string) bool {
	// Remove '+'
	phone = strings.TrimPrefix(phone, "+")

	// Check prefix
	if strings.HasPrefix(phone, "852") {
		// Check other parts
		numberPart := phone[3:]

		return len(numberPart) == 8
	}

	return false
}
