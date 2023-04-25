package entities

// "encoding/json"
// "io"

type MstDownloadUser struct {
	ClientID int64   `json:"clientid"`
	OrgID    int64   `json:"mstorgnhirarchyid"`
	Groupid  []int64 `json:"groupid"`
}
