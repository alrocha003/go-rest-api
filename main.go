package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

type Person struct{
	Id int32
	Name string
}

type About struct {
	Message string	
}

type ResponseMessage struct{
	Status int32
	Message string
}

const GET = "GET"
const API_ROUTE = "/api"

func Sum(w http.ResponseWriter, r *http.Request){
	num, err := strconv.Atoi(r.URL.Query().Get("num")) 
	num2, err := strconv.Atoi(r.URL.Query().Get("num2"))
	
	if err != nil{
		return;
	}
	
	result := (num + num2)
	json.NewEncoder(w).Encode(About{Message: strconv.Itoa(result)})
}

func Sub(writer http.ResponseWriter, request *http.Request){
	num, err := strconv.Atoi(request.URL.Query().Get("num"))
	num2, err := strconv.Atoi(request.URL.Query().Get("num2"))
	
	if err != nil{
		json.NewEncoder(writer).Encode(ResponseMessage{Status:500, 
			Message: "Ocorreu um erro na conversão: " + err.Error()})
			}else{
				result := num - num2
				json.NewEncoder(writer).Encode(ResponseMessage{ Status: 200, Message: strconv.Itoa(result)})
			}
		}
		
func Mult(writer http.ResponseWriter, request *http.Request){
	num, err := strconv.Atoi(request.URL.Query().Get("num"))
	num2, err := strconv.Atoi(request.URL.Query().Get("num2"))
			
	if err != nil{
		json.NewEncoder(writer).Encode(ResponseMessage{Status: 500, Message: "Ocorreu um erro na conversão: " + err.Error()})
			}else{
				result := num - num2
				json.NewEncoder(writer).Encode(ResponseMessage{Status: 200, Message: strconv.Itoa(result)})
			}
		}
			
			func Div(writer http.ResponseWriter, request *http.Request){
				num, err := strconv.Atoi(request.URL.Query().Get("num"))
				num2, err := strconv.Atoi(request.URL.Query().Get("num2"))
				
				if err != nil{
					json.NewEncoder(writer).Encode(ResponseMessage{Status: 500, Message: "Ocorreu um erro na conversão: " + err.Error()})
					}else{
						result := num / num2
						json.NewEncoder(writer).Encode(ResponseMessage{Status: 200, Message: strconv.Itoa(result)})
					}
				}
				
				func main(){
					router := mux.NewRouter()
					router.HandleFunc(API_ROUTE+"/divv", Div).Methods(GET)
					router.HandleFunc(API_ROUTE+"/mult", Mult).Methods(GET)
					router.HandleFunc(API_ROUTE+"/sum", Sum).Methods(GET)
					router.HandleFunc(API_ROUTE+"/sub", Sub).Methods(GET)
					port := ":8080"
					
					fmt.Println("Aplicação executando em http://localhost", port)
					log.Fatal(http.ListenAndServe(port, router))
				}
				