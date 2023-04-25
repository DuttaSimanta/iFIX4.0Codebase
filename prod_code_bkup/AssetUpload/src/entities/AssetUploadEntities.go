package entities

type TrnAsset struct {
	Id                 int64
	ClientId           int64
	MstOrgnHirarchyId  int64
	MstDifftypeid      int64
	AssetId            string
	AdditionalAttr     string
	DeleteFlag         int64
	ActiveFlag         int64
	AuditTransactionId int64
}
type MapAssetDifferentiation struct {
	Id                       int64  `json:"id"`
	Clientid                 int64  `json:"clientid"`
	Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
	Mstdifferentiationtypeid int64  `json:"mstdifferentiationtypeid"`
	Mstdifferentiationid     int64  `json:"mstdifferentiationid"`
	Trnassetid               int64  `json:"trnassetid"`
	Value                    string `json:"value"`
	Deleteflg                int64  `json:"deleteflg"`
	Activeflg                int64  `json:"activeflg"`
	AuditTransactionId       int64  `json:"audittransactionid"`
}
