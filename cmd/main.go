package main

import (
	"github.com/nomenclate/forzago"
	"github.com/nomenclate/forzago/inputs"
	"github.com/nomenclate/forzago/outputs"
)

func main() {

	i := inputs.NewUdpListener(2323)
	q := forzago.NewChannelQueue()
	o := &outputs.StdoutOutputter{}
	p := forzago.NewDataIn(i, q, o)

	p.Start()
}
