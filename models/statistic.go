package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"homework/utils"
)

type  ChildInfo struct {
	Child_id       int            `bson:"child_id" json:"child_id"`
	Child_name     string         `bson:"child_name" json:"child_name"`
	Finished       bool           `bson:"finished" json:"finished"`
	Solution_id    int            `bson:"solution_id" json:"solution_id"`
	Praise_number  int            `bson:"praise_number" json:"praise_number"`
}

// for save
type Statistic struct {
	Statistic_id    int            `bson:"statistic_id" json:"statistic_id"`
	Task_id         int            `bson:"task_id" json:"task_id"`
	Comment         string         `bson:"comment" json:"comment"`
	All_child       []ChildInfo    `bson:"all_child" json:"all_child"`

}

// for query
type  KidInfo struct {
	Id_            bson.ObjectId  `bson:"_id"`
	Name           string         `bson:"name"`
	Kid_id         int            `bson:"kid_id"`
}
// just for query
type ClaseInfo struct {
	Id_            bson.ObjectId     `bson:"_id"`
	Name           string            `bson:"name"`
	Clase_id       int               `bson:"clase_id"`
	Kid_ids        []bson.ObjectId   `bson:"kid_ids"`
}

type StatisticComment struct{
	Comment		string		`json:"comment"`
}


func init() {

}

//  add statistic for task model

func AddStatistic(task_id int,class_ids []int) (id int, err error) {
	utils.WriteLog(utils.LOG_INFO,"enter AddStatistic,task_id:",task_id,"  class_ids:",class_ids)
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("statistics")


	count,err := collection.Find(bson.M{"task_id": task_id}).Count()

	if err !=nil {
		utils.WriteLog(utils.LOG_ERROR,err.Error())
		return -1,err
	}

	if count>0 {
		s := Statistic{}
		err := collection.Find(bson.M{"task_id": task_id}).One(&s)
		if err == nil {
			utils.WriteLog(utils.LOG_INFO,"find statistic by task_id:",task_id,"return it : ",s)
		} else {
			utils.WriteLog(utils.LOG_ERROR,"find statistic by task_id:",task_id," error : ",err.Error())
		}

		return s.Statistic_id,err
	}

	// 查询班级
	clases := []ClaseInfo{}
	clase_c := mgoSession.DB("ks_production").C("clases")
	err = clase_c.Find(bson.M{"clase_id":bson.M{"$in": class_ids}}).All(&clases)

	if err != nil {
		utils.WriteLog(utils.LOG_ERROR,err.Error())
		return -1,err
	}

	if 0 == len(clases) {
		utils.WriteLog(utils.LOG_ERROR,"not find any class")
		return -1,err
	}

	childs := []ChildInfo{}
	kids_c := mgoSession.DB("ks_production").C("kids")
	for _, class_info := range clases {

		for _,kid_objectid:=range class_info.Kid_ids{
			// query kid
			k := KidInfo{}
			err = kids_c.Find(bson.M{"_id": kid_objectid}).One(&k)
			if err == nil {
				child := ChildInfo{}
				child.Child_id = k.Kid_id
				child.Child_name = k.Name
				childs = append(childs,child)
			}
		}
	}



	s := Statistic{}

	// use task_id for statistic id
	s.Statistic_id = task_id
	s.Comment = ""
	s.Task_id = task_id
	s.All_child = childs

	collection.Insert(&s)

	utils.WriteLog(utils.LOG_INFO,"add statistic : ",s)

	return task_id, err
}

// GetStatisticById retrieves Statistic by Id. Returns error if
// Id doesn't exist
func GetStatisticById(id int) (v *Statistic, err error) {
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("statistics")
	s := Statistic{}
	err = collection.Find(bson.M{"statistic_id": id}).One(&s)

	if err == nil{
		return &s,err
	}
	return nil, err
}

// 更新统计评论

func UpdateStatisticCommentById(id int,comment *StatisticComment) (err error){
	if len(comment.Comment) == 0 {
		return errors.New("COMMIT DATA ERROR!")
	}
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("statistics")
	err = collection.Update(bson.M{"statistic_id": id}, bson.M{"$set": bson.M{"comment": comment.Comment}})
	if nil != err {
		utils.WriteLog(utils.LOG_ERROR,err.Error())
	}
	return err
}

// 更新小孩作业完成情况,提供给solution模型调用

func UpdateStatisticChildSolutionFinish(mgoSession *mgo.Session, task_id int , child_id int) (err error){

	if mgoSession == nil {
		mgoSession = GetMgoSession()
	}

	collection := mgoSession.DB("ks_production").C("statistics")
	err = collection.Update(bson.M{"task_id": task_id,"all_child.child_id":child_id}, bson.M{"$set": bson.M{"all_child.$.finished": true}})
	if nil != err{
		utils.WriteLog(utils.LOG_ERROR,"update child finished error task_id:",task_id," child_id",child_id,err.Error())
	}
	return err
}

// 更新小孩作业完成的solution_id
func UpdateStatisticChildSolutionId(mgoSession *mgo.Session, task_id int , child_id int, solution_id int) (err error){

	if mgoSession == nil {
		mgoSession = GetMgoSession()
	}

	collection := mgoSession.DB("ks_production").C("statistics")
	err = collection.Update(bson.M{"task_id": task_id,"all_child.child_id":child_id}, bson.M{"$set": bson.M{"all_child.$.solution_id": solution_id}})
	if nil != err{
		utils.WriteLog(utils.LOG_ERROR,"update child solution_id error task_id:",task_id," child_id",child_id,err.Error())
	}
	return err
}

// 更新 点赞数
func UpdateStatisticChildSolutionPraiseNum(mgoSession *mgo.Session, task_id int , child_id int, praise_number int) (err error){

	if mgoSession == nil {
		mgoSession = GetMgoSession()
	}

	collection := mgoSession.DB("ks_production").C("statistics")
	err = collection.Update(bson.M{"task_id": task_id,"all_child.child_id":child_id}, bson.M{"$set": bson.M{"all_child.$.praise_number": praise_number}})
	if nil != err{
		utils.WriteLog(utils.LOG_ERROR,"update child praise_number error task_id:",task_id," child_id",child_id,err.Error())
	}
	return err
}

// DeleteStatistic deletes Statistic by Id and returns error if
// the record to be deleted doesn't exist
func DeleteStatistic(mgoSession *mgo.Session,task_id int) (err error) {
	/*if mgoSession == nil {
		mgoSession = GetMgoSession()
	}
	collection := mgoSession.DB("ks_production").C("statistics")
	err = collection.Remove(bson.M{"task_id": task_id})
	if nil != err{
		utils.WriteLog(utils.LOG_ERROR,"delete statistic by task_id:",task_id," failed", err.Error())
	}*/
	return
}
