package entities

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
