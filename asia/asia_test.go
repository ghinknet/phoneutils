package asia

import "testing"

func TestAsiaProcessor_IsChinaMainland(t *testing.T) {
	if !Asia.IsChinaMainland("8619537806666") {
		t.Error("example 8619537806666 can not passed the test")
		return
	}
	if Asia.IsChinaMainland("8629537806666") {
		t.Error("example 8629537806666 passed the test unexpectedly")
		return
	}
	if Asia.IsChinaMainland("869537806666") {
		t.Error("example 869537806666 passed the test unexpectedly")
		return
	}
}

func TestAsiaProcessor_IsHK(t *testing.T) {
	if !Asia.IsHK("85290608282") {
		t.Error("example 85290608282 can not passed the test")
		return
	}
	if Asia.IsHK("8520608282") {
		t.Error("example 8520608282 passed the test unexpectedly")
		return
	}
}

func TestAsiaProcessor_IsMO(t *testing.T) {
	if !Asia.IsMO("85390608282") {
		t.Error("example 85390608282 can not passed the test")
		return
	}
	if Asia.IsMO("8530608282") {
		t.Error("example 8530608282 passed the test unexpectedly")
		return
	}
}

func TestAsiaProcessor_IsTW(t *testing.T) {
	if !Asia.IsTW("886912345678") {
		t.Error("example 886912345678 can not passed the test")
		return
	}
	if !Asia.IsTW("88612345678") {
		t.Error("example 886912345678 can not passed the test")
		return
	}
	if Asia.IsTW("8861234567") {
		t.Error("example 8861234567 passed the test unexpectedly")
		return
	}
}

func TestAsiaProcessor_IsJP(t *testing.T) {
	if !Asia.IsJP("8107031659666") {
		t.Error("example 8107031659666 can not passed the test")
		return
	}
	if !Asia.IsJP("817031659666") {
		t.Error("example 817031659666 can not passed the test")
		return
	}
	if Asia.IsJP("81031659666") {
		t.Error("example 81031659666 passed the test unexpectedly")
		return
	}
}
