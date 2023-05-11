package router

import (
	"airland-server/docs"
	"airland-server/src/config"
	"airland-server/src/cross_cutting/models"
	"airland-server/src/infrastructure/middlewares"

	"net/http"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routes struct {
	Router  *gin.Engine
	Version *gin.RouterGroup
	Public  *gin.RouterGroup
	Private *gin.RouterGroup
	Admin   *gin.RouterGroup
}

func NewRouter() (routes *Routes) {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.C.Server.Host, config.C.Server.Port)
	router := gin.Default()
	router.Use(cors.New(getCorsConfig()))
	version := router.Group("v1/")
	auth := middlewares.Bootstrap()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	private := version.Group("private/")
	private.Use(auth.CheckJWT)

	routes = &Routes{
		Router:  router,
		Version: version,
		Public:  version,
		Private: private,
		Admin:   version.Group("admin/"),
	}

	routes.swagger()
	routes.Private.GET("/health", healthChecker)
	return routes
}

func (r *Routes) swagger() {
	docs.SwaggerInfo.BasePath = "/v1"
	r.Public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func getCorsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	return config
}

// Health Checker godoc
// @Summary Check the api health
// @param Authorization header string false "Should be Bearer ..."
// @security Authorization
// @Schemes
// @Description Check if the api is working properly
// @Produce json
// @Success 200 {object} models.ResultWrapper[any]
// @Router /private/health [get]
func healthChecker(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.NewResultWrapper[any]("health.checker.success", true, nil, nil))
}
