package entities

import (
	"encoding/json"
	"io"
)

type MstbusinessmatrixEntity struct {
	Id                                   int64  `json:"id"`
	Clientid                             int64  `json:"clientid"`
	Mstorgnhirarchyid                    int64  `json:"mstorgnhirarchyid"`
	Mstrecorddifferentickettypeid        int64  `json:"mstrecorddifferentickettypeid"`
	Mstrecorddifferentiationtickettypeid int64  `json:"mstrecorddifferentiationtickettypeid"`
	Mstrecorddifferentiationcatid        int64  `json:"mstrecorddifferentiationcatid"`
	Mstrecorddifferentiationimpactid     int64  `json:"mstrecorddifferentiationimpactid"`
	Mstrecorddifferentiationurgencyid    int64  `json:"mstrecorddifferentiationurgencyid"`
	Mstrecorddifferentiationpriorityid   int64  `json:"mstrecorddifferentiationpriorityid"`
	Activeflg                            int64  `json:"activeflg"`
	Offset                               int64  `json:"offset"`
	Limit                                int64  `json:"limit"`
	Clientname                           string `json:"clientname"`
	Mstorgnhirarchyname                  string `json:"mstorgnhirarchyname"`
	Tickettype                           string `json:"tickettype"`
	Mstrecordtickettypedifftypeid        int64  `json:"mstrecordtickettypedifftypeid"`
	Categoryname                         string `json:"categoryname"`
	Mstrecordcatlabelid                  int64  `json:"mstrecordcatlabelid"`
	Impactname                           string `json:"impactname"`
	Mstrecordimpactdifftypeid            int64  `json:"mstrecordimpactdifftypeid"`
	Urgencyname                          string `json:"urgencyname"`
	Mstrecordurgencydifftypeid           int64  `json:"mstrecordurgencydifftypeid"`
	Priorityname                         string `json:"Priorityname"`
	Mstrecordprioritydifftypeid          int64  `json:"mstrecordprioritydifftypeid"`
	Catname                              string
	Parentcatname                        string
	Estimatedeffort                      string `json:"estimatedeffort"`
	Slacompliance                        string `json:"slacompliance"`
	ChangeType                           string `json:"changetype"`
}

type MstbusinessmatrixEntities struct {
	Total  int64                     `json:"total"`
	Values []MstbusinessmatrixEntity `json:"values"`
}

type MstlastlevelEntity struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	Lastcategorylevelid int64  `json:"lastlevelcategoryid"`
	Catname             string
	Parentcatname       string
}

type MstlastlevelEntityResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []MstlastlevelEntity `json:"details"`
}

type MstbusinessmatrixResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details MstbusinessmatrixEntities `json:"details"`
}

type MstbusinessmatrixconfigurationEntities struct {
	Direction int64 `json:"direction"`
}

type MstbusinessmatrixResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstbusinessmatrixEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
