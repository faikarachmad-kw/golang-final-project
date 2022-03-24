package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateComment(c *gin.Context) {
	db:=database.GetDB()

	userData :=c.MustGet("UserData").(jwt.MapClaims)
	contentType:=helpers.GetContentType(c)

	comment :=models.Comment{}
	userID:=uint(userData["id"].(float64))

	if contentType==appJSON{
		c.ShouldBindJSON(&comment)
	}else{
		c.ShouldBind(&comment)
	}

	comment.UserID=userID

	err:=db.Debug().Create(&comment).Error

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"err":"Bad Request",
			"message":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"id":comment.ID,
		"message":comment.Message,
		"photo_id":comment.PhotoID,
		"user_id":comment.UserID,
		"created_at":comment.CreatedAt,
	})
}

func GetComment(c *gin.Context) {

}
func UpdateComment(c *gin.Context) {

}
func DeleteComment(c *gin.Context) {

}