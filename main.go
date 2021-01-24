package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
Needs following environment variables:

PROCESS_SEEKER_PYTHON_INTERPRETER
PROCESS_SEEKER_SCRIPT_PATH
*/
var (
	pythonInterpreter string
	scriptPath        string
)

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type requestBody struct {
	ProcessName string `json:"processName"`
}

func main() {

	pythonInterpreter = os.Getenv("PROCESS_SEEKER_PYTHON_INTERPRETER")
	if pythonInterpreter == "" {
		pythonInterpreter = "python"
	}

	scriptPath = os.Getenv("PROCESS_SEEKER_SCRIPT_PATH")
	if scriptPath == "" {
		scriptPath = "find_process.py"
	}
	if len(os.Args) <= 1 {
		log.Fatal("usage: process_seeker  <port>")
	}
	port := os.Args[1]

	http.HandleFunc("/search", searchForProcess)
	log.Println("Start server at port: ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Panicln("Server error", err)
	}
}

func searchForProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		resp := response{
			Status:  -1,
			Message: "Bad method",
		}

		err := json.NewEncoder(w).Encode(resp)
		respErr(err)
		return
	}
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("PARSING REQUEST BODY ERROR: ", err)

		resp := response{
			Status:  -1,
			Message: "Parsing body error",
		}

		err := json.NewEncoder(w).Encode(resp)
		respErr(err)
		return
	}
	result := find(body.ProcessName, pythonInterpreter, scriptPath)
	if result == 0 {
		resp := response{
			Status:  0,
			Message: "Process exists",
		}

		err := json.NewEncoder(w).Encode(resp)
		respErr(err)
	} else {
		resp := response{
			Status:  1,
			Message: "Process does not exist",
		}

		err := json.NewEncoder(w).Encode(resp)
		respErr(err)
	}

}

func respErr(err error) {
	if err != nil {
		log.Println("RESPONSE ERROR: ", err)
	}
}
