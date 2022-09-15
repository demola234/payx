package helpers

import (
	"net/http"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

func GetUploadedFile(c *gin.Context) string {

	// Create our instance
	cld, _ := cloudinary.NewFromURL("cloudinary://211576879732455:W6p_HMMIrDZkEfheHRUHIkSTdOo@dcnuiaskr")

	// Add tags
	// fileTags := c.PostForm("tags")
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to upload",
		})
	}

	// Access the filename using a desired file access id.
	result, err := cld.Upload.Upload(c, file, uploader.UploadParams{
		PublicID: "profileName",
	})

	if err != nil {
		c.String(http.StatusNotFound, "We were unable to find the file requested")
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":    "Successfully found the image",
		"secureURL":  result.SecureURL,
		"publicURL":  result.URL,
		"created_at": result.CreatedAt.String(),
	})

	return result.URL

}
