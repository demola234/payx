package helpers

import(

	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

// success

func SuccessResponse (c *gin.Context, data []bson.M){
	
	c.JSON(200, gin.H{"status": "Success",	"data": data[0]})
}

// errors

func InternalError (c *gin.Context, message string){
	
	c.JSON(500, gin.H{"status": "Success",	"message": message})
}