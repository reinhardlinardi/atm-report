package httpresp

import "errors"

var ErrReq = errors.New("bad request")
var ErrServer = errors.New("interal server error")
