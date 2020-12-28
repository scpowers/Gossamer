package Gossamer

// define Neuron interface to get polymorphism-like behavior
type Neuron interface {
	GetSynapses() ([]int, []int)
}

// define baseline neuron struct
type RegularNeuron struct {
	// universal properties
	SynIn  []int
	SynOut []int
}

// define spiking neuron struct
type SpikingNeuron struct {
	RegularNeuron
	Vmem float64 // membrane potential
	Vt   float64 // threshold voltage
}

// regular neuron constructor
func NewNeuron() *RegularNeuron {
	return &RegularNeuron{}
}

// spiking neuron constructor
func NewSpikingNeuron() *SpikingNeuron {
	n := new(SpikingNeuron)
	n.RegularNeuron = RegularNeuron{}
	n.Vmem = 0
	n.Vt = 0
	return n
}

// get synapses of a neuron
func (n *RegularNeuron) GetSynapses() ([]int, []int) {
	return n.SynIn, n.SynOut
}
