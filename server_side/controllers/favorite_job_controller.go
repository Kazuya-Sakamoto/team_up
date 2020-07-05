package controllers

import (
	"app/server_side/models"
	"app/server_side/services"
	"encoding/json"
)

// FavoriteJobController ...
type FavoriteJobController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *FavoriteJobController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create FavoriteJob
// @Param	body		body 	models.FavoriteJob	true		"body for FavoriteJob content"
// @Success 201 {int} models.FavoriteJob
// @Failure 500 body is empty
// @router / [post]
func (c *FavoriteJobController) Post() {
	defer c.HandlePanic()
	var favoriteJob models.FavoriteJob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &favoriteJob)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(services.PostFavoriteJob(favoriteJob))
	c.ServeJSON()
}

// Get FavoriteJob
// @Title Get
// @Description get FavoriteJob by FavoriteJobID
// @Param	FavoriteJobID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.FavoriteJob
// @Failure 500 :FavoriteJobID is empty
// @router /:FavoriteJobID [get]
func (c *FavoriteJobController) Get() {
	defer c.HandlePanic()
	favoriteJobID, err := c.GetInt64(":FavoriteJobID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetFavoriteJob(int64(favoriteJobID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get FavoriteJob
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.FavoriteJob
// @Failure 500
// @router / [get]
func (c *FavoriteJobController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	var userID int64
	c.Ctx.Input.Bind(&userID, "user_id")

	c.getHandle(models.GetAllFavoriteJobs(limit, offset, userID))
	c.ServeJSON()
}

// // Put ...
// // @Title Put
// // @Description update the FavoriteJob
// // @Param	FavoriteJobID		path 	string	true		"The FavoriteJobID you want to update"
// // @Param	body		body 	models.FavoriteJob	true		"body for FavoriteJob content"
// // @Success 200 {object} models.FavoriteJob
// // @Failure 500 :FavoriteJobID is not int
// // @router /:FavoriteJobID [put]
// func (c *FavoriteJobController) Put() {
// 	defer c.HandlePanic()
// 	favoriteJobID, err := c.GetInt64(":FavoriteJobID")
// 	if err != nil {
// 		c.parseErrorHandle(err)
// 	} else {
// 		var favoriteJob models.FavoriteJob
// 		err = json.Unmarshal(c.Ctx.Input.RequestBody, &favoriteJob)
// 		if err != nil {
// 			c.unmarshalErrorHandle(err)
// 		} else {
// 			c.putHandle(nil, models.UpdateFavoriteJob(int64(favoriteJobID), &favoriteJob))
// 		}
// 	}
// 	c.ServeJSON()
// }

//Delete FavoriteJob
// @Title Delete
// @Description delete the FavoriteJob
// @Param  FavoriteJobID        path    string  true        "The FavoriteJobID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 FavoriteJobID is empty
// @router / [delete]
func (c *FavoriteJobController) Delete() {
	defer c.HandlePanic()
	var favoriteJob models.FavoriteJob
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &favoriteJob)
	if err != nil {
		c.unmarshalErrorHandle(err)
	} else {
		c.deleteHandle(nil, services.DeleteFavoriteJobWithUserIDAndJobID(favoriteJob))
	}

	c.ServeJSON()
}
