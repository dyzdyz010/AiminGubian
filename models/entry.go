package models

import (
	"encoding/json"
	// "fmt"
	"strconv"
	"time"
)

type Entry struct {
	Id      string `json:"id"`
	Title   string `json:"title" form:"title"`
	Author  string `json:"author" form:"author"`
	Date    string `json:"date" form:"date"`
	Section string `json:"section" form:"section"`
	Content string `json:"content" form:"content"`
	Status  string `json:"status" form:"status"`
}

func AllEntries() ([]Entry, error) {
	size, err := db.Zsize(zname("all", "entry"))
	if err != nil {
		return nil, err
	}

	result, err := db.Zscan(zname("all", "entry"), "", "", "", size)
	if err != nil {
		return nil, err
	}

	eids := make([]string, 0)
	for i := len(result) - 2; i >= 0; i -= 2 {
		eids = append(eids, result[i])
	}
	if len(eids) == 0 {
		return nil, nil
	}

	result, err = db.Multi_hget(h_entry, eids)
	if err != nil {
		return nil, err
	}

	entries := []Entry{}
	for i := 1; i < len(result); i += 2 {
		entryStr := result[i]
		entry := Entry{}
		json.Unmarshal([]byte(entryStr), &entry)
		entries = append(entries, entry)
	}

	return entries, nil
}

func PublishedEntriesBySection(sid string) ([]Entry, error) {
	size, err := db.Zsize(zname("published", "entry"))
	if err != nil {
		return nil, err
	}
	result, err := db.Zscan(zname("published", "entry"), "", "", "", size)
	if err != nil {
		return nil, err
	}

	eids := make([]string, 0)
	for i := len(result) - 2; i >= 0; i -= 2 {
		eids = append(eids, result[i])
	}
	if len(eids) == 0 {
		return nil, nil
	}

	result, err = db.Multi_hget(h_entry, eids)
	if err != nil {
		return nil, err
	}

	entries := []Entry{}
	for i := 1; i < len(result); i += 2 {
		entryStr := result[i]
		entry := Entry{}
		json.Unmarshal([]byte(entryStr), &entry)
		if entry.Section == sid {
			entries = append(entries, entry)
		}
	}

	return entries, nil
}

func EntriesBySection(sid string) ([]Entry, error) {
	size, err := db.Zsize(zname(sid, "entry"))
	if err != nil {
		return nil, err
	}

	result, err := db.Zscan(zname(sid, "entry"), "", "", "", size)
	if err != nil {
		return nil, err
	}

	eids := make([]string, 0)
	for i := len(result) - 2; i >= 0; i -= 2 {
		eids = append(eids, result[i])
	}
	if len(eids) == 0 {
		return nil, nil
	}

	result, err = db.Multi_hget(h_entry, eids)
	if err != nil {
		return nil, err
	}

	entries := []Entry{}
	for i := 1; i < len(result); i += 2 {
		entryStr := result[i]
		entry := Entry{}
		json.Unmarshal([]byte(entryStr), &entry)
		entries = append(entries, entry)
	}

	return entries, nil
}

func EntryById(eid string) (*Entry, error) {
	estr, err := db.Hget(h_entry, eid)
	if err != nil {
		return nil, err
	}
	entry := &Entry{}
	json.Unmarshal([]byte(estr), entry)

	return entry, nil
}

func AddEntry(e Entry) (string, error) {
	e.Id = Hash(e.Title)

	ebytes, _ := json.Marshal(e)
	err := db.Hset(h_entry, e.Id, string(ebytes))
	if err != nil {
		return "", err
	}

	date, _ := time.Parse("2006-01-02", e.Date)
	tnow := time.Now()
	t := time.Date(date.Year(), date.Month(), date.Day(), tnow.Hour(), tnow.Minute(), tnow.Second(), tnow.Nanosecond(), tnow.Location())
	score := t.Unix()
	err = db.Zset(zname(e.Section, "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	if e.Status == "published" {
		err = db.Zset(zname("published", "entry"), e.Id, score)
		if err != nil {
			return "", err
		}
	}

	err = db.Zset(zname("all", "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	return e.Id, nil
}

func UpdateEntry(e Entry) error {
	oldestr, err := db.Hget(h_entry, e.Id)
	oldE := Entry{}
	err = json.Unmarshal([]byte(oldestr), &oldE)

	ebytes, _ := json.Marshal(e)
	err = db.Hset(h_entry, e.Id, string(ebytes))
	if err != nil {
		return err
	}

	scoreStr, err := db.Zget(zname(oldE.Section, "entry"), e.Id)
	score, _ := strconv.Atoi(scoreStr)
	if oldE.Section != e.Section {
		if err != nil {
			return err
		}
		err = db.Zdel(zname(oldE.Section, "entry"), e.Id)
		if err != nil {
			return err
		}
		err = db.Zset(zname(e.Section, "entry"), e.Id, int64(score))
		if err != nil {
			return err
		}
	}

	if e.Status == "published" && oldE.Status == "draft" {
		err = db.Zset(zname("published", "entry"), e.Id, int64(score))
		if err != nil {
			return err
		}
	}
	if e.Status == "draft" && oldE.Status == "published" {
		err = db.Zdel(zname("published", "entry"), e.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteEntry(eid string) error {
	entry, err := EntryById(eid)
	if err != nil {
		return err
	}

	err = db.Zdel(zname("all", "entry"), eid)
	if err != nil {
		return err
	}
	err = db.Zdel(zname(entry.Section, "entry"), eid)
	if err != nil {
		return err
	}

	if entry.Status == "published" {
		err = db.Zdel(zname("published", "entry"), eid)
		if err != nil {
			return err
		}

	}
	err = db.Hdel(h_entry, eid)
	if err != nil {
		return err
	}

	return nil
}
