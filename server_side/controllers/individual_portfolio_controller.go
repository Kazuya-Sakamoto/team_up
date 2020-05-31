package controllers

import (
	"app/server_side/models"
	"encoding/json"
)

// IndividualPortfolioController ...
type IndividualPortfolioController struct {
	RequiredLoginController
}

// URLMapping ...
func (c *IndividualPortfolioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create IndividualPortfolio
// @Param	body		body 	models.IndividualPortfolio	true		"body for IndividualPortfolio content"
// @Success 201 {int} models.IndividualPortfolio
// @Failure 500 body is empty
// @router / [post]
func (c *IndividualPortfolioController) Post() {
	defer c.HandlePanic()
	var individualPortfolio models.IndividualPortfolio
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &individualPortfolio)
	if err != nil {
		c.unmarshalErrorHandle(err)
		c.ServeJSON()
		return
	}
	c.postHandle(models.CreateIndividualPortfolio(individualPortfolio))
	c.ServeJSON()
}

// Get IndividualPortfolio
// @Title Get
// @Description get IndividualPortfolio by IndividualPortfolioID
// @Param	IndividualPortfolioID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.IndividualPortfolio
// @Failure 500 :IndividualPortfolioID is empty
// @router /:IndividualPortfolioID [get]
func (c *IndividualPortfolioController) Get() {
	defer c.HandlePanic()
	individualPortfolioID, err := c.GetInt64(":IndividualPortfolioID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.getHandle(models.GetIndividualPortfolio(int64(individualPortfolioID)))
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get IndividualPortfolio
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.IndividualPortfolio
// @Failure 500
// @router / [get]
func (c *IndividualPortfolioController) GetAll() {
	defer c.HandlePanic()

	var limit, offset int64
	c.Ctx.Input.Bind(&limit, "limit")
	c.Ctx.Input.Bind(&offset, "offset")

	c.getHandle(models.GetAllIndividualPortfolios(limit, offset))
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the IndividualPortfolio
// @Param	IndividualPortfolioID		path 	string	true		"The IndividualPortfolioID you want to update"
// @Param	body		body 	models.IndividualPortfolio	true		"body for IndividualPortfolio content"
// @Success 200 {object} models.IndividualPortfolio
// @Failure 500 :IndividualPortfolioID is not int
// @router /:IndividualPortfolioID [put]
func (c *IndividualPortfolioController) Put() {
	defer c.HandlePanic()
	individualPortfolioID, err := c.GetInt64(":IndividualPortfolioID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		var individualPortfolio models.IndividualPortfolio
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &individualPortfolio)
		if err != nil {
			c.unmarshalErrorHandle(err)
		} else {
			c.putHandle(nil, models.UpdateIndividualPortfolio(int64(individualPortfolioID), &individualPortfolio))
		}
	}
	c.ServeJSON()
}

//Delete IndividualPortfolio
// @Title Delete
// @Description delete the IndividualPortfolio
// @Param  IndividualPortfolioID        path    string  true        "The IndividualPortfolioID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 IndividualPortfolioID is empty
// @router /:IndividualPortfolioID [delete]
func (c *IndividualPortfolioController) Delete() {
	defer c.HandlePanic()
	individualPortfolioID, err := c.GetInt64(":IndividualPortfolioID")
	if err != nil {
		c.parseErrorHandle(err)
	} else {
		c.deleteHandle(nil, models.DeleteIndividualPortfolio(int64(individualPortfolioID)))
	}
	c.ServeJSON()
}
