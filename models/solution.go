package models

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

/**
{
	"solution_id": "yyyy",
	"task_id": "xxxx",
	"created_at": "20160812",
	"content" : "宝贝讲一段故事",
	"images" : ["http://xx.xx.xx/xxx","http://yy.yy.yy/yyy"],
	"videos" : ["rtmp://xx.xx.xx/xxx","rtmp://yy.yy.yy/yyy"],
	"audios" : ["http://xx.xx.xx/xxx","http://yy.yy.yy/yyy"],
	"praises" : [
		{
            		"role": "teacher",
            		"name": "刘德华",
           		"id": "123456"
		}, {
			"role": "parent",
		    	"name": "张惠妹",
		    	"id": "123221"
		}
	]
	"comment" : "评论"
	"flowers" : 5
	"class_name": "小红花版",
	"class_id": "123456",
	"parent_name": "刘德华",
	"parent_id": "1234",
	"child_name": "刘向蕙",
	"child_id": "1234"
}
*/

//用于存储
type Solution struct {
	SolutionID int       `bson:"solution_id"`
	TaskID     int       `bson:"task_id"`
	CreatedAt  time.Time `bson:"created_at"`
	Content    string    `bson:"content"`
	Images     []string  `bson:"images"`
	Videos     []string  `bson:"videos"`
	Audios     []string  `bson:"audios"`
	Praises    []Praise  `bson:"praises"`
	Comment    string    `bson:"comment"`
	Flowers    int       `bson:"flowers"`
	ClassName  string    `bson:"class_name"`
	ClassID    int       `bson:"class_id"`
	ParentName string    `bson:"parent_name"`
	ParentID   int       `bson:"parent_id"`
	ChildName  string    `bson:"child_name"`
	ChildID    int       `bson:"child_id"`
}

//返回某个方案细节
type SolutionReturnDetail struct {
	SolutionID int       `json:"solution_id"`
	TaskID     int       `json:"task_id"`
	CreatedAt  time.Time `json:"created_at"`
	Content    string    `json:"content"`
	Images     []string  `json:"images"`
	Videos     []string  `json:"videos"`
	Audios     []string  `json:"audios"`
	Praises    []Praise  `json:"praises"`
	Comment    string    `json:"comment"`
	Flowers    int       `json:"flowers"`
	ClassName  string    `json:"class_name"`
	ClassID    int       `json:"class_id"`
	ParentName string    `json:"parent_name"`
	ParentID   int       `json:"parent_id"`
	ChildName  string    `json:"child_name"`
	ChildID    int       `json:"child_id"`
}

//返回某个方案梗概
type SolutionReturnGist struct {
	SolutionID int       `json:"solution_id"`
	TaskID     int       `json:"task_id"`
	CreatedAt  time.Time `json:"created_at"`
	Content    string    `json:"content"`
	Images     []string  `json:"images"`
	Videos     []string  `json:"videos"`
	Audios     []string  `json:"audios"`
}

//用户提交
type SolutionCommit struct {
	TaskID     int      `json:"task_id"`
	Content    string   `json:"content"`
	Images     []string `json:"images"`
	Videos     []string `json:"videos"`
	Audios     []string `json:"audios"`
	ClassName  string   `json:"class_name"`
	ClassID    int      `json:"class_id"`
	ParentName string   `json:"parent_name"`
	ParentID   int      `json:"parent_id"`
	ChildName  string   `json:"child_name"`
	ChildID    int      `json:"child_id"`
}

type Praise struct {
	Role string `bson:"role" json:"role"`
	Name string `bson:"name" json:"name"`
	ID   int    `bson:"id" json:"id"`
}

type Comment struct {
	Content string `json:"content"`
}

type Flowers struct {
	Number int `json:"number"`
}

func init() {

}

//添加家庭作业方案
func AddSolution(m *SolutionCommit) (err error) {
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")
	solution_id, err := GetAutoIncreaseId(mgoSession, "solution_id")

	if len(m.Content) == 0 && len(m.Images) == 0 && len(m.Videos) == 0 && len(m.Audios) == 0 {
		return errors.New("COMMIT DATA ERROR!")
	}

	if err == nil {
		solution := Solution{}
		solution.SolutionID = solution_id
		solution.TaskID = m.TaskID
		solution.CreatedAt = time.Now()
		solution.Content = m.Content
		solution.Images = m.Images
		solution.Videos = m.Videos
		solution.Audios = m.Audios
		solution.ClassName = m.ClassName
		solution.ClassID = m.ClassID
		solution.ParentName = m.ParentName
		solution.ParentID = m.ParentID
		solution.ChildName = m.ChildName
		solution.ChildID = m.ChildID

		err = collection.Insert(solution)

		// 更新作业完成标记
		if err == nil {
			UpdateStatisticChildSolutionFinish(mgoSession, m.TaskID, m.ChildID)
			UpdateStatisticChildSolutionId(mgoSession, m.TaskID, m.ChildID, solution_id)
		}
	}

	return
}

