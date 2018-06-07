package clog

import (
	"github.com/zhaochuanyun/clog"
	"github.com/zhaochuanyun/gmonitor/conf"
	"github.com/zhaochuanyun/namecli/api"
)

func init() {
	clog.AddrFunc = func() (string, error) {
		return api.Name(conf.C.Clog.Addr)
	}
	clog.Init(conf.C.Clog.Name, "", conf.C.Clog.Level, conf.C.Clog.Mode)
}
