package philosopher

import (
	"fmt"
	"sync"
	"time"

	"github.com/Dutesier/pythaGOras/src/main"
)

type eater interface {
	Eat(duration time.Duration)
}

type sleeper interface {
	Sleep(duration time.Duration)
}

type thinker interface {
	Think()
}

type printer interface {
	Print(msg string)
}

// I know this means something else
type dyer interface {
	Die(fckMut *sync.RWMutex)
	WillIDie(lastMeal time.Time, timeToDie, timeToExecute time.Duration) (bool, time.Duration)
}

type philosopher interface {
	eater
	sleeper
	thinker
	dyer
	printer
}

const (
	dead = iota
	eating
	sleeping
	thinking
)

type Times struct {
	timeToDie   time.Duration
	timeToEat   time.Duration
	timeToSleep time.Duration
	creation    time.Time
}

type Philo struct {
	// Identinfier
	name string
	// Forks
	rightForkMut *sync.RWMutex
	rightFork    *bool
	leftForkMut  *sync.RWMutex
	leftFork     *bool
	// Global Status
	fckMut *sync.RWMutex
	fck    *bool
	// Times (shared by all philos)
	durations *Times
	// Personal Time
	lastMeal time.Time
	// Stats
	status     uint
	timesEaten uint
	fullWhen   int
}

func (ph *Philo) Eat(duration time.Duration) {
	ph.Print(ph.name + "is eating")
	time.Sleep(duration)
	ph.lastMeal = time.Now()
	ph.timesEaten++
}

func (ph *Philo) Sleep(duration time.Duration) {
	d, when := ph.WillIDie(ph.lastMeal, ph.durations.timeToDie, duration)
	if d {
		ph.Print(ph.name + "is sleeping")
		time.Sleep(time.Duration(when))
		ph.Die(ph.fckMut)
	}
	ph.Print(ph.name + "is sleeping")
	time.Sleep(duration)
}

func (ph *Philo) Think() {
	ph.Print(ph.name + "is thinking")
	ph.status = thinking
}

func (ph *Philo) WillIDie(lastMeal time.Time, timeToDie, timeToExecute time.Duration) (bool, time.Duration) {
	when := time.Since(lastMeal) + time.Duration(timeToDie)
	if when <= time.Duration(timeToExecute) {
		return true, when
	}
	return false, 0
}

func (ph *Philo) Die(fckMut *sync.RWMutex) {
	fckMut.Lock()
	*ph.fck = true
	fckMut.Unlock()
	ph.status = dead
	ph.Print(ph.name + "has died!")
	main.WG.Done()
}

func (ph *Philo) Print(msg string) {
	ph.fckMut.Lock()
	if *ph.fck == false || ph.status == dead {
		fmt.Println(time.Since(ph.durations.creation), msg)
	}
	ph.fckMut.Unlock()
}
