package server

import "pokerok/internal/provider"

type Provider interface {
	Registration(user provider.User) error
	Auth(user provider.User) error
}
