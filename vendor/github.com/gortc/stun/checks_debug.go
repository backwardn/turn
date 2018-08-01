// +build debug

package stun

import "github.com/gortc/stun/internal/hmac"

// CheckSize returns *AttrLengthError if got is not equal to expected.
func CheckSize(a AttrType, got, expected int) error {
	if got == expected {
		return nil
	}
	return &AttrLengthErr{
		Got:      got,
		Expected: expected,
		Attr:     a,
	}
}

func checkHMAC(got, expected []byte) error {
	if hmac.Equal(got, expected) {
		return nil
	}
	return &IntegrityErr{
		Expected: expected,
		Actual:   got,
	}
}

func checkFingerprint(got, expected uint32) error {
	if got == expected {
		return nil
	}
	return &CRCMismatch{
		Actual:   got,
		Expected: expected,
	}
}

// IsAttrSizeInvalid returns true if error means that attribute size is invalid.
func IsAttrSizeInvalid(err error) bool {
	_, ok := err.(*AttrLengthErr)
	return ok
}
