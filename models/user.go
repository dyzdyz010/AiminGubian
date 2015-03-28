package models

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func UserByName(name string) (*User, error) {
	result, err := db.Hget(h_user, name)
	if err != nil {
		return nil, err
	}
	// fmt.Println(result)
	user := &User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func AddUser(name, password string) (*User, error) {
	user := &User{}
	user.Name = name
	user.Password = Hash(password)
	abytes, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(abytes))
	err = db.Hset(h_user, user.Name, string(abytes))
	if err != nil {
		return nil, err
	}

	return user, nil
}
