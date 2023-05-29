package controller

import (
	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/models"
	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/utils"
	"github.com/gin-gonic/gin"
)

type QueryByKeywordBody struct {
	Query string `json:"query" binding:"required"`
}

func (q QueryByKeywordBody) ToQuery() string {
	return "%" + q.Query + "%"
}

type QueryByKeywordResult struct {
	DisplayName             string
	DisplayManufacturerName string
	Count                   uint
	Manufacturer            uint
	Model                   uint
	Search                  string
}

func (qkr *QueryByKeywordResult) FillRelated() {
	if qkr.Model != 0 {
		mm := &models.ModelModel{}
		mam := &models.ManufacturerModel{}
		mm.QueryModelByID(qkr.Model)
		mam.QueryManufacturerByID(qkr.Manufacturer)
		qkr.DisplayName = mm.DisplayName
		qkr.DisplayManufacturerName = mam.DisplayName
	}
}

// Query By Keywords
func QueryByKeywords(c *gin.Context) {
	var query QueryByKeywordBody
	if c.ShouldBind(&query) != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		c.Abort()
		return
	}
	// todo:Security check
	results := []*QueryByKeywordResult{}
	models.DB.Raw("SELECT count(id) as count, manufacturer, model, search FROM t_firmwares where `search` like ? GROUP BY `model`;", query.ToQuery()).Scan(&results)
	for i := range results {
		results[i].FillRelated()
	}
	JSONResponse(c, 200, results)
}

func QueryByModelID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"message": "Bad Request"})
		c.Abort()
		return
	}
	_id, err := utils.StringToUint(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		c.Abort()
		return
	}
	fm := models.FirmwareModel{}
	fms := fm.QueryFirmwaresByModelID(_id)
	JSONResponse(c, 200, fms)
}
