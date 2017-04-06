package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"homework/models"
	"strconv"
	//"time"
	"fmt"
)

// oprations for Task
type TaskController struct {
	beego.Controller
}

func (c *TaskController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("Index", c.Index)
	c.Mapping("Show", c.Show)
	c.Mapping("Destory", c.Destory)
}

// @router / [get]
func (this *TaskController) Index() {
	class_id, err := strconv.Atoi(this.GetString("class_id"))
	fmt.Println(class_id)
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}

	due_time, err := this.GetInt64("due_time")
	fmt.Println(due_time)
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	ask_query_params := models.TaskQueryParams{
		ClassId: class_id,
		DueTime: due_time,
	}

	tasks, err := models.FindAllTasksByParams(&ask_query_params)
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

// @router / [post]
func (this *TaskController) Create() {
	var task models.Task
	json.Unmarshal(this.Ctx.Input.RequestBody, &task)
	err := models.CreateNewTask(&task)
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "publish failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.GetReturnData(0, "publish success !", nil)
	this.ServeJSON()
}

// @router /:id [delete]
func (this *TaskController) Destory() {
	task_id, err := strconv.Atoi(this.GetString(":id"))
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "DELETE failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	err = models.DestoryTaskByLogic(task_id)
	if err != nil {
		this.Data["json"] = models.GetReturnData(-1, "DELETE failed "+err.Error(), nil)
		this.ServeJSON()
		return
	}
	this.Data["json"] = models.GetReturnData(0, "DELETE success !", nil)
	this.ServeJSON()
}
