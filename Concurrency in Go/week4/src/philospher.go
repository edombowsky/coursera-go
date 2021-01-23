package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var separator = "------------------------------------------------------------------------------------------------------"

// ChopStick structure
// ------------------------------------------------------------------------
// Note:
// Due to the introduction of a regulator agent ("the Regulator")
// only the Chop Stick identifier is included as it will not be
// necessary to implement any mutex in this structure.
// ------------------------------------------------------------------------
type ChopStick struct {
	id int
}

// Regulator structure (The "regulator agent")
// -----------------------------------------------------
type Regulator struct {
	activePhilosophers int            // Number of current Philosophers still candidates to have a turn at eating
	philosophers       []*Philosopher // The list of Philosophers (to gain access to their info)
	signalBack         *chan string   // The shared signal channel where notifications will be received
}

func (regulator Regulator) giveTurn(wg *sync.WaitGroup, mutex *sync.Mutex) {
	i := 0
	defer wg.Done()

	for {

		turn1 := (i) % 5
		turn2 := (i + 2) % 5

		turn1Consumed := 0
		turn2Consumed := 0

		mutex.Lock()
		turn1Consumed = regulator.philosophers[turn1].eatingTurns
		turn2Consumed = regulator.philosophers[turn2].eatingTurns
		mutex.Unlock()

		if turn1Consumed < 3 {
			regulator.philosophers[turn1].signalIn <- "Proceed"
		} else {
			regulator.activePhilosophers--
		}

		if turn2Consumed < 3 {
			regulator.philosophers[turn2].signalIn <- "Proceed"
		} else {
			regulator.activePhilosophers--
		}

		if turn1Consumed < 3 {
			<-*regulator.philosophers[turn1].signalOut
		}

		if turn2Consumed < 3 {
			<-*regulator.philosophers[turn2].signalOut
		}

		if regulator.activePhilosophers <= 0 {
			break
		} else {
			i++
			fmt.Println(separator)
		}

	}
}

// Philosopher structure
// ------------------------------------------------------------------------
type Philosopher struct {
	id              int          // Its identifier
	leftCS, rightCS *ChopStick   // The left and right hands chop sticks assigned
	signalIn        chan string  // The chanel to receive activation signal/permission
	signalOut       *chan string // The chanel to send back ending to the Regulator
	eatingTime      int          // The time units that this philosopher consumes to eat
	eatingTurns     int          // The eating turns already consumed
}

func (philosopher *Philosopher) eat(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for {
		<-philosopher.signalIn
		fmt.Printf("starting to eat  <%d> \t\t( Using ChopSticks: #%d and #%d )\n",
			philosopher.id, philosopher.leftCS.id, philosopher.rightCS.id)
		mutex.Lock()
		philosopher.eatingTurns++
		mutex.Unlock()
		time.Sleep(time.Duration(philosopher.eatingTime) * time.Millisecond)
		fmt.Printf("finishing eating <%d> \t\t( Eating turns consumed: %d. Units of time required: %d ms )\n",
			philosopher.id, philosopher.eatingTurns, philosopher.eatingTime)
		*philosopher.signalOut <- "done"

		if philosopher.eatingTurns >= 3 {
			break
		}
	}
}

// ----------------------------------
// MAIN execution entry point
// ----------------------------------
func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	// ----------------
	// Chop Sticks
	// ----------------
	ChopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = &ChopStick{i}
	}

	// ------------------------------------------------------------------------
	// A channel to receive a signal back when a philosophersopher ends eating
	// ------------------------------------------------------------------------
	signalBack := make(chan string)

	// ------------------------------------------------------------------------
	// The philosophersopher list configuration
	// ------------------------------------------------------------------------
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			i + 1,               // philosopher identifier
			ChopSticks[i],       // left chopstick
			ChopSticks[(i+1)%5], // right chopstick
			make(chan string),   // signal input
			&signalBack,         // signal output
			rand.Intn(1000),     // eating time
			0,                   // eating turns already consumed
		}
	}

	// ------------------------------------------------------------------------
	// The Regulator configuration that intervene give turn permissions
	// ------------------------------------------------------------------------
	Regulator := Regulator{5, philosophers, &signalBack}

	// ------------------------------------------------------------------------
	// The MAIN routine
	// ------------------------------------------------------------------------
	fmt.Println(separator)
	fmt.Println("MISSION BEGINS")
	fmt.Println(separator)

	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup

	wg.Add(6)
	go Regulator.giveTurn(&wg, mutex)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg, mutex)
	}
	wg.Wait()

	fmt.Println(separator)
	fmt.Println("MISSION ACCOMPLISHED")
	fmt.Println(separator)
}
