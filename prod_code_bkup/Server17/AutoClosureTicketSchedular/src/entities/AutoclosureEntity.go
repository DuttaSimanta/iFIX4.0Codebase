package entities

type AutoClosureConfig struct {
	ClientID          int64
	MstorgnhirarchyID int64
	IsAutoclose       int64
	AutocloseFromdate string
	AutocloseTodate   string
}
type RecordInfo struct {
	RecordID             int64
	RecordStageID        int64
	Closuredate          int64
	StatusID             int64
	StatusSeq            int64
	ClientID             int64
	MstorgnhirarchyID    int64
	RecordtypeID         int64
	RecorddifftypeID     int64
	RecorddiffID         int64
	ID                   int64
	CreatedgrpID         int64
	MstuserID            int64
	WorkingDifftypeID    int64
	WorkingDiffID        int64
	PreviousStateID      int64
	TickettypeDiffId     int64
	TickettypeDiffTypeId int64
}

type RequestBody struct {
	ClientID          int64 `json:"clientid"`
	MstorgnhirarchyID int64 `json:"mstorgnhirarchyid"`
	RecorddifftypeID  int64 `json:"recorddifftypeid"`
	RecorddiffID      int64 `json:"recorddiffid"`
	PreviousstateID   int64 `json:"previousstateid"`
	CurrentstateID    int64 `json:"currentstateid"`
	TransactionID     int64 `json:"transactionid"`
	CreatedgroupID    int64 `json:"createdgroupid"`
	MstgroupID        int64 `json:"mstgroupid"`
	MstuserID         int64 `json:"mstuserid"`
	UserID            int64 `json:"userid"`
}

type WorkflowResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details string `json:"details"`
}
type RecordmultiplecommonEntity struct {
	ClientID          int64                   `json:"clientid"`
	Mstorgnhirarchyid int64                   `json:"mstorgnhirarchyid"`
	RecordID          int64                   `json:"recordid"`
	RecordstageID     int64                   `json:"recordstageid"`
	Details           []RecordTermnamesEntity `json:"details"`
	ForuserID         int64                   `json:"foruserid"`
	Userid            int64                   `"json:userid"`
	Recorddifftypeid  int64                   `json:"recorddifftypeid"`
	Recorddiffid      int64                   `json:"recorddiffid"`
	Usergroupid       int64                   `json:"usergroupid"`
}

type RecordTermnamesEntity struct {
	ID              int64  `json:"id"`
	Termname        string `json:"tername"`
	Recordtermvalue string `json:"recordtermvalue"`
	Iscompulsory    int64  `json:"iscompulsory"`
	Termtypename    string `json:"termtypename"`
	Termtypeid      int64  `json:"termtypeid"`
	Insertedvalue   string `json:"insertedvalue"`
	Seq             int64  `json:"seq"`
	Termdescription string `json:"termdescription"`
}

type RecordcommonResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
