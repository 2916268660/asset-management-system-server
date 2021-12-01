package structs


type Response struct {
	ErrNo    int
	ErrMsg   string
	RespData interface{}
}
