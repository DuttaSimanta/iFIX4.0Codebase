package entities

type TicketDispatchEntity struct {
	ClientID     int64  `json:"clientid"`
	OrgID        int64  `json:"mstorgnhirarchyid"`
	TicketID     int64  `json:"transactionid"`
	MstGroupID   int64  `json:"mstgroupid"`
	MstUserID    int64  `json:"mstuserid"`
	WorkingID    int64  `json:"workingcatid"`
	TicketTypeID int64  `json:"tickettypeid"`
	TicketCode   string `json:"ticketcode"`
}

type TicketDispatchEntities struct {
	Values []TicketDispatchEntity `json:"values"`
}

type GroupForwarding struct {
	TicketID       int64 `json:"transactionid"`
	MstGroupID     int64 `json:"mstgroupid"`
	Mstuserid      int64 `json:"mstuserid"`
	Createdgroupid int64 `json:"createdgroupid"`
	Samegroup      bool  `json:"samegroup"`
	UserID         int64 `json:"userid"`
}

type GetRecordDetailsRequest struct {
	ClientID          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	RecordID          int64 `json:"recordid"`
}
type GetStateDetailsRequest struct {
	ClientID          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	RecordID          int64 `json:"recordid"`
	RecordStagedID    int64 `json:"'recordstageid'"`
}
type GetStateBySeqRequest struct {
	ClientID          int64 `json:"clientid"`
	MstorgnhirarchyID int64 `json:"mstorgnhirarchyid"`
	Typeseqno         int64 `json:"typeseqno"`
	SeqNo             int64 `json:"seqno"`
	TransitionID      int64 `json:"transitionid"`
	ProcessID         int64 `json:"processid"`
	UserID            int64 `json:"userid"`
}
type MoveWorkFlowRequest struct {
	ClientID             int64 `json:"clientid"`
	MstorgnhirarchyID    int64 `json:"mstorgnhirarchyid"`
	RecorddifftypeID     int64 `json:"recorddifftypeid"`
	RecordDiffID         int64 `json:"recorddiffid"`
	TransitionID         int64 `json:"transitionid"`
	PreviousstateID      int64 `json:"previousstateid"`
	CurrentstateID       int64 `json:"currentstateid"`
	Manualstateselection int64 `json:"manualstateselection"`
	TransactionID        int64 `json:"transactionid"`
	CreatedgroupID       int64 `json:"createdgroupid"`
	MstgroupID           int64 `json:"mstgroupid"`
	MstuserID            int64 `json:"mstuserid"`
	Issrrequestor        int64 `json:"issrrequestor"`
	UserID               int64 `json:"userid"`
}
type ChangeRecordGroupRequest struct {
	CreatedgroupID int64 `json:"createdgroupid"`
	MstgroupID     int64 `json:"mstgroupid"`
	MstuserID      int64 `json:"mstuserid"`
	Samegroup      bool  `json:"samegroup"`
	TransactionID  int64 `json:"transactionid"`
	UserID         int64 `json:"userid"`
}
type RecordDetailsResponeData struct {
	Status  bool                  `json:"success"`
	Message string                `json:"message"`
	Details []RecordDetailsEntity `json:"details"`
}
type WorkFlowEntity struct {
	WorkFlowID int64 `json:"workflowid"`
	CatID      int64 `json:"catid"`
	CatTypeID  int64 `json:"cattypeid"`
}
type RecordDetailsEntity struct {
	Clientid             int64          `json:"clientid"`
	Mstorgnhirarchyid    int64          `json:"mstorgnhirarchyid"`
	Recordid             int64          `json:"recordid"`
	Title                string         `json:"title"`
	Description          string         `json:"description"`
	RecordTypeDiffTypeID int64          `json:"typedifftypeid"`
	RecordTypeID         int64          `json:"recordtypeid"`
	RecordType           string         `json:"recordtype"`
	GroupLevelID         int64          `json:"grouplevelid"`
	GroupLevel           string         `json:"grouplevel"`
	GroupID              int64          `json:"groupid"`
	Group                string         `json:"group"`
	PriorityID           int64          `json:"priorityid"`
	PriorityTypeID       int64          `json:"prioritytypeid"`
	Priority             string         `json:"priority"`
	StatusID             int64          `json:"statusid"`
	Status               string         `json:"status"`
	StatusSeqNo          int64          `json:"statusseqno"`
	ImpactID             int64          `json:"impactid"`
	Impact               string         `json:"impact"`
	UrgencyID            int64          `json:"urgencyid"`
	Urgency              string         `json:"urgency"`
	SourceType           string         `json:"source"`
	AssigneeID           int64          `json:"assigneeid"`
	Assignee             string         `json:"assignee"`
	RequestorInfo        string         `json:"requestorinfo"`
	AssignedGroupLevelID int64          `json:"assignedgrouplevelid"`
	AssignedGroupLevel   string         `json:"assignedgrouplevel"`
	AssignedGroupID      int64          `json:"assignedgroupid"`
	AssignedGroup        string         `json:"assignedgroup"`
	CreatedBy            string         `json:"createdby"`
	CreatorID            int64          `json:"creatorid"`
	ID                   int64          `json:"id"`
	Code                 string         `json:"code"`
	CreatedDateTime      string         `json:"createddatetime"`
	RecordStageID        int64          `json:"recordstageid"`
	WorkFlowDetails      WorkFlowEntity `json:"workflowdetails"`
	Vipuser              string         `json:"isvip"`
	Duedate              string         `json:"duedate"`
	RequestorName        string         `json:"requestername"`
	RequestorEmail       string         `json:"requesteremail"`
	RequestorMobile      string         `json:"requestermobile"`
	RequestorLocation    string         `json:"requesterlocation"`
	OrgRequestorName     string         `json:"orgrequestername"`
	OrgRequestorEmail    string         `json:"orgrequesteremail"`
	OrgRequestorMobile   string         `json:"orgrequestermobile"`
	OrgRequestorLocation string         `json:"orgrequesterlocation"`
	IsReslBreach         bool           `json:"resobreachcomment"`
	IsRespBreach         bool           `json:"respbreachcomment"`
	OriginalUserID       int64          `json:"originaluserid"`
	TypeSeqNo            int64          `json:"typeseqno"`
	Haspermission        bool           `json:"haspermission"`
}
type StateSeqRespEntity struct {
	Recorddifftypeid int64 `json:"recorddifftypeid"`
	RecordDiffID     int64 `json:"recorddiffid"`
	Mststateid       int64 `json:"'mststateid"`
}

// type StateSeqRespEntities struct {
// 	Values []StateSeqRespEntity `json:"'Values"`
// }
type StateSeqResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []StateSeqRespEntity `json:"details"`
}