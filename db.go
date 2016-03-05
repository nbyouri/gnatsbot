package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type netbsd_bugs struct {
	Number      string
	Category    string
	Synopsis    string
	Severity    string
	Priority    string
	Responsible string
	Class       string
	Originator  string
}

func main() {
	fmt.Println("------- RUNNING ------")
	rand.Seed(time.Now().UTC().UnixNano())

	orm, err := xorm.NewEngine("postgres", "user=pgsql dbname=netbsd_bugs sslmode=disable")
	if err != nil {
		log.Panic(err)
		return
	}
	defer orm.Close()

	// simple insert
	newbug := &netbsd_bugs{
		strconv.Itoa(rand.Intn(100)), "pkg", "failed packages", "diff", "test", "ok", "test", "me",
	}
	_, err = orm.Insert(newbug)
	if err != nil {
		fmt.Println(err)
	}

	// simple select
	var res = []netbsd_bugs{}
	orm.Find(&res)

	// print results
	for _, row := range res {
		row.toString()
	}

	// more complex select
	orm.Where("id = 15").Delete(&netbsd_bugs{})

	fmt.Println("------- END -------")
}

func (bug netbsd_bugs) toString() {
	fmt.Printf("---- NetBSD problem report ----\n"+
		">Number: %s\n>Category: %s\n>Synopsis %s\n>Severity %s\n"+
		">Priority: %s\n>Responsible: %s\n>Class: %s\n>Originator: %s\n",
		bug.Number, bug.Category, bug.Synopsis, bug.Severity, bug.Priority,
		bug.Responsible, bug.Class, bug.Originator)
}
