package requests

import (
	"github.com/pseudoelement/go-rabbitmq/models"
	api_module "github.com/pseudoelement/golang-utils/src/api"
)

func FetchTodoItem() models.JP_Todo_Resp {
	resp, _ := api_module.Get[models.JP_Todo_Resp]("https://jsonplaceholder.typicode.com/todos/1", make(map[string]string), make(map[string]string))
	return resp
}
