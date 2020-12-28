package main

import (
	"fmt"
	gs "github.com/scpowers/Gossamer"
)

func main() {
	fmt.Println("Regular neuron test")
	n1 := gs.NewNeuron(false)
	a, b := n1.GetSynapses()
	fmt.Println(n1.IsSpiking)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("\n\nSpiking neuron test")
	n2 := gs.NewNeuron(true)
	a, b = n2.GetSynapses()
	fmt.Println(n2.IsSpiking)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("\n\nNow testing synapses")
	s := gs.NewSynapse(n1, n2)
	fmt.Println(s.Pre.GetSynapses())
	fmt.Println(s.Post.GetSynapses())

	fmt.Println("\n\nNow testing layers")
	l := gs.NewLayer(5, true)
	fmt.Println(l.IsSpikingLayer)
	for i := 0; i < 5; i++ {
		fmt.Printf("Checking neuron %d\n", i)
		fmt.Println(l.Neurons[i].GetSynapses())
		fmt.Println(l.Neurons[i].Vmem)
	}

}
