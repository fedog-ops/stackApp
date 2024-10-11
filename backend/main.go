package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"stackApp/stackmachine"
)

type ResponseStruct struct {
	Result int    `json:"result"`
	Error  string `json:"error"`
}

func StackMachineAPI(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
        http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var requestBody struct {
		Command string `json:"command"`
	}

	json.NewDecoder(request.Body).Decode(&requestBody)

	result, err := stackmachine.StackMachine(requestBody.Command)
	var response ResponseStruct
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Result = result
	}

	writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(response)

}

func main() {

	http.HandleFunc("/stackmachine", StackMachineAPI)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
