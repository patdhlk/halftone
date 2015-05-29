package parallel

import ()

//we need to be able to send the data to processing to workers, via our dispatcher
//the WorkRequest struct holds our work request which can be send through a go
//channel
type WorkRequest struct {
	Name  string
	X     int
	Y     int
	Value int32
}
