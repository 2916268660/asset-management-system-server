package response

type Function struct {
	ID         int64  `json:"id"`          // 待办事项的id
	Kind       string `json:"kind"`        // 待办事项的种类
	CreateTime int64  `json:"create_time"` // 发起时间
	Status     int    `json:"status"`      // 待办事项状态
}

type AssetReceive struct {
	ID            int64    `json:"id"`            // 主键ID
	UserId        string   `json:"userId"`        // 申请人账号
	UserName      string   `json:"userName"`      // 申请人姓名
	UserPhone     string   `json:"userPhone"`     // 申请人联系方式
	Department    string   `json:"department"`    // 申请人所属部门
	Category      string   `json:"category"`      // 资产品类
	Nums          int      `json:"nums"`          // 申请资产数量
	Days          int      `json:"days"`          // 申请天数
	Assets        []string `json:"assets"`        // 资产的序列号json字符串
	AdminId       string   `json:"adminId"`       // 同意领用管理员的账号
	AdminName     string   `json:"adminName"`     // 同意领用管理员的姓名
	AdminPhone    string   `json:"adminPhone"`    // 同意领用管理员的联系方式
	ProviderId    string   `json:"providerId"`    // 发放资产人的账号
	ProviderName  string   `json:"providerName"`  // 发放资产人的姓名
	ProviderPhone string   `json:"providerPhone"` // 发放资产人的联系方式
	Remake        string   `json:"remake"`        // 备注信息
	Status        int      `json:"status"`        // 任务单状态
	ExpireTime    int64    `json:"expireTime"`    // 到期时间
	ProvideTime   int64    `json:"provideTime"`   // 发放时间
	CreateTime    int64    `json:"createTime"`    // 申请时间
	AuditTime     int64    `json:"auditTime"`     // 审批时间
	UpdateTime    int64    `json:"updateTime"`    // 更新时间
	RollbackTime  int64    `json:"rollbackTime"`  // 撤回时间
}

type AssetRevert struct {
	ID             int64    `json:"id"`             // 主键ID
	UserId         string   `json:"userId"`         // 申请人账号
	UserName       string   `json:"userName"`       // 申请人姓名
	UserPhone      string   `json:"userPhone"`      // 申请人联系方式
	Department     string   `json:"department"`     // 申请人所属部门
	Nums           int      `json:"nums"`           // 归还资产数量
	Assets         []string `json:"assets"`         // 资产的序列号json字符串
	ReclaimerId    string   `json:"reclaimerId"`    // 同意回收管理员的账号
	ReclaimerName  string   `json:"reclaimerName"`  // 同意回收管理员的姓名
	ReclaimerPhone string   `json:"reclaimerPhone"` // 同意回收管理员的联系方式
	Remake         string   `json:"remake"`         // 备注信息
	Status         int      `json:"status"`         // 任务单状态
	CreateTime     int64    `json:"createTime"`     // 申请时间
	RevertTime     int64    `json:"revertTime"`     // 归还时间
	UpdateTime     int64    `json:"updateTime"`     // 更新时间
	RollbackTime   int64    `json:"rollbackTime"`   // 撤回时间
}

type AssetRepairs struct {
	ID            int64    `json:"id"`            // 主键ID
	UserId        string   `json:"userId"`        // 申请人账号
	UserName      string   `json:"userName"`      // 申请人姓名
	UserPhone     string   `json:"userPhone"`     // 申请人联系方式
	Address       string   `json:"address"`       // 地址
	Assets        []string `json:"assets"`        // 资产的序列号json字符串
	Remake        string   `json:"remake"`        // 备注信息
	RepairerName  string   `json:"repairerName"`  // 维修人员的姓名
	RepairerPhone string   `json:"repairerPhone"` // 维修人员的联系方式
	Status        int      `json:"status"`        // 维修单状态
	CreateTime    int64    `json:"createTime"`
	UpdateTime    int64    `json:"updateTime"`
	ReceiveTime   int64    `json:"receiveTime"`  // 接单时间
	RepairedTime  int64    `json:"repairedTime"` // 维修完成时间
	RollbackTime  int64    `json:"rollbackTime"`
}

type AssetsInfo struct {
	SerialId   string `json:"serialId"`
	Status     int    `json:"status"`
	ExpireTime int64  `json:"expireTime"`
}

type AssetDetail struct {
	SerialId   string  `json:"serialId"`
	Category   string  `json:"category"`
	Name       string  `json:"name"`
	Status     int     `json:"status"`
	Price      float64 `json:"price"`
	Provide    string  `json:"provide"`
	CreateTime int64   `json:"createTime"`
}

type AssetVO struct {
	SerialId string `json:"serialId"`
	Category string `json:"category"`
	Name     string `json:"name"`
}

type Asset struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	SerialId string `json:"serialId"`
	Status   int    `json:"status"`
}

type ReceiveUser struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	UserPhone   string `json:"userPhone"`
	ProvideTime int64  `json:"provideTime"`
	ExpireTime  int64  `json:"expireTime"`
}
