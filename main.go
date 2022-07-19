package main

import (
	"elements-service/common"
	"elements-service/component"
	"elements-service/component/uploadprovider"
	"elements-service/middleware"
	"elements-service/modules/attributename/attributenamemodel"
	"elements-service/modules/element/elementmodel"
	"elements-service/modules/element/elementtransport/ginelement"
	"elements-service/modules/recipe/recipemodel"
	"elements-service/modules/recipe/recipetransport/ginrecipe"
	"elements-service/modules/restaurant/restauranttransport/ginrestaurant"
	"elements-service/modules/restaurantlike/transport/ginrestaurantlike"
	"elements-service/modules/upload/uploadtransport/gin_upload"
	"elements-service/modules/user/usermodel"
	"elements-service/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3Secret := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3Secret, s3Domain)

	secretKey := os.Getenv("SYSTEM_SECRET")

	//gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	err = migrateDB(db)
	if err != nil {
		log.Fatalln(err)
	}

	appCtx := component.NewAppContext(db, s3Provider, secretKey)
	err = runService(appCtx)
	if err != nil {
		log.Fatalln(err)
	}

}

func runService(appCtx component.AppContext) error {

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	v1 := r.Group("/api/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))
	v1.POST("/upload", gin_upload.Upload(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
	}

	recipes := v1.Group("/recipes")
	{
		recipes.GET("", middleware.RequiredAuth(appCtx), ginrecipe.ListRecipe(appCtx))
		recipes.POST("", ginrecipe.CreateRecipe(appCtx))
		recipes.GET("/:id", ginrecipe.GetRecipe(appCtx))
		recipes.PUT("/:id", ginrecipe.UpdateRecipe(appCtx))
		recipes.POST("/:id/elements", ginelement.CreateElement(appCtx))
		recipes.GET("/:id/elements", ginelement.ListElement(appCtx))
		recipes.DELETE("/:id", ginrecipe.DeleteRecipe(appCtx))
		recipes.POST("/:id/identifiers", ginrecipe.CreateIdentifierList(appCtx))
		recipes.POST("/:id/attribute-names", ginrecipe.CreateAttributeName(appCtx))
		//
		//recipes.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
	}

	elements := v1.Group("/elements")
	{
		elements.GET("/graph/:id", ginelement.GetElementGraph(appCtx))
		elements.GET("/:id", ginelement.GetElement(appCtx))
		elements.PUT("/:id", ginelement.UpdateElement(appCtx))
		elements.DELETE("/:id", ginelement.DeleteElement(appCtx))
	}

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	return r.Run()
}

func migrateDB(db *gorm.DB) error {
	//db.Migrator().DropTable(usermodel.User{})
	//db.Migrator().DropTable(recipemodel.Recipe{})
	//db.Migrator().DropTable(elementmodel.Element{})

	err := db.AutoMigrate(usermodel.User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(recipemodel.Recipe{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(elementmodel.Element{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(recipemodel.Identifier{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(attributenamemodel.AttributeName{})
	if err != nil {
		return err
	}
	return nil
}
