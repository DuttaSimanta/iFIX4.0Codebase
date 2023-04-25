package entities

import (
	"encoding/json"
	"io"
)

type MstDifferentiationmapEntity struct {
	ID                    int64   `json:"id"`
	FromclientID          int64   `json:"fromclientid"`
	FromorgnID            int64   `json:"fromorgnid"`
	ToclinentID           int64   `json:"toclientid"`
	ToorgnID              []int64 `json:"toorgnid"`
	DifferentiationtypeID int64   `json:"differentiationtypeid"`
	DiffererntiationID    []int64 `json:"differerntiationid"`
	MapID                 int64   `json:"mapid"`
	Offset                int64   `json:"offset"`
	Limit                 int64   `json:"limit"`
}

type MstDifferentiationmapResponseFields struct {
	ID               int64  `json:"id"`
	MapID            int64  `json:"mapid"`
	Fromclientname   string `json:"fromclientname"`
	Fromorgnname     string `json:"fromorgnname"`
	Fromdifftypename string `json:"fromdifftypename"`
	Fromdiffname     string `json:"fromdiffname"`
	Toclientname     string `json:"toclientname"`
	Toorgnname       string `json:"toorgnname"`
	Todifftypename   string `json:"todifftypename"`
	Todiffname       string `json:"todiffname"`
}

type MstDifferentiationmaEntities struct {
	Total  int64                                 `json:"total"`
	Values []MstDifferentiationmapResponseFields `json:"values"`
}

type MstResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details MstDifferentiationmaEntities `json:"details"`
}

func (p *MstDifferentiationmapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type MstDifferentiationmapResponseEntity struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	//Details MstClientUserEntities `json:"details"`
}
