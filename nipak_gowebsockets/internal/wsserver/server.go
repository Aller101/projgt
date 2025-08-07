package wsserver

import "net/http"

type WSServer interface {
	Start() error
}

type wsServer struct {
	srv *http.Server
	mux *http.ServeMux
}

func New(addr string) WSServer {
	mux := http.NewServeMux()
	wsServer := wsServer{
		mux: mux,
		srv: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
	return &wsServer
}

func (s *wsServer) Start() error {
	s.mux.HandleFunc("/test", s.testHandler)
	return s.srv.ListenAndServe()
}

func (s *wsServer) testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test is successful"))
}
