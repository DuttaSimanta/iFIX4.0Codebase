package entities
import (
	"encoding/json"
	"io"
)

type MstScheduledNotificationEntity struct {

     Id                       int64  `json:"id"`
     //HeaderName               string `json:"headername"`
     Clientid                 int64  `json:"clientid"`
     Clientname               string `json:"clienname"`
     Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
     Mstorgnhirarchyname      string `json:"mstorgnhirarchyname"`
     ScheduledEventid         int64  `json:"scheduledeventid"`
     ScheduledEventName       string `json:"scheduledeventname"`
     ChannelType              int64  `json:"channeltype"`
     EmailSub                 string `json:"emailsub"`
     EmailBody                string `json:"emailbody"`
     SendToUseridsArray       []int64 `json:"senduseridsarray"`
     SendToUserids            string `json:"sendtousersid"`
     SendToUserNames          string `json:"sendtousersnames"`
     SendToGroupidsArray      []int64 `json:"sendgroupidsarray"`
     SendToGroupids           string `json:"sendtogroupsid"`
     SendToGroupNames         string `json:"sendtogroupnames"`
     AdditionalRecipintArray  []string `json:"additionalrecipintarray"`
     AdditionalRecipint       string `json:"additionalrecipint"`
     TriggerConditionDays     int64  `json:"triggerconditiondays"`
     ScheduleTime             string `json:"scheduledtime"`
     RecordDiffid             int64  `json:"recorddiffid"`
     RecordDiffName           string `json:"recorddiffname"`
     RecordDiffTypeid         int64  `json:"recorddifftypeid"`
     RecordDiffTypeName       string `json:"recorddifftypename"`
     PrioritySeqNo            int64  `json:"priorityseqno"`
     PrioritySeqName          string `json:"priorityseqname"`
     Activeflg                int64  `json:"activeglg"`
     Offset                   int64  `json:"offset"`
	 Limit                    int64 `json:"limit"`
}
type GetClientAndOrgWiseclientuserEntity struct{
     Clientid                 int64  `json:"clientid"`
     Clientname               string `json:"clienname"`
}
type MstScheduledNotificationEntities struct {
	Total  int64            `json:"total"`
	Values []MstScheduledNotificationEntity `json:"values"`
}

type MstScheduledNotificationResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MstScheduledNotificationEntities  `json:"details"`
}

type GetClientAndOrgWiseclientuserResponse struct {
     Success bool             `json:"success"`
     Message string           `json:"message"`
     Details []GetClientAndOrgWiseclientuserEntity  `json:"details"`
}
type MstScheduledNotificationResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstScheduledNotificationEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
