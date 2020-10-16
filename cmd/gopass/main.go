package main

import (
	"flag"
	"fmt"

	"github.com/servusdei2018/gopass"
)

var (
	PWD string // Master password
	UID string // Unique generation identifier
	LEN int    // Length (<= 66)
)

func init() {
	flag.StringVar(&PWD, "pwd", "", "master password")
	flag.StringVar(&UID, "uid", "", "unique generation identifier")
	flag.IntVar(&LEN, "len", 16, "length (<= 66)")

	flag.Parse()
}

func main() {
	fmt.Println(gopass.GenPass(PWD, UID, LEN))
}
