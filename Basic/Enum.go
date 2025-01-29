package main

import "fmt"

type ServerState int 

const(
	Stopped ServerState = iota
	Running
	Terminated
)

var state = map[ServerState]string{
	Stopped: "Stopped",
	Running: "Running",
	Terminated: "Terminated",
}

func (s ServerState) String() string{
	return state[s]
}

func transitionState(s ServerState) ServerState{
	switch s{
	case Stopped : return Running
	case Running : return Terminated
	case Terminated : return Stopped
	default:
		panic(fmt.Errorf("unknown state: %s", s)) 
	}
}

func main(){
	ns := transitionState(Stopped)
	fmt.Println(ns)
	fmt.Println(transitionState(ns))
	fmt.Println(transitionState(transitionState(ns)))
}