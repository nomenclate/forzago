package forzago

import (
	"fmt"

	"github.com/nomenclate/forzago"
)

type StdoutOutputter struct {
}

func (o *StdoutOutputter) StartOutputting(q forzago.Queue) {
	fmt.Println("Starting output")
	for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
		var data = forzago.DataOutPacket{}
		data.UnmarshalBinary(message)
		fmt.Println(data)
	}
}
