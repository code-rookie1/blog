package main

import (
	"blog/core"
	"blog/global"
	"blog/initialize"
)

func main()  {
	global.GVB_VP = core.Viper()    //初始话Viper 读取配置
	global.GVB_DB = initialize.GormMysql()   //初始化数据库
	initialize.MysqlTables(global.GVB_DB)   //初始化表
	// 程序结束前关闭数据库连接
	db, _ := global.GVB_DB.DB()
	defer db.Close()
	core.RunWindowsServer()
}
