package server

import (
	"backend/middleware"
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter(deps *Dependencies) *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "../frontend/public/assets")
	r.Static("/app", "../frontend/app")
	r.Static("/environment", "../frontend/environment")
	r.Static("/src", "../frontend/src")
	r.Static("/configs", "../configs")
	
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.TimingMiddleware())
	//authMiddleware := deps.Middleware.AuthenticatorMiddleware()


	rg := r.Group("/admin")
	routes.RegisterAdminPages(rg)

	rg = r.Group("/api")
	//rg.Use(authMiddleware)
	routes.SetUserAPIPaths(deps.UserHandler, rg)
	routes.SetTypesAPIPath(deps.TypeHandler, rg)


	r.NoRoute(func(c *gin.Context) { c.File("../frontend/public/pages/ERRORS/404.html") })
	return r
}
