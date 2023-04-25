package entities

import (
	"encoding/json"
	"io"
)

type ReportDownloadListEntity struct {
	Id               int64  `json:"id"`
	Originalfilename string `json:originalfilename`
	Uploadedfilename string `json:uploadedfilename`
	Refuserid        int64  `json:"refuserid"`
	Offset           int64  `json:"offset"`
	Limit            int64  `json:"limit"`
}

type ReportDownloadListEntities struct {
	Total  int64                      `json:"total"`
	Values []ReportDownloadListEntity `json:"values"`
}

type ReportDownloadListResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details ReportDownloadListEntities `json:"details"`
}

type ReportDownloadListResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ReportDownloadListEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
