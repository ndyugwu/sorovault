package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soroworks/sorovault/internal/store"
)

func TestNewPostgresEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		databaseURL string
		opts        store.PoolOptions
		wantError   bool
	}{
		{
			name:        "invalid connection string",
			databaseURL: "not-a-valid-url",
			opts:        store.PoolOptions{},
			wantError:   true,
		},
		{
			name:        "empty connection string",
			databaseURL: "",
			opts:        store.PoolOptions{},
			wantError:   true,
		},
		{
			name:        "min conns exceeds max conns",
			databaseURL: "postgres://user:pass@localhost:5432/db",
			opts:        store.PoolOptions{MinConns: 10, MaxConns: 2},
			wantError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			p, err := store.NewPostgres(ctx, tt.databaseURL, tt.opts)
			if tt.wantError {
				assert.Error(t, err)
				assert.Nil(t, p)
			} else {
				assert.NoError(t, err)
				if p != nil {
					_ = p.Close()
				}
			}
		})
	}
}
