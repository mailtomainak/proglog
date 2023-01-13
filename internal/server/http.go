package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(addr string) *http.Server {
	httpsrv := newHTTPServer()
	r := mux.NewRouter()
	r.HandleFunc("/", httpsrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpsrv.handleConsume).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r}
}

type httpServer struct {
	Log *Log
}

// Produce appends to the logs.
func (s *httpServer) handleProduce(writer http.ResponseWriter, request *http.Request) {
	var req ProduceRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	offset, err := s.Log.Append(req.Record)

	if err == ErrorOffsetNotFound {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ProduceResponse{Offset: offset}
	err = json.NewEncoder(writer).Encode(res)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Consume returns the log at a particular offset.
func (s *httpServer) handleConsume(writer http.ResponseWriter, request *http.Request) {
	var req ConsumeRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	record, err := s.Log.Read(req.Offset)

	if err == ErrorOffsetNotFound {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ConsumeResponse{Record: record}
	err = json.NewEncoder(writer).Encode(res)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func newHTTPServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}
