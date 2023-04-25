package entities

type MstRecordDifferentiation struct {
	Id                  int64
	ClientId            int64
	MstOrgnHirarchyId   int64
	RecordDiffTypeId    int64
	ParentId            int64
	Name                string
	SeqNo               int64
	DeleteFlag          int64
	ActiveFlag          int64
	AuditTransactionId  int64
	ParentCategoryids   string
	ParentCategoryNames string
}

type MstRecordType struct {
	Id                   int64
	ClientId             int64
	MstOrgnHirarchyId    int64
	FromRecordDiffTypeId int64
	FromRecordDiffId     int64
	ToRecordDiffTypeId   int64
	ToRecordDiffId       int64
	DeleteFlag           int64
	ActiveFlag           int64
	AuditTransactionId   int64
}

type MstBusinessMatrix struct {
	Id                                   int64
	ClientId                             int64
	MstOrgnHirarchyId                    int64
	MstRecordDifferentiationTicketTypeId int64
	MstRecordDifferentiationCatId        int64
	MstRecordDifferentiationImpactId     int64
	MstRecordDifferentiationUrgencyId    int64
	MstRecordDifferentiationPriorityId   int64
	DeleteFlag                           int64
	ActiveFlag                           int64
	AuditTransactionId                   int64
}

type MapCategoryWithEstimateTime struct {
	Id                 int64
	ClientId           int64
	MstOrgnHirarchyId  int64
	RecordDiffId       int64
	EstimatedTime      string
	Efficiency         string
	ChangeType         string
	DeleteFlag         int64
	ActiveFlag         int64
	AuditTransactionId int64
}
