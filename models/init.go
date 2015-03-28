package models

import (
	"encoding/json"
	"errors"
	// "fmt"
	"github.com/astaxie/beego/config"
	"github.com/dyzdyz010/AiminGubian/ssdb"
	"io/ioutil"
)

var db *ssdb.Client
var Appconf config.ConfigContainer

// Hash Map Names
var db_prefix = "amgb_"
var h_entry = db_prefix + "entry"
var h_section = db_prefix + "section"
var h_user = db_prefix + "user"
var broadcast = db_prefix + "broadcast"
var Sections []map[string]interface{}

func init() {
	err := errors.New("")

	Appconf, err = config.NewConfig("json", "conf/site.json")
	if err != nil {
		panic(err)
	}

	// fmt.Println(Appconf.Int("database::port"))
	host := Appconf.String("database::host")
	port, _ := Appconf.Int("database::port")
	// fmt.Println(Appconf.DIY("blog::author"))

	db, err = ssdb.Connect(host, port)
	if err != nil {
		panic(err)
	}

	db.Hclear(h_user)

	usersRaw, err := Appconf.DIY("site::admin")
	if err != nil {
		panic(err)
	}
	users := usersRaw.([]interface{})
	for _, a := range users {
		atemp, ok := a.(map[string]interface{})
		if !ok {
			panic(ok)
		}

		user := &User{}
		user.Name = atemp["name"].(string)
		user.Password = atemp["password"].(string)
		AddUser(user.Name, user.Password)
	}

	file, err := ioutil.ReadFile("models/sections.json")
	if err != nil {
		panic(err)
	}
	Sections = make([]map[string]interface{}, 0)
	json.Unmarshal(file, &Sections)
	// fmt.Println(Sections)
}
