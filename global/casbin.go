package global

// 用户权限
const (
	User     = "user"
	Charger  = "charger"
	Provider = "provider"
)

var PoliciesList = [][]string{
	{
		User,
		"/applyReceive",
		"POST",
	},
	{
		User,
		"/applyRevert",
		"POST",
	},
	{
		User,
		"/applyRepair",
		"POST",
	},
	{
		User,
		"/getAsset",
		"GET",
	},
	{
		User,
		"/getReceiveTodo",
		"GET",
	},
	{
		User,
		"/getRevertTodo",
		"GET",
	},
	{
		User,
		"/getRepairsTodo",
		"GET",
	},
	{
		User,
		"/rollback",
		"POST",
	},
	{
		Charger,
		"/applyReceive",
		"POST",
	},
	{
		Charger,
		"/applyRevert",
		"POST",
	},
	{
		Charger,
		"/applyRepair",
		"POST",
	},
	{
		Charger,
		"/getAsset",
		"GET",
	},
	{
		Charger,
		"/getReceiveTodo",
		"GET",
	},
	{
		Charger,
		"/getRevertTodo",
		"GET",
	},
	{
		Charger,
		"/getRepairsTodo",
		"GET",
	},
	{
		Charger,
		"/rollback",
		"POST",
	},
	{
		Charger,
		"/getAuditTodo",
		"GET",
	},
}
