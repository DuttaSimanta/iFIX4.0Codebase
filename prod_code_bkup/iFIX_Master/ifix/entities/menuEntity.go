package entities

import (
	"encoding/json"
	"io"
)

type MenuByUserRequest struct {
	ClientID          int64 `json:"clientid"`
	MstorgnhirarchyID int64 `json:"mstorgnhirarchyid"`
	UserID            int64 `json:"userid"`
	Menu            string `json:"menu"`
}

type MenuHierarchyEntity struct {
	ID     int64                 `json:"id"`
	Label  string                `json:"label"`
	Parent int64                 `json:"parent"`
	Path   string                `json:"path"`
	Items  []MenuHierarchyEntity `json:"items"`
}

//FromJSONMenuByUserRequest is used for convert data into JSON format
func (p *MenuByUserRequest) FromJSONMenuByUserRequest(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type MenuHierarchyResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []MenuHierarchyEntity `json:"details"`
}

func FindMenuChild(memu MenuHierarchyEntity, menulist []MenuHierarchyEntity) (menulistresponse []MenuHierarchyEntity) {
	for _, menus := range menulist {
		if memu.ID == menus.Parent {
			menulistresponse = append(menulistresponse, menus)
		}
	}
	return
}

func GetRootMenu(menulist []MenuHierarchyEntity) (menulistresponse []MenuHierarchyEntity) {
	for _, menu := range menulist {
		if menu.Parent == 0 {
			menulistresponse = append(menulistresponse, menu)
		}
	}
	for i, menu := range menulistresponse {
		menulistresponse[i].Items = MakeMenuTree(menu, menulist)
	}
	return menulistresponse
}

func MakeMenuTree(memu MenuHierarchyEntity, menulist []MenuHierarchyEntity) (menulistresponse []MenuHierarchyEntity) {
	menulistresponse = FindMenuChild(memu, menulist)
	for i, child := range menulistresponse {

		menulistresponse[i].Items = MakeMenuTree(child, menulist)

	}
	return menulistresponse
}

type MenuEntity struct {
	ID                int64  `json:"id"`
	ClientID          int64  `json:"clientid"`
	MstorgnhirarchyID int64  `json:"mstorgnhirarchyid"`
	Urlid             int64  `json:"urlid"`
	Parentmenuid      int64  `json:"parentmenuid"`
	Moduleid          int64  `json:"moduleid"`
	Sequence_no       int64  `json:"sequence_no"`
	Menudesc          string `json:"menudesc"`
	Leafnode          int    `json:"leafnode"`
}
type MenuSingleEntity struct {
	ID       int64  `json:"id"`
	Menudesc string `json:"menudesc"`
}
type MenuEntityResp struct {
	ID                int64  `json:"id"`
	ClientID          int64  `json:"clientid"`
	MstorgnhirarchyID int64  `json:"mstorgnhirarchyid"`
	Urlid             int64  `json:"urlid"`
	Parentmenuid      int64  `json:"parentmenuid"`
	Moduleid          int64  `json:"moduleid"`
	Sequence_no       int64  `json:"sequence_no"`
	Menudesc          string `json:"menudesc"`
	Activeflg         int    `json:"leafnode"`
	Parentmenu        string `json:"parentmenu"`
	Clientname        string `json:"clientname"`
	Orgnname          string `json:"Orgnname"`
	Modulename        string `json:"modulename"`
	Url               string `json:"url"`
}

//MstModuleClientEntities is a entity with two fields
type MenuEntities struct {
	Total  int64            `json:"total"`
	Values []MenuEntityResp `json:"values"`
}

//FromJSON is used for convert data into JSON format
func (p *MenuEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type MenuEntitySingleResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details []MenuSingleEntity `json:"details"`
}
type MenuEntityResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details MenuEntities `json:"details"`
}
