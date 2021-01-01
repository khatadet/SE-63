package controllers
// ของแพร
import (
	"context"
	"fmt"
	"strconv"

	"github.com/PON/app/ent"
	"github.com/PON/app/ent/patientrecord"

	"github.com/gin-gonic/gin"
)

// PatientrecordController defines the struct for the patientrecord controller
type PatientrecordController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePatientrecord handles POST requests for adding patientrecord entities
// @Summary Create patientrecord
// @Description Create patientrecord
// @ID create-patientrecord
// @Accept   json
// @Produce  json
// @Param patientrecord body ent.Patientrecord true "Patientrecord entity"
// @Success 200 {object} ent.Patientrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecords [post]
func (ctl *PatientrecordController) CreatePatientrecord(c *gin.Context) {
	obj := ent.Patientrecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patientrecord binding failed",
		})
		return
	}

	u, err := ctl.client.Patientrecord.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetPatientrecord handles GET requests to retrieve a patientrecord entity
// @Summary Get a patientrecord entity by ID
// @Description get patientrecord by ID
// @ID get-patientrecord
// @Produce  json
// @Param id path int true "Patientrecord ID"
// @Success 200 {object} ent.Patientrecord
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecords/{id} [get]
func (ctl *PatientrecordController) GetPatientrecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Patientrecord.
		Query().
		Where(patientrecord.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListPatientrecord handles request to get a list of patientrecord entities
// @Summary List patientrecord entities
// @Description list patientrecord entities
// @ID list-patientrecord
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patientrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecords [get]
func (ctl *PatientrecordController) ListPatientrecord(c *gin.Context) {
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

	patientrecords, err := ctl.client.Patientrecord.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patientrecords)
}

// DeletePatientrecord handles DELETE requests to delete a patientrecord entity
// @Summary Delete a patientrecord entity by ID
// @Description get patientrecord by ID
// @ID delete-patientrecord
// @Produce  json
// @Param id path int true "Patientrecord ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecords/{id} [delete]
func (ctl *PatientrecordController) DeletePatientrecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Patientrecord.
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

// UpdatePatientrecord handles PUT requests to update a patientrecord entity
// @Summary Update a patientrecord entity by ID
// @Description update patientrecord by ID
// @ID update-patientrecord
// @Accept   json
// @Produce  json
// @Param id path int true "Patientrecord ID"
// @Param patientrecord body ent.Patientrecord true "Patientrecord entity"
// @Success 200 {object} ent.Patientrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecords/{id} [put]
func (ctl *PatientrecordController) UpdatePatientrecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Patientrecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patientrecord binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Patientrecord.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}
// NewPatientrecordController creates and registers handles for the patientrecord controller
func NewPatientrecordController(router gin.IRouter, client *ent.Client) *PatientrecordController {
	uc := &PatientrecordController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
 
 // InitPatientrecordController registers routes to the main engine
 func (ctl *PatientrecordController) register() {
	patientrecords := ctl.router.Group("/patientrecords")
 
	patientrecords.GET("", ctl.ListPatientrecord)
 
	// CRUD
	patientrecords.POST("", ctl.CreatePatientrecord)
	patientrecords.GET(":id", ctl.GetPatientrecord)
	patientrecords.PUT(":id", ctl.UpdatePatientrecord)
	patientrecords.DELETE(":id", ctl.DeletePatientrecord)
 }
 
