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
	db:=database.GetDB()
	userData :=c.MustGet("UserData").(jwt.MapClaims)
	userID:=uint(userData["id"].(float64))
	Photo:=[]models.Photo{}
	resData:=[]map[string]interface{}{}
	_=resData
	err:=db.Preload("User").Where("user_id=?",userID).Find(&Photo).Error
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"err":"Bad Request",
			"message":err.Error(),
		})
		return
	}

	for i:=range Photo{
		nestedData:=map[string]interface{}{
			"email":Photo[i].User.Email,
			"username":Photo[i].User.Username,
		}
		data:=map[string]interface{}{
			"id":Photo[i].ID,
			"title":Photo[i].Title,
			"caption":Photo[i].Caption,
			"photo_url":Photo[i].PhotoURL,
			"user_id":Photo[i].UserID,
			"created_at":Photo[i].CreatedAt,
			"updated_at":Photo[i].UpdatedAt,
			"User":nestedData,
		}

		resData=append(resData,data)
	}

	c.JSON(http.StatusOK,resData)
}

func UpdatePhoto(c *gin.Context){

}

func DeletePhoto(c *gin.Context){

}