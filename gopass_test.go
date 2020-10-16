package gopass_test

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/servusdei2018/gopass"
)

func BenchmarkGenPass(b *testing.B) {
	b.Run("LEN=66", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GenPass("s3cUr3P@sSw0rd", "gmail", 66)
		}
	})
	b.Run("LEN=16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GenPass("s3cUr3P@sSw0rd", "gmail", 16)
		}
	})
}

func ExampleUsage() {
	fmt.Println(GenPass("password", "uid", 66))
	// Output:
	// 0x9f930eec2ce8f715af44c3d1a502a26958f4bc7dbedb4ae4a6f027620b42c60e
}

func TestGenPass(t *testing.T) {
	checks := [][]string{
		{"password", "pass", "16", "0x9379c98db81650"},
		{"password", "pswd", "16", "0x95288f5cb386b7"},
		{"thisbadD", "abcd", "66", "0x007d98ad94acb44b0b43a79584bc96e51b15f78aa681afa8ec7e9fcc5aef1c09"},
		{"Ddabsiht", "dcba", "66", "0xa428aafb6d48ecbaf095ca6f204357e1106574dec4b218a8ee346b1dcc08b7dd"},
	}

	for _, c := range checks {
		LEN, err := strconv.Atoi(c[2])
		if err != nil {
			t.Error("unexpected error on TestGenpass:", err)
		}

		if got := GenPass(c[0], c[1], LEN); got != c[3] {
			t.Error("want ", c[3], "got", got)
		}
	}
}
