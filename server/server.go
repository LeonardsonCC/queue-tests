package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/LeonardsonCC/queue-tests/queue"
)

type Server struct {
	Queue *queue.Queue
}

func NewServer() *Server {
	return &Server{
		Queue: queue.NewQueue(),
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/ping", s.getPing)
	http.HandleFunc("/queue/add", s.addList)
	http.HandleFunc("/queue/pop", s.popList)
	http.HandleFunc("/queue/length", s.lengthOfList)

	return http.ListenAndServe(":8080", nil)
}

func (s *Server) getPing(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("got /ping request")
	io.WriteString(w, `{"message": "pong"}`)
}

// addList add the number 1 as value to the linked list
func (s *Server) addList(w http.ResponseWriter, _ *http.Request) {
	s.Queue.Append(1)

	io.WriteString(w, fmt.Sprintf(`{"message": "added", "length": %d}`, s.Queue.Len()))
}

// popList remove the last item from the linked list
func (s *Server) popList(w http.ResponseWriter, _ *http.Request) {
	s.Queue.Pop()

	io.WriteString(w, fmt.Sprintf(`{"message": "added", "length": %d}`, s.Queue.Len()))
}

// lengthOfList the name is very explicit what it does
func (s *Server) lengthOfList(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, fmt.Sprintf(`{"length": %d}`, s.Queue.Len()))
}
