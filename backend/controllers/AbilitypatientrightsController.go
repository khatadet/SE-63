package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/PON/app/ent"
   "github.com/PON/app/ent/abilitypatientrights"
   "github.com/gin-gonic/gin"
)
 
// AbilitypatientrightsController defines the struct for the abilitypatientrights controller
type AbilitypatientrightsController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateAbilitypatientrights handles POST requests for adding abilitypatientrights entities
// @Summary Create abilitypatientrights
// @Description Create abilitypatientrights
// @ID create-abilitypatientrights
// @Accept   json
// @Produce  json
// @Param abilitypatientrights body ent.Abilitypatientrights true "Abilitypatientrights entity"
// @Success 200 {object} ent.Abilitypatientrights
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /abilitypatientrightss [post]
func (ctl *AbilitypatientrightsController) CreateAbilitypatientrights(c *gin.Context) {
	obj := ent.Abilitypatientrights{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "abilitypatientrights binding failed",
		})
		return
	}
  
	u, err := ctl.client.Abilitypatientrights.
		Create().
		SetOperative(int(obj.Operative)).
		SetMedicalSupplies(int(obj.MedicalSupplies)).
		SetExamine(int(obj.Examine)).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 // GetAbilitypatientrights handles GET requests to retrieve a abilitypatientrights entity
// @Summary Get a abilitypatientrights entity by ID
// @Description get abilitypatientrights by ID
// @ID get-abilitypatientrights
// @Produce  json
// @Param id path int true "Abilitypatientrights ID"
// @Success 200 {object} ent.Abilitypatientrights
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /abilitypatientrightss/{id} [get]
func (ctl *AbilitypatientrightsController) GetAbilitypatientrights(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Abilitypatientrights.
		Query().
		Where(abilitypatientrights.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 // ListAbilitypatientrights handles request to get a list of abilitypatientrights entities
// @Summary List abilitypatientrights entities
// @Description list abilitypatientrights entities
// @ID list-abilitypatientrights
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Abilitypatientrights
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /abilitypatientrightss [get]
func (ctl *AbilitypatientrightsController) ListAbilitypatientrights(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	abilitypatientrightss, err := ctl.client.Abilitypatientrights.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, abilitypatientrightss)
 }
 // DeleteAbilitypatientrights handles DELETE requests to delete a abilitypatientrights entity
// @Summary Delete a abilitypatientrights entity by ID
// @Description get abilitypatientrights by ID
// @ID delete-abilitypatientrights
// @Produce  json
// @Param id path int true "Abilitypatientrights ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /abilitypatientrightss/{id} [delete]
func (ctl *AbilitypatientrightsController) DeleteAbilitypatientrights(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Abilitypatientrights.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
 }
  
// UpdateAbilitypatientrights handles PUT requests to update a abilitypatientrights entity
// @Summary Update a abilitypatientrights entity by ID
// @Description update abilitypatientrights by ID
// @ID update-abilitypatientrights
// @Accept   json
// @Produce  json
// @Param id path int true "Abilitypatientrights ID"
// @Param abilitypatientrights body ent.Abilitypatientrights true "Abilitypatientrights entity"
// @Success 200 {object} ent.Abilitypatientrights
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /abilitypatientrightss/{id} [put]
func (ctl *AbilitypatientrightsController) UpdateAbilitypatientrights(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Abilitypatientrights{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "abilitypatientrights binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Abilitypatientrights.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }
// NewAbilitypatientrightsController creates and registers handles for the abilitypatientrights controller
func NewAbilitypatientrightsController(router gin.IRouter, client *ent.Client) *AbilitypatientrightsController {
	uc := &AbilitypatientrightsController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitAbilitypatientrightsController registers routes to the main engine
 func (ctl *AbilitypatientrightsController) register() {
	abilitypatientrightss := ctl.router.Group("/abilitypatientrightss")
  
	abilitypatientrightss.GET("", ctl.ListAbilitypatientrights)
  
	// CRUD
	abilitypatientrightss.POST("", ctl.CreateAbilitypatientrights)
	abilitypatientrightss.GET(":id", ctl.GetAbilitypatientrights)
	abilitypatientrightss.PUT(":id", ctl.UpdateAbilitypatientrights)
	abilitypatientrightss.DELETE(":id", ctl.DeleteAbilitypatientrights)
 }
   