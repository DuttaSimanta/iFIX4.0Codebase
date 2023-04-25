package dao

import (
	"database/sql"
	"errors"
	Logger "src/logger"
)

func GetOrgName(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, rerecordDiffID int64) (string, string, error) {
	var orgName string
	var ticketTypeName string
	var OrgNameQuery string = "SELECT a.name,b.name FROM mstorgnhierarchy a,mstrecorddifferentiation b  where a.clientid = b.clientid and a.id = b.mstorgnhirarchyid and b.id=? and b.activeflg=1 and b.deleteflg=0"
	OrgNameScanErr := db.QueryRow(OrgNameQuery, rerecordDiffID).Scan(&orgName, &ticketTypeName)
	if OrgNameScanErr != nil {
		Logger.Log.Println(OrgNameScanErr)
		return orgName, ticketTypeName, errors.New("ERROR: Scan Error For GetOrgName")
	}
	return orgName, ticketTypeName, nil
}

func Getheaderr(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, rerecordDiffID int64) ([]string, error) {
	var headerName []string

	var selectHeaderForCategoryQuery string = "select headername from mstexceltemplate where clientid=? and mstorgnhirarchyid=? and templatetypeid=3 and recorddiffid=?  and deleteflg=0 and activeflg=1 order by seqno asc"
	//fetching category header Details and storing into slice
	categoryHeadeResultSet, err := db.Query(selectHeaderForCategoryQuery, clientID, mstOrgnHirarchyId, rerecordDiffID)
	if err != nil {
		Logger.Log.Println(err)

		return headerName, errors.New("ERROR: Unable to fetch categoryHeaderResultSet")
	}
	defer categoryHeadeResultSet.Close()
	for categoryHeadeResultSet.Next() {
		var header string
		//	var  diffTypeId int64
		err = categoryHeadeResultSet.Scan(&header)
		if err != nil {
			Logger.Log.Println(err)

			return headerName, errors.New("ERROR: Unable to scan categoryHeadeResultSet")
		}
		headerName = append(headerName, header)
	}
	return headerName, nil
}

