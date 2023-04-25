package entities

import (
	"encoding/json"
	"io"
)

type WorkflowUtilityEntity struct{
	ID             		 int64    `json:"id"`
	Clientid             int64    `json:"clientid"`
	Mstorgnhirarchyid    int64    `json:"mstorgnhirarchyid"`
	Fieldid    			 int64    `json:"Fieldid"`
	Type				 int 	  `json:"type"`
	Processid			 int64 	  `json:"processid"`
	Seqno			 	 int64 	  `json:"seqno"`
	Typeseqno			 int64 	  `json:"Typeseqno"`
	Recorddifftypeid	 int64 	  `json:"recorddifftypeid"`
	Recorddiffid	 	 int64 	  `json:"recorddiffid"`
	Mststateid	 	 int64 	  		`json:"mststateid"`
	Transitionid	 	 int64 	  `json:"transitionid"`
	Loginname		     string   `json:"loginname"`
}
type TicketUserEntity struct{
	ID             		 int64    `json:"id"`
	Userid             	 int64    `json:"userid"`
	Groupid              int64    `json:"groupid"`
	Username             string    `json:"username"`
	Recordid    		 int64    `json:"recordid"`
	Opendate    		 int64    `json:"opendate"`
}
type StagingUtilityEntity struct{
	ID             		 int64    `json:"id"`
	Clientid             int64    `json:"clientid"`
	Mstorgnhirarchyid    int64    `json:"mstorgnhirarchyid"`
	Assignedgroupid      int64    `json:"assignedgroupid"`
	Assigneduserid       int64    `json:"assigneduserid"`
	Assigneduser         string   `json:"assigneduser"`
	Assignedloginname    string   `json:"assignedloginname"`
	Lastuserid			 int64 	  `json:"lastuserid"`
	Lastuser			 string   `json:"lastuser"`
	Reassigncount		 int64 	  `json:"reassigncount"`
	Recordid			 int64 	  `json:"recordid"`
	Assignedgroup		 string   `json:"assignedgroup"`
}
type StateStatusEntity struct{
	Recorddifftypeid	 int64 	  `json:"recorddifftypeid"`
	Recorddiffid	 	 int64 	  `json:"recorddiffid"`
	Mststateid	 	 int64 	  `json:"mststateid"`
}

type WorkflowSingleEntity struct{
	ID             int64    `json:"id"`
	Seqno          int64    `json:"seqno"`
	Name          string    `json:"name"`
}
type MapStateEntity struct{
	Stateid             int64    `json:"stateid"`
	Statename          	string    `json:"statename"`
}
type MapStateTypeEntity struct{
	Statetypeid         int64    `json:"statetypeid"`
	Statetypename       string    `json:"statetypename"`
	States            []MapStateEntity   `json:"states"`
}
type StateCategory struct{
	Recorddifftypeid	 int64 	  `json:"recorddifftypeid"`
	Recorddiffid	 	 int64 	  `json:"recorddiffid"`
	States              []MapStateTypeEntity   `json:"states"`
}
type StateProcessEntity struct{
	Stateid             int64    `json:"stateid"`
	Statename          	string    `json:"statename"`
	Statetypeid         int64    `json:"statetypeid"`
	Statetypename       string    `json:"statetypename"`
}

type StateProcessResponse struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	StateCategory `json:"details"`
}
type StateStatusResponse struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	[]StateStatusEntity `json:"details"`
}
type WorkflowUtilityResponse struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	[]WorkflowSingleEntity `json:"details"`
}
func (w *WorkflowUtilityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
func (w *TicketUserEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

