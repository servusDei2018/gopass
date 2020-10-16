package gopass

import (
	"crypto/sha256"
	"fmt"
)

// GenPass generates a unique password from PWD and UID of length LEN (max length is 66)
func GenPass(PWD, UID string, LEN int) string {
	if LEN > 66 {
		LEN = 66
	}

	return fmt.Sprintf("%#x", sha256.Sum256([]byte(PWD+UID)))[0:LEN]
}
