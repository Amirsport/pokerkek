package server

import (
	"fmt"
	"pokerok/internal/provider"

	"github.com/gin-gonic/gin"
)

func (s *Server) Registration(c *gin.Context) {
	u := provider.User{}
	if err := c.Bind(&u); err != nil {
		c.JSON(400, provider.Message{Err: err.Error()})
		return
	}
	if u.Login == nil || u.Password == nil || u.Nickname == nil {
		c.JSON(400, provider.Message{Err: "ZAPOLNI UBLUDOK"})
		return
	}
	// функция заноса данных в бд
	if err := s.prov.Registration(u); err != nil {
		c.JSON(500, provider.Message{Err: "SQL error"})
		return
	}
	fmt.Println(*u.Login, *u.Password)
	c.JSON(200, provider.Message{Err: "ok"})
}
