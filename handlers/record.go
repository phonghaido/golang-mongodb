package handlers

import (
	"time"

	"github.com/phonghaido/golang-mongodb/data"
)

func ConvertRecordResponse(records []data.MongoDBRecord) []data.RecordResponse {
	var result []data.RecordResponse
	for _, record := range records {
		result = append(result, data.RecordResponse{
			Key:        record.Key,
			CreatedAt:  record.CreatedAt.Format(time.RFC3339),
			TotalCount: calculateSum(record.Count),
		})
	}
	return result
}

func calculateSum(slice []int) int {
	sum := 0
	for _, el := range slice {
		sum = sum + el
	}
	return sum
}
