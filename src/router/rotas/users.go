package rotas

import (
	"api/src/controllers"
	"net/http"
)

var userRouter = []Rota{
	{
		URI:                "/user",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/user",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchUsers,
		RequerAutenticacao: false,
	},
	{
		URI:                "/user/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SearchUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/user/{UsuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/user/{UsuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: false,
	},
}
