package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var name, subject, v, ans string
	var flag int
	fmt.Println("Welcome to QUIZOG")
	fmt.Printf("\nPlease Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("\nWelcome ", name, ".\nAvailable streams are maths, science and GK.\nAll streams have 10 questions in total.\nEvery question have one mark each")

	stream := [3]string{"maths", "science", "gk"}

	score := 0

	for {
		fmt.Printf("\nSelect quiz subject:  ")
		fmt.Scanln(&subject)
		subject = strings.ToLower(subject)

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

	filename := "subjects/" + subject + ".csv"

	csvFile, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	csvReader := csv.NewReader(csvFile)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(row[0] + " : ")
		fmt.Scanln(&ans)
		if strings.Compare(ans, row[1]) == 0 {
			score += 1
		}
	}

	fmt.Println("\nThe total score is ", score)

}
