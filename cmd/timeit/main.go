package main

import (
	"admodev/timeit/cli"
	"admodev/timeit/domain"
	"admodev/timeit/usecase"
	"fmt"
)

func main() {
	workDuration := 25 * 60 // 25 Minutos - 25 Minutes
	breakDuration := 5 * 60 // 5 Minutos - 5 Minutes
	timer := domain.NewTimer(workDuration, breakDuration)
	pomodoro := usecase.NewPomodoro(timer)

	for {
		cli.PrintMenu()
		choice := cli.GetUserInput()

		switch choice {
		case "1\n":
			fmt.Println("Starting Pomodoro...")
			pomodoro.StartPomodoro()
		case "2\n":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
