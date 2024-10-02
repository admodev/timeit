package domain

type Timer struct {
	WorkDuration  int // En segundos - In seconds
	BreakDuration int // En segundos - In seconds
	RemainingTime int
	IsWorkSession bool
	IsRunning     bool
}

func NewTimer(workDuration, breakDuration int) *Timer {
	return &Timer{
		WorkDuration:  workDuration,
		BreakDuration: breakDuration,
		IsWorkSession: true,
	}
}

func (t *Timer) Start() {
	t.IsRunning = true
	t.RemainingTime = t.WorkDuration
}

func (t *Timer) SwitchSession() {
	t.IsWorkSession = !t.IsWorkSession
	if t.IsWorkSession {
		t.RemainingTime = t.WorkDuration
	} else {
		t.RemainingTime = t.BreakDuration
	}
}

func (t *Timer) Tick() {
	if t.IsRunning && t.RemainingTime > 0 {
		t.RemainingTime--
	}
}
