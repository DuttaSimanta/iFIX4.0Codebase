package entities

import (
	"encoding/json"
	"io"
)

type RecordDetailsEntity struct {
	ClientID          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	// Recorddifftypeid    int64  `json:"recorddifftypeid"`
	// Recorddiffid        int64  `json:"recorddiffid"`
	RecordID int64 `json:"recordid"`
	// Userid              int64  `"json:userid"`
	// Usergroupid         int64  `json:"usergroupid"`
	// Originaluserid      int64  `json:"originaluserid"`
	// Originalusergroupid int64  `json:"originalusergroupid"`
	// Recordname          string `json:"recordname"`
	// Recordesc           string `json:"recordesc"`
	RecordStageid int64 `json:"recordstageid"`
	Requestid     int64 `json:"requestid "`
}
type Workflowentity struct {
	Id                int64 `json:"id"`
	Clientid          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	// Processid                int64                    `json:"processid"`
	// Recordtitle              string                   `json:"recordtitle"`
	Transactionid int64 `json:"transactionid"`
	// Transactionids           []int64                    `json:"transactionids"`
	Currentstateid int64 `json:"currentstateid"`
	// Previousstateid          int64                    `json:"previousstateid"`
	// Previousstateids         []int64                  `json:"previousstateids"`
	// Transitionids            []int64                  `json:"transitionids"`
	// Activities               []int64                  `json:"activities"`
	// Activity                 int64                    `json:"activity"`
	// Recorddifftypeid         int64                    `json:"recorddifftypeid"`
	// Recorddiffid             int64                    `json:"recorddiffid"`
	// Mstgroupid               int64                    `json:"mstgroupid"`
	// Mstuserid                int64                    `json:"mstuserid"`
	// Createduserid            int64                    `json:"createduserid"`
	Createdgroupid int64 `json:"createdgroupid"`
	Userid         int64 `json:"Userid"`
	// Transitionid   int64 `json:"transitionid"`
	// Deleteflg                int64                    `json:"deleteflg"`
	// Recordid                 int64                    `json:"recordid"`
	// Recordstageid            int64                    `json:"recordstageid"`
	// Mstrequestid             int64                    `json:"mstrequestid"`
	// Details                  string                   `json:"details"`
	// Detailsjson              string                   `json:"detailsjson"`
	// Processname              string                   `json:"processname"`
	// Tablename                string                   `json:"tablename"`
	// Loginname                string                   `json:"loginname"`
	// Username                 string                   `json:"username"`
	// Groupname                string                   `json:"groupname"`
	// Manualstateselection     int                      `json:"manualstateselection"`
	// Mstdatadictionaryfieldid int64                    `json:"mstdatadictionaryfieldid"`
	// Dateofrecordchange       int64                    `json:"dateofrecordchange"`
	// Users                    []WorkflowResponseEntity `json:"users"`
	// Starttime                int64                    `json:"starttime"`
	// Endtime                  int64                    `json:"endtime"`
	// Activeflg                int64                    `json:"activeflg"`
	// Audittransactionid       int64                    `json:"audittransactionid"`
	// Iscomplete               int64                    `json:"iscomplete"`
	// Parentid                 int64                    `json:"parentid"`
	Changestatus int64 `json:"changestatus"`
	// Childids                 []int64                  `json:"childids"`
	// Samegroup                bool                     `json:"samegroup"`
	// Recorddiffids            []RecorddiffEntity       `json:"recorddiffids"`
	// Isupdate                 bool                     `json:"isupdate"`
	Issrrequestor int64 `json:"issrrequestor"`
	// Creatorgroupid           int64                     `json:"creatorgroupid"`
	// IsAttaching           int64                       `json:"isattaching"`
	// Nextstateseq           int64                       `json:"nextstateseq"`
}
type StatedetailEntity struct {
	Id             int64
	Stateid        int64
	Statusid       int64
	StatusSeq      int64
	Statusname     string
	Userid         int64
	Createdgroupid int64
	Index          int
	Mstuserid      int64
}
type ErrorEntity struct {
	ClientID          int64
	Mstorgnhirarchyid int64
	Stateid           int64
	RecordID          int64
	Requestjson       string
	Responsejson      string
	Comment           string
	Isdefect          string
}
type RecordstatusResponeData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (p *RecordstatusResponeData) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type ParentchildEntity struct {
	Parentid       int64   `json:"parentid"`
	Childids       []int64 `json:"childids"`
	Userid         int64   `json:"userid"`
	Createdgroupid int64   `json:"createdgroupid"`
	Isupdate       bool    `json:"isupdate"`
	Transactionid  int64   `json:"transactionid"`
	Usergroupid    int64   `json:"usergroupid"`
	IsAttaching    int64   `json:"isattaching"`
}
type WorkflowResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details string `json:"details"`
}
