package main

import (
	"net"

	"github.com/nomenclate/forzago"
	"github.com/nomenclate/forzago/inputs"
	"github.com/nomenclate/forzago/outputs"
)

func main() {
	i := inputs.NewUdpListener(net.ParseIP("10.0.1.172"), 2323)
	q := forzago.NewChannelQueue()
	o := &outputs.StdoutOutputter{}
	p := forzago.NewPipeline(i, q, o)

	p.Start()
}
