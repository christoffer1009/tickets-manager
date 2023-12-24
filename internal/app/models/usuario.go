package models

import "github.com/google/uuid"

// Usuario é uma interface para representar usuários.
type Usuario interface {
	GetID() uuid.UUID
	GetNome() string
	GetEmail() string
	GetSetorLocacao() string
	ToString() string
}
