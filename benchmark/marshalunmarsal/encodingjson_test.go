package marshalunmarshal

import (
	"benchmark/internal/service/data"
	"benchmark/proto/flightsearch"
	"encoding/json"
	"log"
	"testing"
)

func BenchmarkEncodingJson(b *testing.B) {
	dataByte := data.MockDataAsString()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var response *flightsearch.SearchFlightsResponse

		if err := json.Unmarshal(dataByte, &response); err != nil {
			log.Fatalf("Failed to unmarshal JSON: %s", err)
		}

		if _, err := json.Marshal(response); err != nil {
			log.Fatalf("Failed to marshal JSON: %s", err)
		}
	}
}
