package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios { // recebe um ponteiro de banco de dados e usuarios
	return &Usuarios{db}
}

// insere um usuário no banco de dados
// retorna o uint64 ou error
// respositorio é o respositorio de Usuarios
// Criar é o nome do método e recebe um modelo de usuario
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senhas) value(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
