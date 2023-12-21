package main

import (
	"context"
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
	"github.com/AkifhanIlgaz/key-aero-api/controllers"
	"github.com/AkifhanIlgaz/key-aero-api/db"
	"github.com/AkifhanIlgaz/key-aero-api/middleware"
	"github.com/AkifhanIlgaz/key-aero-api/routes"
	"github.com/AkifhanIlgaz/key-aero-api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.TODO()

	config, err := cfg.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not read environment variables", err)
	}

	databases, err := db.ConnectToDatabases(config)
	if err != nil {
		log.Fatal("Could not connect to databases: ", err)
	}

	defer databases.Postgres.Close()
	defer databases.Redis.Close()

	server := gin.Default()
	setCors(server)

	userService := services.NewUserService(ctx, databases.Postgres)
	tokenService := services.NewTokenService(ctx, config, databases.Redis)

	authController := controllers.NewAuthController(config, userService, tokenService)

	userMiddleware := middleware.NewUserMiddleware(userService, tokenService)

	authRouteController := routes.NewAuthRouteController(authController)

	router := server.Group("/api")
	router.GET("/health-checker", userMiddleware.ExtractUser(), userMiddleware.HasRole("zozak"), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "API is healthy"})
	})

	authRouteController.AuthRoute(router)

	log.Fatal(server.Run(":" + config.Port))
}

func setCors(server *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
}
