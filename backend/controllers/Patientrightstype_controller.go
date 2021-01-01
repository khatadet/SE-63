package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/PON/app/ent"
   "github.com/PON/app/ent/patientrightstype"
   "github.com/gin-gonic/gin"

   "github.com/PON/app/ent/abilitypatientrights"
)
 
// PatientrightstypeController defines the struct for the patientrightstype controller
type PatientrightstypeController struct {
   client *ent.Client
   router gin.IRouter
}

// Patientrightstype defines the struct for the Patientrightstype
type Patientrightstype struct {
	Permission   		string
	PermissionArea		string
	Responsible			string
	Abilitypatientrights 	int
	
	
	
}

// CreatePatientrightstype handles POST requests for adding patientrightstype entities
// @Summary Create patientrightstype
// @Description Create patientrightstype
// @ID create-patientrightstype
// @Accept   json
// @Produce  json
// @Param patientrightstype body ent.Patientrightstype true "Patientrightstype entity"
// @Success 200 {object} ent.Patientrightstype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightstypes [post]
func (ctl *PatientrightstypeController) CreatePatientrightstype(c *gin.Context) {
	obj := Patientrightstype{}
	if err := c.ShouldBind(&obj); 
	err != nil {
		c.JSON(400, gin.H{
			"error": "patientrightstype binding failed",
		})
		return
	}


	Abilitypatientrights, err := ctl.client.Abilitypatientrights.
		Query().
		Where(abilitypatientrights.IDEQ(int(obj.Abilitypatientrights))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Abilitypatientrights not found",
		})
		return
	}
  
	u, err := ctl.client.Patientrightstype.
		Create().
		SetPermission(obj.Permission).
		SetPermissionArea(obj.PermissionArea).
		SetResponsible(obj.Responsible).
		SetPatientrightstypeAbilitypatientrights(Abilitypatientrights).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 // GetPatientrightstype handles GET requests to retrieve a patientrightstype entity
// @Summary Get a patientrightstype entity by ID
// @Description get patientrightstype by ID
// @ID get-patientrightstype
// @Produce  json
// @Param id path int true "Patientrightstype ID"
// @Success 200 {object} ent.Patientrightstype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightstypes/{id} [get]
func (ctl *PatientrightstypeController) GetPatientrightstype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	
  
	u, err := ctl.client.Patientrightstype.
		Query().
		Where(patientrightstype.IDEQ(int(id))).
		WithPatientrightstypeAbilitypatientrights().
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 // ListPatientrightstype handles request to get a list of patientrightstype entities
// @Summary List patientrightstype entities
// @Description list patientrightstype entities
// @ID list-patientrightstype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patientrightstype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightstypes [get]
func (ctl *PatientrightstypeController) ListPatientrightstype(c *gin.Context) {
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
  
	patientrightstypes, err := ctl.client.Patientrightstype.
		Query().
		
		WithPatientrightstypeAbilitypatientrights().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, patientrightstypes)
 }
 // DeletePatientrightstype handles DELETE requests to delete a patientrightstype entity
// @Summary Delete a patientrightstype entity by ID
// @Description get patientrightstype by ID
// @ID delete-patientrightstype
// @Produce  json
// @Param id path int true "Patientrightstype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightstypes/{id} [delete]
func (ctl *PatientrightstypeController) DeletePatientrightstype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Patientrightstype.
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
  
// UpdatePatientrightstype handles PUT requests to update a patientrightstype entity
// @Summary Update a patientrightstype entity by ID
// @Description update patientrightstype by ID
// @ID update-patientrightstype
// @Accept   json
// @Produce  json
// @Param id path int true "Patientrightstype ID"
// @Param patientrightstype body ent.Patientrightstype true "Patientrightstype entity"
// @Success 200 {object} ent.Patientrightstype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightstypes/{id} [put]
func (ctl *PatientrightstypeController) UpdatePatientrightstype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Patientrightstype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patientrightstype binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Patientrightstype.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }
// NewPatientrightstypeController creates and registers handles for the patientrightstype controller
func NewPatientrightstypeController(router gin.IRouter, client *ent.Client) *PatientrightstypeController {
	uc := &PatientrightstypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitPatientrightstypeController registers routes to the main engine
 func (ctl *PatientrightstypeController) register() {
	patientrightstypes := ctl.router.Group("/patientrightstypes")
  
	patientrightstypes.GET("", ctl.ListPatientrightstype)
  
	// CRUD
	patientrightstypes.POST("", ctl.CreatePatientrightstype)
	patientrightstypes.GET(":id", ctl.GetPatientrightstype)
	patientrightstypes.PUT(":id", ctl.UpdatePatientrightstype)
	patientrightstypes.DELETE(":id", ctl.DeletePatientrightstype)
 }
   