//获取某个家庭作业方案
func GetSolutionById(solution_id int64) (v *SolutionReturnDetail, err error) {
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")
	solution := Solution{}
	err = collection.Find(bson.M{"solution_id": solution_id}).One(&solution)

	if err == nil {
		v = &SolutionReturnDetail{}
		v.SolutionID = solution.SolutionID
		v.TaskID = solution.TaskID
		v.CreatedAt = solution.CreatedAt
		v.Content = solution.Content
		v.Images = solution.Images
		v.Videos = solution.Videos
		v.Audios = solution.Audios

		v.Praises = solution.Praises
		v.Comment = solution.Comment
		v.Flowers = solution.Flowers

		v.ClassName = solution.ClassName
		v.ClassID = solution.ClassID
		v.ParentName = solution.ParentName
		v.ParentID = solution.ParentID
		v.ChildName = solution.ChildName
		v.ChildID = solution.ChildID
		return
	}
	return
}

//获取所有家庭作业方案
func GetAllSolution(task_id int, the_due_time int64) (ml []SolutionReturnGist, err error) {
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")
	solutions := []Solution{}
	if the_due_time != 0 {
		due_time := time.Unix(the_due_time, 0)
		fmt.Println("timestamp:", the_due_time)
		err = collection.Find(bson.M{"task_id": task_id, "created_at": bson.M{"$lt": due_time}}).Limit(10).Sort("-created_at").All(&solutions)
	} else {
		err = collection.Find(bson.M{"task_id": task_id}).Sort("-created_at").All(&solutions)
	}
	if err == nil {
		if len(solutions) == 0 {
			err = errors.New("NO EXIST")
			return
		}
		for _, solution := range solutions {
			srg := SolutionReturnGist{}
			srg.SolutionID = solution.SolutionID
			srg.TaskID = solution.TaskID
			srg.CreatedAt = solution.CreatedAt
			srg.Content = solution.Content
			srg.Images = solution.Images
			srg.Videos = solution.Videos
			srg.Audios = solution.Audios
			ml = append(ml, srg)
		}
		return
	}
	return
}

//删除
func DeleteSolution(solution_id int64) (err error) {
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")
	err = collection.Remove(bson.M{"solution_id": solution_id})
	return
}

// 更新
func UpdateSolutionById(solution_id int64, m *SolutionCommit) (err error) {
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")

	if len(m.Content) == 0 || len(m.Images) == 0 || len(m.Videos) == 0 || len(m.Audios) == 0 {
		return errors.New("COMMIT DATA ERROR!")
	}

	solution := Solution{}

	solution.TaskID = m.TaskID
	solution.CreatedAt = time.Now()
	solution.Content = m.Content
	solution.Images = m.Images
	solution.Videos = m.Videos
	solution.Audios = m.Audios
	solution.ClassName = m.ClassName
	solution.ClassID = m.ClassID
	solution.ParentName = m.ParentName
	solution.ParentID = m.ParentID
	solution.ChildName = m.ChildName
	solution.ChildID = m.ChildID

	err = collection.Update(bson.M{"solution_id": solution_id}, solution)

	return
}

// 评论
func UpdateSolutionCommentById(solution_id int64, m *Comment) (err error) {
	fmt.Println("Comment:" + m.Content)
	if len(m.Content) == 0 {
		return errors.New("COMMIT DATA ERROR!")
	}
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")
	err = collection.Update(bson.M{"solution_id": solution_id}, bson.M{"$set": bson.M{"comment": m.Content}})
	return
}

// 小红花
func UpdateSolutionflowersById(solution_id int64, m *Flowers) (err error) {
	if m.Number == 0 {
		return errors.New("COMMIT DATA ERROR!")
	}
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")
	err = collection.Update(bson.M{"solution_id": solution_id}, bson.M{"$set": bson.M{"flowers": m.Number}})
	return
}

// 点赞
func UpdateSolutionPraisesById(solution_id int64, m *Praise) (err error) {
	if len(m.Role) == 0 || len(m.Name) == 0 || m.ID == 0 {
		return errors.New("COMMIT DATA ERROR!")
	}
	mgoSession := GetMgoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB("ks_production").C("solutions")
	err = collection.Update(bson.M{"solution_id": solution_id}, bson.M{"$addToSet": bson.M{"praises": m}})
	if err == nil {
		solution := Solution{}
		err = collection.Find(bson.M{"solution_id": solution_id}).One(&solution)
		count := len(solution.Praises)
		UpdateStatisticChildSolutionPraiseNum(mgoSession, solution.TaskID, solution.ChildID, count)
	}
	return
}

// 查看某个家长是否完成作业
func CheckSolutionFinishStateByTaskId(task_id, child_id int64, mgoSession *mgo.Session) (count int, err error) {
	// defer mgoSession.Close()
	fmt.Println("task_id:", task_id, "&child_id:", child_id)
	collection := mgoSession.DB("ks_production").C("solutions")
	count, err = collection.Find(bson.M{"task_id": task_id, "child_id": child_id}).Count()
	fmt.Println(count)
	return
}
