package parallel

import (
	"log"
)

//declareing and initializing our WorkerQueue which is the buffered channel that
//holds the work channels from each worker
//remembering that the worker is responsible for adding itself into the workers
//queue
var WorkerQueue chan chan WorkRequest

//
func StartDispatcher() {
	nworkers := 10

	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan WorkRequest, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		log.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				log.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					log.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
