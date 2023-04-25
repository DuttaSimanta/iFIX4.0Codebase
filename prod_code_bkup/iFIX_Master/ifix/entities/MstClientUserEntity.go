package entities

import (
	"encoding/json"
	"io"
)

//MstClientUserEntity contains all required data fields
type MstClientUserEntity struct {
	ID                int64  `json:"id"`
	Userid            int64  `json:"userid"`
	ClientID          int64  `json:"clientid"`
	MstorgnhirarchyID int64  `json:"mstorgnhirarchyid"`
	Roleid            int64  `json:"roleid"`
	Relmanagerid      int64  `json:"relmanagerid"`
	Relmanager        string `json:"relmanager"`
	Loginname         string `json:"loginname"`
	Clientname        string `json:"clientname"`
	Orgname           string `json:"orgname"`
	Name              string `json:"name"`
	Useremail         string `json:"useremail"`
	Usermobileno      string `json:"usermobileno"`
	Password          string `json:"password"`
	Offset            int64  `json:"offset"`
	Limit             int64  `json:"limit"`
	Secondaryno       string `json:"secondaryno"`
	Division          string `json:"division"`
	Brand             string `json:"brand"`
	City              string `json:"city"`
	Designation       string `json:"designation"`
	Branch            string `json:"branch"`
	Vipuser           string `json:"vipuser"`
	Usertype          string `json:"usertype"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Color             string `json:"color"`
	Type              string `json:"type"`
	Groupids          []int  `json:"groupids"`
	Groupid           int64  `json:"groupid"`
	Groupname         string `json:"groupname"`
	Mfa               int64  `json:"mfa"`
	MfaName           string `json:"mfaname"`
	Createtype        int64 `json:"createtype"`
}

type MstGetUserByRecordidEntity struct {
	ClientID          int64 `json:"clientid"`
	MstorgnhirarchyID int64 `json:"mstorgnhirarchyid"`
	RecordID          int64 `json:"recordid"`
	ID                int64 `json:"id"`
}

//MstClientUserEntity contains search data fields
type MstUserSearchEntity struct {
	ID           int64  `json:"id"`
	Loginname    string `json:"loginname"`
	Lastname     string `json:"lastname"`
	Firstname    string `json:"firstname"`
	Name         string `json:"name"`
	Vipuser      string `json:"vipuser"`
	Useremail    string `json:"useremail"`
	Usermobileno string `json:"usermobileno"`
	Branch       string `json:"branch"`
}

type LoginNameSearchEntity struct {
	Loginname string `json:"loginname"`
}
type NameSearchEntity struct {
	Name string `json:"name"`
}
type BranchSearchEntity struct {
	Branch string `json:"branch"`
}
type LoginnameAndNameEntity struct {
	Loginname string `json:"loginname"`
	Name      string `json:"name"`
	Id        int64  `json:"id"`
	Groupid   int64  `json:"groupid"`
	Groupname string `json:"groupname"`
}

/*func (p *LoginNameSearchEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func (p *NameSearchEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func (p *BranchSearchEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}*/

//FromJSON is used for convert data into JSON format
func (p *MstClientUserEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//FromJSON is used for convert data into JSON format
func (p *MstGetUserByRecordidEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstClientUserEntities is a entity with two fields
type MstClientUserEntities struct {
	Total  int64                 `json:"total"`
	Values []MstClientUserEntity `json:"values"`
}

//MstClientUserEntityResponse is a response with all details
type MstClientUserEntityResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details MstClientUserEntities `json:"details"`
}
type MstClientUserEntityResp struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []MstClientUserEntity `json:"details"`
}

//MstSearchUserEntityResponse is a response with all details
type MstSearchUserEntityResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []MstUserSearchEntity `json:"details"`
}

//MstClientUserEntityResponseInt is a response with int
type MstClientUserEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

//MstSearchUserEntityResponse is a response with all details
type MstRecordwiseuserinfoEntityResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []MstClientUserEntity `json:"details"`
}
type LoginNameSearchEntityResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details []LoginNameSearchEntity `json:"details"`
}
type NameSearchEntityResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details []NameSearchEntity `json:"details"`
}
type BranchSearchEntityResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []BranchSearchEntity `json:"details"`
}
type LoginnameAndNameSearchEntityResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Details []LoginnameAndNameEntity `json:"details"`
}
