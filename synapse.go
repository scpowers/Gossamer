package Gossamer

// define synapse struct
type Synapse struct {
	weight float64
	// pre and post neurons require anything
	// that implements the Neuron interface
	Pre  Neuron
	Post Neuron
}

// synapse constructor
func NewSynapse(pre, post Neuron) *Synapse {
	s := Synapse{
		weight: 0,
		Pre:    pre,
		Post:   post,
	}
	return &s
}
