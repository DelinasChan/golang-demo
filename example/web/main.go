package main

import(
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func getQuery(s socketio.Conn,key string) string{
	url := s.URL()
	return url.Query().Get(key)
}

func main(){
	r := gin.Default();

	server := socketio.NewServer(nil)

	// server.OnEvent(namespace, event, func...)
	server.OnConnect("/chat", func(s socketio.Conn) error {
		s.SetContext("")
		url := s.URL()
		room := url.Query().Get("room")
		// server.JoinRoom("chat",room,s)
		s.Join(room)
		payload :=  map[string]interface{}{
			"id":55688,
		}

		fmt.Printf("Conneect URL: %v\n",url)

		s.Emit("joinRes",payload)
		return nil
	})


	server.OnEvent("/chat","test",func(s socketio.Conn)error{
		url := s.URL()

		room := url.Query().Get("room");
		userId := url.Query().Get("userId");

		fmt.Printf("URL: %v\n",url)
		fmt.Printf("Room: %v UserId: %v\n",room , userId)

		return nil
	})

	server.OnEvent("/chat", "join", func(s socketio.Conn, data string) string {
		return "msg Response"
	})

	//發送訊息
	server.OnEvent("/chat", "message", func(s socketio.Conn, data map[string]interface{}) map[string]interface{} {

		result := map[string]interface{}{
			"message":data["message"],
		}

		room := getQuery(s,"room")
		np := s.Namespace()
		server.BroadcastToRoom(np,room, "res", result)
		return result
	})

	server.OnEvent("/chat", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/chat", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("Connect dIsocount")
	})

	r.GET("/socket-page",func(c *gin.Context){
		c.HTML(200,"socket.html",nil)
	})

	go server.Serve()
	defer server.Close()
	
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	//設定本地靜態頁面資料夾
    r.LoadHTMLGlob("page/*")
	r.Run(":7000")
}