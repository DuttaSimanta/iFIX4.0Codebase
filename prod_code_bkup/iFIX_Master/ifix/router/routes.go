//***************************//
// Package Name: Router
// Date Of Creation: 09/01/2021
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: Routes are defined here
// Functions: PostRequestHandler,main
// API: /addmstclient Inputs:  {"code": <string>,"name":<string>,"description":<string>,"ClientAuditFlg":<single character>}
// API: /getmstclientbyid Inputs:  {"id":<Number>}
// API: /delmstclientbyid Inputs:  {"id":<Number>}
// API: /updatemstclientbyid Inputs:  {"id":<Number>,code": <string>,"name":<string>,"description":<string>,"ClientAuditFlg":<single character>}
// API: /getmstclientall Inputs:  N/A
//
// API: /addmstcountry Inputs:  {"code": <string>,"name":<string>,"description":<string>,"ClientAuditFlg":<single character>}
// API: /getmstcountrybyid Inputs:  {"id":<Number>}
// API: /getmstcountryall Inputs:  {"id":<Number>}
// API: /delmstcountrybyid Inputs:  {"id":<Number>,code": <string>,"name":<string>,"description":<string>,"ClientAuditFlg":<single character>}
// API: /updatemstcountrybyid Inputs:  N/A
// Global Variable: N/A
// Version: 1.0.0
//***************************//

package router

import (
	"iFIX/ifix/handlers"
	"net/http"
)

