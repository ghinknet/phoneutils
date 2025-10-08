package asia

import "strings"

// IsTW determines whether the phone number belongs to Taiwan
func (AsiaProcessor) IsTW(phone string) bool {
	// Remove '+'
	phone = strings.TrimPrefix(phone, "+")

	// Check prefix
	if strings.HasPrefix(phone, "886") {
		// Check other parts
		numberPart := phone[3:]

		return len(numberPart) >= 8 && len(numberPart) <= 9
	}

	return false
}
