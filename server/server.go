package server

import "github.com/twwch/gin-sdk/constant"

type Server interface {
	New(address string) Server
}

func New(typ, address string) Server {
	switch typ {
	case constant.HTTPProtocal:
		return defaultHTTP.New(address)
	default:
		panic("what do you want?")
	}
}
