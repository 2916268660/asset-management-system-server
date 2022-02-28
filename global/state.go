package global

// 申请领用状态
const (
	WaitAudit     = iota + 1 // 审批中
	WaitProvide              // 待发放
	Rejected                 // 已驳回
	ProvideFailed            // 发放失败
	ProvideDone              // 发放成功

	Reverting // 归还中
	Reverted  // 已归还

	Maintaining // 维护中

	WaitReceive // 待接单
	WaitRepair  // 待维修
	RepairDone  // 维修完成

	CanApply // 可领用
	Applied  // 已领用

	Rollback // 已撤销
)

var StatusMap = map[int]string{
	WaitAudit:     "审批中",
	WaitProvide:   "待发放",
	Rejected:      "已驳回",
	ProvideFailed: "发放失败",
	ProvideDone:   "发放成功",

	Reverting: "归还中",
	Reverted:  "已归还",

	Maintaining: "维护中",

	WaitReceive: "待接单",
	WaitRepair:  "待维修",
	RepairDone:  "维修完成",

	CanApply: "可领用",
	Applied:  "已领用",

	Rollback: "已撤销",
}
