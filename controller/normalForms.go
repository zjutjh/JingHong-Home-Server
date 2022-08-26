package controller

import (
	"log"
	"zjutjh/Join-Us/db/model"
	"zjutjh/Join-Us/utility"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func NewNormalForm(c *gin.Context) {
	postData := model.NormalForm{}
	err := c.ShouldBindBodyWith(&postData, binding.JSON)
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
	// utility.SendEmail(*form)
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

func GetNormalForm(c *gin.Context) {
	stu_id := c.Query("stu_id")
	if stu_id == "" {
		utility.ResponseError(c, "No Param")
		return
	}
	form, err := model.GetNormalFormByStuID(stu_id)
	if err != nil {
		utility.ResponseError(c, "Server Error")
		return
	}
	utility.ResponseSuccess(c, gin.H{"data": form})
}

func ExportAllNormalFormExcel(c *gin.Context) {
	file := utility.GenerateExcel()
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=ExportForms.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(file)
}
