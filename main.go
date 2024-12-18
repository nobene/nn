package main

import (
	"strings"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/ssh/terminal"
//	"os"
	"fmt"
)


func main() {
        salt := "here"
	fmt.Println("Your input ?")
	input,_ := terminal.ReadPassword(0)
	if input[len(input) - 1] == 47 {// if last character of token is "/", then we will print token on stdout
	input = input[:len(input) - 1]
	fmt.Println("                               ", string(input), "\n")
	}

//	fmt.Println(string(input))
//	fmt.Println("Salt: ", salt)
	key := argon2.IDKey([]byte(input), []byte(salt), 16, 512*1024, 3, 47)
	re := base64.StdEncoding.EncodeToString(key)
	re = strings.Replace(re, "+", "", -1)
	re = strings.Replace(re, "/", "", -1)
	re = strings.Replace(re, "=", "", -1)
	fmt.Println("\n            ", re, "    ", len(re), "\n\n")
}
