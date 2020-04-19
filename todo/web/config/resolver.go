package config

import (
	infs "github.com/thanhtuan260593/todo/core/interfaces"
	"github.com/thanhtuan260593/todo/infrastructure/database"
	"github.com/thanhtuan260593/todo/infrastructure/usecases"
	"github.com/thanhtuan260593/todo/web/handler"
)

//Resolver struct
type Resolver struct {
	Config  Config
	db      *database.Database
	todoAPI *handler.TodoAPI
	todoUC  infs.ITodoUsecase
}

//NewResolver this function should be call once per server run
func NewResolver() *Resolver {
	var r Resolver
	r.Config = LoadConfig()
	return &r
}

//ResolveDb resolve db di
func (r *Resolver) ResolveDb() *database.Database {
	if r.db == nil {
		r.db = database.New(r.Config.DbURL)
	}
	return r.db
}

//ResolveTodoUsecase todo usecase
func (r *Resolver) ResolveTodoUsecase() infs.ITodoUsecase {
	if r.todoUC == nil {
		r.todoUC = usecases.NewTodo(r.ResolveDb())
	}
	return r.todoUC
}

//ResolveTodoAPI resolve todoAPI
func (r *Resolver) ResolveTodoAPI() *handler.TodoAPI {
	if r.todoAPI == nil {
		r.todoAPI = handler.NewTodoAPI(
			r.ResolveTodoUsecase(),
		)
	}
	return r.todoAPI
}
