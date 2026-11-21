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
		name    string
		url     string
		opts    store.PoolOptions
		wantErr bool
	}{
		{
			name:    "empty database URL",
			url:     "",
			opts:    store.PoolOptions{},
			wantErr: true,
		},
		{
			name:    "invalid URL scheme",
			url:     "not-a-url://",
			opts:    store.PoolOptions{},
			wantErr: true,
		},
		{
			name:    "invalid pool options min greater than max",
			url:     "postgres://localhost:5432/db",
			opts:    store.PoolOptions{MaxConns: 1, MinConns: 5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s, err := store.NewPostgres(context.Background(), tt.url, tt.opts)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, s)
			} else {
				assert.NoError(t, err)
				if s != nil {
					_ = s.Close()
				}
			}
		})
	}
}
