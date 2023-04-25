package entities

import (
	"encoding/json"
	"io"
)

type MsttemplateEntity struct {
	ID                      int64                   `json:"id"`
	Clientid                int64                   `json:"clientid"`
	Clientname              string                  `json:"clientname"`
	Mstorgnhirarchyid       int64                   `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname     string                  `json:"mstorgnhirarchyname"`
	Templatetype            int64                   `json:"templatetype"`
	Templatename            string                  `json:"templatename"`
	Templatecontent         string                  `json:"templatecontent"`
	Autoclosuretime         int64                   `json:"autoclosuretime"`
	Activeflg               int64                   `json:"activeflg"`
	MaptemplatediffEntities []MaptemplatediffEntity `json:"maptemplatediff"`
	Offset                  int64                   `json:"offset"`
	Limit                   int64                   `json:"limit"`
}
type MstNotificationRecipientEntity struct {
	ID                     int64  `json:"id"`
	NotificationTemplateID int64  `json:"notificationtemplateid"`
	RecipientType          string `json:"recipienttype"`
	GroupID                int64  `json:"groupid"`
	GroupName              string `json:"groupname"`
	UserID                 int64  `json:"userid"`
	UserName               string `json:"username"`
}
type NotificationEventParamsEntity struct {
	StatusID        int64 `json:"statusid"`
	PriorityID      int64 `json:"priorityid"`
	NoOfCount       int64 `json:"noofcount"`
	ProcessID       int64 `json:"processid"`
	ProcessComplete int64 `json:"processcomplete"`
	NoOfDays        int64 `json:"noofdays"`
}
type MstNotificationTemplateEntity struct {
	ID                        int64                            `json:"id"`
	ClientID                  int64                            `json:"clientid"`
	ClientName                string                           `json:"clientname"`
	MstOrgnHirarchyID         int64                            `json:"mstorgnhirarchyid"`
	MstOrgnHirarchyName       string                           `json:"mstorgnhirarchyname"`
	RecordTypeTypeID          int64                            `json:"recordtypetypeid"`
	RecordTypeType            string                           `json:"recordtypetype"`
	RecordTypeID              int64                            `json:"recordtypeid"`
	RecordType                string                           `json:"recordtype"`
	WorkingCategoryTypeID     int64                            `json:"workingcategorytypeid"`
	WorkingCategoryType       string                           `json:"workingcategorytype"`
	WorkingCategoryID         int64                            `json:"workingcategoryid"`
	WorkingCategories         []int64                          `json:"workingcategories"`
	WorkingCategory           string                           `json:"workingcategory"`
	ChannelTypeID             int64                            `json:"channeltypeid"`
	ChannelType               string                           `json:"channeltype"`
	SubjectOrTitle            string                           `json:"subjectortitle"`
	Body                      string                           `json:"body"`
	AdditionalRecipient       string                           `json:"additionalrecipient"`
	SendToCreator             int64                            `json:"sendtocreator"`
	SendToOrgCreator          int64                            `json:"sendtoorgcreator"`
	SendToAssignee            int64                            `json:"sendtoassignee"`
	SendToAssigneeGroup       int64                            `json:"sendtoassigneegroup"`
	SendToAssigneeGroupMember int64                            `json:"sendtoassigneegroupmember"`
	EventTypeID               int64                            `json:"eventtypeid"`
	EventType                 string                           `json:"eventtype"`
	EventParams               NotificationEventParamsEntity    `json:"eventparams"`
	Recipients                []MstNotificationRecipientEntity `json:"recipients"`
	Activeflg                 int64                            `json:"activeflg"`
	Offset                    int64                            `json:"offset"`
	Limit                     int64                            `json:"limit"`
	SmsTemplateID             string                           `json:"smstemplateid"`
	SmsTypeID                 int64                            `json:"smstypeid"`
	SmsType                   string                           `json:"smstype"`
	StatusName                string                           `json:"statusname"`
	Isconverted               int64                            `json:"isconverted"`
	Convertedtype             string                           `json:"convertedtype"`
}

type MstNotificationTemplateEntities struct {
	Total  int64                           `json:"total"`
	Values []MstNotificationTemplateEntity `json:"values"`
}

type MaptemplatediffEntity struct {
	ID                               int64  `json:"id"`
	Clientid                         int64  `json:"clientid"`
	Mstorgnhirarchyid                int64  `json:"mstorgnhirarchyid"`
	Templateid                       int64  `json:"templateid"`
	Recorddifftypeid                 int64  `json:"recorddifftypeid"`
	Mstrecorddifferentiationtypename string `json:"mstrecorddifferentiationtypename"`
	Recorddiffid                     int64  `json:"recorddiffid"`
	Mstrecorddifferentiationname     string `json:"mstrecorddifferentiationname"`
	Activeflg                        int64  `json:"activeflg"`
	RecorddifftypeParentid           int64  `json:"recorddifftypeParentid"`
}

type MsttemplateEntities struct {
	Total  int64               `json:"total"`
	Values []MsttemplateEntity `json:"values"`
}
type MstNotificationVariable struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type MstNotificationEvent struct {
	ID   int64  `json:"id"`
	Name string `json:"eventname"`
}
type MstNotificationVariableResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details []MstNotificationVariable `json:"details"`
}
type MstNotificationEventResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details []MstNotificationEvent `json:"details"`
}
type MsttemplateResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Details MsttemplateEntities `json:"details"`
}

type MsttemplateResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
type MsttemplateResponseIntArr struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Details []int64 `json:"details"`
}
type MstNotificationTemplateResponse struct {
	Success bool                            `json:"success"`
	Message string                          `json:"message"`
	Details MstNotificationTemplateEntities `json:"details"`
}

func (w *MsttemplateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
func (w *MstNotificationTemplateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
