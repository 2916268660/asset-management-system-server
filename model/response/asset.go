package response

import "time"

type Function struct {
	ID        int64     `json:"id"`         // 待办事项的id
	Kind      string    `json:"kind"`       // 待办事项的种类
	CreatTime time.Time `json:"creat_time"` // 发起时间
	Status    string    `json:"status"`     // 待办事项状态
}