/*func GetLasRecorddifftypeid(db *sql.DB, clientID int64, mstorgnhirarchyid int64) (int64, error) {
	var lasRecorddifftypeid int64
	var selectLasRecorddifftypeidQuery = "select id from  mstrecorddifferentiationtype c where c.clientid=? and c.mstorgnhirarchyid=? and c.parentid =1 and c.deleteflg=0 and c.activeflg=1 order by seqno asc limit 1 offset 4"
	lasRecorddifftypeidResult, err := db.Query(selectLasRecorddifftypeidQuery, clientID, mstorgnhirarchyid)
	if err != nil {
		Logger.Log.Println(err)
		return lasRecorddifftypeid, err
	}
	for lasRecorddifftypeidResult.Next() {
		//var lasRecorddifftypeid int64
		//	var  diffTypeId int64
		err = lasRecorddifftypeidResult.Scan(&lasRecorddifftypeid)
		if err != nil {
			Logger.Log.Println(err)

			return lasRecorddifftypeid, err
		}
		//headerName = append(headerName, header)
	}
	return lasRecorddifftypeid, nil
}*/
func GetParentcatagory(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, rerecordDiffID int64) ([]string, []string, []string, []string, []string, []string, []string, error) {
	var parentCategoryNames []string
	var impactNames []string
	var urgencyNames []string
	var priorityNames []string
	var estimatedTimes []string
	var efficiencies []string
	var changeTypes []string

	var selectParentcategoryQuery = " SELECT a.id as iD,a.parentcategorynames as parentCatagoryNames,e.name as impactNames, " +
		" f.name as urgencyNames,g.name as priorityNames,h.estimatedtime as estimatedTimes,h.efficiency efficiencies,h.changetype changetype FROM " +
		" mstrecorddifferentiation a,mstbusinessmatrix d,mstrecorddifferentiation e,mstrecorddifferentiation f,  " +
		" mstrecorddifferentiation g,mapcategorywithestimatetime h where a.clientid=? and a.mstorgnhirarchyid=? and " +
		" a.id=d.mstrecorddifferentiationcatid and d.mstrecorddifferentiationimpactid=e.id and d.mstrecorddifferentiationurgencyid=f.id " +
		" and d.mstrecorddifferentiationpriorityid=g.id and a.id=h.recorddiffid and  a.recorddifftypeid in  " +
		" (SELECT max(id) FROM mstrecorddifferentiationtype where parentid=1 and  id in (SELECT torecorddifftypeid FROM mstrecordtype where " +
		" clientid=a.clientid and mstorgnhirarchyid=a.mstorgnhirarchyid and torecorddiffid=0 and fromrecorddiffid=d.mstrecorddifferentiationtickettypeid) and deleteflg=0 " +
		" and activeflg=1 ) and d.clientid=a.clientid and d.mstorgnhirarchyid=a.mstorgnhirarchyid and  d.mstrecorddifferentiationtickettypeid=? and d.activeflg=1 and  " +
		" d.deleteflg=0 and e.clientid=a.clientid and e.mstorgnhirarchyid=a.mstorgnhirarchyid and e.activeflg=1 and  " +
		" e.deleteflg=0 and f.clientid=a.clientid and f.mstorgnhirarchyid=a.mstorgnhirarchyid and f.activeflg=1 and f.deleteflg=0" +
		" and g.clientid=a.clientid and g.mstorgnhirarchyid=a.mstorgnhirarchyid and g.activeflg=1 and g.deleteflg=0 and  " +
		" h.clientid=a.clientid and h.mstorgnhirarchyid=a.mstorgnhirarchyid and h.activeflg=1 and h.deleteflg=0 and a.activeflg=1 " +
		" and a.deleteflg=0"
	parentcategoryResult, err := db.Query(selectParentcategoryQuery, clientID, mstOrgnHirarchyId, rerecordDiffID)
	if err != nil {
		Logger.Log.Println(err)

		return parentCategoryNames, impactNames, urgencyNames, priorityNames, estimatedTimes, efficiencies, changeTypes, err
	}
	defer parentcategoryResult.Close()
	for parentcategoryResult.Next() {
		var id int64
		var parentCategoryName string
		var impactName string
		var urgencyName string
		var priorityName string
		var estimatedTime string
		var efficienci string
		var changeType string
		//	var  diffTypeId int64
		err = parentcategoryResult.Scan(&id, &parentCategoryName, &impactName, &urgencyName, &priorityName, &estimatedTime, &efficienci, &changeType)
		if err != nil {
			Logger.Log.Println(err)

			return parentCategoryNames, impactNames, urgencyNames, priorityNames, estimatedTimes, efficiencies, changeTypes, err
		}
		//headerName = append(headerName, header)
		parentCategoryNames = append(parentCategoryNames, parentCategoryName)
		impactNames = append(impactNames, impactName)
		urgencyNames = append(urgencyNames, urgencyName)
		priorityNames = append(priorityNames, priorityName)
		estimatedTimes = append(estimatedTimes, estimatedTime)
		efficiencies = append(efficiencies, efficienci)
		changeTypes = append(changeTypes, changeType)

	}
	return parentCategoryNames, impactNames, urgencyNames, priorityNames, estimatedTimes, efficiencies, changeTypes, err
}

