package unmarshal

import (
	"benchmark/internal/service/data"
	"benchmark/proto/flightsearch"
	"github.com/json-iterator/go"
	"log"
	"testing"
)

var jsoniterdefault = jsoniter.ConfigDefault

func BenchmarkEncodingJsonIterDefault(b *testing.B) {
	dataByte := data.MockDataAsString()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var response *flightsearch.SearchFlightsResponse

		if err := jsoniterdefault.Unmarshal(dataByte, &response); err != nil {
			log.Println(err)
		}
	}
}
