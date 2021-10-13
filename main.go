package main

import (
	"myblog/dao"
	"myblog/models"
	"myblog/routers"
)

func main() {
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close() //用完记得关闭数据库,不要放在函数里面（大坑）！
	//绑定模型
	dao.DB.AutoMigrate(&models.Todo{}) //创建的表名为todos，如果手动创建表名一定要一致

	r := routers.SetRouters() //设置路由
	r.Run(":8080")

}
