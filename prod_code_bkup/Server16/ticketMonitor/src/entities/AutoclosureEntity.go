package entities

type RecordInfo struct {
	RecordID          int64
	RecordStageID     int64
	Closuredate       int64
	StatusID          int64
	StatusSeq         int64
	ClientID          int64
	MstorgnhirarchyID int64
	RecordtypeID      int64
	RecorddifftypeID  int64
	RecorddiffID      int64
	ID                int64
	CreatedgrpID      int64
	MstuserID         int64
	WorkingDifftypeID int64
	WorkingDiffID     int64
	PreviousStateID   int64
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
}

type WorkflowResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details string `json:"details"`
}
