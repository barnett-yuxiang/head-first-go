package chapter02

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Test03() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make a guess")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if guess < target {
		fmt.Println("Low")
	} else if guess > target {
		fmt.Println("High")
	}

}
