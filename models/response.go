package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	IsSuccess   bool        `json:"isSuccess"`
	Status      int         `json:"status"`
	StatusCode  string      `json:"statusCode"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Metadata    []string    `json:"metadata"`
	contentType string
	w           http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		w:           w,
		contentType: "application/json",
	}
}

func (resp *Response) Send() {
	resp.w.Header().Set("Content-Type", resp.contentType)
	resp.w.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.w, string(output))
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Message = "Resource Not Found"
}

func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

func (r *Response) NotProcesableEntity() {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Invalid Input"
}

func SendNotProcesableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotProcesableEntity()
	response.Send()
}