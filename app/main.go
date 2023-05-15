package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", os.Getenv("REDIS")+":6379")
	HandleError(err)
	fmt.Printf("Redis responded at " + os.Getenv("REDIS") + ":6379\n")
	return c
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "This page intentionally left blank."
		jsonResp, err := json.Marshal(resp)
		HandleError(err)
		w.Write(jsonResp)
		log.Printf("blank called\n")

	})

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		conn := RedisConnect()
		_, err := conn.Do("INCR", "hello")
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Hello, World!"
		s, err := redis.String(conn.Do("GET", "hello"))
		resp["requestCount"] = s
		jsonResp, err := json.Marshal(resp)
		HandleError(err)
		w.Write(jsonResp)
		log.Printf("hello called\n")
	})
	fmt.Printf("Server running (port=8080)\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
