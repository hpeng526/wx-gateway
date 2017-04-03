package service

import (
	"database/sql"
	"fmt"
	"github.com/hpeng526/wx-gateway/po"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type UserService struct {
	DataSource string
}

func (service *UserService) FindAllUser() ([]po.User, error) {
	db, err := sql.Open("sqlite3", service.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "select id, wx_id, template_id, create_time from t_user"

	rows, err := db.Query(query)
	var users []po.User
	for rows.Next() {
		var id int64
		var wxId string
		var tempId string
		var createTime time.Time
		err = rows.Scan(&id, &wxId, &tempId, &createTime)
		if err != nil {
			return nil, err
		}
		user := po.User{id, wxId, tempId, createTime}
		users = append(users, user)
	}
	return users, err

}

func (service *UserService) FindUserById(id int64) (*po.User, error) {
	db, err := sql.Open("sqlite3", service.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "select id, wx_id, template_id, create_time from t_user where id = ?"

	var u po.User
	err = db.QueryRow(query, id).Scan(&u.UserId, &u.UserWXId, &u.TemplateId, &u.CreateTime)

	if err != nil {
		return &po.User{}, err
	} else {
		return &u, nil
	}

}

func (service *UserService) InsertUser(u *po.User) int64 {
	db, err := sql.Open("sqlite3", service.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt, err := tx.Prepare("insert into t_user(id, wx_id, template_id, create_time) values(?,?,?,?)")
	defer sqlStmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user is %v\n", u)

	res, err := sqlStmt.Exec(u.UserId, u.UserWXId, u.TemplateId, u.CreateTime)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	affect, err := res.RowsAffected()
	fmt.Println("row affect %d", affect)
	return affect
}

func (service *UserService) DeleteUser(userId int64) (int64, error) {
	db, err := sql.Open("sqlite3", service.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt, err := db.Prepare("delete from t_user where id=?")

	res, err := sqlStmt.Exec(userId)

	if err != nil {
		return 0, nil
	}
	affect, err := res.RowsAffected()

	if err != nil {
		return 0, nil
	}

	return affect, nil
}
