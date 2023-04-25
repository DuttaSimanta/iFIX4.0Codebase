package entities
import (
	//"encoding/json"
	//"io"
)


type MstExcelTemplateTypeEntity struct {

     Id                       int64  `json:"id"`
     TypeName                 string `json:"typename"`
}
/*type MstExcelTemplateTypeEntities struct {
	//Total  int64            `json:"total"`
	Values []MstExcelTemplateTypeEntity `json:"values"`
}
*/
type MstExcelTemplateTypeResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details []MstExcelTemplateTypeEntity  `json:"details"`
}