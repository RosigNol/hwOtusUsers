package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"embed"
	"time"
	"math/rand"
	"github.com/gin-gonic/gin"
	"otusHWUsers/config"
	"otusHWUsers/routes"
	"otusHWUsers/controllers"
	dbConn "otusHWUsers/db/sqlc"
	_ "github.com/lib/pq"
	ginMonitor "github.com/bancodobrasil/gin-monitor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	server *gin.Engine
	db     *dbConn.Queries

	UserController controllers.UserController
	UserRoutes     routes.UserRoutes
	embedMigrations embed.FS
)

func init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	psqlInfo := "postgres://"+ config.User + ":"+ config.Password + "@"+ config.Host +":" + config.PortPG+ "/" + config.DBName + "?sslmode=disable"

	conn, err := sql.Open(config.PostgreDriver, psqlInfo)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	UserController = *controllers.NewUserController(db)
	UserRoutes = routes.NewUserRoutes(UserController)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	monitor, err := ginMonitor.New("v1.0.0", ginMonitor.DefaultErrorMessageKey, []float64{0.5, 0.95, 0.99})
	if err != nil {
	    panic(err)
	}

	server.Use(monitor.Prometheus())

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		time.Sleep(time.Duration(rand.Intn(100000-1) + 1) * time.Millisecond)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})


	UserRoutes.UserRoute(router)

	log.Fatal(server.Run(":" + config.Port))
}
