package main

import (
	"fmt"
	logger2 "github.com/MaxKudIT/messkudi/internal/logger"
	auth2 "github.com/MaxKudIT/messkudi/internal/services/auth"
	chat2 "github.com/MaxKudIT/messkudi/internal/services/chat"
	contact2 "github.com/MaxKudIT/messkudi/internal/services/contact"
	group2 "github.com/MaxKudIT/messkudi/internal/services/group"
	chat_message2 "github.com/MaxKudIT/messkudi/internal/services/message/chat_message"
	group_message2 "github.com/MaxKudIT/messkudi/internal/services/message/group_message"
	session2 "github.com/MaxKudIT/messkudi/internal/services/session"
	user2 "github.com/MaxKudIT/messkudi/internal/services/user"
	"github.com/MaxKudIT/messkudi/internal/storage"
	"github.com/MaxKudIT/messkudi/internal/storage/auth"
	"github.com/MaxKudIT/messkudi/internal/storage/chat"
	"github.com/MaxKudIT/messkudi/internal/storage/contact"
	"github.com/MaxKudIT/messkudi/internal/storage/group"
	"github.com/MaxKudIT/messkudi/internal/storage/message/chat_message"
	group_message1 "github.com/MaxKudIT/messkudi/internal/storage/message/group_message"
	chat3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/chat"
	group3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/group"
	group_message3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/message/group_message"
	chat4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/chat"
	group4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/group"
	group_message4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/message/group_message"

	contact3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/contact"

	"github.com/MaxKudIT/messkudi/internal/storage/session"
	"github.com/MaxKudIT/messkudi/internal/storage/user"
	auth4 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/auth"
	chat_message3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/message/chat_message"
	user3 "github.com/MaxKudIT/messkudi/internal/transport/web/handlers/user"
	"github.com/MaxKudIT/messkudi/internal/transport/web/handlers/websocket"
	auth3 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/auth"
	contact4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/contact"
	chat_message4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/message/chat_message"
	user4 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/user"
	websocket2 "github.com/MaxKudIT/messkudi/internal/transport/web/routers/websocket"
	"github.com/MaxKudIT/messkudi/internal/transport/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
)

func main() {

	if err := godotenv.Load("/home/max/Рабочий стол/te/messkudi/.env"); err != nil {
		log.Println(err.Error())
	}
	logger := logger2.New(slog.LevelInfo)
	lstor := logger.With("Layer", "Storage")
	lserv := logger.With("Layer", "Service")
	lhand := logger.With("Layer", "Handlers")
	dbstruct := storage.NewDatabase(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_NAME")))
	db, err := dbstruct.ConnectionDB()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	ust := user.New(db, lstor) //user
	us := user2.New(ust, lserv)
	ush := user3.New(us, lhand)
	usr := user4.New(ush)
	ast := auth.New(db, lstor) //auth
	as := auth2.New(ast, lserv)

	sst := session.New(db, lstor) //session
	ss := session2.New(sst, lserv)

	ash := auth4.New(as, ss, lhand) //auth
	asr := auth3.New(ash)

	cst := contact.New(db, lstor) //contact
	cs := contact2.New(cst, lserv)
	ch := contact3.New(cs, lhand)
	cr := contact4.New(ch)

	wsh := websocket.New(cs, lhand) //websocket
	wsr := websocket2.New(wsh)

	chst := chat.New(db, lstor) //chat
	csv := chat2.New(chst, lserv)
	chh := chat3.New(csv, lhand)
	chr := chat4.New(chh)

	grst := group.New(db, lstor)
	gsv := group2.New(grst, lserv)
	gh := group3.New(gsv, lhand)
	gr := group4.New(gh)

	cms := chat_message.New(db, lstor) //chatmessage
	cmsv := chat_message2.New(cms, lserv)
	cmh := chat_message3.New(cmsv, lhand)
	cmr := chat_message4.New(cmh)

	gms := group_message1.New(db, lstor) //groupmessage
	gmsv := group_message2.New(gms, lserv)
	gmh := group_message3.New(gmsv, lhand)
	gmr := group_message4.New(gmh)

	server := server.New(usr, asr, wsr, cr, cmr, gmr, chr, gr) //server
	router := server.Create()
	router.Run(":3000")
}
