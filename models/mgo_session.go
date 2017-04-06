package models

import (
        "gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

var (
	mgoSession *mgo.Session
)

func init() {

	dba := beego.AppConfig.String("DataBaseAddress")
	var err error
	mgoSession, err = mgo.Dial(dba)
	if err != nil {
	    panic(err)
	}
}

func GetMgoSession() *mgo.Session {
    if mgoSession == nil {
	    dba := beego.AppConfig.String("DataBaseAddress")
	    var err error
	    mgoSession, err = mgo.Dial(dba)
	    if err != nil {
		    panic(err)
	    }
    }
    return mgoSession.Clone()
}
