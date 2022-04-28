package request

type ApplyReceiveForm struct {
	Category string `json:"category" binding:"required"`           // 资产品类
	Nums     int    `json:"nums" binding:"required",gte=1,lte=10`  // 申请领用资产数量
	Days     int    `json:"days" binding:"required",gte=1,lte=365` // 申请天数
	Remake   string `json:"remake" binding:"max=600"`              // 备注信息
}

type ApplyRevertForm struct {
	Assets []string `json:"assets" binding:"required"` // 申请归还资产的序列号数组
	Remake string   `json:"remake" binding:"max=600"`  // 备注信息
}

type ApplyRepairForm struct {
	Address string   `json:"address" binding:"required,max=100"`
	Assets  []string `json:"assets" binding:"required"`
	Remake  string   `json:"remake" binding:"max=600"`
}

type AssetInfo struct {
	Category string  `json:"category" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price"`
	Provide  string  `json:"provide"`
	Nums     int     `json:"nums" binding:"required",gte=1`
}

type AuditStatus struct {
	ID     int64 `json:"id" binding:"required"`
	Status int   `json:"status" binding:"required"`
}

type RollbackInfo struct {
	ID       int64  `json:"id" binding:"required"`
	Category string `json:"category" binding:"required"`
}

type UpdateAssetInfo struct {
	SerialId string  `json:"serialId" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price"`
	Provide  string  `json:"provide"`
}

type ReceiveForm struct {
	ID             int64    `json:"id"`         // 主键ID
	UserId         string   `json:"userId"`     // 申请人账号
	UserName       string   `json:"userName"`   // 申请人姓名
	UserPhone      string   `json:"userPhone"`  // 申请人联系方式
	Department     string   `json:"department"` // 申请人所属部门
	Nums           int      `json:"nums"`       // 归还资产数量
	Days           int      `json:"days"`
	Assets         []string `json:"assets"`         // 资产的序列号json字符串
	ReclaimerId    string   `json:"reclaimerId"`    // 同意回收管理员的账号
	ReclaimerName  string   `json:"reclaimerName"`  // 同意回收管理员的姓名
	ReclaimerPhone string   `json:"reclaimerPhone"` // 同意回收管理员的联系方式
	Remake         string   `json:"remake"`         // 备注信息
	Status         int      `json:"status"`
}
