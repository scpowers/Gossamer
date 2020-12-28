package main

import (
	"fmt"
	gs "github.com/scpowers/Gossamer"
)

func main() {
	fmt.Println("Regular neuron test")
	n1 := gs.NewNeuron()
	a, b := n1.GetSynapses()
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("\n\nRegular neuron test")
	n2 := gs.NewSpikingNeuron()
	a, b = n2.GetSynapses()
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("\n\nnow testing synapses")
	s := gs.NewSynapse(n1, n2)
	fmt.Println(s.Pre.GetSynapses())
	fmt.Println(s.Post.GetSynapses())

}
