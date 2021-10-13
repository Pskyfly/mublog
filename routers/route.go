package routers

import (
	"myblog/controller"

	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	//告诉gin框架在哪里找静态文件
	r.Static("/layui", "static/layui")
	r.Static("/css", "static/css")
	r.Static("/acm/layui", "static/layui")
	r.Static("/static","static")
	//告诉gin框架在哪里找hmtl文件
	r.Static("/img", "img")
	r.LoadHTMLGlob("templates/*")
	r.GET("/test", controller.TestHandler)
	r.GET("/blog", controller.BlogHandler)
	r.GET("/homepage", controller.HomePageHandler)
	r.GET("/", controller.HomePageHandler)
	r.GET("/acm", controller.ACMPageHandler)
	r.GET("/acm/week1", controller.ACMWeek1PageHandler)
	r.GET("/acm/week2", controller.ACMWeek2PageHandler)
	r.GET("/acm/week3", controller.ACMWeek3PageHandler)
	r.GET("/acm/week4", controller.ACMWeek4PageHandler)
	r.GET("/acm/week5", controller.ACMWeek5PageHandler)
	r.GET("/linux", controller.LinuxContestHandler)
	r.GET("/problems", controller.ProblemsHandler)
	r.GET("/layui", controller.LayUIContestHandler)
	r.GET("/talk",controller.TalkHandler)
	v1group := r.Group("v1")
	{
		//代办事项
		//添加
		v1group.POST("/todo", controller.CreateATodo)
		//查看
		//查看所有的代办事项
		v1group.GET("/todo", controller.ShowTodoList)
		//修改(前端会返回一个带id的url)
		v1group.PUT("todo/:id", controller.UpdateATodo)
		//删除，前端会把删除信息传入到DELETE接口上，所以需要用DELETE接口来接收（前面几个也一样）
		v1group.DELETE("todo/:id", controller.DeleteATodo)
	}
	return r
}
