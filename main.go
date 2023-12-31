package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var supportedMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

func main() {
	http.HandleFunc("/", handleRequest)

	var port string = "8080"
	if len(os.Args) >= 3 {
		if os.Args[1] == "-p" || os.Args[1] == "--port" {
			port = os.Args[2]
		}
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodPut:
		handlePut(w, r)
	case http.MethodPatch:
		handlePatch(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	case http.MethodHead:
		handleHead(w, r)
	case http.MethodOptions:
		handleOptions(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		httpError(w, http.StatusBadRequest, "URL query parameter is required")
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to fetch the URL: "+err.Error())
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		httpError(w, http.StatusBadRequest, "URL query parameter is required")
		return
	}

	clientRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpError(w, http.StatusBadRequest, "Failed to read request body: "+err.Error())
		return
	}
	defer r.Body.Close()

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(clientRequestBody))
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to create new request: "+err.Error())
		return
	}
	copyHeaders(req.Header, r.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to perform POST request: "+err.Error())
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		httpError(w, http.StatusBadRequest, "URL query parameter is required")
		return
	}

	clientRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpError(w, http.StatusBadRequest, "Failed to read request body: "+err.Error())
		return
	}
	defer r.Body.Close()

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(clientRequestBody))
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to create new request: "+err.Error())
		return
	}
	copyHeaders(req.Header, r.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to perform PUT request: "+err.Error())
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func handlePatch(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		httpError(w, http.StatusBadRequest, "URL query parameter is required")
		return
	}

	clientRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpError(w, http.StatusBadRequest, "Failed to read request body: "+err.Error())
		return
	}
	defer r.Body.Close()

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(clientRequestBody))
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to create new request: "+err.Error())
		return
	}
	copyHeaders(req.Header, r.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to perform PATCH request: "+err.Error())
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		httpError(w, http.StatusBadRequest, "URL query parameter is required")
		return
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to create new request: "+err.Error())
		return
	}
	copyHeaders(req.Header, r.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to perform DELETE request: "+err.Error())
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func handleHead(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		httpError(w, http.StatusBadRequest, "URL query parameter is required")
		return
	}

	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to create new request: "+err.Error())
		return
	}
	copyHeaders(req.Header, r.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to perform HEAD request: "+err.Error())
		return
	}
	defer resp.Body.Close()

	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		httpError(w, http.StatusBadRequest, "URL query parameter is required")
		return
	}

	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to create new request: "+err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "Failed to perform OPTIONS request: "+err.Error())
		return
	}
	defer resp.Body.Close()

	targetSupportedMethods := parseAllowHeader(resp.Header.Get("Allow"))
	allowedMethods := intersectMethods(targetSupportedMethods, supportedMethods)

	allowedMethodsStr := strings.Join(allowedMethods, ", ")
	w.Header().Set("Allow", allowedMethodsStr)
	w.Header().Set("Access-Control-Allow-Methods", allowedMethodsStr)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusNoContent)
}

func httpError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "{\"error\": \"%s\"}\n", message)
	log.Printf("Error %d: %s", statusCode, message)
}

func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func parseAllowHeader(header string) []string {
	return strings.Split(header, ",")
}

func intersectMethods(a, b []string) []string {
	m := make(map[string]bool)
	var intersection []string

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; ok {
			intersection = append(intersection, item)
		}
	}

	return intersection
}
