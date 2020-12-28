package Gossamer

// define baseline neuron struct
type Neuron struct {
	IsSpiking bool
	SynIn     []Synapse
	SynOut    []Synapse
	Vmem      float64 // membrane potential
	Vt        float64 // threshold voltage
}

// regular neuron constructor
func NewNeuron(Spiking bool) *Neuron {
	n := new(Neuron)
	n.IsSpiking = Spiking
	return n
}

// get synapses of a neuron
func (n *Neuron) GetSynapses() ([]Synapse, []Synapse) {
	return n.SynIn, n.SynOut
}
