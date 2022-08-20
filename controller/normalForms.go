package controller

import (
	"log"
	"zjutjh/Join-Us/db/model"
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
)

func NewNormalForm(c *gin.Context) {
	postData := model.NormalForm{}
	err := c.ShouldBindJSON(&postData)
	if err != nil {
		log.Println("New Normal:", err)
		utility.ResponseError(c, "Post Data Error")
		return
	}
	var form *model.NormalForm
	form, _ = model.GetNormalFormByStuID(postData.StuID)

	if (form == &model.NormalForm{}) {
		err = model.NewNormalForm(&postData)
	} else {
		err = model.UpdateNormalForm(&postData)
	}
	if err != nil {
		utility.ResponseError(c, "Internal Server Error")
		return
	}
	utility.ResponseSuccess(c, nil)
}

func GetNormalFormTotal(c *gin.Context) {
	res := model.GetNormalFormTotal()
	utility.ResponseSuccess(c, gin.H{"data": res})
}

func GetAllNormalFormsBrief(c *gin.Context) {
	forms, err := model.GetAllNormalForms()
	if err != nil {
		utility.ResponseError(c, "Internal Server Error")
	}
	utility.ResponseSuccess(c, gin.H{"data": forms})
}
