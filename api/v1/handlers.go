package v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/phonghaido/golang-mongodb/data"
	"github.com/phonghaido/golang-mongodb/db"
	"github.com/phonghaido/golang-mongodb/handlers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var InMemory = data.InMemory{
	Data: make(map[string]string),
}

func HandlePostMongoDB(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var reqPayload data.MongoDBRequestPayload
	err = json.Unmarshal(body, &reqPayload)
	if err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	startDate, err := time.Parse("2006-01-02", reqPayload.StartDate)
	if err != nil {
		http.Error(w, "Error parsing start date", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", reqPayload.EndDate)
	if err != nil {
		http.Error(w, "Error parsing end date", http.StatusBadRequest)
		return
	}

	pipeLine := bson.A{
		bson.M{
			"$project": bson.M{
				"key":       1,
				"createdAt": 1,
				"counts":    1,
				"totalCount": bson.M{
					"$sum": "$counts",
				},
				"_id": 0,
			},
		},
		bson.M{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": reqPayload.MinCount,
					"$lt": reqPayload.MaxCount,
				},
				"createdAt": bson.M{
					"$gte": primitive.NewDateTimeFromTime(startDate),
					"$lte": primitive.NewDateTimeFromTime(endDate),
				},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := db.ConnectToMongoDB(ctx)
	records := db.FetchRecords(ctx, &client, pipeLine)
	responsePayload := data.MongoDBResponsePayload{
		Code:   0,
		Msg:    "Success",
		Record: handlers.ConvertRecordResponse(records),
	}
	jsonResp, err := json.Marshal(responsePayload)
	if err != nil {
		http.Error(w, "Error parsing response payload", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func HandleInMemory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePostInMemory(w, r)
	case http.MethodGet:
		handleGetInMemory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func handlePostInMemory(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error parsing request payload", http.StatusBadRequest)
		return
	}
	var reqPayload data.InMemoryPayload
	err = json.Unmarshal(body, &reqPayload)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusInternalServerError)
		return
	}
	InMemory.Data[reqPayload.Key] = reqPayload.Value
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Success"))
}

func handleGetInMemory(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	key := queryParams.Get("key")

	var data data.InMemoryPayload
	data.Key = key
	data.Value = InMemory.Data[key]

	jsonResp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error parsing response payload", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
