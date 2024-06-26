package handler

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/markovk1n/webTodo/docs"

	"github.com/markovk1n/webTodo/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	tmpl     *template.Template
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		tmpl:     template.Must(template.ParseGlob("./templates/*.html")),
		services: service,
	}
}

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.SetHTMLTemplate(h.tmpl)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", h.home)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)

			}
		}
		items := api.Group("items")
		{
			items.GET("/:id", h.getItem)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}

func (h *Handler) home(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", "HELLO")

}
