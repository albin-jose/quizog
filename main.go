package main

import (
	"fmt"
	"strings"
)

func main() {
	var name, subject, v string
	var flag int
	fmt.Println("Welcome to QUIZOG")
	fmt.Printf("\nPlease Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("\nWelcome ", name, ".\nAvailable streams are maths, science and GK.\nAll streams have 10 questions in total.\nEvery question have one mark each")

	stream := [3]string{"maths", "science", "gk"}

	//score := 0

	for {
		fmt.Printf("\nSelect quiz subject:  ")
		fmt.Scanln(&subject)

		for _, v = range stream {
			if strings.Compare(subject, v) == 0 {
				flag = 1
				break
			}

		}
		if flag == 1 {
			break
		} else {
			fmt.Println("\nYou had entered a wrong choice.Please choose Maths, Science or GK")
		}
	}
}
