package Gossamer

// define layer struct
type Layer struct {
	Spiking bool
	Neurons []Neuron
}

// layer constructor
func NewLayer(size int, isSpiking bool) *Layer {
	l := Layer{Spiking: isSpiking}
	var nlist []Neuron
	if isSpiking == true {
		for i := 0; i < size; i++ {
			nlist = append(nlist, &SpikingNeuron{})
		}
		l.Neurons = nlist
	} else {
		for i := 0; i < size; i++ {
			nlist = append(nlist, &RegularNeuron{})
		}
		l.Neurons = nlist
	}
	return &l
}
