package handlers

import (
	"log"
	"net/http"
)

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

type Goodbye struct {
	l *log.Logger
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Good bye!"))
}
