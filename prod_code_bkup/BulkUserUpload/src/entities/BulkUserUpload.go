package entities

type MapClientUserRoleUser struct {
	Id                 int64
	ClientId           int64
	MstOrgnHirarchyId  int64
	RoleId             int64
	UserId             int64
	ActiveFlg          int64
	DeleteFlg          int64
	AuditTransactionId int64
}

type MstGroupMember struct {
	Id                 int64
	ClientId           int64
	MstOrgnHirarchyId  int64
	GroupId            int64
	UserId             int64
	ActiveFlg          int64
	DeleteFlg          int64
	AuditTransactionId int64
}

type MstClientUser struct {
	Id                   int64
	ClientId             int64
	MstOrgnHirarchyId    int64
	LoginName            string
	FirstName            string
	LastName             string
	Name                 string
	UserEmail            string
	UserMobileNo         int64
	Password             string
	Passwordactivatedate string
	SecondaryNo          string
	Division             string
	Brand                string
	City                 string
	Designation          string
	Branch               string
	UserType             string
	VipUser              string
	Activeflg            int64
	Deleteflg            int64
	Audittransactionid   int64
}

type MstUser struct {
	Id                   int64
	ClientId             int64
	MstOrgnHirarchyId    int64
	ExternalUserID       int64
	LoginName            string
	FirstName            string
	LastName             string
	Name                 string
	UserEmail            string
	UserMobileNo         int64
	Password             string
	Passwordactivatedate string
	SecondaryNo          string
	Division             string
	Brand                string
	City                 string
	Designation          string
	Branch               string
	UserType             string
	VipUser              string
	Activeflg            int64
	Deleteflg            int64
	Audittransactionid   int64
}
