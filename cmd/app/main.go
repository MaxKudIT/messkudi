package main

import (
	"fmt"
	logger2 "github.com/MaxKudIT/messkudi/internal/logger"
	auth2 "github.com/MaxKudIT/messkudi/internal/services/auth"
	user2 "github.com/MaxKudIT/messkudi/internal/services/user"
	"github.com/MaxKudIT/messkudi/internal/storage"
	"github.com/MaxKudIT/messkudi/internal/storage/auth"
	"github.com/MaxKudIT/messkudi/internal/storage/user"
	auth4 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/auth"
	user3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/user"
	auth3 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/auth"
	user4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/user"
	"github.com/MaxKudIT/messkudi/internal/transport/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
)

func main() {

	if err := godotenv.Load("/home/max/messkudi/.env"); err != nil {
		log.Println(err.Error())
	}
	logger := logger2.New(slog.LevelInfo)
	lstor := logger.With("Layer", "Storage")
	lserv := logger.With("Layer", "Service")
	lhand := logger.With("Layer", "Handlers")
	dbstruct := storage.NewDatabase(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_NAME")))
	db, err := dbstruct.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}

	ust := user.New(db, lstor)
	us := user2.New(ust, lserv)
	ush := user3.New(us, lhand)
	usr := user4.New(ush)
	ast := auth.New(db, lstor)
	as := auth2.New(ast, lserv)
	ash := auth4.New(as, lhand)
	asr := auth3.New(ash)

	server := server.New(usr, asr)
	router := server.Create()
	router.Run(":3000")
}
