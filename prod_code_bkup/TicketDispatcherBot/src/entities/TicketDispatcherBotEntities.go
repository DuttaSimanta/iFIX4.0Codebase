package entities

type TicketDispatchEntity struct {
	ClientID     int64 `json:"clientid"`
	OrgID        int64 `json:"mstorgnhirarchyid"`
	TicketID     int64 `json:"transactionid"`
	MstGroupID   int64 `json:"mstgroupid"`
	MstUserID    int64 `json:"mstuserid"`
	WorkingID    int64 `json:"workingcatid"`
	TicketTypeID int64 `json:"tickettypeid"`
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
