package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/sorovault/internal/store"
)

func TestNewMemory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		check func(t *testing.T, m *store.Memory)
	}{
		{
			name: "initial state has no contracts",
			check: func(t *testing.T, m *store.Memory) {
				ctx := context.Background()
				stats, err := m.Stats(ctx)
				require.NoError(t, err)
				assert.Equal(t, 0, stats.Contracts)
				assert.Equal(t, 0, stats.Versions)
			},
		},
		{
			name: "ping and close succeed on fresh memory store",
			check: func(t *testing.T, m *store.Memory) {
				ctx := context.Background()
				assert.NoError(t, m.Ping(ctx))
				assert.NoError(t, m.Close())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := store.NewMemory()
			tt.check(t, m)
		})
	}
}
