package tproduct

import (
	"github.com/gin-gonic/gin"
	"go_server/model"
	"go_server/controller/common"

)

func FindPdtById(c *gin.Context){
	var pdt []model.TProduct
	if err := model.DB.Where("guid = ?", c.Param("id")).Find(&pdt).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	common.SendResponse(pdt,c)
	return
}

func FindHotPdtList(c *gin.Context){
	var pdt []model.TProduct
	if err := model.DB.Limit(5).Where(map[string]interface{}{"isDelete": "FALSE", "sellerId": c.Query("id"),"status":c.Query("status"),"checkStatus":"CHECKSUCCESS"}).Order("updateTime desc").Find(&pdt).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	common.SendResponse(pdt,c)
	return
}

func FindBySellerId(c *gin.Context){
	var pdt []model.TProduct
	if err := model.DB.Offset(c.Query("pageNum")).Limit(c.Query("pageSize")).Where(map[string]interface{}{"isDelete": "FALSE", "sellerId": c.Query("id"),"status":c.Query("status"),"checkStatus":"CHECKSUCCESS"}).Order("updateTime desc").Find(&pdt).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	common.SendResponse(pdt,c)
	return
}

func FindNpassBySellerId(c *gin.Context){
	var pdt []model.TProduct
	if err := model.DB.Offset(c.Query("pageNum")).Limit(c.Query("pageSize")).Where(map[string]interface{}{"isDelete": "FALSE", "sellerId": c.Query("id"),"checkStatus":"CHECKFAIL"}).Order("updateTime desc").Find(&pdt).Error; err != nil {
		common.SendErrorMsg(err.Error(),c)
		return
	}
	var resultList []interface{}
	for _ ,pdtitem := range pdt{
		var pdtAmsg model.TProductAuth
		resultmap := make(map[string]interface{})
		resultmap["pdtName"] = pdtitem.Title
		resultmap["pdtCheckstatus"] = pdtitem.Checkstatus
		if err := model.DB.Limit(1).Where(map[string]interface{}{"pId": pdtitem.Guid}).Order("checkTime desc").Find(&pdtAmsg).Error; err != nil {
			resultmap["noPassMsg"] = pdtAmsg.Checkinfo
		}
		resultList = append(resultList, resultmap)
	}
	common.SendResponse(resultList,c)
	return
}