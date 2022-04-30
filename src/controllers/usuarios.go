package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	// erro = errors.New("Deu erro")
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// buscar todos os usuário no banco de dados
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

// buscar um usuário especifico no banco
func SearchUser(w http.ResponseWriter, r *http.Request) {
	// nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	// db, erro := banco.Conectar()
	// if erro != nil {
	// 	respostas.Erro(w, http.StatusInternalServerError, erro)
	// }
	// defer db.Close()

	// repository := repositorios.NovoRepositorioDeUsuarios(db)
	// user, erro := repositorio.
	w.Write([]byte("Buscando Usuário!"))
}

// altera as informações de um usuário no banco
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// excluí as informações de um usuário no banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
