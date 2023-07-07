package handler

import (
	"os"

	"github.com/fydhfzh/fp-4/database"
	"github.com/fydhfzh/fp-4/docs"
	"github.com/fydhfzh/fp-4/repository/category_repository/category_pg"
	"github.com/fydhfzh/fp-4/repository/product_repository/product_pg"
	"github.com/fydhfzh/fp-4/repository/transaction_repository/transaction_pg"
	"github.com/fydhfzh/fp-4/repository/user_repository/user_pg"
	"github.com/fydhfzh/fp-4/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	var port = os.Getenv("PORT")

	database.InitializedDatabase()
	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewPGUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	categoryRepo := category_pg.NewPGCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	productRepo := product_pg.NewPGProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	transactionRepo := transaction_pg.NewPGTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo, productRepo, userRepo, categoryRepo)
	transactionHandler := NewTransactionHandler(transactionService)

	authService := service.NewAuthService(userRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
		userRoute.PATCH("/topup", authService.Authentication(), userHandler.TopUp)
	}

	categoryRoute := route.Group("/categories")
	{
		categoryRoute.Use(authService.Authentication(), authService.AdminAuthorization())
		categoryRoute.POST("/", categoryHandler.CreateCategory)
		categoryRoute.GET("/", categoryHandler.GetAllCategories)
		categoryRoute.PATCH("/:categoryId", categoryHandler.UpdateCategory)
		categoryRoute.DELETE("/:categoryId", categoryHandler.DeleteCategory)
	}

	productRoute := route.Group("/products")
	{
		productRoute.Use(authService.Authentication())
		productRoute.GET("/", productHandler.GetAllProducts)
		productRoute.POST("/", authService.AdminAuthorization(), productHandler.CreateProduct)
		productRoute.PUT("/:productId", authService.AdminAuthorization(), productHandler.UpdateProduct)
		productRoute.DELETE("/:productId", authService.AdminAuthorization(), productHandler.DeleteProduct)
	}

	transactionRoute := route.Group("/transactions")
	{
		transactionRoute.Use(authService.Authentication())
		transactionRoute.POST("/", transactionHandler.CreateTransaction)
		transactionRoute.GET("/my-transactions", transactionHandler.GetMyTransactions)
		transactionRoute.GET("/user-transactions", authService.AdminAuthorization(), transactionHandler.GetUsersTransactions)
	}

	docs.SwaggerInfo.Title = "API Toko Belanja"
	docs.SwaggerInfo.Description = "Ini adalah server API Toko Belanja."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "fp-4-production-e413.up.railway.app"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	route.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.Run(":" + port)
}
