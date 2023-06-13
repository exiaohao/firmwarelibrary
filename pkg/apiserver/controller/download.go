package controller

import (
	"os"

	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/models"
	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/utils"
	"github.com/gin-gonic/gin"
)

func DownloadFirmwareByID(c *gin.Context) {
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
	firmwareModel := new(models.FirmwareModel)
	firmwareModel.QueryFirmwareByID(_id)

	firmwareModel.FillRelated()
	basePath, _ := os.Getwd()
	filePath := basePath + "/files/" + firmwareModel.FileExtraJson.Name
	if _, err := os.OpenFile(filePath, os.O_RDONLY, 0666); err != nil {
		c.JSON(400, gin.H{"message": "No such file" + filePath})
		c.Abort()
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+firmwareModel.FileExtraJson.Name)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath)
}
