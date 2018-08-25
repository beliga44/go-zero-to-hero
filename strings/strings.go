package main

import "fmt"

func main() {
	printBytes("HelloWorld")
	printChar("HelloWorld")
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}
}

func printChar(s string) {
	fmt.Println("************")
	fmt.Println(s[0])
	fmt.Println("************")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
}
