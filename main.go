package main

import (
	"fmt"
	"log"

	"github.com/mmirolim/exp-with-set-get/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	var p models.Person
	p.Name = "Alexsandr"
	p.Phone = "+998937777777"
	p.Created("Mirolim")

	err = c.Insert(&p)
	if err != nil {
		log.Fatal(err)
	}

	result := models.Person{}
	err = c.Find(bson.M{"base.createdby": "Mirolim"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("person %+v \n", result)
}
