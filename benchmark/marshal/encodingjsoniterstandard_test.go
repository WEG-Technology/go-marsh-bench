package marshal

import (
	"benchmark/internal/service/data"
	"github.com/json-iterator/go"
	"testing"
)

var jsonstandard = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkEncodingJsonIterStandard(b *testing.B) {
	result := data.MockData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := jsonstandard.Marshal(result); err != nil {
			b.Error(err)
		}
	}
}
