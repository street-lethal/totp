package main

import (
	"fmt"
	"time"
	"totp/src"
)

func main() {
	secret, err := src.LoadFromFile("./data/secret.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bin, err := secret.Base32()
	if err != nil {
		fmt.Println(err)
		return
	}

	totp := src.NewTOTP()
	fmt.Println(totp.Calc(bin))

	fmt.Printf("現在: %02d秒\n", time.Now().Second())
}
