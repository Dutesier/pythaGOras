package main

import (
	"sync"
	"time"
)

var WG = sync.WaitGroup{}

const ammount = 3

func main() {

	t := Times{
		timeToDie:   time.Duration(200 * time.Millisecond),
		timeToEat:   time.Duration(100 * time.Millisecond),
		timeToSleep: time.Duration(50 * time.Millisecond),
		creation:    time.Now(),
	}

	var ph Philo
	forkMuts := make(map[int]*sync.RWMutex, ammount)
	forks := make(map[int]*bool, ammount)
	for i := 0; i < ammount; i++ {
		forkMuts[i] = &sync.RWMutex{}
		forks[i] = new(bool)
	}
	fck := new(bool)
	fckMut := &sync.RWMutex{}
	WG.Add(ammount)
	for i := 0; i < ammount; i++ {
		switch i {
		case 0:
			ph.leftFork = forks[ammount-1]
			ph.leftForkMut = forkMuts[ammount-1]
			ph.rightFork = forks[i]
			ph.rightForkMut = forkMuts[i]
		case ammount - 1:
			ph.leftFork = forks[i-1]
			ph.leftForkMut = forkMuts[i-1]
			ph.rightFork = forks[0]
			ph.rightForkMut = forkMuts[0]
		default:
			ph.leftFork = forks[i-1]
			ph.leftForkMut = forkMuts[i-1]
			ph.rightFork = forks[i]
			ph.rightForkMut = forkMuts[i]
		}
		ph.wait = WG
		ph.lastMeal = time.Now()
		ph.durations = &t
		ph.name = string(i)
		ph.status = thinking
		ph.fck = fck
		ph.fckMut = fckMut
		go dinner(ph)
	}
	WG.Wait()
}

func dinner(ph Philo) {
	for {
		if ph.fullWhen > 0 {
			if ph.fullWhen >= ph.timesEaten {
				break
			}
		}
		if time.Since(ph.lastMeal) >= ph.durations.timeToDie {
			ph.Die(ph.fckMut)
			break
		}
		if ph.status == thinking {
			ph.TryEat()
		}
		if ph.status == eating {
			ph.Sleep(ph.durations.timeToSleep)
		}
		if ph.status != thinking && ph.status != dead {
			ph.Think()
		}

	}
}
