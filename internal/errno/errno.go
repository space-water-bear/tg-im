package errno

type Errno struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Errno) Error() string {
	return e.Msg
}

func NewDefaultError(code int32) error {
	return &Errno{Code: code, Msg: ErrMsg[code]}
}

func NewParamsError(err error) error {
	return &Errno{Code: ErrBind, Msg: err.Error()}
}
