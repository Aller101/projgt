package wsserver

import (
	"net/http"
	"sync"
	"time"

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
	srv       *http.Server
	mux       *http.ServeMux
	upgrader  *websocket.Upgrader
	wsClients map[*websocket.Conn]struct{}
	mutex     *sync.RWMutex
	broadcast chan *wsMessage
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
		wsClients: map[*websocket.Conn]struct{}{},
		mutex:     &sync.RWMutex{},
		broadcast: make(chan *wsMessage),
	}
	return &wsServer
}

func (s *wsServer) Start() error {
	s.mux.Handle("/", http.FileServer(http.Dir(templDir)))
	s.mux.HandleFunc("/test", s.testHandler)
	s.mux.HandleFunc("/ws", s.wsHandler)
	go s.writeToClients()
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
	log.Infof("Client with address %s\n", conn.RemoteAddr().String())
	s.mutex.Lock()
	s.wsClients[conn] = struct{}{}
	s.mutex.Unlock()

	go s.readFromClient(conn)
}

func (s *wsServer) readFromClient(conn *websocket.Conn) {
	for {
		msg := new(wsMessage)
		err := conn.ReadJSON(msg)
		if err != nil {
			log.Errorf("Error with reading from ws: %v", err)
			break
		}
		msg.IPAddress = conn.RemoteAddr().String()
		msg.Time = time.Now().Format("15:00")
		s.broadcast <- msg
	}
	s.mutex.Lock()
	delete(s.wsClients, conn)
	s.mutex.Unlock()
}

func (s *wsServer) writeToClients() {
	for msg := range s.broadcast {
		s.mutex.RLock()
		for client := range s.wsClients {

			// if err := client.WriteJSON(msg); err != nil {
			// 	log.Errorf("Error with writing msg: %v", err)
			// 	continue
			// }

			func() {
				if err := client.WriteJSON(msg); err != nil {
					log.Errorf("Error with writing msg: %v", err)
				}
			}()
		}
		s.mutex.RUnlock()
	}
}
