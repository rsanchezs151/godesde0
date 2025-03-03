package users

import (
	"fmt"
	"time"

	"github.com/rsanchezs151/godesde0/models"
)

func AltaUsuario() {
	u := new(models.User)
	u.AddUser(1, "Roberto", time.Now(), true)
	fmt.Println(u)
}
