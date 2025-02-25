//go:build unit

package solver

import (
	"testing"
	"time"
)

func TestIsTimestampRecent(t *testing.T) {
	tests := []struct {
		name      string
		timestamp int
		diff      int
		want      bool
	}{
		{
			name:      "current timestamp should be recent",
			timestamp: int(time.Now().UnixMilli()),
			diff:      1000, // 1 second
			want:      true,
		},
		{
			name:      "timestamp from future within diff should be recent",
			timestamp: int(time.Now().UnixMilli() + 500),
			diff:      1000,
			want:      true,
		},
		{
			name:      "timestamp from past within diff should be recent",
			timestamp: int(time.Now().UnixMilli() - 500),
			diff:      1000,
			want:      true,
		},
		{
			name:      "timestamp too far in future should not be recent",
			timestamp: int(time.Now().UnixMilli() + 2000),
			diff:      1000,
			want:      false,
		},
		{
			name:      "timestamp too far in past should not be recent",
			timestamp: int(time.Now().UnixMilli() - 2000),
			diff:      1000,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isTimestampRecent(tt.timestamp, tt.diff)
			if got != tt.want {
				t.Errorf("isTimestampRecent() = %v, want %v", got, tt.want)
			}
		})
	}
}
