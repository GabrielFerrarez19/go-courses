package main

import (
	"fmt"
	"sync"
)

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Printf("Worker %d done\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	const tasks = 5

// 	wg.Add(tasks)
// 	for i := 0; i <= tasks; i++ {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Main done\n")
// }

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://aws.com",
		"https://mercadolivre.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg)
	}
	wg.Wait()
	fmt.Println("main done\n")
}
