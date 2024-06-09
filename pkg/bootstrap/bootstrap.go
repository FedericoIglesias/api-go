package bootstrap

import (
	"api-go/internal/domain"
	"api-go/internal/user"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
		Users: []domain.User{
			{
				ID:        1,
				FirstName: "Pepe",
				LastName:  "Coco",
				Email:     "@algo",
			},
			{
				ID:        2,
				FirstName: "Cacho",
				LastName:  "Goxila",
				Email:     "@otraCosa",
			},
			{
				ID:        3,
				FirstName: "Armando",
				LastName:  "Banquito",
				Email:     "@otroGato",
			},
		},
		MaxUserID: 3,
	}
}
