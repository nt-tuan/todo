package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhtuan260593/todo/core/entities"
	infs "github.com/thanhtuan260593/todo/core/interfaces"
	model "github.com/thanhtuan260593/todo/web/models"
)

//TodoAPI struct todo
type TodoAPI struct {
	uc infs.ITodoUsecase
}

//NewTodoAPI func
func NewTodoAPI(uc infs.ITodoUsecase) *TodoAPI {
	handler := &TodoAPI{
		uc: uc,
	}
	return handler
}

//CreateItem api
func (h *TodoAPI) CreateItem(c *gin.Context) {
	var body struct {
		Title string
	}
	if err := JSONError(c, c.BindJSON(&body)); err != nil {
		return
	}
	item, err := h.uc.CreateItem(body.Title)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, GetError(err))
	}
	c.JSON(http.StatusOK, model.ResponseItem{item.ID, item.Title, item.IsDone})
}

//GetItem api
func (h *TodoAPI) GetItem(c *gin.Context) {
	var id model.ActionItem
	if err := JSONError(c, c.BindUri(&id)); err != nil {
		return
	}

	if item, err := h.uc.GetByID(id.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, model.NewResponseItem(item))
	}
}

//GetItems api
func (h *TodoAPI) GetItems(c *gin.Context) {
	var page = model.Pagable{0, 5, "", "ASC"}
	if err := JSONError(c, c.ShouldBindQuery(&page)); err != nil {
		return
	}
	if items, err := h.uc.ListItems(page.ToEntity()); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		rs := []*model.ResponseItem{}
		for _, i := range items {
			rs = append(rs, model.NewResponseItem(&i))
		}
		c.JSON(http.StatusOK, rs)
	}
}

//UpdateItem an item
func (h *TodoAPI) UpdateItem(c *gin.Context) {
	var id model.ActionItem
	var body model.CreateItem
	if err1, err2 := JSONError(c, c.BindUri(&id)), JSONError(c, c.BindJSON(&body)); err1 != nil || err2 != nil {
		return
	}
	update, err := h.uc.UpdateItem(id.ID, &entities.Item{
		Title: body.Title,
	})
	if err := JSONError(c, err); err != nil {
		return
	}
	c.JSON(http.StatusOK, model.NewResponseItem(update))
}

//ToggleItem api
func (h *TodoAPI) ToggleItem(c *gin.Context) {
	var id model.ActionItem
	if err := JSONError(c, c.BindUri(&id)); err != nil {
		return
	}

	if ok, err := h.uc.ToggleItem(id.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, &model.ResponseItem{ok.ID, ok.Title, ok.IsDone})
	}
}

//DeleteItem api
func (h *TodoAPI) DeleteItem(c *gin.Context) {
	var id model.ActionItem
	if err := JSONError(c, c.BindUri(&id)); err != nil {
		return
	}

	if err := JSONError(c, h.uc.DeleteItem(id.ID)); err != nil {
		return
	}

	c.JSON(http.StatusOK, true)
}
