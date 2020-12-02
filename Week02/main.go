package main

import (
	"Go-000/Week02/service"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {
	//直接模拟调用
	userlist, err := service.UserService.GetUserByAge(10)
	if err != nil {
		log.Fatalf("%v", errors.Cause(err))
	}
	fmt.Println(userlist)
	user, err := service.UserService.GetUserInfo(11234434)
	if err != nil {
		log.Fatalf("%v", err)
	} else {
		fmt.Println(user)
	}

}
