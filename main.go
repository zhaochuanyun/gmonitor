// 进程监控服务
package main

import (
	"fmt"
	"time"

	"github.com/simplejia/clog"
	"github.com/simplejia/utils"

	_ "github.com/zhaochuanyun/gmonitor/clog"
	"github.com/zhaochuanyun/gmonitor/comm"
	"github.com/zhaochuanyun/gmonitor/conf"
	"github.com/zhaochuanyun/gmonitor/svr"
)

func request(command string, service string) {
	url := fmt.Sprintf("http://%s:%d", utils.LocalIp, conf.C.Port)
	params := map[string]string {
		"command": command,
		"service": service,
	}
	gpp := &utils.GPP {
		Uri:     url,
		Timeout: time.Second * 8,
		Params:  params,
	}
	body, err := utils.Get(gpp)
	if err != nil {
		fmt.Printf("Error: [gmonitor maybe down!] %v, %s\n", err, body)
		return
	}

	fmt.Println(string(body))
	return
}

func main() {
	switch {
	case conf.Start != "":
		request(comm.START, conf.Start)
	case conf.Stop != "":
		request(comm.STOP, conf.Stop)
	case conf.Restart != "":
		request(comm.RESTART, conf.Restart)
	case conf.GraceRestart != "":
		request(comm.GRESTART, conf.GraceRestart)
	case conf.Status != "":
		request(comm.STATUS, conf.Status)
	default:
		clog.Info("main() StartSvr")
		svr.StartSvr()
	}
}
