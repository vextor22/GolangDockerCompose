package restservice

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
}

// RegisterRedisEndpoints registers redis endpoints to the router
func RegisterRedisEndpoints(r *mux.Router) {
	r.HandleFunc("/redisCountInc", redisCountInc).Methods("GET")
	r.HandleFunc("/redisView", redisCountView).Methods("GET")
}
func redisCountInc(w http.ResponseWriter, r *http.Request) {
	res, err := redisClient.Incr("counter").Result()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "This has been seen and incremented: %d times", res)
}

func redisCountView(w http.ResponseWriter, r *http.Request) {
	res, err := redisClient.Get("counter").Result()

	if err != nil {
		http.Error(w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
				err.Error()), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "This has been seen by the incrementer endpoint: %s times", res)
}
