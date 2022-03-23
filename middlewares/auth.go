package middlewares

import (
	"final-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
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

// func ProductAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetDB()
// 		productId, err := strconv.Atoi(c.Param("productId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Bad Request",
// 				"message": "Invalid Params",
// 			})
// 			return
// 		}

// 		userData := c.MustGet("UserData").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		Product := models.Product{}

// 		err = db.Select("user_id").First(&Product, uint(productId)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Data Not Found",
// 				"message": "Not Exist",
// 			})
// 			return
// 		}

// 		if Product.UserID != userID {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "Not Authorized",
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }
