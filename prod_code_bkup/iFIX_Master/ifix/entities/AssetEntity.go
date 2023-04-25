package entities

import (
	"encoding/json"
	"io"
)

type AssetEntity struct {
	Id                         int64  `json:"id"`
	Clientid                   int64  `json:"clientid"`
	Mstorgnhirarchyid          int64  `json:"mstorgnhirarchyid"`
	Mstdifferentiationtypeid   int64  `json:"mstdifferentiationtypeid"`
	Assetid                    string `json:"assetid"`
	Additionalattr             string `json:"additionalattr"`
	Activeflg                  int64  `json:"activeflg"`
	Audittransactionid         int64  `json:"audittransactionid"`
	Offset                     int64  `json:"offset"`
	Limit                      int64  `json:"limit"`
	Clientname                 string `json:"clientname"`
	Mstorgnhirarchyname        string `json:"mstorgnhirarchyname"`
	Mstdifferentiationtypename string `json:"mstdifferentiationtypename"`
}

type AssetIDEntity struct {
	ID         int64                `json:"id"`
	Assetid    string               `json:"assetid"`
	History    string               `json:"assethistory"`
	Attributes []AssetEntityDiffVal `json:"attributes"`
}

type AssetSearchEntity struct {
	Clientid                 int64  `json:"clientid"`
	Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
	Mstdifferentiationtypeid int64  `json:"mstdifferentiationtypeid"`
	Mstdifferentiationid     int64  `json:"mstdifferentiationid"`
	Value                    string `json:"value"`
}

type AssetSearchResEntity struct {
	AssetAttributes []Assettype     `json:"assetattributes"`
	AssetValues     []AssetIDEntity `json:"assetvales"`
}

type AssetEntityDiffVal struct {
	TypeId   int64  `json:"typeid"`
	TypeName string `json:"typename"`
	AttrId   int64  `json:"attrid"`
	AttrName string `json:"attrname"`
	Value    string `json:"value"`
}

type AssetEntityDiffValUpdate struct {
	Clientid                 int64                `json:"clientid"`
	Mstorgnhirarchyid        int64                `json:"mstorgnhirarchyid"`
	Mstdifferentiationtypeid int64                `json:"mstdifferentiationtypeid"`
	Assetid                  int64                `json:"trnassetid"`
	Attributes               []AssetEntityDiffVal `json:"attributes"`
}

type AssetEntityDiffVals struct {
	Values []AssetEntityDiffVal `json:"values"`
}

type AssetEntityByType struct {
	Id      int64  `json:"id"`
	Assetid string `json:"assetid"`
}

type AssetMapWithRecordType struct {
	Clientid            int64  `json:"clientid"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	ID                  int64  `json:"id"`
	DiffTypeID          int64  `json:"difftypeid"`
	DiffTypeName        string `json:"difftypename"`
	DiffTypeParent      int64  `json:"difftypeparent"`
	DiffID              int64  `json:"diffid"`
	DiffName            string `json:"diffname"`
}

type AssetMapWithRecordTypes struct {
	Total  int64                    `json:"total"`
	Values []AssetMapWithRecordType `json:"values"`
}

type AssetEntitiesByType struct {
	Values []AssetEntityByType `json:"values"`
}

type AssetEntities struct {
	Total  int64         `json:"total"`
	Values []AssetEntity `json:"values"`
}

type Assettype struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Parent int64  `json:"parent"`
	Seqno  int64  `json:"seqno"`
}

type AssetSearchResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details AssetSearchResEntity `json:"details"`
}
type AssettypeResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Details []Assettype `json:"details"`
}

type AssetResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Details AssetEntities `json:"details"`
}

type AssetResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type AssetByTypeResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Details AssetEntitiesByType `json:"details"`
}

type AssetEntityDiffValResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Details AssetEntityDiffVals `json:"details"`
}

type AssetMapWithRecordTypeResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details AssetMapWithRecordTypes `json:"details"`
}

func (w *AssetEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *AssetEntityDiffValUpdate) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *AssetSearchEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
