package controllers

import (
	// "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/AiminGubian/models"
	"github.com/qiniu/api/rs"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Prepare() {
	checkLogin(c)
}

func checkLogin(c *AdminController) {
	name := c.GetSession("user")
	if name == nil && c.Ctx.Request.URL.String() != "/admin/login" {
		c.Redirect("/admin/login", 302)
	} else if name != nil && c.Ctx.Request.URL.String() == "/admin/login" {
		c.Redirect("/admin", 302)
	}
}

func (c *AdminController) Login() {

	renderTemplate(c.Ctx, "views/admin/login.amber", c.Data)
}

func (c *AdminController) PostLogin() {
	user := models.User{}
	err := c.ParseForm(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	u, err := models.UserByName(user.Name)
	if err != nil {
		c.Data["Message"] = err.Error()
	} else if models.Hash(user.Password) != u.Password {
		c.Data["Message"] = "Wrong password."
	} else {
		c.SetSession("user", user.Name)
		c.Redirect("/admin", 302)
	}

	renderTemplate(c.Ctx, "views/admin/login.amber", c.Data)
}

func (c *AdminController) Index() {
	entries, err := models.AllEntries()
	if err != nil {
		panic(err)
	}
	c.Data["Entries"] = entries
	sectionsMap := make(map[string]interface{})
	for _, section := range models.Sections {
		sectionsMap[section["id"].(string)] = section["name"]
		fmt.Println(section)
	}
	c.Data["Sections"] = models.Sections
	bc, err := models.GetBroadcast()
	if err != nil {
		panic(err)
	}
	fmt.Println(bc)
	c.Data["Broadcast"] = bc

	renderTemplate(c.Ctx, "views/admin/index.amber", c.Data)
}

func (c *AdminController) Section() {
	renderTemplate(c.Ctx, "views/admin/section.amber", c.Data)
}

func (c *AdminController) Entry() {
	eid := c.Ctx.Input.Param(":id")
	if eid != "new" {
		entry, _ := models.EntryById(eid)
		if entry == nil {
			c.Abort("404")
			return
		} else {
			c.Data["Entry"] = entry
			eid = entry.Id
		}
	}

	fmt.Println(models.Sections)
	c.Data["PostId"] = eid
	c.Data["Sections"] = models.Sections

	renderTemplate(c.Ctx, "views/admin/entry.amber", c.Data)
}

func (c *AdminController) PostEntry() {
	fmt.Println(string(c.Ctx.Input.RequestBody))
	entry := models.Entry{}
	err := c.ParseForm(&entry)
	if err != nil {
		panic(err)
	}
	eid := c.Ctx.Input.Param(":id")
	err = nil
	if eid == "new" {
		eid, err = models.AddEntry(entry)
	} else {
		entry.Id = eid
		err = models.UpdateEntry(entry)
	}
	if err != nil {
		panic(err)
	}
	c.Data["PostId"] = eid
	entry.Id = eid
	c.Data["Entry"] = entry
	c.Data["Sections"] = models.Sections

	renderTemplate(c.Ctx, "views/admin/entry.amber", c.Data)
}

func (c *AdminController) DeleteEntry() {
	err := models.DeleteEntry(c.GetString("id"))
	if err != nil {
		panic(err)
	}

	c.Redirect("/admin", 302)
}

func (c *AdminController) PostBroadcast() {
	// fmt.Println(c.GetString("broadcast"))
	bc := c.GetString("broadcast")
	err := models.SetBroadcast(bc)
	if err != nil {
		panic(err)
	}
	c.Data["Broadcast"] = bc

	c.Redirect("/admin", 302)
}

func (this *AdminController) QiniuToken() {
	bucket := models.Appconf.String("qiniu::bucket")
	putPolicy := rs.PutPolicy{
		Scope: bucket,
	}

	this.Data["json"] = putPolicy.Token(nil)
	this.ServeJson()
}
