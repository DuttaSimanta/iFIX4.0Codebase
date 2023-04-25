package entity

import (
	"encoding/xml"
	"log"
	"net/http"
)

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	ValidateARCOSUser ValidateARCOSUser `xml:"ValidateARCOSUser"`
}

type ValidateARCOSUser struct {
	ApiKey             string `xml:"ApiKey"`
	ClientCode         string `xml:"ClientCode"`
	UserId             string `xml:"UserId"`
	OrgCode            string `xml:"ScenarioName"`
	TicketNo           string `xml:"TicketNo"`
	SrTaskNumber       string `xml:"srTaskNumber"`
	CRNo               string `xml:"CRNo"`
	TaskNo             string `xml:"TaskNo"`
	ServerIP           string `xml:"ServerIP"`
	IsValidateServerIP string `xml:"IsValidateServerIP"`
	//UserGroup          string `xml:"UserGroup"`

}

type SuccessResponseXML struct {
	Status string `xml:"Success"`
	//Message string `xml:"Message"`
}

type ErrorResponseXML struct {
	Status  string `xml:"Failure"`
	Message string `xml:"Message"`
}

func ThrowSuccessResponseXML(response SuccessResponseXML, w http.ResponseWriter) {
	//log.Println(response)
	xmlResponse, xmlError := xml.Marshal(response)
	if xmlError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "Text/xml")
	w.Write(xmlResponse)
}

func ThrowErrorResponseXML(response ErrorResponseXML, w http.ResponseWriter) {
	//log.Println(response)
	xmlResponse, xmlError := xml.Marshal(response)
	if xmlError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "Text/xml")
	w.Write(xmlResponse)
}
