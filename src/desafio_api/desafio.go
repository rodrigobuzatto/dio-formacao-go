package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Client struct {
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Cidade   string `json:"cidade"`
}

type Clients struct {
	Clientes []Client `json:"clientes"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler).Methods("GET", "POST")
	router.HandleFunc("/{name}", handler).Methods("PUT", "DELETE")

	http.ListenAndServe(":8000", router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getClients(w, r)
	case http.MethodPost:
		saveClient(w, r)
	case http.MethodPut:
		updateClient(w, r)
	case http.MethodDelete:
		deleteClient(w, r)
	}
}

func getClientsFromJsonFile() Clients {
	jsonFile, err := os.Open("clients.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var clients Clients

	json.Unmarshal(byteValue, &clients)

	return clients
}

func saveClientsToJsonFile(clients Clients) {
	file, _ := json.MarshalIndent(clients, "", " ")

	_ = ioutil.WriteFile("clients.json", file, 0644)
}

func getClients(w http.ResponseWriter, r *http.Request) {
	clients := getClientsFromJsonFile()

	json.NewEncoder(w).Encode(clients)
}

func saveClient(w http.ResponseWriter, r *http.Request) {
	clients := getClientsFromJsonFile()

	reqBody, _ := ioutil.ReadAll(r.Body)

	var client Client

	json.Unmarshal(reqBody, &client)

	newClients := append(clients.Clientes, client)

	var savedClients Clients

	savedClients.Clientes = newClients

	saveClientsToJsonFile(savedClients)

	json.NewEncoder(w).Encode("Cliente salvo com sucesso!")
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	clients := getClientsFromJsonFile()

	vars := mux.Vars(r)
	name := vars["name"]

	for index, client := range clients.Clientes {
		if client.Nome == name {
			clients.Clientes = append(clients.Clientes[:index], clients.Clientes[index+1:]...)
		}
	}

	saveClientsToJsonFile(clients)

	json.NewEncoder(w).Encode("Cliente removido com sucesso")
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	clients := getClientsFromJsonFile()

	message := "Cliente não encontrado"

	vars := mux.Vars(r)

	name := vars["name"]

	for index, client := range clients.Clientes {
		if client.Nome == name {
			clients.Clientes = append(clients.Clientes[:index], clients.Clientes[index+1:]...)

			var updatedClient Client

			json.NewDecoder(r.Body).Decode(&updatedClient)

			clients.Clientes = append(clients.Clientes, updatedClient)

			saveClientsToJsonFile(clients)

			message = "Cliente atualizado com sucesso"
			break
		}
	}

	json.NewEncoder(w).Encode(message)
}
