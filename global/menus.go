package global

import "server/model/response"

// RoleMenusMap 菜单列表
var RoleMenusMap = map[string][]*response.Menu{
	User: []*response.Menu{
		&response.Menu{
			ID:       1,
			AuthName: "首页",
			Path:     "/home",
		},
		&response.Menu{
			ID:       10,
			AuthName: "我的资产",
			Path:     "/asset",
		},
		&response.Menu{
			ID:       2,
			AuthName: "我的待办",
			Path:     "",
			Children: []*response.Menu{
				&response.Menu{
					ID:       3,
					AuthName: "领用待办",
					Path:     "/receiveTodo",
				},
				&response.Menu{
					ID:       4,
					AuthName: "归还待办",
					Path:     "/revertTodo",
				},
				&response.Menu{
					ID:       5,
					AuthName: "维修待办",
					Path:     "/repairsTodo",
				},
			},
		},
	},
	Charger: []*response.Menu{
		&response.Menu{
			ID:       1,
			AuthName: "首页",
			Path:     "/home",
		},
		&response.Menu{
			ID:       10,
			AuthName: "我的资产",
			Path:     "/asset",
		},
		&response.Menu{
			ID:       2,
			AuthName: "我的待办",
			Path:     "",
			Children: []*response.Menu{
				&response.Menu{
					ID:       6,
					AuthName: "审批待办",
					Path:     "/auditTodo",
				},
				&response.Menu{
					ID:       3,
					AuthName: "领用待办",
					Path:     "/receiveTodo",
				},
				&response.Menu{
					ID:       4,
					AuthName: "归还待办",
					Path:     "/revertTodo",
				},
				&response.Menu{
					ID:       5,
					AuthName: "维修待办",
					Path:     "/repairsTodo",
				},
			},
		},
	},
	Provider: []*response.Menu{
		&response.Menu{
			ID:       1,
			AuthName: "首页",
			Path:     "/home",
		},
		&response.Menu{
			ID:       2,
			AuthName: "我的待办",
			Path:     "/todo",
		},
		&response.Menu{
			ID:       7,
			AuthName: "用户管理",
			Path:     "/user",
		},
		&response.Menu{
			ID:       11,
			AuthName: "资产管理",
			Path:     "/assetForAdmin",
		},
	},
}
