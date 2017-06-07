package service

import (
	"github.com/hpeng526/wx-gateway/po"
	"log"
	"testing"
	"time"
)

func TestUserService_InsertUser(t *testing.T) {

	user := po.User{
		UserId:     2,
		UserWXId:   "oXymQwcLPX",
		TemplateId: "TY33t4IkXbyo",
		CreateTime: time.Now()}

	us := UserService{}

	row := us.InsertUser(&user)
	log.Printf("afrow is %d", row)
}

func TestUserService_FindUserById(t *testing.T) {
	us := UserService{DataSource: "../gateway.sqlite"}
	user, err := us.FindUserById(1)
	if err != nil {
		log.Printf("error %v\n", err)
	}
	log.Println(user)
}

func TestUserService_FindAllUser(t *testing.T) {
	us := UserService{DataSource: "../gateway.sqlite"}
	users, err := us.FindAllUser()
	if err != nil {
		log.Printf("error %v\n", err)
	}
	log.Println(users)
}

func TestUserService_FindUsersByGroup(t *testing.T) {
	us := UserService{DataSource: "../gateway.sqlite"}
	users, err := us.FindUsersByGroup(100)
	if err != nil {
		log.Printf("error %v\n", err)
	}
	log.Println(users)
}

func TestUserService_DeleteUser(t *testing.T) {
	us := UserService{DataSource: "../gateway.sqlite"}

	af, _ := us.DeleteUser(2)
	log.Println("delete ", af)
}
