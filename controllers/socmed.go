package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateSocmed(c *gin.Context) {
	db:=database.GetDB()

	userData :=c.MustGet("UserData").(jwt.MapClaims)
	contentType:=helpers.GetContentType(c)

	socMed :=models.SocialMedia{}
	userID:=uint(userData["id"].(float64))

	if contentType==appJSON{
		c.ShouldBindJSON(&socMed)
	}else{
		c.ShouldBind(&socMed)
	}

	socMed.UserID=userID

	err:=db.Debug().Create(&socMed).Error

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"err":"Bad Request",
			"message":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"id":socMed.ID,
		"name":socMed.Name,
		"social_media_url":socMed.SocialMediaURL,
		"user_id":socMed.UserID,
		"created_at":socMed.CreatedAt,
	})
}
func GetSocmed(c *gin.Context) {
	db:=database.GetDB()
	userData :=c.MustGet("UserData").(jwt.MapClaims)
	userID:=uint(userData["id"].(float64))
	socMed:=[]models.SocialMedia{}
	resData:=[]map[string]interface{}{}
	_=resData
	err:=db.Preload("User").Where("user_id=?",userID).Find(&socMed).Error
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"err":"Bad Request",
			"message":err.Error(),
		})
		return
	}

	for i:=range socMed{
		nestedData:=map[string]interface{}{
			"id":socMed[i].User.ID,
			"username":socMed[i].User.Username,
			"profile_image_url":"place holder string, di spesifikasi tabel tidak ada satupun kolom profile_image_url, disini tiba2 ada",
		}
		data:=map[string]interface{}{
			"id":socMed[i].ID,
			"name":socMed[i].Name,
			"social_media_url":socMed[i].SocialMediaURL,
			"user_id":socMed[i].UserID,
			"created_at":socMed[i].CreatedAt,
			"updated_at":socMed[i].UpdatedAt,
			"User":nestedData,
		}

		resData=append(resData,data)
	}

	c.JSON(http.StatusOK,gin.H{
		"social_medias":resData})
}
func UpdateSocmed(c *gin.Context) {

}
func DeleteSocmed(c *gin.Context) {

}