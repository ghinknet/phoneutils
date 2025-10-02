package asia

import "strings"

// IsChinaMainland determines whether the phone number belongs to China Mainland
func (AsiaProcessor) IsChinaMainland(phone string) bool {
	// Remove '+'
	phone = strings.TrimPrefix(phone, "+")

	// Check prefix
	if !strings.HasPrefix(phone, "86") {
		return false
	}

	// Check other parts
	numberPart := phone[2:]

	// Check length
	if len(numberPart) != 11 {
		return false
	}

	// Check unit
	if numberPart[0] != '1' {
		return false
	}

	return true
}
