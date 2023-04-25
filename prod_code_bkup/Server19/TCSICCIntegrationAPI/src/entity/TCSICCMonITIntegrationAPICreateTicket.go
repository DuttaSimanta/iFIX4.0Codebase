package entity

//RecordData is a details value of Recordsets entity
type RecordData struct {
	ID  int64 `json:"id"`
	Val int64 `json:"val"`
}

//RecordSet is a details value of Recordsets entity
type RecordSet struct {
	ID   int64        `json:"id"`
	Type []RecordData `json:"type"`
	Val  int64        `json:"val"`
}

//RecordAdditional is a details value of RecordAdditional entity
type RecordAdditional struct {
	ID      int64  `json:"id"`
	Termsid int64  `json:"termsid"`
	Val     string `json:"val"`
}

//RecordField is a details value of Recordsets entity
type RecordField struct {
	// DifftypeID int64        `json:"difftypeid"`
	// DiffID     int64        `json:"diffid"`
	// Terms      []RecordTerm `json:"terms"`
	TermID int64        `json:"termid"`
	Val    []RecordTerm `json:"val"`
}

//RecordTerm is a details value of Recordsets entity
type RecordTerm struct {
	OriginalName string `json:"originalName"`
	FileName     string `json:"fileName"`
}

type RecordEntity struct {
	//ID                  int64              `json:"id"`
	ClientID            int64              `json:"clientid"`
	Mstorgnhirarchyid   int64              `json:"mstorgnhirarchyid"`
	Requesterinfo       string             `json:"requesterinfo"`
	RecordTypeSeq       int64              `json:"recordtypeseq"`
	RecordTypeID        int64              `json:"recordtypeid"`
	Recordname          string             `json:"recordname"`
	Recordesc           string             `json:"recordesc"`
	RecordPriorityID    string             `json:"recordpriorityid"`
	Recordattachpath    string             `json:"recordattachpath"`
	Recordcategorydtls  string             `json:"recordcategorydtls"`
	Recordpsdetails     string             `json:"recordpsdetails"`
	Recordsourcetype    string             `json:"recordsourcetype"`
	RecordcreatedID     int64              `json:"recordcreatedid"`
	RecordoriginalID    int64              `json:"recordoriginalid"`
	Recordrequestinfo   string             `json:"recordrequestinfo"`
	RecordSets          []RecordSet        `json:"recordsets"`
	Recordfields        []RecordField      `json:"recordfields"`
	ResponsedifftypeID  []int64            `json:"responsedifftypeid"`
	Createduserid       int64              `json:"createduserid"`
	Createdusergroupid  int64              `json:"createdusergroupid"`
	Userid              int64              `json:"userid"`
	Originaluserid      int64              `json:"originaluserid"`
	Originalusergroupid int64              `json:"originalusergroupid"`
	AssetIds            []int64            `json:"assetIds"`
	Workingcatlabelid   int64              `json:"workingcatlabelid"`
	Additionalfields    []RecordAdditional `json:"additionalfields"`
	Requestername       string             `json:"requestername"`
	Requesteremail      string             `json:"requesteremail"`
	Requestermobile     string             `json:"requestermobile"`
	Requesterlocation   string             `json:"requesterlocation"`
	Source              string             `json:"source"`
}

type PriorityChange struct {
	ClientID            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	PriorityDiffTypeID  int64  `json:"recorddifftypeid"`
	PriorityDiffID      int64  `json:"recorddiffid"`
	RecordID            int64  `json:"recordid"`
	Createduserid       int64  `json:"userid"`
	Createdusergroupid  int64  `json:"usergroupid"`
	Originaluserid      int64  `json:"originaluserid"`
	Originalusergroupid int64  `json:"originalusergroupid"`
	Recordname          string `json:"recordname"`
	Recordesc           string `json:"recordesc"`
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
