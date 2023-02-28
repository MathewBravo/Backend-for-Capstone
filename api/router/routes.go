package routes

import (
	listings "Backend_for_Capstone/api"
	users "Backend_for_Capstone/api"
	"github.com/gin-gonic/gin"
)

func RouteFetch(router *gin.Engine) {
	router.GET("/hello", listings.GetMessage)
	router.GET("/search", listings.GetListingByParam)
	router.GET("/search/all", listings.GetListingsForSearchDisplay)
	router.GET("/users/all", users.GetAllUsers)
	router.GET("/admin/pending", users.GetListingsForAdminPending)
}
