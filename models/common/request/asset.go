package request

type ApplyReceiveForm struct {
	Category int    // 资产品类
	Nums     int    // 申请领用资产数量
	Days     int    // 申请天数
	Remake   string // 备注信息
}

type ApplyRevertForm struct {
	Assets []string // 申请归还资产的序列号数组
	Remake string   // 备注信息
}

type ApplyRepairForm struct {
	Address string
	Assets  []string
	Remake  string
}

type AssetInfo struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Provide  string  `json:"provide"`
}
