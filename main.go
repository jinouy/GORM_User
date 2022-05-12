package main

import (
	"gorm_tcp/controller"
	"net/http"
)

func main() {

	http.HandleFunc("/getUsers", controller.GetUsers)
	http.HandleFunc("/getUserName", controller.GetUserName)
	http.HandleFunc("/getUserPage", controller.GetUserPage)
	http.HandleFunc("/addUser", controller.AddUser)
	http.HandleFunc("/deleteUserName", controller.DelUser)
	http.HandleFunc("/updateUserName", controller.DpdUserName)
	//http.HandleFunc("/deleteUserID", controller.DelUserID)

	http.ListenAndServe(":8080", nil)

}
