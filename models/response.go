package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
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
		IsSuccess:   true,
		Status:      http.StatusOK,
		StatusCode:  "OK",
		Message:     "Users retrieved successfully",
		Metadata:    []string{},
		w:           w,
		contentType: "application/json",
	}
}

func (resp *Response) Send(format string) {
	resp.w.Header().Set("Content-Type", resp.contentType)
	resp.w.WriteHeader(resp.Status)

	var output []byte

	switch format {
	case "xml":
		output, _ = xml.Marshal(&resp)
		resp.w.Header().Set("Content-Type", "text/xml")
	case "yaml":
		output, _ = yaml.Marshal(&resp)
	default:
		output, _ = json.Marshal(&resp)
		resp.w.Header().Set("Content-Type", "application/json")
	}

	fmt.Fprintln(resp.w, string(output))
}

func SendData(w http.ResponseWriter, data interface{}, format string) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send(format)
}

func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Message = "Resource Not Found"
}

func SendNotFound(w http.ResponseWriter, format string) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send(format)
}

func (r *Response) NotProcesableEntity() {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Invalid Input"
}

func SendNotProcesableEntity(w http.ResponseWriter, format string) {
	response := CreateDefaultResponse(w)
	response.NotProcesableEntity()
	response.Send(format)
}
