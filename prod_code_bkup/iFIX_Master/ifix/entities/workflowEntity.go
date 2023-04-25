package entities

import (
	"encoding/json"
	"io"
)

type Workflowentity struct {
	Id                       int64                    `json:"id"`
	Clientid                 int64                    `json:"clientid"`
	Mstorgnhirarchyid        int64                    `json:"mstorgnhirarchyid"`
	Processid                int64                    `json:"processid"`
	Recordtitle              string                   `json:"recordtitle"`
	Transactionid            int64                    `json:"transactionid"`
	Transactionids           []int64                    `json:"transactionids"`
	Currentstateid           int64                    `json:"currentstateid"`
	Previousstateid          int64                    `json:"previousstateid"`
	Previousstateids         []int64                  `json:"previousstateids"`
	Transitionids            []int64                  `json:"transitionids"`
	Activities               []int64                  `json:"activities"`
	Activity                 int64                    `json:"activity"`
	Recorddifftypeid         int64                    `json:"recorddifftypeid"`
	Recorddiffid             int64                    `json:"recorddiffid"`
	Mstgroupid               int64                    `json:"mstgroupid"`
	Mstuserid                int64                    `json:"mstuserid"`
	Createduserid            int64                    `json:"createduserid"`
	Createdgroupid           int64                    `json:"createdgroupid"`
	Userid                   int64                    `json:"Userid"`
	Transitionid             int64                    `json:"transitionid"`
	Deleteflg                int64                    `json:"deleteflg"`
	Recordid                 int64                    `json:"recordid"`
	Recordstageid            int64                    `json:"recordstageid"`
	Mstrequestid             int64                    `json:"mstrequestid"`
	Details                  string                   `json:"details"`
	Detailsjson              string                   `json:"detailsjson"`
	Processname              string                   `json:"processname"`
	Tablename                string                   `json:"tablename"`
	Loginname                string                   `json:"loginname"`
	Username                 string                   `json:"username"`
	Groupname                string                   `json:"groupname"`
	Manualstateselection     int                      `json:"manualstateselection"`
	Mstdatadictionaryfieldid int64                    `json:"mstdatadictionaryfieldid"`
	Dateofrecordchange       int64                    `json:"dateofrecordchange"`
	Users                    []WorkflowResponseEntity `json:"users"`
	Starttime                int64                    `json:"starttime"`
	Endtime                  int64                    `json:"endtime"`
	Activeflg                int64                    `json:"activeflg"`
	Audittransactionid       int64                    `json:"audittransactionid"`
	Iscomplete               int64                    `json:"iscomplete"`
	Parentid                 int64                    `json:"parentid"`
	Changestatus             int64                    `json:"changestatus"`
	Childids                 []int64                  `json:"childids"`
	Samegroup                bool                     `json:"samegroup"`
	Recorddiffids            []RecorddiffEntity       `json:"recorddiffids"`
	Isupdate                 bool                     `json:"isupdate"`
	Issrrequestor            int64                     `json:"issrrequestor"`
	Creatorgroupid           int64                     `json:"creatorgroupid"`
	IsAttaching           	 int64                       `json:"isattaching"`
	Nextstateseq             int64                       `json:"nextstateseq"`
	Levelid           		 int64                       `json:"levelid"`
}
type WorkflowStateEntity struct {
	Templatetransitionid int64 `json:"templatetransitionid"`
	Currentstateid       int64 `json:"currentstateid"`
	Previousstateid      int64 `json:"previousstateid"`
	Currentseq           int64 `json:"currentseq"`
	Previousseq          int64 `json:"previousseq"`
}

type WorkflowResponseEntity struct {
	Id          int64  `json:"id"`
	Details     string `json:"details"`
	Detailsjson string `json:"detailsjson"`
	Iscomplete  int64  `json:"iscomplete"`
	Mstgroupid  int64  `json:"mstgroupid"`
	Mstuserid   int64  `json:"mstuserid"`
	Activityid  int64  `json:"activityid"`
	Loginname   string `json:"loginname"`
	Groupname   string `json:"groupname"`
	Processname string `json:"processname"`
}
type WorkflowTransitionEntity struct {
	Currentstateid   int64  `json:"currentstateid"`
	Currentstate     string `json:"currentstate"`
	Recorddifftypeid int64  `json:"recorddifftypeid"`
	Recorddiffid     int64  `json:"recorddiffid"`
}
type WorkflowStateResponseEntity struct {
	Activityids []int64                  `json:"activityids"`
	Groups      []WorkflowResponseEntity `json:"groups"`
}

type WorkflowTransitionEntityResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details []WorkflowTransitionEntity `json:"details"`
}
type WorkflowStateEntityResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details WorkflowStateResponseEntity `json:"details"`
}
type WorkflowEntityResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Details []WorkflowResponseEntity `json:"details"`
}
type WorkflowResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
type StateResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Details []StateEntity `json:"details"`
}

func (w *Workflowentity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
