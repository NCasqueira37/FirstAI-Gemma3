package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Model  string
	Prompt string
	Stream bool
}

type Response struct {
	Model                string
	Created_At           string
	Response             string
	Done                 bool
	Context              []int
	Total_Duration       int
	Load_Duration        int
	Prompt_Eval_Count    int
	Prompt_Eval_Duration int
	Eval_Count           int
	Eval_Duration        int
}

func main() {
	// REST API
	url := "http://localhost:11434/api/generate"

	// Get prompt from user
	fmt.Println("Send a message to the AI")
	reader := bufio.NewReader(os.Stdin)
	prompt, _ := reader.ReadString('\n')
	prompt = strings.TrimSpace(prompt)

	// Let user know that the response is generating
	fmt.Println("Generating...")

	// Create the payload
	// Encode the data into JSON
	payload := Request{Model: "gemma3", Prompt: prompt}
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create a client and issue a POST
	client := http.Client{}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	// Use the response and convert it to a slice of bytes
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println(string(bodyBytes))

	// Decode the JSON response
	var result Response
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Use the response
	fmt.Println(result.Response)

}
