package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/Unknwon/com"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"github.com/CardInfoLink/log"
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
	category := &Category{Title:name, Created:time.Now(), TopicTime:time.Now()}
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

func GetAllTopics(des bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if des {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{Title:title, Content:content, Created:time.Now(),
		Updated:time.Now(), ReplyTime:time.Now() }
	_, err := o.Insert(topic)
	return err
}

func DelTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id:tidNum}
	_, err = o.Delete(topic)
	return err;
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	log.Debug(tidNum)
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	topic := new(Topic)

	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err;
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id:tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
}

