package entities

import (
	"encoding/json"
	"io"
)

type Ldapattrentity struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type LdapAttrEntityResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details []Ldapattrentity `json:"details"`
}
type LoginEntityReq struct {
	ID                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	OrgnTypeId        int64  `json:"orgnTypeId"`
	Loginname         string `json:"loginname"`
	Password          string `json:"password"`
	UserEmail         string `json:"useremail"`
	Oldpassword       string `json:"Oldpassword"`
	Code              string `json:"code"`
	OrgMFA            int64  `json:"orgmfa"`
	UserMFA           int64  `json:"usermfa"`
	Totp              string `json:"totp"`
	Secretkey         string `json:"secretkey"`
	Type              string `json:"type"`
	OTP              string `json:"otp"`
}

type LoginEntityResp struct {
	ID          int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Userid            int64  `json:"userid"`
	Token             string `json:"token"`
	Org               string `json:"org"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	OrgnTypeId        int64  `json:"orgnTypeId"`
	Logintypeid       int64  `json:"logintypeid"`
	Loginname         string `json:"loginname"`
	Roleid            int64  `json:"Roleid"`
	IsAdmin           int64  `json:"isAdmin"`
	Islocallogin      int64  `json:"islocallogin"`
	Dashboardurl      string `json:"dashboardurl"`
	Externalurl       string `json:"externalurl"`
	Url               string `json:"url"`
	Rolename          string `json:"rolename"`
	OriginalFileName  string `json:"originalfilename"`
	UploadedFileName  string `json:"uploadedfilename"`
	Password          string `json:"password"`
	OrgMFA            int64  `json:"orgmfa"`
	UserMFA           int64  `json:"usermfa"`
	OTP               string `json:"otp"`
	Time               string `json:"time"`
}
type Usergroupentity struct {
	ID           int64  `json:"id"`
	Groupname    string `json:"groupname"`
	Hascatalog   string `json:"hascatalog"`
	Levelid      int64  `json:"levelid"`
	IsManagement string `json:"ismanagement"`
}

type Userurlentity struct {
	Urlkey string `json:"urlkey"`
	Url    string `json:"url"`
}
type UserEntity struct {
	Clientid          int64             `json:"clientid"`
	Mstorgnhirarchyid int64             `json:"mstorgnhirarchyid"`
	Deafultgroup      int64             `json:"deafultgroup"`
	Loginname         string            `json:"loginname"`
	Lastname          string            `json:"lastname"`
	Firstname         string            `json:"firstname"`
	Username          string            `json:"username"`
	Branch            string            `json:"branch"`
	Vipuser           string            `json:"vipuser"`
	Email             string            `json:"email"`
	Mobile            string            `json:"mobile"`
	Mstorgnname       string            `json:"mstorgnname"`
	Clientname        string            `json:"clientname"`
	Createtype        int64             `json:"createtype"`
	Orgntypeid        int               `json:"orgntypeid"`
	Logintypeid       int               `json:"logintypeid"`
	Userid            int64             `json:"userid"`
	Roleid            int64             `json:"Roleid"`
	IsAdmin           int64             `json:"isAdmin"`
	Rolename          string            `json:"rolename"`
	Color             string            `json:"color"`
	Uploadedbgimage   string            `json:"uploadedbgimage"`
	Uploadedlogoimage string            `json:"uploadedlogoimage"`
	Add               bool              `json:"addFlag"`
	Delete            bool              `json:"deleteFlag"`
	Edit              bool              `json:"editFlag"`
	View              bool              `json:"viewFlag"`
	Group             []Usergroupentity `json:"group"`
	Propertyids       []int64           `json:"propertyids"`
	Urls              []Userurlentity   `json:"urls"`
}

type LoginResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Details []LoginEntityResp `json:"details"`
}
type UserResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details []UserEntity `json:"details"`
}

func (w *LoginEntityReq) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
func (w *LoginEntityResp) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
func (w *UserEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
