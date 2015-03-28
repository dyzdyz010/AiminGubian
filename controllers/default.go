package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/AiminGubian/models"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {

	file, err := ioutil.ReadFile("models/sections.json")
	if err != nil {
		panic(err)
	}
	sections := make([]map[string]interface{}, 0)
	json.Unmarshal(file, &sections)
	// fmt.Println(sections)
	c.Data["Sections"] = sections

	entriesArr := make([][]models.Entry, 0)
	for _, section := range sections {
		sid := section["id"].(string)
		// fmt.Println(section, entriesMap)
		entries, err := models.PublishedEntriesBySection(sid)
		if err != nil {
			panic(err)
		}
		entriesArr = append(entriesArr, entries)
	}
	fmt.Println(entriesArr)
	c.Data["Entries"] = entriesArr

	bc, err := models.GetBroadcast()
	if err != nil {
		panic(err)
	}
	c.Data["Broadcast"] = bc

	renderTemplate(c.Ctx, "views/index.amber", c.Data)
}

func (c *MainController) Section() {
	sid := c.Ctx.Input.Param(":id")
	entries, err := models.PublishedEntriesBySection(sid)
	if err != nil {
		panic(err)
	}
	file, err := ioutil.ReadFile("models/sections.json")
	if err != nil {
		panic(err)
	}
	c.Data["Entries"] = entries
	sections := make([]map[string]interface{}, 0)
	json.Unmarshal(file, &sections)
	for _, section := range sections {
		sectionId := section["id"].(string)
		// fmt.Println(section)
		if sectionId == sid {
			c.Data["SectionTitle"] = section["name"]
		}
	}
	// fmt.Println(c.Data)

	renderTemplate(c.Ctx, "views/section.amber", c.Data)
}

func (c *MainController) Entry() {
	eid := c.Ctx.Input.Param(":id")
	entry, err := models.EntryById(eid)
	if err != nil {
		panic(err)
	}
	c.Data["Entry"] = entry

	renderTemplate(c.Ctx, "views/entry.amber", c.Data)
}

func (c *MainController) HostIntro() {

	file, err := ioutil.ReadFile("models/host.json")
	if err != nil {
		panic(err)
	}
	express := make(map[string]interface{})
	json.Unmarshal(file, &express)
	c.Data["Data"] = express

	renderTemplate(c.Ctx, "views/host-index.amber", c.Data)
}

func (c *MainController) ExpressIntro() {

	renderTemplate(c.Ctx, "views/express-index.amber", c.Data)
}

func (c *MainController) Charts() {

	file, err := ioutil.ReadFile("models/charts.json")
	if err != nil {
		panic(err)
	}
	c.Data["ChartsData"] = string(file)
	renderTemplate(c.Ctx, "views/charts-index.amber", c.Data)
}

func (c *MainController) Gallery() {
	file, err := ioutil.ReadFile("models/gallery.json")
	if err != nil {
		panic(err)
	}
	photos := make(map[string]interface{})
	json.Unmarshal(file, &photos)
	c.Data["Photos"] = photos["progress"]

	renderTemplate(c.Ctx, "views/gallery.amber", c.Data)
}

func (c *MainController) TimeLine() {

	file, err := ioutil.ReadFile("models/gallery.json")
	if err != nil {
		panic(err)
	}
	express := make(map[string]interface{})
	json.Unmarshal(file, &express)
	c.Data["Preparation"] = express["begin"]
	c.Data["Progress"] = express["progress"]

	renderTemplate(c.Ctx, "views/timeline-index.amber", c.Data)
}

func (c *MainController) Propaganda() {

	file, err := ioutil.ReadFile("models/news.json")
	if err != nil {
		panic(err)
	}
	news := make([]map[string]interface{}, 0)
	json.Unmarshal(file, &news)
	c.Data["News"] = news

	renderTemplate(c.Ctx, "views/propaganda.amber", c.Data)
}
