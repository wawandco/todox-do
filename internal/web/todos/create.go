package todos

import (
	"net/http"
	"paganotoni/todox/internal/models"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
)

func Create(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}
	err := form.Decode(r, &todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo.Completed = false
	todos := r.Context().Value("todoService").(models.TodoService)

	err = todos.Create(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("list", list)

	err = rw.RenderClean("todos/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
