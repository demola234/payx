package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var v = CreateAccountDetails()
		fmt.Sprintln(v)
	}
}

func UploadProfileImage() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func HashPassword() {}

func VerifyPassword() {

}
