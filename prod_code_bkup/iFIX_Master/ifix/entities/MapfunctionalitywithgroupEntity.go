package entities

import (
	"encoding/json"
	"io"
)

type MapfunctionalitywithgroupEntity struct {
	Id                     int64   `json:"id"`
	Clientid               int64   `json:"clientid"`
	Mstorgnhirarchyid      int64   `json:"mstorgnhirarchyid"`
	Recorddifftypeid       int64   `json:"recorddifftypeid"`
	Recorddiffid           int64   `json:"recorddiffid"`
	Mstfunctionailyid      int64   `json:"mstfunctionailyid"`
	Diffid                 []int64 `json:"diffid"`
	Groupid                []int64 `json:"groupid"`
	Refuserid              []int64 `json:"refuserid"`
	Activeflg              int64   `json:"activeflg"`
	Offset                 int64   `json:"offset"`
	Limit                  int64   `json:"limit"`
	Recorddifftypestatusid int64   `json:"recorddifftypestatusid"`
	Recorddiffstatusid     []int64 `json:"recorddiffstatusid"`
}

type MapfunctionalitywithgroupResponseEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Recorddifftypeid    int64  `json:"recorddifftypeid"`
	Recorddiffid        int64  `json:"recorddiffid"`
	Mstfunctionailyid   int64  `json:"mstfunctionailyid"`
	Diffid              int64  `json:"diffid"`
	Groupid             int64  `json:"groupid"`
	Refuserid           int64  `json:"refuserid"`
	Activeflg           int64  `json:"activeflg"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Recorddifftypname   string `json:"recorddifftypname"`
	Recorddiffname      string `json:"recorddiffname"`
	Mstfunctionailyname string `json:"mstfunctionailyname"`
	Diffname            string `json:"diffname"`
	Refusername         string `json:"refusername"`
	Grpname             string `json:"grpname"`

	Recorddifftypestatusid int64 `json:"recorddifftypestatusid"`
	Recorddiffstatusid     int64 `json:"recorddiffstatusid"`

	Recorddifftypestatusname string `json:"recorddifftypestatusname"`
	Recorddiffstatusname     string `json:"recorddiffstatusname"`
}

type Organizationgrpname struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Isworkflow string `json:"isworkflow"`
}

type MapfunctionalitywithgroupEntities struct {
	Total  int64                                     `json:"total"`
	Values []MapfunctionalitywithgroupResponseEntity `json:"values"`
}

type OrganizationgrpnameEntities struct {
	Total  int64                 `json:"total"`
	Values []Organizationgrpname `json:"values"`
}

type MapfunctionalitywithgroupResponse struct {
	Success bool                              `json:"success"`
	Message string                            `json:"message"`
	Details MapfunctionalitywithgroupEntities `json:"details"`
}

type OrganizationgrpnameResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details OrganizationgrpnameEntities `json:"details"`
}

type MapfunctionalitywithgroupResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MapfunctionalitywithgroupEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
