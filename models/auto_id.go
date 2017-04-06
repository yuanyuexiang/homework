package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
	"errors"
	"fmt"
	"time"
)

type SystemConfig struct
{
	Id_  bson.ObjectId        `bson:"_id"`
	Name string               `bson:"name"`
	Nval int                  `bson:"nval"`
	Updated_at time.Time      `bson:"updated_at"`
	Created_at time.Time      `bson:"created_at"`
}


func GetAutoIncreaseId(mgoSession *mgo.Session, field_name string) (int,error)  {

	// get db
	dbn := beego.AppConfig.String("DataBaseName")
	if len(dbn) == 0 {
		panic(errors.New("No DataBaseName"))
	}

	db := mgoSession.DB(dbn)

	// get collection
	c := db.C("system_configs")

	count,err := c.Find(bson.M{"name": field_name}).Count()
	if err != nil {
		return -1,err
	}

	if count == 0 {
		fmt.Println("not find this field")
		// insert one record
		sc := SystemConfig{}
		sc.Id_ =  bson.NewObjectId()
		sc.Name = field_name
		sc.Nval = 1
		sc.Created_at = time.Now()
		sc.Updated_at = time.Now()
		insert_err := c.Insert(&sc)
		if insert_err !=nil{
			fmt.Println("insert system_config record failed")
			return -1,insert_err
		}
		fmt.Println(sc)
		return sc.Nval,nil
	}

	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"nval": 1}},
		Upsert: true,
		ReturnNew: true,
	}

	doc := struct{ Nval int }{}

	_,err = c.Find(bson.M{"name": field_name}).Apply(change, &doc)

	if err!=nil {
		return -1,err
	}

	return doc.Nval,nil;
}