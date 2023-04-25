package dao


import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var sqlgetcred = "SELECT id, clientid, mstorgnhirarchyid, credentialtypeid, credentialaccount, credentialpassword, credentialkey, activeflg FROM mstclientcredential WHERE clientid = ? AND mstorgnhirarchyid = ? AND credentialtypeid=1 AND activeflg = 1 AND deleteflg = 0 "


func (dbc DbConn) GetCredentialById(page *entities.FileuploadEntity) (entities.FileuploadEntity, error) {
    logger.Log.Println("In side GetCredentialById")
    value := entities.FileuploadEntity{}
    rows, err := dbc.DB.Query(sqlgetcred, page.Clientid, page.Mstorgnhirarchyid)
    defer rows.Close()
    if err != nil {
        logger.Log.Println("GetAllMsttermtype Get Statement Prepare Error", err)
        return value, err
    }
    for rows.Next() {
        rows.Scan(&value.Id,&value.Clientid,&value.Mstorgnhirarchyid,&value.Credentialtype,&value.Credentialaccount,&value.Credentialpassword,&value.Credentialkey,&value.Activeflg)
    }
    return value, nil
}
