package main

import (
	"fmt"
	"log"
	"todoapp/app/models"
	"todoapp/config"
)

func main()  {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.Logfile)

	// データ作成用
	log.Println("tets")
	fmt.Println(models.Db)
	u := &models.User{}
	u.Name  ="Ochiai"
	u.Email  ="hiromitsua@gmail.com"
	u.Password = "testtest"
	fmt.Println(u)

	// u.CreateUser()

	 //get確認用
	// u,_:= models.GetUser("Tanaka")
	// fmt.Println("getAllUser : ",u)

	//update用 
	// user := &models.User{}
	// user.UpdateUser(1, map[string]interface{}{"email": "newemailnewEmail@gmail.com"})
	// u,_:= models.GetAllUser()
	// fmt.Println("getAllUser : ",u)

	//delete確認用
	// user := &models.User{}
	// user.DeleteUser(4)
	// u, _ := models.GetAllUser()
	// fmt.Println("getAllUser :",u)

}