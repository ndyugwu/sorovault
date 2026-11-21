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
			name:        "empty connection string",
			databaseURL: "",
			opts:        store.PoolOptions{},
			wantError:   true,
		},
		{
			name:        "invalid connection string scheme",
			databaseURL: "not-a-postgres-url",
			opts:        store.PoolOptions{},
			wantError:   true,
		},
		{
			name:        "invalid pool bounds min greater than max",
			databaseURL: "postgres://localhost:5432/db",
			opts:        store.PoolOptions{MinConns: 10, MaxConns: 5},
			wantError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			_, err := store.NewPostgres(ctx, tt.databaseURL, tt.opts)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
