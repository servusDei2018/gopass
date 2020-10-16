# gopass
gopass is a package which exports a method to generate unique, repeatable passwords.

The idea behind `gopass` is to have a password manager that doesn't store your passwords -- ever. The only information ever stored will reside in the user.

To avoid having the user remember each and every password, we generate needed passwords on-the-fly by using the secure SHA-256 hash algorithm to hash a string which is the combination of a master password and a unique identifier. The result is a unique and secure password that he can regenerate on-the-fly.

## Features
 - Fast
 - Lightweight (no external dependencies)
 - 100% Cross Platform

## Installation
To install gopass, either download one of the prebuilt binaries from the [releases page](https://github.com/servusDei2018/gopass/releases) or run:

`go install -i github.com/servusDei2018/gopass/cmd/gopass`


## Usage
gopass may be used through the command line or as a library.

#### Command Line
gopass is perfectly usable from the commandline with a speedy-fast, cross-platform and lightweight binary distribution.

Let's say our master password is "password" and we want to generate a password for GitHub:

```
$ gopass -pwd password -uid github
0x58f867aa51d2af
```

There you have a unique password that'll be generated everytime you run gopass with that same password and unique identifier. The only thing you have to remember is your master password, and then use a standard naming convention for your UID's. Voilla: secure passwords stored only in your head.

While 16 characters should be enough, **you may increase the length of the password up to 66 characters** (maximum security) with the `--len` flag:

```
$ gopass -pwd password -uid github -len 66
0x58f867aa51d2af9b2818e76fe7ccaeccfcfaf53d09e83013f7080225f2da0713
```

#### Library
It's simple to add gopass functionality to your Go program:

```
package main

import (
	"fmt"

	"github.com/servusdei2018/gopass"
)

func main() {
	// Generate 66-character password
	fmt.Println(gopass.GenPass("password", "uid", 66))
}
```