/*func GetImpactUrgencyPriorityNames(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, lastRecorddifftypeid int64) ([]string, []string, []string, error) {

	var impactids []int64
	var impactNames []string
	var urgencyIds []int64
	var urgencyNames []string
	var priorityIds []int64
	var priorityNames []string
	//Getting impact Datails
	var getimpact = "SELECT c.id ,a.name from mstrecorddifferentiation a,mstbusinessmatrix b,mstrecorddifferentiation c where c.recorddifftypeid=? and c.id=b.mstrecorddifferentiationcatid and mstrecorddifferentiationimpactid=a.id and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and  a.activeflg=1 and a.deleteflg=0 and  b.activeflg=1 and b.deleteflg=0 and c.activeflg=1 and c.deleteflg=0"
	impactResult, err := db.Query(getimpact, lastRecorddifftypeid, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println(err)

		return impactNames, urgencyNames, priorityNames, err
	}
	for impactResult.Next() {
		var impactid int64
		var impactName string
		//	var  diffTypeId int64
		err = impactResult.Scan(&impactid, &impactName)
		if err != nil {
			Logger.Log.Println(err)

			return impactNames, urgencyNames, priorityNames, err
		}
		//headerName = append(headerName, header)
		impactNames = append(impactNames, impactName)
		impactids = append(impactids, impactid)
	}
	//Getting urgency details
	var geturgency = "SELECT c.id ,a.name from mstrecorddifferentiation a,mstbusinessmatrix b,mstrecorddifferentiation c where c.recorddifftypeid=? and c.id=b.mstrecorddifferentiationcatid and b.mstrecorddifferentiationurgencyid=a.id and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and  a.activeflg=1 and a.deleteflg=0 and  b.activeflg=1 and b.deleteflg=0 and c.activeflg=1 and c.deleteflg=0"
	urgenyResult, err := db.Query(geturgency, lastRecorddifftypeid, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println(err)

		return impactNames, urgencyNames, priorityNames, err
	}
	for urgenyResult.Next() {
		var urgencyId int64
		var urgencyName string
		//	var  diffTypeId int64
		err = urgenyResult.Scan(&urgencyId, &urgencyName)
		if err != nil {
			Logger.Log.Println(err)

			return impactNames, urgencyNames, priorityNames, err
		}
		//headerName = append(headerName, header)
		urgencyNames = append(urgencyNames, urgencyName)
		urgencyIds = append(urgencyIds, urgencyId)
	}
	//getting priority details
	var getpriority = "SELECT c.id ,a.name from mstrecorddifferentiation a,mstbusinessmatrix b,mstrecorddifferentiation c where c.recorddifftypeid=? and c.id=b.mstrecorddifferentiationcatid and b.mstrecorddifferentiationpriorityid=a.id and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.clientid and a.mstorgnhirarchyid=b.mstorgnhirarchyid and a.clientid=c.clientid and a.mstorgnhirarchyid=c.mstorgnhirarchyid and  a.activeflg=1 and a.deleteflg=0 and  b.activeflg=1 and b.deleteflg=0 and c.activeflg=1 and c.deleteflg=0"
	priorityResult, err := db.Query(getpriority, lastRecorddifftypeid, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println(err)

		return impactNames, urgencyNames, priorityNames, err
	}
	for priorityResult.Next() {
		var priorityId int64
		var priorityName string
		//	var  diffTypeId int64
		err = priorityResult.Scan(&priorityId, &priorityName)
		if err != nil {
			Logger.Log.Println(err)

			return impactNames, urgencyNames, priorityNames, err
		}
		//headerName = append(headerName, header)
		priorityNames = append(priorityNames, priorityName)
		priorityIds = append(priorityIds, priorityId)
	}
	return impactNames, urgencyNames, priorityNames, err
}
func GetEstimatedtimesEfficiencies(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, lastRecorddifftypeid int64) ([]string, []string, error) {
	var estimatedtimes []string
	var efficiencies []string
	var getestimatedtimesEfficiencies = "SELECT a.recorddiffid,a.estimatedtime,a.efficiency FROM mapcategorywithestimatetime a where a.clientid=? and a.mstorgnhirarchyid=? and a.recorddiffid in (SELECT id FROM mstrecorddifferentiation b where b.recorddifftypeid =? and b.clientid=a.clientid and b.mstorgnhirarchyid=a.mstorgnhirarchyid and b.deleteflg=0 and b.activeflg=1) and a.activeflg=1 and a.deleteflg=0"
	timeEfficiencyResult, err := db.Query(getestimatedtimesEfficiencies, clientID, mstOrgnHirarchyId, lastRecorddifftypeid)
	if err != nil {
		Logger.Log.Println(err)

		return estimatedtimes, efficiencies, err
	}
	for timeEfficiencyResult.Next() {
		var recordDiffId int64
		var estimatedtime string
		var efficiencie string
		//	var  diffTypeId int64
		err = timeEfficiencyResult.Scan(&recordDiffId, &estimatedtime, &efficiencie)
		if err != nil {
			Logger.Log.Println(err)

			return estimatedtimes, efficiencies, err
		}
		//headerName = append(headerName, header)
		estimatedtimes = append(estimatedtimes, estimatedtime)
		efficiencies = append(efficiencies, efficiencie)
	}
	return estimatedtimes, efficiencies, err

}
*/
