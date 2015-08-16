package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	name  string
	phone string
	err   error
}

const (
	TooManyArgs = "too many arguments"
)

func (p *Person) Error() error {
	return p.err
}

func (p *Person) Name(str ...string) string {

	switch len(str) {
	case 0:
		return p.name
	case 1:
		p.name = str[0]
		return ""
	}
	p.err = errors.New(TooManyArgs)
	return ""
}

func (p *Person) Phone(str ...string) string {

	switch len(str) {
	case 0:
		return p.phone
	case 1:
		p.phone = str[0]
		return ""
	}

	p.err = errors.New(TooManyArgs)
	return ""
}

func (p Person) GetBSON() (interface{}, error) {
	obj := struct {
		Name  string
		Phone string
	}{

		p.Name(),
		p.Phone(),
	}

	return obj, nil
}

func (p *Person) SetBSON(raw bson.Raw) error {
	var obj struct {
		Name  string
		Phone string
	}

	err := raw.Unmarshal(&obj)

	if err != nil {
		return err
	}
	p.Name(obj.Name)
	p.Phone(obj.Phone)

	return p.Error()
}
