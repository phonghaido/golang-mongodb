package data

import "time"

type MongoDBRecord struct {
	Key       string    `bson:"key"`
	CreatedAt time.Time `bson:"createdAt"`
	Count     []int     `bson:"counts"`
}

type MongoDBRequestPayload struct {
	StartDate string `json: "startDate"`
	EndDate   string `json: "endDate"`
	MinCount  int    `json: "minCount"`
	MaxCount  int    `json: "maxCount"`
}

type MongoDBResponsePayload struct {
	Code   int              `json:"code"`
	Msg    string           `json:"msg"`
	Record []RecordResponse `json:"records"`
}

type RecordResponse struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int    `json:"totalCount"`
}

type InMemory struct {
	Data map[string]string `json:"data"`
}

type InMemoryPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
