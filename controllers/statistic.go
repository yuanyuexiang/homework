package controllers

import (
	//"encoding/json"
	"homework/models"
	"strconv"
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

// oprations for Statistic
type StatisticController struct {
	beego.Controller
}

func (c *StatisticController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// @Title Get
// @Description get Statistic by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Statistic
// @Failure 403 :id is empty
// @router /:id [get]
func (c *StatisticController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetStatisticById(id)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"GET FAIL " + err.Error(),nil)
	} else {
		c.Data["json"] = models.GetReturnData(0,"GET OK",v)
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the statitic comment
// @Param	id		path 	string	true		        "The id you want to update"
// @Param	body		body 			                "body for statistic content"
// @Success 200 {object} models.Solution
// @Failure 403 :id is not int
// @router /:id/comment [put]
func (c *StatisticController) PutStatisticComment() {
	fmt.Println("body:" + string(c.Ctx.Input.RequestBody))
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.StatisticComment{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1,"JSON DATA ERROR " + err.Error(),nil)
		c.ServeJSON()
		return
	}
	//v.Content = "asdasdasdasdasds"
	//fmt.Println("Comment:" + v.Content)
	if len(v.Comment) < 1 {
		c.Data["json"] = models.GetReturnData(-1,"COMMENT DATA SHORT ",nil)
		c.ServeJSON()
		return
	}

	if err = models.UpdateStatisticCommentById(id,&v); err == nil {
		c.Data["json"] = models.GetReturnData(0,"UPDATE SUCCESS",nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1,"UPDATE FAIL "+err.Error(),nil)
	}
	c.ServeJSON()
}
