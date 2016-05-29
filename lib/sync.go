package lib

import (
	"encoding/json"
	"fmt"
)

func (ns *NetworkSync) Recv() *PrimeCalc {
	//socket sync with master

	work := &PrimeCalc{}
	var recvContent string
	json := json.Unmarshal([]byte(recvContent), work)
	fmt.Println(json)

	return nil
}

func (ns *NetworkSync) Send(p *PrimeCalc) {
}
