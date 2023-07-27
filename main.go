package main

import (
	"instant_messaging_system/router"
)

func main() {
	panic(router.Router().Run(":8080"))
}

// 只会执行一次
//func main() {
//	fmt.Println(num)
//	fmt.Println(num)
//	fmt.Println(num)
//}
//
//var num = hello()
//
//func hello() int {
//	println("执行了")
//	return 1
//}

//func main() {
//	db.MySQLClient.AutoMigrate(models.Room{}, models.UserRoom{}, models.Message{})
//}
