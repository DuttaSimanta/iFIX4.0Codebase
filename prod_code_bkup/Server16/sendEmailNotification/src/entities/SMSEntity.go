package entities

type SMSServiceEntity struct {
	CustomerID         string   `json:"customerId"`
	DestinationAddress []string `json:"destinationAddress"`
	Message            string   `json:"message"`
	SourceAddress      string   `json:"sourceAddress"`
	MessageType        string   `json:"messageType"`
	DltTemplateID      string   `json:"dltTemplateId"`
	EntityID           string   `json:"entityId"`
}
