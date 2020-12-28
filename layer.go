package Gossamer

// define layer struct
type Layer struct {
	IsSpikingLayer bool
	Neurons        []*Neuron
}

// layer constructor
func NewLayer(size int, isSpiking bool) *Layer {
	l := Layer{IsSpikingLayer: isSpiking}
	var nlist []*Neuron
	for i := 0; i < size; i++ {
		a := NewNeuron(isSpiking)
		nlist = append(nlist, a)
	}
	l.Neurons = nlist
	return &l
}
