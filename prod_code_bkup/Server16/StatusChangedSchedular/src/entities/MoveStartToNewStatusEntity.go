package entities

type RequestBody struct {
	ClientID             int64 `json:"clientid"`
	MstorgnhirarchyID    int64 `json:"mstorgnhirarchyid"`
	RecorddifftypeID     int64 `json:"recorddifftypeid"`
	RecorddiffID         int64 `json:"recorddiffid"`
	PreviousstateID      int64 `json:"previousstateid"`
	CurrentstateID       int64 `json:"currentstateid"`
	TransactionID        int64 `json:"transactionid"`
	CreatedgroupID       int64 `json:"createdgroupid"`
	MstgroupID           int64 `json:"mstgroupid"`
	MstuserID            int64 `json:"mstuserid"`
	Manualstateselection int64 `json:"manualstateselection"`
	//Samegroup         bool  `json:"samegroup"`
	UserID int64 `json:"userid"`
}
type WorkflowResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details string `json:"details"`
}
