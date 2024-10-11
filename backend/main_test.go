package main 

import (
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"fmt"
)

func TestAPI_CanOnlyAcceptPUT(t *testing.T){
	req, err := http.NewRequest("GET", "/stackmachine", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StackMachineAPI)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	
}
func TestAPI_NormalSum(t *testing.T){
	body := []byte(`{"command": "3 4 3 5 5 1 1 1 SUM"}`)
	req, err := http.NewRequest("POST", "/stackmachine", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StackMachineAPI)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := struct {
        Result int    `json:"result"`
        Error  string `json:"error"`
    }{
        Result: 23,
        Error:  "",
    }

    var actual struct {
        Result int    `json:"result"`
        Error  string `json:"error"`
    }
    err = json.Unmarshal(rr.Body.Bytes(), &actual)
    if err != nil {
        t.Fatalf("Could not parse response: %v", err)
    }

    if actual != expected {
        t.Errorf("handler returned unexpected body: got %+v want %+v", actual, expected)
    }
}

func TestAPI_AcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    int
		shouldError bool
	}{
		{name: "empty error", commands: "", expected: 0, shouldError: true},
		{name: "add overflow", commands: "50000 DUP +", expected: 0, shouldError: true},
		{name: "too few add", commands: "99 +", expected: 0, shouldError: true},
		{name: "too few minus", commands: "99 -", expected: 0, shouldError: true},
		{name: "too few multiply", commands: "99 *", expected: 0, shouldError: true},
		{name: "empty stack", commands: "99 CLEAR", expected: 0, shouldError: true},
		{name: "sum single value", commands: "99 SUM", expected: 99, shouldError: false},
		{name: "sum empty", commands: "SUM", expected: 0, shouldError: true},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22, shouldError: false},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0, shouldError: true},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5, shouldError: false},
		{name: "single integer", commands: "9876", expected: 9876, shouldError: false},
		{name: "invalid command", commands: "DOGBANANA", expected: 0, shouldError: true},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60, shouldError: false},
		{name: "minus", commands: "2 5 -", expected: 3, shouldError: false},
		{name: "underflow minus", commands: "5 2 -", expected: 0, shouldError: true},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000, shouldError: false},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000, shouldError: false},
		{name: "overflow plus", commands: "50000 1 +", expected: 0, shouldError: true},
		{name: "overflow single value", commands: "50001", expected: 0, shouldError: true},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0, shouldError: false},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0, shouldError: true},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100, shouldError: false},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0, shouldError: false},
		{name: "clear empties stack", commands: "CLEAR", expected: 0, shouldError: true},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23, shouldError: false},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, shouldError: false},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0, shouldError: true},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120, shouldError: false},
	}

	for _, test := range tests {
		fmt.Println(test.name)
		requestBody, _ := json.Marshal(map[string]string{
			"command": test.commands,
		})

		req, err := http.NewRequest("POST", "/stackmachine", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(StackMachineAPI)

		handler.ServeHTTP(rr, req)

		var response struct {
			Result int    `json:"result"`
			Error  string `json:"error"`
		}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Could not parse response: %v", err)
		}

		if test.shouldError {
			if response.Error == "" {
				t.Errorf("%s (%s) Expected error, but got none", test.name, test.commands)
			}
		} else if response.Result != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, response.Result, test.expected)
		}
	}
}
