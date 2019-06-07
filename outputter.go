package forzago

import (
	"fmt"
)

type StdoutOutputter struct {
}

func (o *StdoutOutputter) StartOutputting(q Queue) {
	fmt.Println("Starting output")
	for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
		var data = DataOutPacket{}
		data.UnmarshalBinary(message)
		fmt.Println(data)
	}
}
