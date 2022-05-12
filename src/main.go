package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var WG = sync.WaitGroup{}

func main() {
	fw := 0

	if len(os.Args) != 5 && len(os.Args) != 6 {
		fmt.Printf("Usage ./pythaGOras [Number philosophers] [time to die] [time to eat] [time to sleep] (Full when)")
		return
	}
	ammount, err := strconv.Atoi((os.Args[1]))
	if err != nil {
		fmt.Printf("Usage ./pythaGOras [Number philosophers] [time to die] [time to eat] [time to sleep] (Full when)")
		return
	}
	ttd, err := strconv.Atoi((os.Args[2]))
	if err != nil {
		fmt.Printf("Usage ./pythaGOras [Number philosophers] [time to die] [time to eat] [time to sleep] (Full when)")
		return
	}
	tte, err := strconv.Atoi((os.Args[3]))
	if err != nil {
		fmt.Printf("Usage ./pythaGOras [Number philosophers] [time to die] [time to eat] [time to sleep] (Full when)")
		return
	}
	tts, err := strconv.Atoi((os.Args[4]))
	if err != nil {
		fmt.Printf("Usage ./pythaGOras [Number philosophers] [time to die] [time to eat] [time to sleep] (Full when)")
		return
	}
	if len(os.Args) == 6 {
		fw, err = strconv.Atoi((os.Args[5]))
		if err != nil {
			fw = 0
		}
	}
	t := Times{
		timeToDie:   time.Duration(time.Duration(ttd) * time.Millisecond),
		timeToEat:   time.Duration(time.Duration(tte) * time.Millisecond),
		timeToSleep: time.Duration(time.Duration(tts) * time.Millisecond),
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
		if fw != 0 {
			ph.fullWhen = uint(fw)
		}
		ph.wait = WG
		ph.lastMeal = time.Now()
		ph.durations = &t
		ph.name = strconv.Itoa(i)
		ph.status = thinking
		ph.fck = fck
		ph.fckMut = fckMut
		go dinner(ph)
	}
	WG.Wait()
}

func dinner(ph Philo) {
	for {
		if ph.status == dead {
			break
		}
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
			if ph.status == dead {
				break
			}
			ph.TryEat()
		}
		if ph.status == eating {
			if ph.status == dead {
				break
			}
			ph.Sleep(ph.durations.timeToSleep)
		}
		if ph.status != thinking {
			if ph.status == dead {
				break
			}
			ph.Think()
		}
	}
	WG.Done()
}
