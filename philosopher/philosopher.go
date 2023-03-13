package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{ sync.Mutex } // sharing chopsticks

type Philosopher struct {
	// one chopstick between each adjacent pair of philosophers
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	// Each philosopher is numbered, 1 through N
	id int
	//Each philosopher should eat only M times
	eat_count int
}

func (p *Philosopher) Eat(eat_signal chan bool, done_signal *sync.WaitGroup) {
	defer done_signal.Done()
	// The philosophers pick up the chopsticks in any order
	for i := 0; i < p.eat_count; i++ {
		// get permission from a host
		<-eat_signal
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("Starting to eat %d #%d\n", p.id, i)
		fmt.Printf("Finished eating %d #%d\n", p.id, i)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
		// release permission to eat
		eat_signal <- true
	}
}

type Dinner struct {
	// chopsticks
	chopsticks []*Chopstick
	// dinner guests managed by host
	philosophers []*Philosopher
	// Number of philospohers host allows to eat concurrently
	max_concurrent_eats int
	// Number of philospohers host EATING concurrently
	concurrent_eats_count int
}

func main() {

	var num_of_philosophers int = 5      // number of philosphers
	var max_concurrent_eats int = 2      // max number of philosphers eating at the same time
	var number_of_eats_per_philo int = 3 // number of times to eat
	chopsticks := make([]*Chopstick, num_of_philosophers)
	for i := 0; i < num_of_philosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, num_of_philosophers)
	for i := 0; i < num_of_philosophers; i++ {
		philosophers[i] = &Philosopher{id: i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%num_of_philosophers],
			eat_count:      number_of_eats_per_philo, // number of times to eat
		}

	}
	// use channel to communicate between host and philos
	var wg sync.WaitGroup
	eat_signal := make(chan bool, max_concurrent_eats)
	wg.Add(len(philosophers))
	for _, philo := range philosophers {
		go (*philo).Eat(eat_signal, &wg)
	}

	// kick off first two eaters
	for i := 0; i < max_concurrent_eats; i++ {
		eat_signal <- true
	}

	wg.Wait()
}
