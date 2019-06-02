package restservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type measurement struct {
	SensorName string `json:"sensor"`
	Value      int    `json:"value"`
	Location   string `json:"location"`
}

var mongoServer = "mongodb://mongo:27017"

// RegisterMongoEndpoints registers mongo endpoints to the router
func RegisterMongoEndpoints(r *mux.Router) {
	r.HandleFunc("/mongoView", mongoView).Methods("GET")
	r.HandleFunc("/mongoAddReading", mongoPostReading).Methods("POST")
}

func mongoView(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	defer cancel()
	if err != nil {
		http.Error(w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
				err.Error()), http.StatusInternalServerError)
		return
	}

	collection := client.Database("SensorData").Collection("test_measurements")
	findOptions := options.Find()
	var results []*measurement
	cur, err := collection.Find(ctx, bson.D{{"sensorname", "snake"}}, findOptions)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
				err.Error()), http.StatusInternalServerError)
		return
	}

	for cur.Next(ctx) {
		var elem measurement
		err := cur.Decode(&elem)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
					err.Error()), http.StatusInternalServerError)
			return
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		http.Error(w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
				err.Error()), http.StatusInternalServerError)
		return
	}
	cur.Close(ctx)
	j, err := json.Marshal(results)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
				err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func mongoPostReading(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	defer cancel()
	if err != nil {
		http.Error(w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
				err.Error()), http.StatusInternalServerError)
		return
	}

	var m measurement
	json.Unmarshal(b, &m)

	collection := client.Database("SensorData").Collection("test_measurements")

	res, err := collection.InsertOne(ctx, m)

	if err != nil {
		http.Error(w,
			fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError),
				err.Error()), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("%s", res.InsertedID)))
}
