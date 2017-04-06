package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Task struct {
	TaskId            int      `bson:"task_id" json:"task_id"`
	Content           string   `bson:"content"  json:"content"`
	Images            []string `bson:"images"  json:"images"`
	Audios            []string `bson:"audios"  json:"audios"`
	ClassName         string   `bson:"class_name"  json:"class_name"`
	ClassId           int      `bson:"class_id"  json:"class_id,string"`
	TeacherName       string   `bson:"teacher_name"  json:"teacher_name"`
	TeacherId         int      `bson:"teacher_id"  json:"teacher_id,string"`
	AuthorityClassIds []int    `bson:"authority_class_ids"  json:"authority_class_ids"`
	// 0正常，1删除
	Status      int `bson:"status" json:"status"`
	CommonModel `bson:",inline"`
}

type TaskResult struct {
	TaskId    int       `bson:"task_id" json:"task_id"`
	Content   string    `bson:"content"  json:"content"`
	Images    []string  `bson:"images"  json:"images"`
	Audios    []string  `bson:"audios"  json:"audios"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type TaskParentResult struct {
	TaskId    int       `bson:"task_id" json:"task_id"`
	Content   string    `bson:"content"  json:"content"`
	Images    []string  `bson:"images"  json:"images"`
	Audios    []string  `bson:"audios"  json:"audios"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	Finished  bool      `json:"finished"`
}

type TaskQueryParams struct {
	ClassId int   `json:"class_id"`
	DueTime int64 `json:"due_time"`
}

func CreateNewTask(task *Task) error {
	mgs := GetMgoSession()
	defer mgs.Close()
	taskCollection := mgs.DB("ks_production").C("tasks")
	autoId, err := GetAutoIncreaseId(mgs, "task_id")
	if err != nil {
		return err
	}
	task.TaskId = autoId
	t := time.Now()
	task.CreatedAt = t
	task.UpdatedAt = t
	err = taskCollection.Insert(&task)
	if err == nil {
		AddStatistic(task.TaskId, task.AuthorityClassIds)
	}
	return err
}

func FindAllTasksByParams(query_params *TaskQueryParams) ([]TaskResult, error) {
	mgs := GetMgoSession()
	defer mgs.Close()
	var tasks []TaskResult
	taskCollection := mgs.DB("ks_production").C("tasks")
	due_time := time.Unix(query_params.DueTime, 0)
	err := taskCollection.Find(bson.M{"authority_class_ids": query_params.ClassId, "status": 0, "created_at": bson.M{"$lt": due_time}}).Limit(10).Sort("-created_at").All(&tasks)
	return tasks, err
}

func IsFinishedByUser(user_id, task_id int, mgoSession *mgo.Session) (b bool, err error) {
	count, err := CheckSolutionFinishStateByTaskId(int64(task_id), int64(user_id), mgoSession)
	b = (count > 0)
	return
}

func FindAllTasksByParentParams(user_id int, query_params *TaskQueryParams) (*[]TaskParentResult, error) {
	mgs := GetMgoSession()
	defer mgs.Close()
	taskResults := []TaskParentResult{}
	taskCollection := mgs.DB("ks_production").C("tasks")
	due_time := time.Unix(query_params.DueTime, 0)
	err := taskCollection.Find(bson.M{"authority_class_ids": query_params.ClassId, "status": 0, "created_at": bson.M{"$lt": due_time}}).Limit(10).Sort("-created_at").All(&taskResults)
	if err != nil {
		return &taskResults, err
	}
	for i, taskResult := range taskResults {
		finished, _ := IsFinishedByUser(user_id, taskResult.TaskId, mgs)
		taskResults[i].Finished = finished
	}
	return &taskResults, err
}

func GetTaskById(task_id int) (*Task, error) {
	mgs := GetMgoSession()
	defer mgs.Close()
	task := Task{}
	taskCollection := mgs.DB("ks_production").C("tasks")
	err := taskCollection.Find(bson.M{"task_id": task_id, "status": 0}).One(&task)
	return &task, err
}

func DestoryTaskByLogic(task_id int) error {
	mgs := GetMgoSession()
	defer mgs.Close()
	taskCollection := mgs.DB("ks_production").C("tasks")
	err := taskCollection.Update(bson.M{"task_id": task_id, "status": 0}, bson.M{"$set": bson.M{"status": 1}})
	return err
}
