package lilymetrics

import (
	"log"
	"net/http"

	// "github.com/elazarl/goproxy/transport"
	// "github.com/elazarl/goproxy/transport"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	transport "github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/gorilla/mux"
)

func setupRoutes() {
	server = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	defer server.Close()
	//setup postgress notification listener
	listener := initializeListener()
	go handleNotifications(server, listener)

	//setup socket routes
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		// log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		go runCowSay(msg)
	})

	server.OnEvent("/", "task", func(s socketio.Conn, msg string) {
		go runTask(msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		log.Println("chat:", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("socket error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("disconnected/closed", reason)
	})

	//start the socket server
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	router := mux.NewRouter()
	//messages comming in from socket front end get routed to socket server
	router.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})

	router.HandleFunc("/metrics-dashboard/log", handleEvent).Methods("POST") // Only keep the event handler route
	router.HandleFunc("/metrics-dashboard/job", handleJobEvent).Methods("POST")
	router.HandleFunc("/metrics-dashboard/status", handleStatusEvent).Methods("POST")
	router.HandleFunc("/metrics-dashboard/matcher", handleMatcherEvent).Methods("POST")
	router.HandleFunc("/metrics-dashboard/result", handleJobResult).Methods("POST")
	router.HandleFunc("/log-updates", handleGetLogUpdates).Methods("GET")
	router.HandleFunc("/job-updates", handleGetJobUpdates).Methods("GET")

	router.PathPrefix("/files").Handler(http.StripPrefix("/files", http.FileServer(http.Dir("/tmp/lilypad/data/downloaded-files/"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("dashboard/build"))))
	http.Handle("/", http.FileServer(http.Dir("dashboard/build")))
	log.Fatal(http.ListenAndServe(":8000", router))
}
