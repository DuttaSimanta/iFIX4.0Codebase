package entities

import (
	"encoding/json"
	"io"
)

type ReportDownloadEntity struct {
	Id               int64  `json:"id"`
	Originalfilename string `json:originalfilename`
	Uploadedfilename string `json:uploadedfilename`
	Refuserid        int64  `json:"refuserid"`
	Offset           int64  `json:"offset"`
	Limit            int64  `json:"limit"`
}

type ReportDownloadEntities struct {
	Total  int64                  `json:"total"`
	Values []ReportDownloadEntity `json:"values"`
}

type ReportDownloadResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details ReportDownloadEntities `json:"details"`
}

type ReportDownloadResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ReportDownloadEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