//Route is a basic sturct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"404 Not Found",
		"POST",
		"/",
		ThrowBlankResponse,
	},
	Route{
		"addmstclient",
		"POST",
		"/addmstclient",
		handlers.AddMstClient,
	},
	Route{
		"getmstclientbyid",
		"POST",
		"/getmstclientbyid",
		handlers.GetMstClientByID,
	},
	Route{
		"getmstclientall",
		"POST",
		"/getmstclientall",
		handlers.GetMstClientAll,
	},
	Route{
		"delmstclientbyid",
		"POST",
		"/delmstclientbyid",
		handlers.DelMstClientByID,
	},
	Route{
		"updatemstclientbyid",
		"POST",
		"/updatemstclientbyid",
		handlers.UpdateMstClientByID,
	},
	Route{
		"addmstcountry",
		"POST",
		"/addmstcountry",
		handlers.AddMstCountry,
	},
	Route{
		"getmstcountrybyid",
		"POST",
		"/getmstcountrybyid",
		handlers.GetMstCountryByID,
	},
	Route{
		"getmstcountryall",
		"POST",
		"/getmstcountryall",
		handlers.GetMstCountryAll,
	},
	Route{
		"delmstcountrybyid",
		"POST",
		"/delmstcountrybyid",
		handlers.DelMstCountryByID,
	},
	Route{
		"updatemstcountrybyid",
		"POST",
		"/updatemstcountrybyid",
		handlers.UpdateMstCountryByID,
	}, Route{
		"insertProcessDelegateUser",
		"POST",
		"/insertProcessDelegateUser",
		handlers.InsertProcessDelegateUser,
	}, Route{
		"moveWorkflow",
		"POST",
		"/moveWorkflow",
		handlers.MoveWorkflow,
	}, Route{
		"insertModule",
		"POST",
		"/insertModule",
		handlers.InsertModule,
	}, Route{
		"GetAllModules",
		"POST",
		"/getAllModules",
		handlers.GetAllModules,
	}, Route{
		"deleteModule",
		"POST",
		"/deleteModule",
		handlers.DeleteModule,
	}, Route{
		"updateModule",
		"POST",
		"/updateModule",
		handlers.UpdateModule,
	}, Route{
		"insertUrl",
		"POST",
		"/insertUrl",
		handlers.InsertUrl,
	}, Route{
		"getAllUrls",
		"POST",
		"/getAllUrls",
		handlers.GetAllUrls,
	}, Route{
		"getAllModuleUrls",
		"POST",
		"/getAllModuleUrls",
		handlers.GetAllModuleUrls,
	}, Route{
		"deleteUrl",
		"POST",
		"/deleteUrl",
		handlers.DeleteUrl,
	}, Route{
		"deleteModuleUrl",
		"POST",
		"/deleteModuleUrl",
		handlers.DeleteModuleUrl,
	}, Route{
		"updateUrl",
		"POST",
		"/updateUrl",
		handlers.UpdateUrl,
	},
	Route{
		"createrole",
		"POST",
		"/createrole",
		handlers.AddRole,
	},
	Route{
		"deleterole",
		"POST",
		"/deleterole",
		handlers.DeleteRole,
	},
	Route{
		"updaterole",
		"POST",
		"/updaterole",
		handlers.UpdateRole,
	},
	Route{
		"getrole",
		"POST",
		"/getrole",
		handlers.GetAllRole,
	},
	Route{
		"createclient",
		"POST",
		"/createclient",
		handlers.AddClient,
	},
	Route{
		"deleteclient",
		"POST",
		"/deleteclient",
		handlers.DeleteClient,
	},
	Route{
		"updateclient",
		"POST",
		"/updateclient",
		handlers.UpdateClient,
	},
	Route{
		"getclient",
		"POST",
		"/getclient",
		handlers.GetAllClients,
	},
	Route{
		"createclientuser",
		"POST",
		"/createclientuser",
		handlers.AddClientUser,
	},
	Route{
		"deleteclientuser",
		"POST",
		"/deleteclientuser",
		handlers.DeleteClientUser,
	},
	Route{
		"updateclientuser",
		"POST",
		"/updateclientuser",
		handlers.UpdateClientUser,
	},
	Route{
		"getclientuser",
		"POST",
		"/getclientuser",
		handlers.GetAllClientUsers,
	}, Route{
		"login",
		"POST",
		"/login",
		handlers.Login,
	},
	Route{
		"getUserDetailsById",
		"POST",
		"/getUserDetailsById",
		handlers.GetUserDetailsById,
	},
	Route{
		"getRecordDiffType",
		"POST",
		"/getRecordDiffType",
		handlers.GetRecordDiffType,
	},
	Route{
		"insertRecordDiff",
		"POST",
		"/insertRecordDiff",
		handlers.InsertRecordDiff,
	},
	Route{
		"getAllRecordDiff",
		"POST",
		"/getAllRecordDiff",
		handlers.GetAllRecordDiff,
	},
	Route{
		"UpdateRecordDiff",
		"POST",
		"/updateRecordDiff",
		handlers.UpdateRecordDiff,
	},
	Route{
		"deleteRecordDiff",
		"POST",
		"/deleteRecordDiff",
		handlers.DeleteRecordDiff,
	},
	Route{
		"createclientuserrole",
		"POST",
		"/createclientuserrole",
		handlers.AddClientUserRole,
	},
	Route{
		"deleteclientuserrole",
		"POST",
		"/deleteclientuserrole",
		handlers.DeleteClientUserRole,
	},
	Route{
		"updateclientuserrole",
		"POST",
		"/updateclientuserrole",
		handlers.UpdateClientUserRole,
	},
	Route{
		"getclientuserrole",
		"POST",
		"/getclientuserrole",
		handlers.GetAllClientUserRole,
	}, Route{
		"getDistinctUrl",
		"POST",
		"/getDistinctUrl",
		handlers.GetDistinctUrl,
	}, Route{
		"GetRemainingUrl",
		"POST",
		"/getRemainingUrl",
		handlers.GetRemainingUrl,
	}, Route{
		"AddModuleClient",
		"POST",
		"/addModuleClient",
		handlers.AddModuleClient,
	},
	Route{
		"DeleteModuleClient",
		"POST",
		"/deleteModuleClient",
		handlers.DeleteModuleClient,
	}, Route{
		"UpdateModuleClient",
		"POST",
		"/updateModuleClient",
		handlers.UpdateModuleClient,
	}, Route{
		"GetAllModuleClients",
		"POST",
		"/getAllModuleClients",
		handlers.GetAllModuleClients,
	}, Route{
		"GetModuleByOrgId",
		"POST",
		"/getModuleByOrgId",
		handlers.GetModuleByOrgId,
	},
	//09.02.2021
	Route{
		"addorganization",
		"POST",
		"/addorganization",
		handlers.AddOrganization,
	},
	Route{
		"updateorganization",
		"POST",
		"/updateorganization",
		handlers.UpdateOrganization,
	},
	Route{
		"getorganization",
		"POST",
		"/getorganization",
		handlers.GetAllOrganization,
	},
	Route{
		"getorganizationclientwise",
		"POST",
		"/getorganizationclientwise",
		handlers.GetAllOrganizationClientWise,
	},
	Route{
		"SearchUser",
		"POST",
		"/searchUser",
		handlers.SearchUser,
	}, Route{
		"addroleaction",
		"POST",
		"/addroleaction",
		handlers.AddActionRole,
	},
	Route{
		"updateroleaction",
		"POST",
		"/updateroleaction",
		handlers.UpdateRoleAction,
	},
	Route{
		"getroleactionforclient",
		"POST",
		"/getroleactionforclient",
		handlers.GetAllRoleActionForClient,
	},
	Route{
		"deleteroleaction",
		"POST",
		"/deleteroleaction",
		handlers.DeleteRoleAction,
	},
	Route{
		"getroleaction",
		"POST",
		"/getroleaction",
		handlers.GetAllRoleAction,
	},
	Route{
		"getaction",
		"POST",
		"/getaction",
		handlers.GetAllAction,
	},
	Route{
		"SearchUserByOrgnId",
		"POST",
		"/searchUserByOrgnId",
		handlers.SearchUserByOrgnId,
	},
	Route{
		"GetRoleWiseAction",
		"POST",
		"/getRoleWiseAction",
		handlers.GetRoleWiseAction,
	}, //10.02.2021
	Route{
		"adduserroleaction",
		"POST",
		"/adduserroleaction",
		handlers.AddUserActionRole,
	},
	Route{
		"updateuserroleaction",
		"POST",
		"/updateuserroleaction",
		handlers.UpdateUserRoleAction,
	},
	Route{
		"getuserroleactionforclient",
		"POST",
		"/getuserroleactionforclient",
		handlers.GetAllUserRoleActionForClient,
	},
	Route{
		"deleteuserroleaction",
		"POST",
		"/deleteuserroleaction",
		handlers.DeleteUserRoleAction,
	},
	Route{
		"getuserroleaction",
		"POST",
		"/getuserroleaction",
		handlers.GetAllUserRoleAction,
	}, Route{
		"Getrolebyorgid",
		"POST",
		"/getrolebyorgid",
		handlers.Getrolebyorgid,
	}, Route{
		"Searchzone",
		"POST",
		"/searchzone",
		handlers.Searchzone,
	},
	Route{
		"GetRoleUserWiseAction",
		"POST",
		"/getRoleUserWiseAction",
		handlers.GetRoleUserWiseAction,
	}, Route{
		"addmodulerolemap",
		"POST",
		"/addmodulerolemap",
		handlers.InsertModulerolemap,
	},
	Route{
		"updatemodulerolemap",
		"POST",
		"/updatemodulerolemap",
		handlers.UpdateModulerolemap,
	},
	Route{
		"getmodulerolemap",
		"POST",
		"/getmodulerolemap",
		handlers.GetAllModulerolemap,
	},
	Route{
		"deletemodulerolemap",
		"POST",
		"/deletemodulerolemap",
		handlers.DeleteModulerolemap,
	},
	Route{
		"InsertMenu",
		"POST",
		"/insertmenu",
		handlers.InsertMenu,
	},
	Route{
		"Getparentmenu",
		"POST",
		"/getparentmenu",
		handlers.Getparentmenu,
	},
	Route{
		"Getmenudetails",
		"POST",
		"/getmenudetails",
		handlers.Getmenudetails,
	}, Route{
		"addmodulerolemapuser",
		"POST",
		"/addmodulerolemapuser",
		handlers.InsertClientmoduleurlroleuser,
	},
	Route{
		"updatemodulerolemapuser",
		"POST",
		"/updatemodulerolemapuser",
		handlers.UpdateClientmoduleurlroleuser,
	},
	Route{
		"getmodulerolemapuser",
		"POST",
		"/getmodulerolemapuser",
		handlers.GetAllClientmoduleurlroleuser,
	},
	Route{
		"deletemodulerolemapuser",
		"POST",
		"/deletemodulerolemapuser",
		handlers.DeleteClientmoduleurlroleuser,
	}, Route{
		"addnonurl",
		"POST",
		"/addnonurl",
		handlers.InsertNonmenuurl,
	},
	Route{
		"updatenonurl",
		"POST",
		"/updatenonurl",
		handlers.UpdateNonmenuurl,
	},
	Route{
		"getnonurl",
		"POST",
		"/getnonurl",
		handlers.GetAllNonmenuurl,
	},
	Route{
		"deletenonurl",
		"POST",
		"/deletenonurl",
		handlers.DeleteNonmenuurl,
	},
	Route{
		"geturlkey",
		"POST",
		"/geturlkey",
		handlers.GetAllUrlkey,
	}, Route{
		"UpdateMenu",
		"POST",
		"/updatemenu",
		handlers.UpdateMenu,
	},
	Route{
		"DeleteMenu",
		"POST",
		"/deletemenu",
		handlers.DeleteMenu,
	}, Route{
		"Geturlmenudetails",
		"POST",
		"/geturlmenudetails",
		handlers.Geturlmenudetails,
	}, Route{
		"addrecordconfig",
		"POST",
		"/addrecordconfig",
		handlers.InsertRecordconfig,
	},
	Route{
		"updaterecordconfig",
		"POST",
		"/updaterecordconfig",
		handlers.UpdateRecordconfig,
	},
	Route{
		"getrecordconfig",
		"POST",
		"/getrecordconfig",
		handlers.GetAllRecordconfig,
	},
	Route{
		"deleterecordconfig",
		"POST",
		"/deleterecordconfig",
		handlers.DeleteRecordconfig,
	},
	Route{
		"GetRecordByDiffType",
		"POST",
		"/getrecordbydifftype",
		handlers.GetRecordByDiffType,
	}, Route{
		"addlabel",
		"POST",
		"/addlabel",
		handlers.InsertRecorddifferentiationtype,
	},
	Route{
		"updatelabel",
		"POST",
		"/updatelabel",
		handlers.UpdateRecorddifferentiationtype,
	},
	Route{
		"getlabel",
		"POST",
		"/getlabel",
		handlers.GetAllRecorddifferentiationtype,
	},
	Route{
		"getalllabel",
		"POST",
		"/getalllabel",
		handlers.GetRecorddifferentiationtype,
	},
	Route{
		"deletelabel",
		"POST",
		"/deletelabel",
		handlers.DeleteRecorddifferentiationtype,
	}, Route{
		"addrecordcategory",
		"POST",
		"/addrecordcategory",
		handlers.InsertRecorddifferentiation,
	},
	Route{
		"updaterecordcategory",
		"POST",
		"/updaterecordcategory",
		handlers.UpdateRecorddifferentiation,
	},
	Route{
		"getrecordcategory",
		"POST",
		"/getrecordcategory",
		handlers.GetAllRecorddifferentiation,
	},
	Route{
		"deleterecordcategory",
		"POST",
		"/deleterecordcategory",
		handlers.DeleteRecorddifferentiation,
	},
	Route{
		"addworingkdiff",
		"POST",
		"/addworingkdiff",
		handlers.InsertWorkdifferentiation,
	},
	Route{
		"updateworingkdiff",
		"POST",
		"/updateworingkdiff",
		handlers.UpdateWorkdifferentiation,
	},
	Route{
		"getworkingdiff",
		"POST",
		"/getworkingdiff",
		handlers.GetAllWorkdifferentiation,
	},
	Route{
		"deleteworkingdiff",
		"POST",
		"/deleteworkingdiff",
		handlers.DeleteWorkdifferentiation,
	}, Route{
		"addrecordtypemap",
		"POST",
		"/addrecordtypemap",
		handlers.InsertRecordtype,
	},
	Route{
		"updaterecordtypemap",
		"POST",
		"/updaterecordtypemap",
		handlers.UpdateRecordtype,
	},
	Route{
		"getrecordtypemap",
		"POST",
		"/getrecordtypemap",
		handlers.GetAllRecordtype,
	},
	Route{
		"deleterecordtypemap",
		"POST",
		"/deleterecordtypemap",
		handlers.DeleteRecordtype,
	}, Route{
		"getdifferentiationname",
		"POST",
		"/getdifferentiationname",
		handlers.GetRecorddifferentiationname,
	}, Route{
		"GetCategoryLevel",
		"POST",
		"/getcategorylevel",
		handlers.GetCategoryLevel,
	}, Route{
		"GetAllRecordDiffTypeByClient",
		"POST",
		"/getallrecorddifftypebyclient",
		handlers.GetAllRecordDiffTypeByClient,
	}, Route{
		"GetRecordDiff",
		"POST",
		"/getrecorddiff",
		handlers.GetRecordDiff,
	}, Route{
		"getfunctionality",
		"POST",
		"/getfunctionality",
		handlers.Getfunctionality,
	}, Route{
		"addcatelog",
		"POST",
		"/addcatelog",
		handlers.InsertCatalog,
	},
	Route{
		"updatecatelog",
		"POST",
		"/updatecatelog",
		handlers.UpdateCatalog,
	},
	Route{
		"getcatelog",
		"POST",
		"/getcatelog",
		handlers.GetAllCatalog,
	},
	Route{
		"deletecatelog",
		"POST",
		"/deletecatelog",
		handlers.DeleteCatalog,
	},
	Route{
		"updateerror",
		"POST",
		"/updateerror",
		handlers.UpdateErrormesg,
	},
	Route{
		"geterror",
		"POST",
		"/geterror",
		handlers.GetAllErrormesg,
	},
	Route{
		"Insertfuncmapping",
		"POST",
		"/insertfuncmapping",
		handlers.Insertfuncmapping,
	}, Route{
		"Getfuncmappingbytype",
		"POST",
		"/getfuncmappingbytype",
		handlers.Getfuncmappingbytype,
	}, Route{
		"Getfuncmappingbycatalogtype",
		"POST",
		"/getfuncmappingbycatalogtype",
		handlers.Getfuncmappingbycatalogtype,
	}, Route{
		"addcatelogmap",
		"POST",
		"/addcatelogmap",
		handlers.InsertCatalogwithcategory,
	},
	Route{
		"updatecatelogmap",
		"POST",
		"/updatecatelogmap",
		handlers.UpdateCatalogwithcategory,
	},
	Route{
		"getcatelogmap",
		"POST",
		"/getcatelogmap",
		handlers.GetAllCatalogwithcategory,
	},
	Route{
		"deletecatelogmap",
		"POST",
		"/deletecatelogmap",
		handlers.DeleteCatalogwithcategory,
	}, Route{
		"getfuncmappingdetails",
		"POST",
		"/getfuncmappingdetails",
		handlers.Getfuncmappingdetails,
	}, Route{
		"Deletefunctionmapping",
		"POST",
		"/deletefunctionmapping",
		handlers.Deletefunctionmapping,
	}, Route{
		"addhigherkey",
		"POST",
		"/addhigherkey",
		handlers.InsertRecorddifferentiationhigherkeyEntity,
	},
	Route{
		"updatehigherkey",
		"POST",
		"/updatehigherkey",
		handlers.UpdateRecorddifferentiationhigherkeyEntity,
	},
	Route{
		"gethigherkey",
		"POST",
		"/gethigherkey",
		handlers.GetAllRecorddifferentiationhigherkeyEntity,
	},
	Route{
		"deletehigherkey",
		"POST",
		"/deletehigherkey",
		handlers.DeleteRecorddifferentiationhigherkeyEntity,
	}, Route{
		"UserMenu",
		"POST",
		"/getmenubyuser",
		handlers.GetMenuByUser,
	}, Route{
		"Getmenubymodule",
		"POST",
		"/getmenubymodule",
		handlers.Getmenubymodule,
	}, Route{
		"getuserwiseaction",
		"POST",
		"/getuserwiseaction",
		handlers.Getuserwiseaction,
	}, Route{
		"adddayofweek",
		"POST",
		"/adddayofweek",
		handlers.InsertClientdayofweek,
	},
	Route{
		"getdayofweek",
		"POST",
		"/getdayofweek",
		handlers.GetAllClientdayofweek,
	},
	Route{
		"deletedayofweek",
		"POST",
		"/deletedayofweek",
		handlers.DeleteClientdayofweek,
	}, Route{
		"adddashboardquery",
		"POST",
		"/adddashboardquery",
		handlers.InsertDashboarddtls,
	},
	Route{
		"getdashboardquery",
		"POST",
		"/getdashboardquery",
		handlers.GetAllDashboarddtls,
	},
	Route{
		"deletedashboardquery",
		"POST",
		"/deletedashboardquery",
		handlers.DeleteDashboarddtls,
	}, Route{
		"deletedashboardquery",
		"POST",
		"/deleteurlfrommenu",
		handlers.DeleteUrlFromMenu,
	}, Route{
		"addsupportgrp",
		"POST",
		"/addsupportgrp",
		handlers.InsertClientsupportgroup,
	},
	Route{
		"getsupportgrp",
		"POST",
		"/getsupportgrp",
		handlers.GetAllClientsupportgroup,
	},
	Route{
		"deletesupportgrp",
		"POST",
		"/deletesupportgrp",
		handlers.DeleteClientsupportgroup,
	},
	Route{
		"updatesupportgrp",
		"POST",
		"/updatesupportgrp",
		handlers.UpdateClientsupportgroup,
	}, Route{
		"addclientsupportgrpholiday",
		"POST",
		"/addclientsupportgrpholiday",
		handlers.InsertClientsupportgroupholiday,
	},
	Route{
		"getclientsupportgrpholiday",
		"POST",
		"/getclientsupportgrpholiday",
		handlers.GetAllClientsupportgroupholiday,
	},
	Route{
		"updateclientsupportgrpholiday",
		"POST",
		"/updateclientsupportgrpholiday",
		handlers.UpdateClientsupportgroupholiday,
	},
	Route{
		"deleteclientsupportgrpholiday",
		"POST",
		"/deleteclientsupportgrpholiday",
		handlers.DeleteClientsupportgroupholiday,
	},

	Route{
		"getsupportgrpname",
		"POST",
		"/getsupportgrpname",
		handlers.GetAllSupportgrpname,
	}, Route{
		"getcities",
		"GET",
		"/getcities",
		handlers.GetAllCity,
	}, Route{
		"getcountries",
		"GET",
		"/getcountries",
		handlers.GetAllCountry,
	},
	Route{
		"addgrpusermap",
		"POST",
		"/addgrpusermap",
		handlers.InsertGroupmember,
	},
	Route{
		"getgrpusermap",
		"POST",
		"/getgrpusermap",
		handlers.GetAllGroupmember,
	},
	Route{
		"updategrpusermap",
		"POST",
		"/updategrpusermap",
		handlers.UpdateGroupmember,
	},
	Route{
		"deletegrpusermap",
		"POST",
		"/deletegrpusermap",
		handlers.DeleteGroupmember,
	},
	Route{
		"addclientholiday",
		"POST",
		"/addclientholiday",
		handlers.InsertClientholiday,
	},
	Route{
		"getclientholiday",
		"POST",
		"/getclientholiday",
		handlers.GetAllClientholiday,
	},
	Route{
		"updateclientholiday",
		"POST",
		"/updateclientholiday",
		handlers.UpdateClientholiday,
	},
	Route{
		"deleteclientholiday",
		"POST",
		"/deleteclientholiday",
		handlers.DeleteClientholiday,
	}, Route{
		"Getgroupbyorgid",
		"POST",
		"/getgroupbyorgid",
		handlers.Getgroupbyorgid,
	}, Route{
		"getsupportgrplevel",
		"GET",
		"/getsupportgrplevel",
		handlers.GetAllPortgrouplevel,
	}, Route{
		"addasset",
		"POST",
		"/addasset",
		handlers.InsertAsset,
	},
	Route{
		"getasset",
		"POST",
		"/getasset",
		handlers.GetAllAsset,
	},
	Route{
		"updateasset",
		"POST",
		"/updateasset",
		handlers.UpdateAsset,
	},
	Route{
		"deleteasset",
		"POST",
		"/deleteasset",
		handlers.DeleteAsset,
	},
	Route{
		"addassetvalidation",
		"POST",
		"/addassetvalidation",
		handlers.InsertAssetvalidate,
	},
	Route{
		"getassetvalidation",
		"POST",
		"/getassetvalidation",
		handlers.GetAllAssetvalidate,
	},
	Route{
		"updateassetvalidation",
		"POST",
		"/updateassetvalidation",
		handlers.UpdateAssetvalidate,
	},
	Route{
		"deleteassetvalidation",
		"POST",
		"/deleteassetvalidation",
		handlers.DeleteAssetvalidate,
	},
	Route{
		"addassetdifferentiation",
		"POST",
		"/addassetdifferentiation",
		handlers.InsertAssetdifferentiation,
	},
	Route{
		"getassetdifferentiation",
		"POST",
		"/getassetdifferentiation",
		handlers.GetAllAssetdifferentiation,
	},
	Route{
		"updateassetdifferentiation",
		"POST",
		"/updateassetdifferentiation",
		handlers.UpdateAssetdifferentiation,
	},
	Route{
		"deleteassetdifferentiation",
		"POST",
		"/deleteassetdifferentiation",
		handlers.DeleteAssetdifferentiation,
	}, Route{
		"getworkinglevel",
		"POST",
		"/getworkinglevel",
		handlers.GetWorkinglevel,
	},

	Route{
		"addcategorysupporgrpmap",
		"POST",
		"/addcategorysupporgrpmap",
		handlers.InsertRecorddifferentiongroup,
	},

	Route{
		"getcategorysupporgrpmap",
		"POST",
		"/getcategorysupporgrpmap",
		handlers.GetAllRecorddifferentiongroup,
	},

	Route{
		"updatecategorysupporgrpmap",
		"POST",
		"/updatecategorysupporgrpmap",
		handlers.UpdateRecorddifferentiongroup,
	},
	Route{
		"deletecategorysupporgrpmap",
		"POST",
		"/deletecategorysupporgrpmap",
		handlers.DeleteRecorddifferentiongroup,
	},
	Route{
		"GetAllCategoryLevel",
		"POST",
		"/getallcategorylevel",
		handlers.GetAllCategoryLevel,
	}, Route{
		"GetRecordDiffByOrg",
		"POST",
		"/getrecorddiffbyorg",
		handlers.GetRecordDiffByOrg,
	}, Route{
		"getassetbytype",
		"POST",
		"/getassetbytype",
		handlers.GetAssetBYType,
	}, Route{
		"getassetdiffval",
		"POST",
		"/getassetdiffval",
		handlers.GetAssetDiffVal,
	}, Route{
		"updateassetdiffval",
		"POST",
		"/updateassetdiffval",
		handlers.UpdateAssetDiffVal,
	}, Route{
		"Getworkdifferentiationvalue",
		"POST",
		"/getworkdifferentiationvalue",
		handlers.Getworkdifferentiationvalue,
	}, Route{
		"addmapfunctionalitygrp",
		"POST",
		"/addmapfunctionalitygrp",
		handlers.InsertMapfunctionalitywithgroup,
	},
	Route{
		"deletemapfunctionalitygrp",
		"POST",
		"/deletemapfunctionalitygrp",
		handlers.DeleteMapfunctionalitywithgroup,
	},

	Route{
		"updatemapfunctionalitygrp",
		"POST",
		"/updatemapfunctionalitygrp",
		handlers.UpdateMapfunctionalitywithgroup,
	},

	Route{
		"getmapfunctionalitygrp",
		"POST",
		"/getmapfunctionalitygrp",
		handlers.GetAllMapfunctionalitywithgroup,
	},

	Route{
		"getorgnameclientwise",
		"POST",
		"/getorgnameclientwise",
		handlers.GetAllOrganizationgrpnames,
	},
	Route{
		"addmstrecordterms",
		"POST",
		"/addmstrecordterms",
		handlers.InsertMstrecordterms,
	}, Route{
		"getmstrecordterms",
		"POST",
		"/getmstrecordterms",
		handlers.GetAllMstrecordterms,
	}, Route{
		"listmstrecordterms",
		"POST",
		"/listmstrecordterms",
		handlers.GetListMstrecordterms,
	}, Route{
		"updatemstrecordterms",
		"POST",
		"/updatemstrecordterms",
		handlers.UpdateMstrecordterms,
	}, Route{
		"deletemstrecordterms",
		"POST",
		"/deletemstrecordterms",
		handlers.DeleteMstrecordterms,
	}, Route{
		"addmststateterms",
		"POST",
		"/addmststateterms",
		handlers.InsertMststateterm,
	}, Route{
		"getmststateterms",
		"POST",
		"/getmststateterms",
		handlers.GetAllMststateterm,
	}, Route{
		"updatemststateterms",
		"POST",
		"/updatemststateterms",
		handlers.UpdateMststateterm,
	}, Route{
		"deletemststateterms",
		"POST",
		"/deletemststateterms",
		handlers.DeleteMststateterm,
	}, Route{
		"getmsttermtype",
		"POST",
		"/getmsttermtype",
		handlers.GetAllMsttermtype,
	}, Route{
		"Getworklowutilitylist",
		"POST",
		"/getworklowutilitylist",
		handlers.Getworklowutilitylist,
	}, Route{
		"addprocessstatemap",
		"POST",
		"/addprocessstatemap",
		handlers.InsertMapprocessstate,
	},
	Route{
		"updateprocessstatemap",
		"POST",
		"/updateprocessstatemap",
		handlers.UpdateMapprocessstate,
	},
	Route{
		"deleteprocessstatemap",
		"POST",
		"/deleteprocessstatemap",
		handlers.DeleteMapprocessstate,
	},
	Route{
		"getprocessstatemap",
		"POST",
		"/getprocessstatemap",
		handlers.GetAllMapprocessstate,
	}, Route{
		"addstate",
		"POST",
		"/addstate",
		handlers.InsertMststate,
	},
	Route{
		"updatestate",
		"POST",
		"/updatestate",
		handlers.UpdateMststate,
	},
	Route{
		"deletestate",
		"POST",
		"/deletestate",
		handlers.DeleteMststate,
	},
	Route{
		"getstate",
		"POST",
		"/getstate",
		handlers.GetAllMststate,
	},
	Route{
		"addstatetype",
		"POST",
		"/addstatetype",
		handlers.InsertMststatetype,
	},

	Route{
		"updatestatetype",
		"POST",
		"/updatestatetype",
		handlers.UpdateMststatetype,
	},

	Route{
		"deletestatetype",
		"POST",
		"/deletestatetype",
		handlers.DeleteMststatetype,
	},

	Route{
		"getstatetype",
		"POST",
		"/getstatetype",
		handlers.GetAllMststatetype,
	},
	Route{
		"createprocess",
		"POST",
		"/createprocess",
		handlers.InsertMstprocess,
	},

	Route{
		"deleteprocess",
		"POST",
		"/deleteprocess",
		handlers.DeleteMstprocess,
	},

	Route{
		"updateprocess",
		"POST",
		"/updateprocess",
		handlers.UpdateMstprocess,
	},

	Route{
		"getprocess",
		"POST",
		"/getprocess",
		handlers.GetAllMstprocess,
	},
	Route{
		"addprocessadmin",
		"POST",
		"/addprocessadmin",
		handlers.InsertMstprocessadmin,
	},

	Route{
		"updateprocessadmin",
		"POST",
		"/updateprocessadmin",
		handlers.UpdateMstprocessadmin,
	},
	Route{
		"deleteprocessadmin",
		"POST",
		"/deleteprocessadmin",
		handlers.DeleteMstprocessadmin,
	},
	Route{
		"getprocessadmin",
		"POST",
		"/getprocessadmin",
		handlers.GetAllMstprocessadmin,
	}, Route{
		"searchworkflowuser",
		"POST",
		"/searchworkflowuser",
		handlers.Searchworkflowuser,
	}, Route{
		"getutilitydatabyfield",
		"POST",
		"/getutilitydatabyfield",
		handlers.Getutilitydatabyfield,
	}, Route{
		"addmaprecordstatedifferentiation",
		"POST",
		"/addmaprecordstatedifferentiation",
		handlers.InsertMaprecordstatetodifferentiation,
	},
	Route{
		"updatemaprecordstatedifferentiation",
		"POST",
		"/updatemaprecordstatedifferentiation",
		handlers.UpdateMaprecordstatetodifferentiation,
	},
	Route{
		"deletemaprecordstatedifferentiation",
		"POST",
		"/deletemaprecordstatedifferentiation",
		handlers.DeleteMaprecordstatetodifferentiation,
	},
	Route{
		"getmaprecordstatedifferentiation",
		"POST",
		"/getmaprecordstatedifferentiation",
		handlers.GetAllMaprecordstatetodifferentiation,
	}, Route{
		"fileupload",
		"POST",
		"/fileupload",
		handlers.UploadFile,
	}, Route{
		"addassigncommontiles",
		"POST",
		"/addassigncommontiles",
		handlers.InsertMapcommontileswithgroup,
	},

	Route{
		"updateassigncommontiles",
		"POST",
		"/updateassigncommontiles",
		handlers.UpdateMapcommontileswithgroup,
	},

	Route{
		"deleteassigncommontiles",
		"POST",
		"/deleteassigncommontiles",
		handlers.DeleteMapcommontileswithgroup,
	},

	Route{
		"getassigncommontiles",
		"POST",
		"/getassigncommontiles",
		handlers.GetAllMapcommontileswithgroup,
	},
	Route{
		"adddocuments",
		"POST",
		"/adddocuments",
		handlers.InsertMstdocumentdtls,
	},

	Route{
		"getdocuments",
		"POST",
		"/getdocuments",
		handlers.GetAllMstdocumentdtls,
	},

	Route{
		"deletedocuments",
		"POST",
		"/deletedocuments",
		handlers.DeleteMstdocumentdtls,
	},

	Route{
		"updatedocuments",
		"POST",
		"/updatedocuments",
		handlers.UpdateMstdocumentdtls,
	},
	Route{
		"addmstrecordfield",
		"POST",
		"/addmstrecordfield",
		handlers.InsertMstrecordfield,
	}, Route{
		"getmstrecordfield",
		"POST",
		"/getmstrecordfield",
		handlers.GetAllMstrecordfield,
	}, Route{
		"updatemstrecordfield",
		"POST",
		"/updatemstrecordfield",
		handlers.UpdateMstrecordfield,
	}, Route{
		"deletemstrecordfield",
		"POST",
		"/deletemstrecordfield",
		handlers.DeleteMstrecordfield,
	}, Route{
		"addquestionanswer",
		"POST",
		"/addquestionanswer",
		handlers.InsertQuestionanswer,
	},

	Route{
		"getquestionanswer",
		"POST",
		"/getquestionanswer",
		handlers.GetAllquestionsanswer,
	},

	Route{
		"deletequestionanswer",
		"POST",
		"/deletequestionanswer",
		handlers.Deletequestionsanswer,
	}, Route{
		"GetWorkinglabelname",
		"POST",
		"/getworkinglabelname",
		handlers.GetWorkinglabelname,
	}, Route{
		"addtemplatevariable",
		"POST",
		"/addtemplatevariable",
		handlers.InsertMsttemplatevariable,
	},

	Route{
		"gettemplatevariable",
		"POST",
		"/gettemplatevariable",
		handlers.GetAllMsttemplatevariable,
	},

	Route{
		"deletetemplatevariable",
		"POST",
		"/deletetemplatevariable",
		handlers.DeleteMsttemplatevariable,
	},

	Route{
		"updatetemplatevariable",
		"POST",
		"/updatetemplatevariable",
		handlers.UpdateMsttemplatevariable,
	},
	Route{
		"addmstclientsla",
		"POST",
		"/addmstclientsla",
		handlers.InsertMstclientsla,
	},

	Route{
		"getmstclientsla",
		"POST",
		"/getmstclientsla",
		handlers.GetAllMstclientsla,
	},

	Route{
		"deletemstclientsla",
		"POST",
		"/deletemstclientsla",
		handlers.DeleteMstclientsla,
	},

	Route{
		"updatemstclientsla",
		"POST",
		"/updatemstclientsla",
		handlers.UpdateMstclientsla,
	},

	Route{
		"getslanames",
		"POST",
		"/getslanames",
		handlers.GetSlanames,
	},

	Route{
		"addslastate",
		"POST",
		"/addslastate",
		handlers.InsertMstslastate,
	},

	Route{
		"getslastate",
		"POST",
		"/getslastate",
		handlers.GetAllMstslastate,
	},

	Route{
		"deletslastate",
		"POST",
		"/deleteslastate",
		handlers.DeleteMstslastate,
	},

	Route{
		"updateslastate",
		"POST",
		"/updateslastate",
		handlers.UpdateMstslastate,
	},

	Route{
		"addslatimezone",
		"POST",
		"/addslatimezone",
		handlers.InsertMstslatimezone,
	},

	Route{
		"getslatimezone",
		"POST",
		"/getslatimezone",
		handlers.GetAllMstslatimezone,
	},

	Route{
		"deleteslatimezone",
		"POST",
		"/deleteslatimezone",
		handlers.DeleteMstslatimezone,
	},
	Route{
		"updateslatimezone",
		"POST",
		"/updateslatimezone",
		handlers.UpdateMstslatimezone,
	}, Route{
		"Getstatebyprocess",
		"POST",
		"/getstatebyprocess",
		handlers.Getstatebyprocess,
	}, Route{
		"addmstslaentity",
		"POST",
		"/addmstslaentity",
		handlers.InsertMstslaentity,
	},
	Route{
		"getmstslaentity",
		"POST",
		"/getmstslaentity",
		handlers.GetAllMstslaentity,
	},
	Route{
		"deletemstslaentity",
		"POST",
		"/deletemstslaentity",
		handlers.DeleteMstslaentity,
	},
	Route{
		"updatemstslaentity",
		"POST",
		"/updatemstslaentity",
		handlers.UpdateMstslaentity,
	}, Route{
		"addmsttemplate",
		"POST",
		"/addmsttemplate",
		handlers.InsertMsttemplate,
	}, Route{
		"gettemplate",
		"POST",
		"/gettemplate",
		handlers.GetAllMsttemplate,
	}, Route{
		"updatetemplate",
		"POST",
		"/updatetemplate",
		handlers.UpdateMsttemplate,
	}, Route{
		"deletetemplate",
		"POST",
		"/deletetemplate",
		handlers.DeleteMsttemplate,
	}, Route{
		"addmstslacriteria",
		"POST",
		"/addmstslacriteria",
		handlers.InsertMstslafullfillmentcriteria,
	},
	Route{
		"getslacriteria",
		"POST",
		"/getslacriteria",
		handlers.GetAllMstslafullfillmentcriteria,
	},
	Route{
		"deleteslacriteria",
		"POST",
		"/deleteslacriteria",
		handlers.DeleteMstslafullfillmentcriteria,
	},
	Route{
		"updateslacriteria",
		"POST",
		"/updateslacriteria",
		handlers.UpdateMstslafullfillmentcriteria,
	}, Route{
		"SearchUserByGroupId",
		"POST",
		"/searchuserbygroupid",
		handlers.SearchUserByGroupId,
	}, Route{
		"Getprocessbydiffid",
		"POST",
		"/getprocessbydiffid",
		handlers.Getprocessbydiffid,
	}, Route{
		"Getprocessbydiffid",
		"POST",
		"/insertprocess",
		handlers.Insertprocess,
	}, Route{
		"Getprocessbydiffid",
		"POST",
		"/getprocessdetails",
		handlers.Getprocessdetails,
	}, Route{
		"Gettransitionstatedetails",
		"POST",
		"/gettransitionstatedetails",
		handlers.Gettransitionstatedetails,
	}, Route{
		"addslapauseindicator",
		"POST",
		"/addslapauseindicator",
		handlers.InsertMstslafcrecorddiff,
	},
	Route{
		"getslapauseindicator",
		"POST",
		"/getslapauseindicator",
		handlers.GetAllMstslafcrecorddiff,
	},
	Route{
		"deleteslapauseindicator",
		"POST",
		"/deleteslapauseindicator",
		handlers.DeleteMstslafcrecorddiff,
	},
	Route{
		"updateslapauseindicator",
		"POST",
		"/updateslapauseindicator",
		handlers.UpdateMstslafcrecorddiff,
	}, Route{
		"addbusinessdirection",
		"POST",
		"/addbusinessdirection",
		handlers.InsertMstbusinessdirection,
	},
	Route{
		"getbusinessdirection",
		"POST",
		"/getbusinessdirection",
		handlers.GetAllMstbusinessdirection,
	},

	Route{
		"deletebusinessdirection",
		"POST",
		"/deletebusinessdirection",
		handlers.DeleteMstbusinessdirection,
	},

	Route{
		"updatebusinessdirection",
		"POST",
		"/updatebusinessdirection",
		handlers.UpdateMstbusinessdirection,
	}, Route{
		"addbusinessmatrix",
		"POST",
		"/addbusinessmatrix",
		handlers.InsertMstbusinessmatrix,
	},

	Route{
		"getbusinessmatrix",
		"POST",
		"/getbusinessmatrix",
		handlers.GetAllMstbusinessmatrix,
	},

	Route{
		"deletebusinessmatrix",
		"POST",
		"/deletebusinessmatrix",
		handlers.DeleteMstbusinessmatrix,
	},

	Route{
		"updatebusinessmatrix",
		"POST",
		"/updatebusinessmatrix",
		handlers.UpdateMstbusinessmatrix,
	},

	Route{
		"checkmatrixconfig",
		"POST",
		"/checkmatrixconfig",
		handlers.Checkbusinessmatrixconfig,
	},
	Route{
		"getlastlevelcatname",
		"POST",
		"/getlastlevelcatname",
		handlers.Getlastlevelcategoryname,
	}, Route{
		"checkprocessdelete",
		"POST",
		"/checkprocessdelete",
		handlers.Checkprocessdelete,
	}, Route{
		"addslaresponsesupportgrp",
		"POST",
		"/addslaresponsesupportgrp",
		handlers.InsertMstslaresponsiblesupportgroup,
	},

	Route{
		"getslaresponsesupportgrp",
		"POST",
		"/getslaresponsesupportgrp",
		handlers.GetAllMstslaresponsiblesupportgroup,
	},

	Route{
		"deleteslaresponsesupportgrp",
		"POST",
		"/deleteslaresponsesupportgrp",
		handlers.DeleteMstslaresponsiblesupportgroup,
	},

	Route{
		"updateslaresponsesupportgrp",
		"POST",
		"/updateslaresponsesupportgrp",
		handlers.UpdateMstslaresponsiblesupportgroup,
	},

	Route{
		"getsupportgrpenableslanames",
		"POST",
		"/getsupportgrpenableslanames",
		handlers.GetAllSlanames,
	},

	Route{
		"getfullfillmentcriteriaid",
		"POST",
		"/getfullfillmentcriteriaid",
		handlers.GetFullfillmentcriteriaid,
	},
	Route{
		"addcategorytaskmap",
		"POST",
		"/addcategorytaskmap",
		handlers.InsertMstcategorytaskmap,
	},

	Route{
		"getcategorytaskmap",
		"POST",
		"/getcategorytaskmap",
		handlers.GetAllMstcategorytaskmap,
	},
	Route{
		"deletecategorytaskmap",
		"POST",
		"/deletecategorytaskmap",
		handlers.DeleteMstcategorytaskmap,
	},
	Route{
		"updatecategorytaskmap",
		"POST",
		"/updatecategorytaskmap",
		handlers.UpdateMstcategorytaskmap,
	}, Route{
		"Createtransition",
		"POST",
		"/createtransition",
		handlers.Createtransition,
	}, Route{
		"Upserttransitiondetails",
		"POST",
		"/upserttransitiondetails",
		handlers.Upserttransitiondetails,
	}, Route{
		"Getlabelbydiffid",
		"POST",
		"/getlabelbydiffid",
		handlers.Getlabelbydiffid,
	}, Route{
		"Deletetransitionstate",
		"POST",
		"/deletetransitionstate",
		handlers.Deletetransitionstate,
	}, Route{
		"addslastartendcriteria",
		"POST",
		"/addslastartendcriteria",
		handlers.InsertMstslastartendcriteria,
	}, Route{
		"getslastartendcriteria",
		"POST",
		"/getslastartendcriteria",
		handlers.GetAllMstslastartendcriteria,
	}, Route{
		"deleteslastartendcriteria",
		"POST",
		"/deleteslastartendcriteria",
		handlers.DeleteMstslastartendcriteria,
	}, Route{
		"updateslastartendcriteria",
		"POST",
		"/updateslastartendcriteria",
		handlers.UpdateMstslastartendcriteria,
	}, Route{
		"getslanameagainstworkflowid",
		"POST",
		"/getslanameagainstworkflowid",
		handlers.GetSlanameagainstworkflowid,
	},
	Route{
		"getclietwiseasset",
		"POST",
		"/getclietwiseasset",
		handlers.GetClietWiseAsset,
	}, Route{
		"getassettypes",
		"POST",
		"/getassettypes",
		handlers.GetAssetTypes,
	}, Route{
		"getassetattributes",
		"POST",
		"/getassetattributes",
		handlers.GetAssetAttributes,
	}, Route{
		"getassetbytypenvalue",
		"POST",
		"/getassetbytypenvalue",
		handlers.GetAssetByTypeNAtrrValue,
	}, Route{
		"addactivity",
		"POST",
		"/addactivity",
		handlers.InsertMstactivity,
	}, Route{
		"getactivity",
		"POST",
		"/getactivity",
		handlers.GetAllMstactivity,
	}, Route{
		"updateactivity",
		"POST",
		"/updateactivity",
		handlers.UpdateMstactivity,
	}, Route{
		"deleteactivity",
		"POST",
		"/deleteactivity",
		handlers.DeleteMstactivity,
	}, Route{
		"getactiontypenames",
		"POST",
		"/getactiontypenames",
		handlers.GetActiontypenames,
	}, Route{
		"Getstatedetails",
		"POST",
		"/getstatedetails",
		handlers.Getstatedetails,
	}, Route{
		"Getnextstatedetails",
		"POST",
		"/getnextstatedetails",
		handlers.Getnextstatedetails,
	}, Route{
		"Checkworkflow",
		"POST",
		"/checkworkflow",
		handlers.Checkworkflow,
	}, Route{
		"Gettransitiongroupdetails",
		"POST",
		"/gettransitiongroupdetails",
		handlers.Gettransitiongroupdetails,
	}, Route{
		"Changerecordgroup",
		"POST",
		"/changerecordgroup",
		handlers.Changerecordgroup,
	},
	Route{
		"gettilesnames",
		"POST",
		"/gettilesnames",
		handlers.GetTilesnames,
	}, Route{
		"Getlablelmappingbydifftype",
		"POST",
		"/getlablelmappingbydifftype",
		handlers.Getlablelmappingbydifftype,
	}, Route{
		"Getactivitywithtype",
		"POST",
		"/getactivitywithtype",
		handlers.Getactivitywithtype,
	}, Route{
		"Gettransitionbyprocess",
		"POST",
		"/gettransitionbyprocess",
		handlers.Gettransitionbyprocess,
	}, {
		"recordwiseuserinfo",
		"POST",
		"/recordwiseuserinfo",
		handlers.GetRecordwiseuserinfo,
	}, {
		"clientwisedayofweek",
		"POST",
		"/clientwisedayofweek",
		handlers.GetClientwisedayofweek,
	}, Route{
		"useridwiseuserinfo",
		"POST",
		"/useridwiseuserinfo",
		handlers.GetIDwiseuserinfo,
	}, Route{
		"getSLAmeternames",
		"GET",
		"/getSLAmeternames",
		handlers.GetSLAmetertype,
	},
	Route{
		"getSLAtermsnames",
		"POST",
		"/getSLAtermsnames",
		handlers.GetSLAtermnames,
	}, Route{
		"Getcategorybycatalog",
		"POST",
		"/getcategorybycatalog",
		handlers.Getcategorybycatalog,
	}, Route{
		"GetRecorddifferentiationbyrecursive",
		"POST",
		"/getrecorddifferentiationbyrecursive",
		handlers.GetRecorddifferentiationbyrecursive,
	}, Route{
		"Deleteprocessdetails",
		"POST",
		"/deleteprocessdetails",
		handlers.Deleteprocessdetails,
	}, Route{
		"addrecordreleationwithterm",
		"POST",
		"/addrecordreleationwithterm",
		handlers.InsertMstrcordtremswisereleationconfig,
	},
	Route{
		"getrecordreleationwithterm",
		"POST",
		"/getrecordreleationwithterm",
		handlers.GetAllMstrcordtremswisereleationconfig,
	},
	Route{
		"updaterecordreleationwithterm",
		"POST",
		"/updaterecordreleationwithterm",
		handlers.UpdateMstrcordtremswisereleationconfig,
	},
	Route{
		"deleterecordreleationwithterm",
		"POST",
		"/deleterecordreleationwithterm",
		handlers.DeleteMstrcordtremswisereleationconfig,
	},
	Route{
		"getrecordreleationnames",
		"POST",
		"/getrecordreleationnames",
		handlers.GetRecordreleationnames,
	},
	Route{
		"getrecordtermnames",
		"POST",
		"/getrecordtermnames",
		handlers.GetRecordtermnames,
	}, Route{
		"GetRecorddifferentiationbyparent",
		"POST",
		"/getRecorddifferentiationbyparent",
		handlers.GetRecorddifferentiationbyparent,
	}, Route{
		"Searchcategory",
		"POST",
		"/searchcategory",
		handlers.Searchcategory,
	}, Route{
		"Getcatelogrecord",
		"POST",
		"/getcatelogrecord",
		handlers.Getcatelogrecord,
	}, Route{
		"addmstsupportgrp",
		"POST",
		"/addmstsupportgrp",
		handlers.InsertMstsupportgrptermmap,
	},
	Route{
		"getmstsupportgrp",
		"POST",
		"/getmstsupportgrp",
		handlers.GetAllMstsupportgrptermmap,
	},
	Route{
		"deletemstsupportgrp",
		"POST",
		"/deletemstsupportgrp",
		handlers.DeleteMstsupportgrptermmap,
	},
	Route{
		"updatemstsupportgrp",
		"POST",
		"/updatemstsupportgrp",
		handlers.UpdateMstsupportgrptermmap,
	}, Route{
		"Updatechildstatus",
		"POST",
		"/updatechildstatus",
		handlers.Updatechildstatus,
	}, Route{
		"Gethopcount",
		"POST",
		"/gethopcount",
		handlers.Gethopcount,
	}, Route{
		"Getstatebyseq",
		"POST",
		"/getstatebyseq",
		handlers.Getstatebyseq,
	}, Route{
		"Changepassword",
		"POST",
		"/changepassword",
		handlers.Changepassword,
	}, Route{
		"Updateusercolor",
		"POST",
		"/updateusercolor",
		handlers.Updateusercolor,
	}, Route{
		"Updateusercolor",
		"POST",
		"/generatetoken",
		handlers.Generatetoken,
	}, Route{
		"Updateusercolor",
		"POST",
		"/validateusertoken",
		handlers.Validateusertoken,
	}, {
		"addbanner",
		"POST",
		"/addbanner",
		handlers.InsertBanner,
	},
	Route{
		"getbanner",
		"POST",
		"/getbanner",
		handlers.GetAllBanner,
	},
	Route{
		"updatebanner",
		"POST",
		"/updatebanner",
		handlers.UpdateBanner,
	},
	Route{
		"deletebanner",
		"POST",
		"/deletebanner",
		handlers.DeleteBanner,
	}, {
		"getbannermessage",
		"POST",
		"/getbannermessage",
		handlers.GetAllMessage,
	}, {
		"updatebannersequence",
		"POST",
		"/updatebannersequence",
		handlers.UpdateBannerSequence,
	}, Route{
		"searchloginname",
		"POST",
		"/searchloginname",
		handlers.SearchLoginName,
	},
	Route{
		"searchname",
		"POST",
		"/searchname",
		handlers.SearchName,
	},
	Route{
		"searchbranch",
		"POST",
		"/searchbranch",
		handlers.SearchBranch,
	}, Route{
		"Inpersonatelogin",
		"POST",
		"/inpersonatelogin",
		handlers.Inpersonatelogin,
	}, Route{
		"searchloginnamebygroupids",
		"POST",
		"/searchloginnamebygroupids",
		handlers.SearchLoginamebyGroupids,
	}, Route{
		"getnotificationevents",
		"POST",
		"/getnotificationevents",
		handlers.GetNotificationEvents,
	}, Route{
		"insertnotificationtemplate",
		"POST",
		"/insertnotificationtemplate",
		handlers.InsertNotificationTemplate,
	}, Route{
		"getallnotificationtemplates",
		"POST",
		"/getallnotificationtemplates",
		handlers.GetAllNotificationTemplates,
	}, Route{
		"updatenotificationtemplate",
		"POST",
		"/updatenotificationtemplate",
		handlers.UpdateNotificationTemplate,
	}, Route{
		"deletenotificationtemplate",
		"POST",
		"/deletenotificationtemplate",
		handlers.DeleteNotificationTemplate,
	}, Route{
		"Gettimeformat",
		"GET",
		"/gettimeformat",
		handlers.Gettimeformat,
	}, Route{
		"getlogintype",
		"GET",
		"/getlogintype",
		handlers.GetLogintype,
	}, {
		"getallnotificationvariables",
		"POST",
		"/getallnotificationvariables",
		handlers.GetAllNotificationVariables,
	}, Route{
		"updatemstldapcertificate",
		"POST",
		"/updatemstldapcertificate",
		handlers.UpdateMstldapCertificate,
	},
	Route{
		"updatemstldap",
		"POST",
		"/updatemstldap",
		handlers.UpdateMstldap,
	},
	Route{
		"deletemstldap",
		"POST",
		"/deletemstldap",
		handlers.DeleteMstldap,
	},
	Route{
		"getallmstldap",
		"POST",
		"/getallmstldap",
		handlers.GetAllMstldap,
	},
	Route{
		"addmstldap",
		"POST",
		"/addmstldap",
		handlers.AddMstldap,
	}, {
		"updatemapldapgrouprole",
		"POST",
		"/updatemapldapgrouprole",
		handlers.Updatemapldapgrouprole,
	},
	Route{
		"deletemapldapgrouprole",
		"POST",
		"/deletemapldapgrouprole",
		handlers.Deletemapldapgrouprole,
	},
	Route{
		"getallmapldapgrouprole",
		"POST",
		"/getallmapldapgrouprole",
		handlers.GetAllmapldapgrouprole,
	},
	Route{
		"addmapldapgrouprole",
		"POST",
		"/addmapldapgrouprole",
		handlers.Insertmapldapgrouprole,
	}, Route{
		"Getldapattributes",
		"POST",
		"/getldapattributes",
		handlers.Getldapattributes,
	}, Route{
		"Gettabledetails",
		"POST",
		"/gettabledetails",
		handlers.Gettabledetails,
	}, Route{
		"updatemstsupportgroup",
		"POST",
		"/updatemstsupportgroup",
		handlers.Updatemstsupportgrp,
	},

	Route{
		"getmstsupportgroup",
		"POST",
		"/getmstsupportgroup",
		handlers.GetAllmstsupportgrp,
	},
	Route{
		"insertmstsupportgrp",
		"POST",
		"/insertmstsupportgrp",
		handlers.Insertmstsupportgrp,
	}, Route{
		"deletemapexternalattributes",
		"POST",
		"/deletemapexternalattributes",
		handlers.DeleteMapexternalattributes,
	},
	Route{
		"getAllmapexternalattributes",
		"POST",
		"/getAllmapexternalattributes",
		handlers.GetAllMapexternalattributes,
	},
	Route{
		"insertmapexternalattributes",
		"POST",
		"/insertmapexternalattributes",
		handlers.InsertMapexternalattributes,
	}, Route{
		"GetMappedattributes",
		"POST",
		"/getmappedattributes",
		handlers.GetMappedattributes,
	},
	Route{
		"deletemstsupportgroup",
		"POST",
		"/deletemstsupportgroup",
		handlers.Deletemstsupportgrp,
	},

	Route{
		"filedownload",
		"POST",
		"/filedownload",
		handlers.DownloadFile,
	},

	Route{
		"getmstsupportgroupbycopyable",
		"POST",
		"/getmstsupportgroupbycopyable",
		handlers.GetAllmstsupportgrpbycopyable,
	},

	Route{
		"updateclientsupportgroupnew",
		"POST",
		"/updateclientsupportgroupnew",
		handlers.UpdateClientsupportgroupnew,
	},
	Route{
		"deleteclientsupportgroupnew",
		"POST",
		"/deleteclientsupportgroupnew",
		handlers.DeleteClientsupportgroupnew,
	},
	Route{
		"getallclientsupportgroupnew",
		"POST",
		"/getallclientsupportgroupnew",
		handlers.GetAllClientsupportgroupnew,
	},
	Route{
		"addclientsupportgroupnew",
		"POST",
		"/addclientsupportgroupnew",
		handlers.InsertClientsupportgroupnew,
	},

	//22.07.2021

	Route{
		"createdifferentiationmap",
		"POST",
		"/createdifferentiationmap",
		handlers.Createdifferentiationmap,
	},
	Route{
		"getalldifferentiationmap",
		"POST",
		"/getalldifferentiationmap",
		handlers.GetAllDifferentiationDtls,
	},

	Route{
		"deletedifferentiationmap",
		"POST",
		"/deletedifferentiationmap",
		handlers.Deletedifferentiationmap,
	},
	Route{
		"getAllclientsupportgroupbyclient",
		"POST",
		"/getAllclientsupportgroupbyclient",
		handlers.GetAllClientsupportgroupbyclient,
	},
	Route{
		"insertclientsupportgroupfromto",
		"POST",
		"/insertclientsupportgroupfromto",
		handlers.InsertClientsupportgroupfromto,
	},

	Route{
		"getallgrpmember",
		"POST",
		"/getallgrpmember",
		handlers.GetAllGrpmember,
	},
	Route{
		"addgroupmember",
		"POST",
		"/addgroupmember",
		handlers.AddGroupmember,
	}, Route{
		"InsertMstprocesstemplate",
		"POST",
		"/insertmstprocesstemplate",
		handlers.InsertMstprocesstemplate,
	}, Route{
		"GetAllMstprocesstemplate",
		"POST",
		"/getallmstprocesstemplate",
		handlers.GetAllMstprocesstemplate,
	}, Route{
		"DeleteMstprocesstemplate",
		"POST",
		"/deletemstprocesstemplate",
		handlers.DeleteMstprocesstemplate,
	}, Route{
		"UpdateMstprocesstemplate",
		"POST",
		"/updatemstprocesstemplate",
		handlers.UpdateMstprocesstemplate,
	}, Route{
		"recordtermscopy",
		"POST",
		"/recordtermscopy",
		handlers.Createrecordtermsmap,
	}, Route{
		"InsertMapprocesstemplatestate",
		"POST",
		"/insertmapprocesstemplatestate",
		handlers.InsertMapprocesstemplatestate,
	}, Route{
		"DeleteMapprocesstemplatestate",
		"POST",
		"/deletemapprocesstemplatestate",
		handlers.DeleteMapprocesstemplatestate,
	}, Route{
		"GetAllMapprocesstemplatestate",
		"POST",
		"/getallmapprocesstemplatestate",
		handlers.GetAllMapprocesstemplatestate,
	}, Route{
		"UpdateMapprocesstemplatestate",
		"POST",
		"/updatemapprocesstemplatestate",
		handlers.UpdateMapprocesstemplatestate,
	}, Route{
		"Getstatebyprocesstemplate",
		"POST",
		"/getstatebyprocesstemplate",
		handlers.Getstatebyprocesstemplate,
	}, Route{
		"Getprocesstemplatedetails",
		"POST",
		"/getprocesstemplatedetails",
		handlers.Getprocesstemplatedetails,
	}, Route{
		"Createprocesstemplatetransition",
		"POST",
		"/createprocesstemplatetransition",
		handlers.Createprocesstemplatetransition,
	}, Route{
		"insertprocesstemplate",
		"POST",
		"/insertprocesstemplate",
		handlers.Insertprocesstemplate,
	}, Route{
		"Deletetemplatetransitionstate",
		"POST",
		"/deletetemplatetransitionstate",
		handlers.Deletetemplatetransitionstate,
	}, Route{
		"Upserttemplatetransitiondetails",
		"POST",
		"/upserttemplatetransitiondetails",
		handlers.Upserttemplatetransitiondetails,
	}, Route{
		"Gettemplatetransitionstatedetails",
		"POST",
		"/gettemplatetransitionstatedetails",
		handlers.Gettemplatetransitionstatedetails,
	}, Route{
		"Deleteprocesstemplatedetails",
		"POST",
		"/deleteprocesstemplatedetails",
		handlers.Deleteprocesstemplatedetails,
	}, Route{
		"Getprocesstemplate",
		"POST",
		"/getprocesstemplate",
		handlers.Getprocesstemplate,
	}, Route{
		"Mapprocesstemplate",
		"POST",
		"/mapprocesstemplate",
		handlers.Mapprocesstemplate,
	},
	Route{
		"getorganizationclientwisenew",
		"POST",
		"/getorganizationclientwisenew",
		handlers.GetAllOrganizationClientWisenew,
	}, Route{
		"checkworkflowstate",
		"POST",
		"/checkworkflowstate",
		handlers.Checkworkflowstate,
	}, Route{
		"checkworkflowstate",
		"POST",
		"/getprocessgroupbyorgid",
		handlers.Getprocessgroupbyorgid,
	}, Route{
		"addtaskmap",
		"POST",
		"/addtaskmap",
		handlers.InsertMaptask,
	},
	Route{
		"gettaskmap",
		"POST",
		"/gettaskmap",
		handlers.GetAllMaptask,
	},
	Route{
		"deletetaskmap",
		"POST",
		"/deletetaskmap",
		handlers.DeleteMaptask,
	}, Route{
		"getdiffdetailsbyseq",
		"POST",
		"/getdiffdetailsbyseq",
		handlers.Getdiffdetailsbyseq,
	}, Route{
		"Updatetaskstatus",
		"POST",
		"/updatetaskstatus",
		handlers.Updatetaskstatus,
	}, Route{
		"updaterecorddifftypeandrecordtype",
		"POST",
		"/updaterecorddifftypeandrecordtype",
		handlers.UpdateRecorddifftypeAndRecordtype,
	},
	Route{
		"deleterecorddifftypeandrecordtype",
		"POST",
		"/deleterecorddifftypeandrecordtype",
		handlers.DeleteRecorddifftypeAndRecordtype,
	},
	Route{
		"addrecorddifftypeandrecordtype",
		"POST",
		"/addrecorddifftypeandrecordtype",
		handlers.InsertRecorddifftypeAndRecordtype,
	},
	Route{
		"updatemstrecorddiffpriority",
		"POST",
		"/updatemstrecorddiffpriority",
		handlers.UpdateMstRecorddiffpriority,
	},
	Route{
		"deleteMstRecorddiffpriority",
		"POST",
		"/deletemstrecorddiffpriority",
		handlers.DeleteMstRecorddiffpriority,
	},
	Route{
		"getallmstrecorddiffpriority",
		"POST",
		"/getallmstrecorddiffpriority",
		handlers.GetAllMstRecorddiffpriority,
	},
	Route{
		"addmstrecorddiffpriority",
		"POST",
		"/addmstrecorddiffpriority",
		handlers.AddMstRecorddiffpriority,
	}, Route{
		"getcategorieslevel",
		"POST",
		"/getcategorieslevel",
		handlers.GetCategoriesLevel,
	}, Route{
		"gettabbuttonnames",
		"POST",
		"/gettabbuttonnames",
		handlers.GetTabsButtonnames,
	}, Route{
		"Getcategorybyparentname",
		"POST",
		"/getcategorybyparentname",
		handlers.Getcategorybyparentname,
	}, Route{
		"Getfromtypebydiffname",
		"POST",
		"/getfromtypebydiffname",
		handlers.Getfromtypebydiffname,
	}, Route{
		"Getmappeddiffbyseq",
		"POST",
		"/getmappeddiffbyseq",
		handlers.Getmappeddiffbyseq,
	}, Route{
		"GetUserByGroupId",
		"POST",
		"/GetUserByGroupId",
		handlers.GetUserByGroupId,
	}, Route{
		"updatemstexceltemplate",
		"POST",
		"/updatemstexceltemplate",
		handlers.UpdateMstExcelTemplate,
	},
	Route{
		"deletemstexceltemplate",
		"POST",
		"/deletemstexceltemplate",
		handlers.DeleteMstExcelTemplate,
	},
	Route{
		"getallmstexceltemplate",
		"POST",
		"/getallmstexceltemplate",
		handlers.GetAllMstExcelTemplate,
	},
	Route{
		"addmstexceltemplate",
		"POST",
		"/addmstexceltemplate",
		handlers.AddMstExcelTemplate,
	}, Route{
		"getallmstexceltemplatetype",
		"POST",
		"/getallmstexceltemplatetype",
		handlers.GetAllMstExcelTemplateType,
	}, Route{
		"getallclientnames",
		"POST",
		"/getallclientnames",
		handlers.GetAllClientsnames,
	}, Route{
		"getlabelbydiffseq",
		"POST",
		"/getlabelbydiffseq",
		handlers.Getlabelbydiffseq,
	}, Route{
		"SearchMenuByUser",
		"POST",
		"/searchmenubyuser",
		handlers.SearchMenuByUser,
	}, Route{
		"updatemstclientcredential",
		"POST",
		"/updatemstclientcredential",
		handlers.UpdateMstClientCredential,
	},
	Route{
		"deletemstclientcredential",
		"POST",
		"/deletemstclientcredential",
		handlers.DeleteMstClientCredential,
	},
	Route{
		"getallmstclientcredential",
		"POST",
		"/getallmstclientcredential",
		handlers.GetAllMstClientCredential,
	},
	Route{
		"insertmstclientcredential",
		"POST",
		"/insertmstclientcredential",
		handlers.InsertMstClientCredential,
	}, Route{
		"getallmstclientcredentialtype",
		"POST",
		"/getallmstclientcredentialtype",
		handlers.GetAllMstClientCredentialType,
	},
	Route{
		"updatemsttemplatevariable",
		"POST",
		"/updatemsttemplatevariable",
		handlers.UpdateMstTemplateVariable,
	},
	Route{
		"deletemsttemplatevariable",
		"POST",
		"/deletemsttemplatevariable",
		handlers.DeleteMstTemplateVariable,
	},
	Route{
		"getallmsttemplatevariable",
		"POST",
		"/getallmsttemplatevariable",
		handlers.GetAllMstTemplateVariable,
	},
	Route{
		"addmsttemplatevariable",
		"POST",
		"/addmsttemplatevariable",
		handlers.AddMstTemplateVariable,
	}, Route{
		"updatemstschedulednotification",
		"POST",
		"/updatemstschedulednotification",
		handlers.UpdateMstScheduledNotification,
	},
	Route{
		"deletemstschedulednotification",
		"POST",
		"/deletemstschedulednotification",
		handlers.DeleteMstScheduledNotification,
	},
	Route{
		"getmstschedulednotification",
		"POST",
		"/getmstschedulednotification",
		handlers.GetMstScheduledNotification,
	},
	Route{
		"addmstschedulednotification",
		"POST",
		"/addmstschedulednotification",
		handlers.AddMstScheduledNotification,
	}, Route{
		"getclientandorgwiseclientuser",
		"POST",
		"/getclientandorgwiseclientuser",
		handlers.GetClientAndOrgWiseclientuser,
	},
	Route{
		"downloadifixdatainexcel",
		"POST",
		"/downloadifixdatainexcel",
		handlers.JsonToExcelConverter,
	}, Route{
		"downloadgridresult",
		"POST",
		"/downloadgridresult",
		handlers.GridResultToExcelConverter,
	}, Route{
		"getorgassignedcustomer",
		"POST",
		"/getorgassignedcustomer",
		handlers.GetAllTicketCustomer,
	}, Route{
		"Searchuserdetailsbygroupid",
		"POST",
		"/searchuserdetailsbygroupid",
		handlers.Searchuserdetailsbygroupid,
	}, Route{
		"getadditionaltab",
		"GET",
		"/getadditionaltab",
		handlers.GetAdditionalTab,
	},
	Route{
		"deleterecordtermadditionalmap",
		"POST",
		"/deleterecordtermadditionalmap",
		handlers.DeleteRecordTermAdditionalMap,
	},
	Route{
		"getrecordtermadditionalmap",
		"POST",
		"/getrecordtermadditionalmap",
		handlers.GetAllRecordTermAdditionalMap,
	},
	Route{
		"addrecordtermadditionalmap",
		"POST",
		"/addrecordtermadditionalmap",
		handlers.InsertRecordTermAdditionalMap,
	}, Route{
		"verifyToTP",
		"POST",
		"/verifytotp",
		handlers.VerifyTOTP,
	}, Route{
		"updaterecordconfigincrement",
		"POST",
		"/updaterecordconfigincrement",
		handlers.UpdateRecordConfigIncrement,
	},
	Route{
		"deleterecordconfigincrement",
		"POST",
		"/deleterecordconfigincrement",
		handlers.DeleteRecordConfigIncrement,
	},
	Route{
		"getallrecordconfigincrement",
		"POST",
		"/getallrecordconfigincrement",
		handlers.GetAllRecordConfigIncrement,
	},
	Route{
		"addrecordconfigincrement",
		"POST",
		"/addrecordconfigincrement",
		handlers.InsertRecordConfigIncrement,
	}, Route{
		"deletedashboardquerycopy",
		"POST",
		"/deletedashboardquerycopy",
		handlers.DeleteDashboardQueryCopy,
	},
	Route{
		"getalldashboardquerycopy",
		"POST",
		"/getalldashboardquerycopy",
		handlers.GetAllDashboardQueryCopy,
	},
	Route{
		"adddashboardquerycopy",
		"POST",
		"/adddashboardquerycopy",
		handlers.AddDashboardQueryCopy,
	}, Route{
		"updatedashboardquery",
		"POST",
		"/updatedashboardquery",
		handlers.UpdateDashboardQuery,
	},
	Route{
		"insertdashboardquery",
		"POST",
		"/insertdashboardquery",
		handlers.AddDashboardQuery,
	}, Route{
		"deleteslatermentry",
		"POST",
		"/deleteslatermentry",
		handlers.DeleteSlaTermEntry,
	},
	Route{
		"getallslatermentry",
		"POST",
		"/getallslatermentry",
		handlers.GetAllSlaTermEntry,
	},
	Route{
		"addslatermentry",
		"POST",
		"/addslatermentry",
		handlers.AddSlaTermEntry,
	},
	Route{
		"updatemapcategorywithkeyword",
		"POST",
		"/updatemapcategorywithkeyword",
		handlers.UpdateMapCategoryWithKeyword,
	},
	Route{
		"deletemapcategorywithkeyword",
		"POST",
		"/deletemapcategorywithkeyword",
		handlers.DeleteMapCategoryWithKeyword,
	},
	Route{
		"getallmapcategorywithkeyword",
		"POST",
		"/getallmapcategorywithkeyword",
		handlers.GetAllMapCategoryWithKeyword,
	},
	Route{
		"insertmapcategorywithkeyword",
		"POST",
		"/insertmapcategorywithkeyword",
		handlers.InsertMapCategoryWithKeyword,
	}, Route{
		"updateuidgen",
		"POST",
		"/updateuidgen",
		handlers.UpdateUidGen,
	},
	Route{
		"deleteuidgen",
		"POST",
		"/deleteuidgen",
		handlers.DeleteUidGen,
	},
	Route{
		"getalluidgen",
		"POST",
		"/getalluidgen",
		handlers.GetAllUidGen,
	},
	Route{
		"adduidgen",
		"POST",
		"/adduidgen",
		handlers.InsertUidGen,
	}, Route{
		"getsupportgroupbyorg",
		"POST",
		"/getsupportgroupbyorg",
		handlers.Getsupportgroupbyorg,
	}, Route{
		"Getprocessgroupbyorgids",
		"POST",
		"/getprocessgroupbyorgids",
		handlers.Getprocessgroupbyorgids,
	}, Route{
		"groupbyuserwise",
		"POST",
		"/groupbyuserwise",
		handlers.Groupbyuserwise,
	}, Route{
		"Searchuserbyclientid",
		"POST",
		"/searchuserbyclientid",
		handlers.Searchuserbyclientid,
	}, Route{
		"addmstrecordactivitycopy",
		"POST",
		"/addmstrecordactivitycopy",
		handlers.AddMstRecordActivityCopy,
	},
	Route{
		"updatemstrecordactivity",
		"POST",
		"/updatemstrecordactivity",
		handlers.UpdateMstRecordActivity,
	},
	Route{
		"deletemstrecordactivity",
		"POST",
		"/deletemstrecordactivity",
		handlers.DeleteMstRecordActivity,
	},
	Route{
		"getallmstrecordactivity",
		"POST",
		"/getallmstrecordactivity",
		handlers.GetAllMstRecordActivity,
	},
	Route{
		"addmstrecordactivity",
		"POST",
		"/addmstrecordactivity",
		handlers.AddMstRecordActivity,
	}, Route{
		"getorgwiseactivitydesc",
		"POST",
		"/getorgwiseactivitydesc",
		handlers.GetOrgWiseActivitydesc,
	}, Route{
		"getorgwiseactivitydesc",
		"POST",
		"/workflowgroupbyuserwise",
		handlers.Workflowgroupbyuserwise,
	}, Route{
		"Getallcreatedsupportgrp",
		"POST",
		"/getallcreatedsupportgrp",
		handlers.Getallcreatedsupportgrp,
	}, Route{
		"Geturlbykey",
		"POST",
		"/geturlbykey",
		handlers.Geturlbykey,
	}, Route{
		"Getcatalogorgwise",
		"POST",
		"/getcatalogorgwise",
		handlers.Getcatalogorgwise,
	}, Route{
		"getfuncmappingbytypeforquery",
		"POST",
		"/getfuncmappingbytypeforquery",
		handlers.Getfuncmappingbytypeforquery,
	}, Route{
		"getorgtoolscode",
		"POST",
		"/getorgtoolscode",
		handlers.GetAllMstorgcode,
	}, Route{
		"addorgtoolscode",
		"POST",
		"/addorgtoolscode",
		handlers.InsertMstorgcode,
	}, Route{
		"updateorgtoolscode",
		"POST",
		"/updateorgtoolscode",
		handlers.UpdateMstorgcode,
	},
	Route{
		"deleteorgtoolscode",
		"POST",
		"/deleteorgtoolscode",
		handlers.DeleteMstorgcode,
	},
	Route{
		"deletemstuserdefaultsupportgroup",
		"POST",
		"/deletemstuserdefaultsupportgroup",
		handlers.DeleteMstUserDefaultSupportGroup,
	},
	Route{
		"updatemstuserdefaultsupportgroup",
		"POST",
		"/updatemstuserdefaultsupportgroup",
		handlers.UpdateMstUserDefaultSupportGroup,
	},
	Route{
		"getallmstuserdefaultsupportgroup",
		"POST",
		"/getallmstuserdefaultsupportgroup",
		handlers.GetAllMstUserDefaultSupportGroup,
	},
	Route{
		"insertmstuserdefaultsupportgroup",
		"POST",
		"/insertmstuserdefaultsupportgroup",
		handlers.InsertMstUserDefaultSupportGroup,
	}, Route{
		"getallcategoryvalue",
		"POST",
		"/getallcategoryvalue",
		handlers.GetAllCategoryvalue,
	}, Route{
		"Getorgname",
		"POST",
		"/getorgname",
		handlers.Getorgname,
	}, Route{
		"mstusersupportgroupchange",
		"POST",
		"/mstusersupportgroupchange",
		handlers.MstUserSupportGroupChange,
	}, Route{
		"addmsttemplatevariablecopy",
		"POST",
		"/addmsttemplatevariablecopy",
		handlers.AddMstTemplateVariablecopy,
	},
	Route{
		"gettemplatevariablelist",
		"POST",
		"/gettemplatevariablelist",
		handlers.GetAllMstTemplateVariableList,
	},
	Route{
		"deletesupportgrpworkhours",
		"POST",
		"/deletesupportgrpworkhours",
		handlers.DeleteMstSupportGroupWorkingHours,
	},
	Route{
		"updatesupportgrpworkhours",
		"POST",
		"/updatesupportgrpworkhours",
		handlers.UpdateMstSupportGroupWorkingHours,
	},
	Route{
		"getsupportgrpworkhours",
		"POST",
		"/getsupportgrpworkhours",
		handlers.GetAllMstSupportGroupWorkingHoursk,
	},
	Route{
		"insertsupportgrpworkhours",
		"POST",
		"/insertsupportgrpworkhours",
		handlers.InsertMstSupportGroupWorkingHours,
	},
	Route{
		"UpdateAssetRecordDiff",
		"POST",
		"/updateassetrecorddiff",
		handlers.UpdateAssetRecordDiff,
	},

	Route{
		"GetAssetRecordDiffByOrg",
		"POST",
		"/getassetrecorddiffbyorg",
		handlers.GetAssetRecordDiffByOrg,
	}, Route{
		"gettable",
		"POST",
		"/gettable",
		handlers.Gettable,
	},
	Route{
		"gettypedescription",
		"POST",
		"/gettypedescription",
		handlers.Gettypedescription,
	},
	Route{
		"updatetransporttable",
		"POST",
		"/updatetransporttable",
		handlers.UpdateTransporttable,
	},
	Route{
		"deletetransporttable",
		"POST",
		"/deletetransporttable",
		handlers.DeleteTransporttable,
	},
	Route{
		"getalltransporttable",
		"POST",
		"/getalltransporttable",
		handlers.GetAllTransporttable,
	},
	Route{
		"inserttransporttable",
		"POST",
		"/inserttransporttable",
		handlers.InsertTransporttable,
	},
	Route{
		"gettypefortransport",
		"POST",
		"/gettypefortransport",
		handlers.Gettypefortransport,
	},
	Route{
		"bulkusergroupcategorydownload",
		"POST",
		"/bulkusergroupcategorydownload",
		handlers.BulkUserWithGroupAndCategoryDownload,
	},
	Route{
		"bulkusergroupcategoryupload",
		"POST",
		"/bulkusergroupcategoryupload",
		handlers.UserWithGroupAndCategoryUpload,
	},

	Route{
		"deleteusergroupcategory",
		"POST",
		"/deleteusergroupcategory",
		handlers.DeleteUserWithGroupAndCategory,
	},
	Route{
		"getusergroupcategory",
		"POST",
		"/getusergroupcategory",
		handlers.GetAllUserWithGroupAndCategory,
	},
	Route{
		"addusergroupcategory",
		"POST",
		"/addusergroupcategory",
		handlers.InsertUserWithGroupAndCategory,
	}, Route{
		"updateusergroupcategory",
		"POST",
		"/updateusergroupcategory",
		handlers.UpdateUserWithGroupAndCategory,
	}, Route{
		"gettoolscode",
		"POST",
		"/gettoolscode",
		handlers.GetAlltoolvalue,
	}, Route{
		"getorgcode",
		"POST",
		"/getorgcode",
		handlers.GetAllorgvalue,
	}, Route{
		"getorganizationwithorgtype",
		"POST",
		"/getorganizationwithorgtype",
		handlers.GetAllOrganizationwithOrgtype,
	},
	Route{
		"getlastcategorylist",
		"POST",
		"/getlastcategorylist",
		handlers.GetLastCategoryList,
	},
	Route{
		"getdelimiter",
		"POST",
		"/getdelemiter",
		handlers.GetDelimiter,
	},
	Route{
		"listAllData",
		"POST",
		"/getserviceuser",
		handlers.GetServiceUser,
	},
	Route{
		"insertemailticket",
		"POST",
		"/saveemailticketconfiguration",
		handlers.SaveEmailTicketConfiguration,
	},
	Route{
		"listAllData",
		"POST",
		"/getemailticketconfigurations",
		handlers.GetEmailTicketConfigurations,
	},

	Route{
		"DeleteRow",
		"POST",
		"/deleteemailticketconfiguration",
		handlers.DeleteEmailTicketConfiguration,
	},
	Route{
		"Updatecoloumn",
		"POST",
		"/updateemailticketconfiguration",
		handlers.UpdateEmailTicketConfiguration,
	},
	Route{
		"addemailbaseconfig",
		"POST",
		"/addemailbaseconfig",
		handlers.AddEmailBaseConfig,
	},
	Route{
		"getdelimiterforallclient",
		"POST",
		"/getdelimiterforallclient",
		handlers.GetDelimiterForAllClient,
	},
	Route{
		"deleteemailbaseconfig",
		"POST",
		"/deleteemailbaseconfig",
		handlers.DeleteEmailTicketConfigu,
	}, Route{
		"getrecordnames",
		"POST",
		"/getrecordnames",
		handlers.Getrecordname,
	}, Route{
		"recordfulldetailsdownload",
		"POST",
		"/recordfulldetailsdownload",
		handlers.RecordfulldetailsToExcelConverter,
	},
	Route{
		"getrecordfulldetails",
		"POST",
		"/getrecordfulldetails",
		handlers.Getrecordfulldetails,
	},
	Route{
		"getclientwiseattribute",
		"POST",
		"/getclientwiseattribute",
		handlers.GetMstAttribute,
	},
	Route{
		"addmstadfsattribute",
		"POST",
		"/addmstadfsattribute",
		handlers.AddMstAttribute,
	},
	Route{
		"getallmstadfsattribute",
		"POST",
		"/getallmstadfsattribute",
		handlers.GetAllMstAttribute,
	},
	Route{
		"updatemstadfsattribute",
		"POST",
		"/updatemstadfsattribute",
		handlers.UpdateMstAttribute,
	},
	Route{
		"deletemstadfsattribute",
		"POST",
		"/deletemstadfsattribute",
		handlers.DeleteMstAttribute,
	}, Route{
		"adfslogin",
		"POST",
		"/adfslogin",
		handlers.Adfslogin,
	}, Route{
		"getrecordbydifftypeofmultiorg",
		"POST",
		"/getrecordbydifftypeofmultiorg",
		handlers.GetRecordByDiffTypeOfMultiOrg,
	}, Route{
		"Insertuserticket",
		"POST",
		"/insertuserticket",
		handlers.Insertuserticket,
	}, Route{
		"Deleteuserticket",
		"POST",
		"/deleteuserticket",
		handlers.Deleteuserticket,
	}, Route{
		"getallopenticket",
		"POST",
		"/getallopenticket",
		handlers.GetOpenTicket,
	}, Route{
		"deleteopenticket",
		"POST",
		"/deleteopenticket",
		handlers.DeleteOpenTicket,
	}, Route{
		"addrepoertdownloadlist",
		"POST",
		"/addrepoertdownloadlist",
		handlers.ReportDownloadList,
	}, Route{
		"detachchildticket",
		"POST",
		"/detachchildticket",
		handlers.Detachchildticket,
	}, Route{
		"searchanalystorgwise",
		"POST",
		"/searchanalystorgwise",
		handlers.SearchAnalystOrgWise,
	}, Route{
		"bulkapprovalfortickets",
		"POST",
		"/bulkapprovalfortickets",
		handlers.Bulkapprovalfortickets,
	}, Route{
		"logout",
		"POST",
		"/logout",
		handlers.Logout,
	}, Route{
		"getUserPropertyName",
		"POST",
		"/getUserPropertyName",
		handlers.GetUserPropertyName,
	}, Route{
		"insertUserRoleProperty",
		"POST",
		"/insertUserRoleProperty",
		handlers.InsertUserRoleProperty,
	}, Route{
		"getUserRoleProperty",
		"POST",
		"/getUserRoleProperty",
		handlers.GetAllUserRoleProperty,
	}, Route{
		"updateUserRoleProperty",
		"POST",
		"/updateUserRoleProperty",
		handlers.UpdateUserPropertyName,
	}, Route{
		"deleteUserRoleProperty",
		"POST",
		"/deleteUserRoleProperty",
		handlers.DeleteUserPropertyName,
	},Route{
		"getallorganizationfield",
		"POST",
		"/getallorganizationfield",
		handlers.GetAllOrganizationField,
	},
	Route{
		"updateorganizationfield",
		"POST",
		"/updateorganizationfield",
		handlers.UpdateOrganizationField,
	},Route{
		"Heartbeat",
		"POST",
		"/heartbeat",
		handlers.Heartbeat,
	},Route{
		"Getallserver",
		"POST",
		"/getallserver",
		handlers.Getallserver,
	},Route{
		"getprocessbyserver",
		"POST",
		"/getprocessbyserver",
		handlers.Getprocessbyserver,
	},Route{
		"Getstatusbyprocess",
		"POST",
		"/getstatusbyprocess",
		handlers.Getstatusbyprocess,
	},Route{
		"RemoveWorkflowUsers",
		"POST",
		"/removeworkflowusers",
		handlers.RemoveWorkflowUsers,
	},Route{
		"SendOTP",
		"POST",
		"/sendOTP",
		handlers.SendOTP,
	},Route{
		"ForgetPassword",
		"POST",
		"/forgetPassword",
		handlers.ForgetPassword,
	},
}
