package wsserver

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	templDir = "./web/templates/html"
)

type WSServer interface {
	Start() error
}

type wsServer struct {
	srv      *http.Server
	mux      *http.ServeMux
	upgrader *websocket.Upgrader
}

func New(addr string) WSServer {
	mux := http.NewServeMux()
	wsServer := wsServer{
		mux: mux,
		srv: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
	return &wsServer
}

func (s *wsServer) Start() error {
	s.mux.Handle("/", http.FileServer(http.Dir(templDir)))
	s.mux.HandleFunc("/test", s.testHandler)
	s.mux.HandleFunc("/ws", s.wsHandler)
	return s.srv.ListenAndServe()
}

func (s *wsServer) testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test is successful"))
}

func (s *wsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Error with ws connection: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Infof(conn.RemoteAddr().String())
}
