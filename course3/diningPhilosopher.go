package course3

import (
	"fmt"
	"sync"
	"time"
)

type chopStick struct {
	sync.Mutex
	id int
}

type philosopherInterface interface {
	ID() int
	Lock()
	Eat()
	Count() int
	Unlock()
	Status(status string)
}

type philosopherImpl struct {
	num, count  int
	left, right *chopStick
}

type permissionCh struct {
	id   int
	done chan bool
}

func (p *philosopherImpl) ID() int {
	return p.num
}

func (p *philosopherImpl) Lock() {
	// Attempt to get a lock on both chopsticks, else unlock left and try again!
	for {
		p.left.Lock()
		if p.right.TryLock() {
			break
		}
		p.left.Unlock()
	}
}

func (p *philosopherImpl) Eat() {
	p.Status("started")
	p.count++
	p.Status("finished")
}

func (p *philosopherImpl) Count() int {
	return p.count
}

func (p *philosopherImpl) Unlock() {
	// Unlock both chopsticks
	p.left.Unlock()
	p.right.Unlock()
}

func (p *philosopherImpl) Status(status string) {
	fmt.Printf("Philosopher %d %s eating using chopsticks %v, and %v | Eat Count: %d\n", p.ID(), status, p.left.id, p.right.id, p.count)
}

func PhilosopherEat(p philosopherInterface, n int, wg *sync.WaitGroup, permission chan permissionCh) {
	grant := make(chan bool)
	for p.Count() < n {
		// Get a lock on both chopsticks
		p.Lock()
		// Ask permission from host
		permission <- permissionCh{p.ID(), grant}
		// Wait for host to allow to eat
		<-grant
		p.Eat()
		p.Unlock()
	}
	close(grant)
	wg.Done()
}

func Host(wg *sync.WaitGroup, permission chan permissionCh, eatCount, tableSize int) {
	counts := eatCount * tableSize
	for counts > 1 {
		if len(permission) == 2 {
			// Get two concurrent philosopher requests
			philosopher1 := <-permission
			philosopher2 := <-permission
			fmt.Printf("Philosopher %v and %v eating together\n", philosopher1.id, philosopher2.id)
			// Give permission to them
			philosopher1.done <- true
			philosopher2.done <- true
			counts -= 2
			time.Sleep(10 * time.Millisecond)
		}
	}
	if counts == 1 {
		fmt.Println("Last philosopher remaining. Odd counts of eat")
		lastPhilosopher := <-permission
		fmt.Printf("Last philosopher %v eating\n", lastPhilosopher.id)
		lastPhilosopher.done <- true
		counts -= 1
	}
	wg.Done()
}

func DiningPhilosopher() {
	var wg sync.WaitGroup
	eatCount := 3
	tableSize := 5
	numChopsticks := make([]*chopStick, tableSize)
	for i := 0; i < tableSize; i++ {
		numChopsticks[i] = &chopStick{id: i + 1}
	}

	numPhilosophers := make([]*philosopherImpl, tableSize)
	for i := 0; i < tableSize; i++ {
		numPhilosophers[i] = &philosopherImpl{i + 1, 0, numChopsticks[i], numChopsticks[(i+1)%tableSize]}
	}

	permission := make(chan permissionCh, 2)
	wg.Add(1)
	go Host(&wg, permission, eatCount, tableSize)

	for _, philosopher := range numPhilosophers {
		wg.Add(1)
		go PhilosopherEat(philosopher, eatCount, &wg, permission)
	}

	wg.Wait()

	fmt.Println("\nStatus Report:")
	for i := 0; i < tableSize; i++ {
		fmt.Printf("Philosopher %v ate %v times\n", i+1, numPhilosophers[i].count)
	}
}
