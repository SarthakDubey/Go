package course3

import "time"

// writer Infinite for loop to update the value stored in the hashed location
func writer(data map[string]int) {
	for {
		data["_key"] = data["_key"] + 1
	}
}

// reader Infinite for loop to read the value stored in the hashed location
func reader(data map[string]int) {
	for {
		var _ int = data["_key"]
	}
}

func RaceCondition() {
	sharedHashMap := make(map[string]int)
	sharedHashMap["_key"] = 1
	go writer(sharedHashMap)
	go reader(sharedHashMap)
	// Start both the goroutines, and wait for 10 seconds to see if they come into a race condition
	// Trying to access the same block of memory for read/write operation at the same time.
	// Running this would lead to fatal error
	time.Sleep(10 * time.Second)
}
