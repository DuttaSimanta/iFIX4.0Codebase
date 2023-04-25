package entities

import (
	"encoding/json"
	"io"
)

type ProcessMonitorEntity struct {
	Id       int64  `json:"id"`
	Serverno string `json:"serverno"`
	Process  string `json:"process"`
	Processlist  []int64 `json:"processlist"`
	Url      string `json:"url"`
	Status   string `json:"status"`
	Time     string `json:"time"`
}
type ProcessMonitorServerEntity struct {
	Serverno string `json:"serverno"`
	Processlist  []ProcessMonitorEntity `json:"processlist"`
}
type ProcessMonitorServerResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []ProcessMonitorServerEntity `json:"details"`
}
type ProcessMonitorResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []ProcessMonitorEntity `json:"details"`
}
func (w *ProcessMonitorEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
