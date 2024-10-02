package cli

import (
	"bufio"
	"fmt"
	"os"
)

func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func PrintMenu() {
	fmt.Println("TimeIT! - Pomodoro Timer")
	fmt.Println("1. Start")
	fmt.Println("2. Quit")
	fmt.Print("Choose an option: ")
}
