package parent

import (
	"github.com/astaxie/beego"
	"homework/models"
	"strconv"
	//"time"
	//"fmt"
)

// oprations for Task
type TaskController struct {
	beego.Controller
}

type TaskQueryResult struct {
	models.Task
	Finished bool `json:"finished"`
}

func (c *TaskController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Show", c.Show)
}

// @router / [get]
func (this *TaskController) Index() {
	class_id, err := strconv.Atoi(this.GetString("class_id"))
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}

	due_time, err := this.GetInt64("due_time")
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	ask_query_params := models.TaskQueryParams{
		ClassId: class_id,
		DueTime: due_time,
	}

	user_id, err := strconv.Atoi(this.GetString("user_id"))
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	tasks, err := models.FindAllTasksByParentParams(user_id, &ask_query_params)
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = models.GetReturnData(0, "success !", tasks)
	this.ServeJSON()
}

// @router /:id [get]
func (this *TaskController) Show() {
	task_id, err := strconv.Atoi(this.GetString(":id"))
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	task, err := models.GetTaskById(task_id)
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.GetReturnData(0, "success !", task)
	this.ServeJSON()
}
