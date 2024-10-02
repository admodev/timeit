package usecase

import (
	"fmt"
	"time"

	"admodev/timeit/domain"
)

type Pomodoro struct {
	Timer *domain.Timer
}

func NewPomodoro(timer *domain.Timer) *Pomodoro {
	return &Pomodoro{
		Timer: timer,
	}
}

func (p *Pomodoro) StartPomodoro() {
	p.Timer.Start()
	for p.Timer.RemainingTime > 0 {
		p.Timer.Tick()
		fmt.Printf("\rTime left: %d seconds", p.Timer.RemainingTime)
		time.Sleep(1 * time.Second)
	}
	p.Timer.SwitchSession()
	fmt.Println("\nSession ended!")
}
