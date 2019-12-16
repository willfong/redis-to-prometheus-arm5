package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("Starting service...")
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/metrics", HandleMetrics)
	http.ListenAndServe(":8080", nil)
}

// HandleRoot root
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "redis-to-prometheus")
}

// HandleMetrics redis to Prometheus
func HandleMetrics(w http.ResponseWriter, r *http.Request) {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		log.Fatal(err)
	}
	for _, k := range keys {
		v, err := redis.String(conn.Do("GET", k))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s %s\n", k, v)
	}
}
