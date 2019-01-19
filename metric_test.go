package graphite

import (
	"fmt"
	"testing"
	"time"
)

func TestNewMetric(t *testing.T) {
	timestamp := time.Now().Unix()
	metric := NewMetric("test1", int(12), timestamp)
	if metric.Timestamp != timestamp {
		t.Fail()
	}

	if tmp, ok := metric.Value.(int); !ok {
		t.Fail()
	} else {
		if tmp != 12 {
			t.Fail()
		}
	}

	if metric.Name != "test1" {
		t.Fail()
	}
}

func TestMetric_ToString(t *testing.T) {
	timestamp := time.Now().Unix()
	metric := NewMetric("test1", int(12), timestamp)
	if metric.ToString() != fmt.Sprintf("test1 12 %d", timestamp) {
		t.Fail()
	}
}
func BenchmarkNewMetric(b *testing.B) {
	b.StopTimer()
	timestamp := time.Now().Unix()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = NewMetric("test1", int(12), timestamp)
	}
}

func BenchmarkMetric_ToString(b *testing.B) {
	b.StopTimer()
	timestamp := time.Now().Unix()
	metric := NewMetric("test1", int(12), timestamp)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = metric.ToString()
	}
}

func BenchmarkMetric_ToByte(b *testing.B) {
	b.StopTimer()
	timestamp := time.Now().Unix()
	metric := NewMetric("test1", int(12), timestamp)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = metric.ToByte()
	}
}
