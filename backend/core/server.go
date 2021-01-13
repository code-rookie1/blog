package core

import (
	"blog/global"
	"blog/initialize"
	"fmt"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer()  {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVB_CONFIG.System.Addr)
	s := initServer(address, Router)
	fmt.Println("server run success on" )
	fmt.Println(s.ListenAndServe().Error())
}
