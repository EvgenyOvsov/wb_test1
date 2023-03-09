package server

import (
	"draft/datastore"
	"net/http"
)

type Server struct {
	ds *datastore.MemCache
}

func NewServer(ds *datastore.MemCache) *Server {
	srv := &Server{
		ds: datastore.NewDatastore(),
	}
	http.HandleFunc("/health", srv.healthCheck)
	http.HandleFunc("/", srv.handler) // Согласно задаче,
	// не знаю почему надо делать один роут, но бог с ним
	return srv
}

func (s *Server) Start() {
	defer func() {
		err := recover()
		if err != nil {
			println(err)
			s.Stop()
		}
		// Тут нужно описать процесс остановки сервера (деструктор)
		// Доп логи, закрытие соединений, правки в базе и т.д.
	}()

	println(http.ListenAndServe(":80", nil).Error())
}

func (s *Server) Stop() {
	// тут описываются условия graceful shutdown,
	// задачи которые нужно выполнить перед остановкой сервера
	// но мы просто играем
}
