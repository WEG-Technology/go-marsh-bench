package unmarshal

import (
	"benchmark/internal/service/data"
	"benchmark/proto/flightsearch"
	"github.com/goccy/go-json"
	"log"
	"testing"
)

func BenchmarkGoJson(b *testing.B) {
	dataByte := data.MockDataAsString()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var response *flightsearch.SearchFlightsResponse

		if err := json.Unmarshal(dataByte, &response); err != nil {
			log.Println(err)
		}
	}
}
