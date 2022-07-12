package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNum := rand.Intn(maxNum)
	//fmt.Println("The secret number is", secretNum)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin) // 标准IO转换为缓存IO，提高效率
	for {
		input, err := reader.ReadString('\n') // 读取一行的输入
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("Your guess is", guess)

		if guess > secretNum {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNum {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
