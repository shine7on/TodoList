package routes

import (
	"github.com/Isshinfunada/TodoList/server/handlers"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/Isshinfunada/TodoList/server/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo, db *models.Queries) {
	// CORS設定を追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},       // フロントエンドのオリジン
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT}, // 許可するHTTPメソッド
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/users", handlers.RegisterUser(db))
	e.POST("/login", handlers.Login(db))

	// TodoHandler のインスタンスを作成
	todoHandler := handlers.TodoHandler{
		TodoService: &services.TodoService{
			Queries: db,
		},
	}

	// TodoHandler の GetTodos メソッドを呼び出す
	e.GET("/todos/:user_id", todoHandler.GetTodos)
	e.POST("/todos", todoHandler.CreateTodo)
	e.POST("/todos/edit", todoHandler.EditTodo)
	e.POST("/todos/delete", todoHandler.DeleteTodo)
	e.POST("/todos/:id/status", todoHandler.UpdateTodoStatus)
}
