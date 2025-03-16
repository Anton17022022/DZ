package middleware

import "net/http"

type WarpperWritter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WarpperWritter) WriteHeader (StatusCode int){
	w.ResponseWriter.WriteHeader(StatusCode)
	w.StatusCode = StatusCode
}