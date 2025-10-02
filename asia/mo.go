package asia

import "strings"

// IsMO determines whether the phone number belongs to Macao
func (AsiaProcessor) IsMO(phone string) bool {
	// Remove '+'
	phone = strings.TrimPrefix(phone, "+")

	// Check prefix
	if strings.HasPrefix(phone, "853") {
		// Check other parts
		numberPart := phone[3:]

		return len(numberPart) == 8
	}

	return false
}
