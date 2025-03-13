package todo

import (
	"context"
	"net/http"
	"time"
)

// Абстракция нахуя-то над сервером
type Server struct {
	httpServer *http.Server
}

// Запуск
func (s *Server) Run(port string, handler http.Handler) error { // Хз, норм ли пробрасывать порты так, а не через конфиг
	s.httpServer = &http.Server{
		Addr:           ":" + port, // Нахуя то
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// GS
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
