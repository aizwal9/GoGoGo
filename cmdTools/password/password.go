package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes = "0123456789"
	specialBytes = "!@#$%^&*()_-+={}[]:;<>,.?/~`"
)

func generatePassword(length int,exclude string) string {

	allBytes := letterBytes + numberBytes + specialBytes

	// Remove excluded characters
	for _, char := range exclude {
		allBytes = strings.ReplaceAll(allBytes, string(char), "")
	}

	password := make([]byte,length)
	for i := range password{
		password[i] = allBytes[rand.Intn(len(allBytes))]
	}
	return string(password)
}

func main(){

	length := flag.Int("length",12,"password length")
	numPasswords := flag.Int("num",1,"number of passwords to generate")
	exclude := flag.String("exclude","","characters to exclude")

	flag.Parse()
	
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<*numPasswords;i++ {
		password := generatePassword(*length,*exclude)
		fmt.Println(password)
	}
}