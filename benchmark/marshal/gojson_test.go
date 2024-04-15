package marshal

import (
	"benchmark/internal/service/data"
	"github.com/goccy/go-json"
	"testing"
)

func BenchmarkGoJson(b *testing.B) {
	result := data.MockData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(result); err != nil {
			b.Error(err)
		}
	}
}
