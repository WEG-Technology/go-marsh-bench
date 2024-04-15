package marshal

import (
	"benchmark/internal/service/data"
	"encoding/json"
	"testing"
)

func BenchmarkEncodingJson(b *testing.B) {
	result := data.MockData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(result); err != nil {
			b.Error(err)
		}
	}
}
