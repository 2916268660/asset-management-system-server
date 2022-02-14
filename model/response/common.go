package response

type Menu struct {
	ID       int     `json:"id"`
	AuthName string  `json:"authName"`
	Path     string  `json:"path"`
	Children []*Menu `json:"children"`
}
