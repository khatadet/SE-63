package controllers
// ของแพร
import (
	"context"
	"fmt"
	"strconv"

	"github.com/PON/app/ent"
	"github.com/PON/app/ent/medicalrecordstaff"

	"github.com/gin-gonic/gin"
)

// MedicalrecordstaffController defines the struct for the medicalrecordstaff controller
type MedicalrecordstaffController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMedicalrecordstaff handles POST requests for adding medicalrecordstaff entities
// @Summary Create medicalrecordstaff
// @Description Create medicalrecordstaff
// @ID create-medicalrecordstaff
// @Accept   json
// @Produce  json
// @Param medicalrecordstaff body ent.Medicalrecordstaff true "Medicalrecordstaff entity"
// @Success 200 {object} ent.Medicalrecordstaff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecordstaffs [post]
func (ctl *MedicalrecordstaffController) CreateMedicalrecordstaff(c *gin.Context) {
	obj := ent.Medicalrecordstaff{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalrecordstaff binding failed",
		})
		return
	}

	u, err := ctl.client.Medicalrecordstaff.
		Create().
		//SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetMedicalrecordstaff handles GET requests to retrieve a medicalrecordstaff entity
// @Summary Get a medicalrecordstaff entity by ID
// @Description get medicalrecordstaff by ID
// @ID get-medicalrecordstaff
// @Produce  json
// @Param id path int true "Medicalrecordstaff ID"
// @Success 200 {object} ent.Medicalrecordstaff
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecordstaffs/{id} [get]
func (ctl *MedicalrecordstaffController) GetMedicalrecordstaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Medicalrecordstaff.
		Query().
		Where(medicalrecordstaff.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListMedicalrecordstaff handles request to get a list of medicalrecordstaff entities
// @Summary List medicalrecordstaff entities
// @Description list medicalrecordstaff entities
// @ID list-medicalrecordstaff
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medicalrecordstaff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecordstaffs [get]
func (ctl *MedicalrecordstaffController) ListMedicalrecordstaff(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	medicalrecordstaffs, err := ctl.client.Medicalrecordstaff.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, medicalrecordstaffs)
}

// DeleteMedicalrecordstaff handles DELETE requests to delete a medicalrecordstaff entity
// @Summary Delete a medicalrecordstaff entity by ID
// @Description get medicalrecordstaff by ID
// @ID delete-medicalrecordstaff
// @Produce  json
// @Param id path int true "Medicalrecordstaff ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecordstaffs/{id} [delete]
func (ctl *MedicalrecordstaffController) DeleteMedicalrecordstaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Medicalrecordstaff.
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

// UpdateMedicalrecordstaff handles PUT requests to update a medicalrecordstaff entity
// @Summary Update a medicalrecordstaff entity by ID
// @Description update medicalrecordstaff by ID
// @ID update-medicalrecordstaff
// @Accept   json
// @Produce  json
// @Param id path int true "Medicalrecordstaff ID"
// @Param medicalrecordstaff body ent.Medicalrecordstaff true "Medicalrecordstaff entity"
// @Success 200 {object} ent.Medicalrecordstaff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecordstaffs/{id} [put]
func (ctl *MedicalrecordstaffController) UpdateMedicalrecordstaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Medicalrecordstaff{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalrecordstaff binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Medicalrecordstaff.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}
// NewMedicalrecordstaffController creates and registers handles for the medicalrecordstaff controller
func NewMedicalrecordstaffController(router gin.IRouter, client *ent.Client) *MedicalrecordstaffController {
	uc := &MedicalrecordstaffController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
 
 // InitMedicalrecordstaffController registers routes to the main engine
 func (ctl *MedicalrecordstaffController) register() {
	medicalrecordstaffs := ctl.router.Group("/medicalrecordstaffs")
 
	medicalrecordstaffs.GET("", ctl.ListMedicalrecordstaff)
 
	// CRUD
	medicalrecordstaffs.POST("", ctl.CreateMedicalrecordstaff)
	medicalrecordstaffs.GET(":id", ctl.GetMedicalrecordstaff)
	medicalrecordstaffs.PUT(":id", ctl.UpdateMedicalrecordstaff)
	medicalrecordstaffs.DELETE(":id", ctl.DeleteMedicalrecordstaff)
 }
 
