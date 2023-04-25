package entities

import (
	"database/sql"
	"encoding/json"
	"iFIX/ifix/logger"
	"io"
	"sync"
)

type RecordfulldetailsRequestEntity struct {
	Clientid           int64   `json:"clientid"`
	Mstorgnhirarchyid  int64   `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyids []int64 `json:"mstorgnhirarchyids"`
	Fromdate           string  `json:"fromdate"`
	Todate             string  `json:"todate"`
	Seqno              int64   `json:"seqno"`
	Tickettypeseq      int64   `json:"tickettypeseq"`
	Diffstatusseqnos   []int64 `json:"diffstatusseqnos"`
}
type RecordDiffOfMultiOrgEntity struct {
	ID                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Name              string `json:"name"`
	Seqno             int64  `json:"seqno"`
}
type Categorydetails struct {
	Label        string `json:"label"`
	Categoryname string `json:"Categoryname"`
}

type ExcelEntity struct {
	Clientid           int64             `json:"clientid"`
	Mstorgnhirarchyid  int64             `json:"mstorgnhirarchyid"`
	Recordid           int64             `json:"recordid"`
	Tickettypeid       int64             `json:"tickettypeid"`
	CustomerName       string            `json:"customername"`
	TicketNo           string            `json:"ticketno"`
	Shortdescription   string            `json:"shortdescription"`
	Status             string            `json:"status"`
	Latestresodatetime string            `json:"latestresodatetime"`
	VendorName         string            `json:"vendorname"`
	VendorTicketid     string            `vendorticketid`
	Category           []Categorydetails `json:"category"`
}

func (w *RecordfulldetailsRequestEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

// type ResultEntity struct {
// 	Recordfullresult []map[string]interface{} `json:"result"`
// 	Order            []string                 `json:"order"`
// }
type ResultEntity struct {
	Recordfullresult []ExcelEntity `json:"result"`
	Order            []string      `json:"order"`
}
type RecordfulldetailsResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details ResultEntity `json:"details"`
}
type RecordDiffOfMultiOrgResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details []RecordDiffOfMultiOrgEntity `json:"details"`
}
type RecordfulldetailsresultEntities struct {
	Result []*interface{}
}

type NullString struct {
	sql.NullString
}

type NullInt64 struct {
	sql.NullInt64
}

type NullFloat64 struct {
	sql.NullFloat64
}

// NewDbFieldBind ...
func NewDbFieldBind() *DbFieldBind {
	return &DbFieldBind{}
}

// FieldBinding is deisgned for SQL rows.Scan() query.
type DbFieldBind struct {
	sync.RWMutex // embedded.  see http://golang.org/ref/spec#Struct_types
	FieldArr     []interface{}
	FieldPtrArr  []interface{}
	FArr         []string
	FieldCount   int64
	FArrTypes    []*sql.ColumnType
	MapFieldToID map[string]int64
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}
func (s NullInt64) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int64)
}
func (s NullFloat64) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Float64)
}
func (fb *DbFieldBind) put(k string, v int64) {
	fb.Lock()
	defer fb.Unlock()
	fb.MapFieldToID[k] = v
}

// Get ...
func (fb *DbFieldBind) Get(k string) interface{} {
	fb.RLock()
	defer fb.RUnlock()
	// TODO: check map key exist and fb.FieldArr boundary.
	return fb.FieldPtrArr[fb.MapFieldToID[k]]
}

// PutFields ...
func (fb *DbFieldBind) PutFields(rows *sql.Rows) error {
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	fb.FArrTypes = colTypes
	fCount := len(colTypes)
	fb.FieldPtrArr = make([]interface{}, fCount)
	fb.MapFieldToID = make(map[string]int64, fCount)

	for k, v := range colTypes {
		switch v.DatabaseTypeName() {
		case "VARCHAR":
			fb.FieldPtrArr[k] = new(NullString)
		case "TEXT":
			fb.FieldPtrArr[k] = new(NullString)
		case "TIMESTAMP":
			fb.FieldPtrArr[k] = new(NullString)
		case "INT":
			fb.FieldPtrArr[k] = new(NullInt64)
		case "FLOAT":
			fb.FieldPtrArr[k] = new(NullFloat64)
		default:
			fb.FieldPtrArr[k] = new(NullString)
			logger.Log.Println("Column Data Type:", v.DatabaseTypeName())
		}
		//fb.FieldPtrArr[k] = new(interface{})
		//	logger.Log.Println("Column Data Type:", v.DatabaseTypeName())
		fb.put(v.Name(), int64(k))
	}
	return nil
}

// GetFieldPtrArr ...
func (fb *DbFieldBind) GetFieldPtrArr() []interface{} {
	//logger.Log.Println("bf", fb.FieldPtrArr)

	return fb.FieldPtrArr
}

// GetFieldArr ...
func (fb *DbFieldBind) GetFieldArr() map[string]interface{} {
	m := make(map[string]interface{}, fb.FieldCount)

	for k, v := range fb.MapFieldToID {

		switch fb.FieldPtrArr[v].(type) {
		case NullString:
			data := fb.FieldPtrArr[v].(NullString)
			aa, _ := data.MarshalJSON()
			logger.Log.Println("------------", aa)
		case NullInt64:
			data := fb.FieldPtrArr[v].(NullInt64)
			aa, _ := data.MarshalJSON()
			logger.Log.Println("------------", aa)
		case NullFloat64:
			data := fb.FieldPtrArr[v].(NullFloat64)
			aa, _ := data.MarshalJSON()
			logger.Log.Println("------------", aa)
		}
		m[k] = fb.FieldPtrArr[v]
	}

	return m
}

// type RecordfulldetailsresultEntities struct {
// 	Id                           interface{}
// 	Clientid                     interface{}
// 	Mstorgnhirarchyid            interface{}
// 	Recordid                     interface{}
// 	Ticketid                     interface{}
// 	Source                       interface{}
// 	Requestorid                  interface{}
// 	Requestorname                interface{}
// 	Requestorlocation            interface{}
// 	Requestorphone               interface{}
// 	Requestoremail               interface{}
// 	Requestorloginid             interface{}
// 	Orgcreatorlocation           interface{}
// 	Orgcreatorphone              interface{}
// 	Orgcreatorid                 interface{}
// 	Orgcreatorname               interface{}
// 	Orgcreatorloginid            interface{}
// 	Orgcreatoremail              interface{}
// 	Tickettypeid                 interface{}
// 	Tickettype                   interface{}
// 	Priorityid                   interface{}
// 	Priority                     interface{}
// 	Statusid                     interface{}
// 	Status                       interface{}
// 	Vipticket                    interface{}
// 	Urgencyid                    interface{}
// 	Urgency                      interface{}
// 	Impactid                     interface{}
// 	Impact                       interface{}
// 	Shortdescription             interface{}
// 	Assigneduserloginid          interface{}
// 	Createddatetime              interface{}
// 	Lastupdateddatetime          interface{}
// 	Assignedgroupid              interface{}
// 	Assignedgroup                interface{}
// 	Assigneduserid               interface{}
// 	Assigneduser                 interface{}
// 	Resogroupid                  interface{}
// 	Resogroup                    interface{}
// 	Resolveduserid               interface{}
// 	Resolveduser                 interface{}
// 	Lastuserid                   interface{}
// 	Lastuser                     interface{}
// 	Reassigncount                interface{}
// 	Respslastatusid              interface{}
// 	Respslabreachstatus          interface{}
// 	Resostatusid                 interface{}
// 	Resolslabreachstatus         interface{}
// 	Reopencount                  interface{}
// 	Reopendatetime               interface{}
// 	Reopenedflag                 interface{}
// 	Prioritycount                interface{}
// 	Followupcount                interface{}
// 	Followupdatetime             interface{}
// 	Followuprespdatetime         interface{}
// 	Followuptimetaken            interface{}
// 	Outboundcount                interface{}
// 	Isparent                     interface{}
// 	Childcount                   interface{}
// 	Parentticketid               interface{}
// 	Responseslameterpercentage   interface{}
// 	Resolutionslameterpercentage interface{}
// 	Worknotenotupdated           interface{}
// 	Respsladuedatetime           interface{}
// 	Resosladuedatetime           interface{}
// 	Categorychangecount          interface{}
// 	Firstresponsedatetime        interface{}
// 	Latestresponsedatetime       interface{}
// 	Firstresodatetime            interface{}
// 	Latestresodatetime           interface{}
// 	Responsetime                 interface{}
// 	Resolutiontime               interface{}
// 	Respclockstatus              interface{}
// 	Resoclockstatus              interface{}
// 	Businessaging                interface{}
// 	Calendaraging                interface{}
// 	Actualeffort                 interface{}
// 	Slaidletime                  interface{}
// 	Respoverduetime              interface{}
// 	Respoverdueperc              interface{}
// 	Resooverduetime              interface{}
// 	Resooverdueperc              interface{}
// 	Pendinguserdatetime          interface{}
// 	Userreplieddatetime          interface{}
// 	Userreplytimetaken           interface{}
// 	Pendingusercount             interface{}
// 	Closedatetime                interface{}
// 	Csatscore                    interface{}
// 	Csatcomment                  interface{}
// 	Activeflg                    interface{}
// 	Deleteflg                    interface{}
// 	Audittransactionid           interface{}
// 	Ifixsysid                    interface{}
// 	Lastupdatedby                interface{}
// 	Lastupdateddate              interface{}
// }
