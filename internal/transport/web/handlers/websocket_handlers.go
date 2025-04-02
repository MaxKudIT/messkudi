package handlers

//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//func WebsocketHandler(c *gin.Context) {
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		log.Fatal(exceptions.ConnectionExcWebsocket().Error())
//	}
//	for {
//		messageType, message, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(exceptions.GetMessageExc())
//		}
//		log.Printf("Получено сообщение: %s", message)
//		if err := conn.WriteMessage(messageType, message); err != nil {
//			log.Println(exceptions.SendMessageErr().Error())
//		}
//	}
//}
