package main

import (
	"github.com/nomenclate/forzago"
)

func main() {

	i := forzago.NewUdpListener(2323)
	q := forzago.NewChannelQueue()
	o := &forzago.StdoutOutputter{}
	p := forzago.NewForzaDataIn(i, q, o)

	p.Start()
}
