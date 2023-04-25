package entities

import (
	"encoding/json"
	"io"
)


type Priorities struct {
	ClientId         int    `json:"clientid"`
	TicketTypeId     int    `json:"tickettypeid"`
	BusiPriorityDesc string `json:"busiprioritydesc"`
	BusiPriorityName string `json:"busipriorityname"`
}

func (p *Priorities) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//curl -v localhost:9090 -d '{"clientid":1,"tickettypeid":2,"busiprioritydesc":"AAA","busipriorityname":"BBB"}'
