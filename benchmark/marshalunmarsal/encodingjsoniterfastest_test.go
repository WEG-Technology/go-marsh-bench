package marshalunmarshal

import (
	"benchmark/internal/service/data"
	"benchmark/proto/flightsearch"
	"github.com/json-iterator/go"
	"testing"
)

var jsonfastest = jsoniter.ConfigFastest

func BenchmarkEncodingJsonIterFastest(b *testing.B) {
	dataByte := data.MockDataAsString()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var response *flightsearch.SearchFlightsResponse

		if err := jsonfastest.Unmarshal(dataByte, &response); err != nil {
			b.Error(err)
		}

		if _, err := jsonfastest.Marshal(response); err != nil {
			b.Fatalf("Failed to marshal JSON: %s", err)
		}
	}
}
