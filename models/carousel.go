package models

import (
	"fmt"
)

type Carousel struct {
	Id       int64  `form:"-"`
	Title    string `form:"title" valid:"Required"`
	Subtitle string `form:"subtitle" valid:"Required"`
	Url      string `form:"url" valid:"Required"`
	PicId    int    `form:"picid" valid:"Required;Min(1)"`
}

func (c *Carousel) Get() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	has, err := Engine.Get(c)
	if !has {
		panic(fmt.Errorf("not exist"))
	}
	return
}

func (c *Carousel) Put() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	_, err = Engine.Insert(c)

	return
}

func (c *Carousel) Delete() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	_, err = Engine.Delete(c)
	return
}

type CarouselSlice []Carousel

func (c *CarouselSlice) Recent(n int) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	err = Engine.Desc("id").Limit(n).Find(c)
	return
}