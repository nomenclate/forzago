package main

import (
	"github.com/nomenclate/forzadata"
)

func main() {

	i := forzadata.NewUdpListener(2323)
	q := forzadata.NewChannelQueue()
	o := &forzadata.StdoutOutputter{}
	p := forzadata.NewForzaDataIn(i, q, o)

	p.Start()
}
