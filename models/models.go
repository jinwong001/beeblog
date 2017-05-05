package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/Unknwon/com"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
)

const (
	_DB_NAME = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time          `orm:"index"`
	Views           int64              `orm:"index"`
	TopicTime       time.Time          `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string            `orm:"size(5000)`
	Attachment      string
	Created         time.Time         `orm:"index"`
	Updated         time.Time         `orm:"index"`
	Views           int64             `orm:"index"`
	Author          string
	ReplyTime       time.Time         `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)

}

func AddCategory(name string) error {
	o := orm.NewOrm()
	category := &Category{Title:name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(category)
	if err == nil {
		return err
	}

	_, err = o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&categories)
	return categories, err
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	category := &Category{Id:cid}
	_, err = o.Delete(category)
	return err;

}

