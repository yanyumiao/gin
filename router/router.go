package router

import "github.com/gin-gonic/gin"
import . "gin-demo/controller" // Notice: ".", 点操作 导入包后 调用这个包的函数时 可以省略包名

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router
	// hello.go
	router.GET("/", Hello)
	router.GET("/hello", Hello)
	router.GET("/ping", Pong)
	router.GET("/async", Async)
	router.GET("/redis/set", Set)
	// person.go
	router.GET("/person/add", AddPerson) // http://127.0.0.1:10086/person/add?firstname=a&lastname=b
	router.GET("/person/get", GetPerson) // http://127.0.0.1:10086/person/get?id=1
	router.GET("/person/all", GetAllPerson)
	router.GET("/person/del", DelPerson)       // ?id=4
	router.GET("/person/update", UpdatePerson) // ?id=4&firstname=a&lastname=b
	// router.StaticFile("/favicon.ico", "./resource/favicon.ico")
	return router
}
