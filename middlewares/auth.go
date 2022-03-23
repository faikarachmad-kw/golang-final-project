package middlewares

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		c.Set("UserData", verifyToken)
		c.Next()
	}
}

func Authorization(input string) gin.HandlerFunc {
return func(c *gin.Context){

}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(c.Param("photoID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Params",
			})
			return
		}

		userData := c.MustGet("UserData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Data Not Found",
				"message": "Not Exist",
			})
			return
		}
		fmt.Println("photo user id",Photo.UserID)
		fmt.Println("current user id",userID)
		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "Not Authorized",
			})
			return
		}

		c.Next()
	}
}
