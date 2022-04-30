package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Retornar router com rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
