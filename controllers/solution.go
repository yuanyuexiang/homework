package controllers

import (
	"encoding/json"
	"fmt"
	"homework/models"
	"strconv"
	"github.com/astaxie/beego"
)

// oprations for Solution
type SolutionController struct {
	beego.Controller
}

func (c *SolutionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *SolutionController) Prepare() {
	/*
		认证代码	
	*/
}

// @Title Post
// @Description create Solution
// @Param	body		body 	models.Solution	true		"body for Solution content"
// @Success 201 {int} models.Solution
// @Failure 403 body is empty
// @router / [post]
func (c *SolutionController) Post() {
	var v models.SolutionCommit
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"JSON DATA ERROR " + err.Error(),nil)
		c.ServeJSON()
		return
	}
	fmt.Println("body:" + string(c.Ctx.Input.RequestBody))
	if err = models.AddSolution(&v); err == nil {
		//c.Ctx.Output.SetStatus(201)
		c.Data["json"] = models.GetReturnData(0,"OK",nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1,"FAIL " + err.Error(),nil)
	}
	c.ServeJSON()
}

// @Title Get
// @Description get Solution by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Solution
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SolutionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetSolutionById(id)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"GET FAIL " + err.Error(),nil)
	} else {
		c.Data["json"] = models.GetReturnData(0,"GET OK",v)
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Solution
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Solution
// @Failure 403
// @router / [get]
func (c *SolutionController) GetAll() {
	task_id,err := c.GetInt("task_id")
	due_time, err2 := c.GetInt64("due_time")
	if err != nil && err2 != nil {
		c.Data["json"] = models.GetReturnData(-1,"PRARM ERROR " + err.Error() + " AND " + err2.Error(),nil)
		c.ServeJSON()
		return
	}

	v, err := models.GetAllSolution(task_id,due_time)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"GET FAIL " + err.Error(),nil)
	} else {
		c.Data["json"] = models.GetReturnData(0,"GET OK",v)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Solution
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Solution	true		"body for Solution content"
// @Success 200 {object} models.Solution
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SolutionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.SolutionCommit{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"JSON DATA ERROR " + err.Error(),nil)
		c.ServeJSON()
		return
	}
	if err = models.UpdateSolutionById(id,&v); err == nil {
		c.Data["json"] = models.GetReturnData(0,"UPDATE SUCCESS",nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1,"UPDATE FAIL "+err.Error(),nil)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Solution
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Solution	true		"body for Solution content"
// @Success 200 {object} models.Solution
// @Failure 403 :id is not int
// @router /:id/comment [put]
func (c *SolutionController) PutComment() {
	fmt.Println("body:" + string(c.Ctx.Input.RequestBody))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Comment{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"JSON DATA ERROR " + err.Error(),nil)
		c.ServeJSON()
		return
	}
	//v.Content = "asdasdasdasdasds"
	//fmt.Println("Comment:" + v.Content)
	if len(v.Content) < 1 {
		c.Data["json"] = models.GetReturnData(-1,"COMMENT DATA SHORT ",nil)
		c.ServeJSON()
		return
	}

	if err = models.UpdateSolutionCommentById(id,&v); err == nil {
		c.Data["json"] = models.GetReturnData(0,"UPDATE SUCCESS",nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1,"UPDATE FAIL "+err.Error(),nil)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Solution
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Solution	true		"body for Solution content"
// @Success 200 {object} models.Solution
// @Failure 403 :id is not int
// @router /:id/flowers [put]
func (c *SolutionController) PutFlowers() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Flowers{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"JSON DATA ERROR " + err.Error(),nil)
		c.ServeJSON()
		return
	}

	if v.Number <= 0 {
		c.Data["json"] = models.GetReturnData(-1,"FLOWER DATA SHORT ",nil)
		c.ServeJSON()
		return
	}

	if err := models.UpdateSolutionflowersById(id,&v); err == nil {
		c.Data["json"] = models.GetReturnData(0,"UPDATE SUCCESS",nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1,"UPDATE FAIL "+err.Error(),nil)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Solution
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Solution	true		"body for Solution content"
// @Success 200 {object} models.Solution
// @Failure 403 :id is not int
// @router /:id/praises [put]
func (c *SolutionController) PutPraises() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Praise{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"JSON DATA ERROR " + err.Error(),nil)
		c.ServeJSON()
		return
	}
	
	if err := models.UpdateSolutionPraisesById(id,&v); err == nil {
		c.Data["json"] = models.GetReturnData(0,"UPDATE SUCCESS",nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1,"UPDATE FAIL "+err.Error(),nil)
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Solution
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SolutionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteSolution(id); err == nil {
		c.Data["json"] = models.GetReturnData(0,"DELETE SUCCESS",nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1,"DELETE FAIL "+err.Error(),nil)
	}
	c.ServeJSON()
}
