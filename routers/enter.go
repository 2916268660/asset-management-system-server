package routers

import "server/routers/user"

type routersGroup struct {
	UserRoutersGroup user.RoutersGroup
}

var RoutersGroupApp = new(routersGroup)
