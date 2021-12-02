package user

type modelGroup struct {
	RegisterModel
	LoginModel
}

var ModelGroupApp = new(modelGroup)