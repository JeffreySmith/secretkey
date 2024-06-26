package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/JeffreySmith/secretkey"
)

func main(){
	nFlag := flag.Int("n", 64, "the number of characters to generate.")
	verbose := flag.Bool("v",false, "display more information.")
	flag.Parse()

	n := *nFlag
	
	if n == 0 {
		panic(errors.New("Zero characters is invalid"))
	}

	if *verbose{
		choices := len(secretkey.Characters)
		fmt.Printf("You are generating %d characters, each with %d possibilities each.\n", n, choices)
		fmt.Println("Here is your secret key: ")
	}
	
	
	secret := secretkey.CreateKey(n)
	fmt.Println(secret)
}
