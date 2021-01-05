package controllers

import (
   "context"
   
   "strconv"
   "github.com/PON/app/ent"
   "github.com/PON/app/ent/insurance"
   "github.com/gin-gonic/gin"
)

// InsuranceController defines the struct for the Insurance controller
type InsuranceController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateInsurance handles POST requests for adding Insurance entities
// @Summary Create Insurance
// @Description Create Insurance
// @ID create-Insurance
// @Accept   json
// @Produce  json
// @Param Insurance body ent.Insurance true "Insurance entity"
// @Success 200 {object} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Insurances [post]
func (ctl *InsuranceController) CreateInsurance(c *gin.Context) {
   obj := ent.Insurance{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "insurance binding failed",
       })
       return
   }

   u, err := ctl.client.Insurance.
       Create().
       SetInsurancecompany(obj.Insurancecompany).

       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{
           "error": "saving failed",
       })
       return
   }

   c.JSON(200, u)
}
// GetInsurance handles GET requests to retrieve a insurance entity
// @Summary Get a Insurance entity by ID
// @Description get Insurance by ID
// @ID get-insurance
// @Produce  json
// @Param id path int true "Insurance ID"
// @Success 200 {object} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances/{id} [get]
func (ctl *InsuranceController) GetInsurance(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }

   u, err := ctl.client.Insurance.
       Query().
       Where(insurance.IDEQ(int(id))).
       Only(context.Background())
   if err != nil {
       c.JSON(404, gin.H{
           "error": err.Error(),
       })
       return
   }

   c.JSON(200, u)
}

// ListInsurance handles request to get a list of insurance entities
// @Summary List insurance entities
// @Description list insurance entities
// @ID list-insurance
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances [get]
func (ctl *InsuranceController) ListInsurance(c *gin.Context) {
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

   insurances, err := ctl.client.Insurance.
       Query().
       Limit(limit).
       Offset(offset).
       All(context.Background())
   	if err != nil {
       c.JSON(400, gin.H{"error": err.Error(),})
       return
   }

   c.JSON(200, insurances)
}

// UpdateInsurance handles PUT requests to update a insurance entity
// @Summary Update a insurance entity by ID
// @Description update insurance by ID
// @ID update-insurance
// @Accept   json
// @Produce  json
// @Param id path int true "Insurance ID"
// @Param insurance body ent.Insurance true "Insurance entity"
// @Success 200 {object} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances/{id} [put]
func (ctl *InsuranceController) UpdateInsurance(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }

   obj := ent.Insurance{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "insurance binding failed",
       })
       return
   }
   obj.ID = int(id)
   u, err := ctl.client.Insurance.
       UpdateOne(&obj).
       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{"error": "update failed",})
       return
   }

   c.JSON(200, u)
}

// NewInsuranceController creates and registers handles for the insurance controller
func NewInsuranceController(router gin.IRouter, client *ent.Client) *InsuranceController {
   uc := &InsuranceController{
       client: client,
       router: router,
   }
   uc.register()
   return uc
}

// InitInsuranceController registers routes to the main engine
func (ctl *InsuranceController) register() {
    insurances := ctl.router.Group("/insurances")

    insurances.GET("", ctl.ListInsurance)

   // CRUD
   insurances.POST("", ctl.CreateInsurance)
   insurances.GET(":id", ctl.GetInsurance)
   insurances.PUT(":id", ctl.UpdateInsurance)
   //insurances.DELETE(":id", ctl.DeleteInsurance)
}


