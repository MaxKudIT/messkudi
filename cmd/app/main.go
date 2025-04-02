package main

import (
	"fmt"
	user2 "github.com/MaxKudIT/messkudi/internal/services/user"
	"github.com/MaxKudIT/messkudi/internal/storage"
	"github.com/MaxKudIT/messkudi/internal/storage/user"
	"github.com/MaxKudIT/messkudi/internal/transport/web"
	user3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/user"
	user4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
)

type newus interface {
	GetUserById()
	SaveUser()
	DeleteUser()
}
type userRouters interface {
	Regrouters(engine *gin.Engine)
}

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(err.Error())
	}

	dbstruct := storage.NewDatabase(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_NAME")))
	db, err := dbstruct.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
	ust := user.New(db)
	us := user2.New(ust, &slog.Logger{})
	ush := user3.New(us)
	usr := user4.New(ush)
	server := web.New(usr)
	server.StartServer(":8080")

}
