package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhtuan260593/todo/web/config"
	"github.com/thanhtuan260593/todo/web/handler"
)

//Server struct
type Server struct {
	resolver *config.Resolver
}

//LoadResolver ..
func (sv *Server) LoadResolver() {
	sv.resolver = config.NewResolver()
}

//New server
func New() *Server {
	server := &Server{}
	server.LoadResolver()
	return server
}

func (sv *Server) reserveProxy(c *gin.Context) {
	proxyURL := sv.resolver.Config.ProxyURL
	if proxyURL == "" {
		return
	}
	//var u, _ = url.Parse(proxyURL)
	//proxy := httputil.NewSingleHostReverseProxy(u)
	//proxy.ServeHTTP(c.Writer, c.Request)
}

//Start server
func (sv *Server) Start() {
	r := gin.Default()
	v1 := r.Group("/todo")
	{
		v1.GET("/items", sv.todoAPI(func(api *handler.TodoAPI) gin.HandlerFunc {
			return api.GetItems
		}))
		v1.POST("/create/item", sv.todoAPI(func(api *handler.TodoAPI) gin.HandlerFunc {
			return api.CreateItem
		}))
		v1.GET("/item/:id", sv.todoAPI(func(api *handler.TodoAPI) gin.HandlerFunc {
			return api.GetItem
		}))
		v1.PUT("/item/:id/update", sv.todoAPI(func(api *handler.TodoAPI) gin.HandlerFunc {
			return api.UpdateItem
		}))
		v1.PUT("/item/:id/toggle", sv.todoAPI(func(api *handler.TodoAPI) gin.HandlerFunc {
			return api.ToggleItem
		}))
		v1.DELETE("/item/:id", sv.todoAPI(func(api *handler.TodoAPI) gin.HandlerFunc {
			return api.DeleteItem
		}))
	}
	r.NoRoute(sv.reserveProxy)
	r.Run(sv.resolver.Config.ServerURL)
	fmt.Println("Running server at", sv.resolver.Config.ServerURL)
}

func (sv *Server) todoAPI(call func(api *handler.TodoAPI) gin.HandlerFunc) gin.HandlerFunc {
	return call(sv.resolver.ResolveTodoAPI())
}
