package main

import (
	"sync"
	"time"

	"github.com/Dutesier/pythaGOras/src/philosopher"
)

var WG = sync.WaitGroup{}

const ammount = 3

func main() {

	t := philosopher.Times{
		timeToDie:   time.Duration(200 * time.Millisecond),
		timeToEat:   time.Duration(100 * time.Millisecond),
		timeToSleep: time.Duration(50 * time.Millisecond),
		creation:    time.Now(),
	}

	var ph philosopher.Philo
	var forkMuts map[int]sync.RWMutex
	var forks map[int]*bool
	for i := 0; i < ammount; i++ {
		forkMuts[i] = sync.RWMutex{}
		*forks[i] = false
	}

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
		ph.Times = t
		ph.name = string(i)
		ph.status = philosopher.thinking
		go dinner(ph)
	}
	WG.Wait()
}

func dinner(ph philosopher.Philo) {
	for {
		if time.Since(ph.lastMeal) >= ph.timeToDie {
			ph.Die()
			break
		}
		if ph.status == philosopher.thinking {
			ph.TryEat()
		}
		if ph.status == philosopher.eating {
			ph.Sleep(ph.Times.timeToSleep)
		}
		if ph.status != philosopher.thinking && ph.status != philosopher.dead {
			ph.Think()
		}

	}
}

func GetWG() sync.WaitGroup {
	return WG
}

// func main() {
// 	t := time.Now()
// 	timeToDieInMilliseconds := 2000

// 	time.Sleep(2 * time.Second)
// 	x := time.Since(t) + time.Duration(timeToDieInMilliseconds)*time.Millisecond
// 	fmt.Printf("%v, %T", x, x)
// }
