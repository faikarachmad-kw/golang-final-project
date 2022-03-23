package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(c *gin.Context){
	db:=database.GetDB()

	userData :=c.MustGet("UserData").(jwt.MapClaims)
	contentType:=helpers.GetContentType(c)

	Photo :=models.Photo{}
	userID:=uint(userData["id"].(float64))

	if contentType==appJSON{
		c.ShouldBindJSON(&Photo)
	}else{
		c.ShouldBind(&Photo)
	}

	Photo.UserID=userID

	err:=db.Debug().Create(&Photo).Error

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"err":"Bad Request",
			"message":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"id":Photo.ID,
		"title":Photo.Title,
		"caption":Photo.Caption,
		"photo_url":Photo.PhotoURL,
		"user_id":Photo.UserID,
		"created_at":Photo.CreatedAt,
	})
}

func GetPhoto(c *gin.Context){

}

func UpdatePhoto(c *gin.Context){

}

func DeletePhoto(c *gin.Context){

}