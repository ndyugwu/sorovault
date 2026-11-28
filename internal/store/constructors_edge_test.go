package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/sorovault/internal/store"
)

// TestNewPostgresEdgeCases covers inputs NewPostgres must refuse before it
// ever reaches a database, so it needs no TEST_DATABASE_URL.
func TestNewPostgresEdgeCases(t *testing.T) {
	t.Parallel()

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	tests := []struct {
		name string
		ctx  context.Context
		url  string
		opts store.PoolOptions
		want string
	}{
		{
			name: "unparseable connection string",
			ctx:  context.Background(),
			url:  "://not a connection string",
			want: "store: parsing postgres connection string",
		},
		{
			name: "minimum above maximum",
			ctx:  context.Background(),
			url:  "postgres://127.0.0.1:1/db",
			opts: store.PoolOptions{MaxConns: 1, MinConns: 4},
			want: "exceeds its maximum",
		},
		{
			name: "cancelled context never connects",
			ctx:  cancelled,
			url:  "postgres://127.0.0.1:1/db",
			want: "store:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s, err := store.NewPostgres(tt.ctx, tt.url, tt.opts)
			require.Error(t, err)
			assert.Nil(t, s)
			assert.Contains(t, err.Error(), tt.want)
		})
	}
}

// TestNewMemoryEdgeCases pins what a freshly built in-memory store holds and
// that two stores never share state.
func TestNewMemoryEdgeCases(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("a new store is empty", func(t *testing.T) {
		t.Parallel()

		m := store.NewMemory()
		_, err := m.GetContract(ctx, "testnet", "CDZZ")
		assert.ErrorIs(t, err, store.ErrNotFound)
		_, err = m.GetState(ctx, "cursor")
		assert.ErrorIs(t, err, store.ErrNotFound)
		networks, err := m.ListNetworks(ctx)
		require.NoError(t, err)
		assert.Empty(t, networks)
	})

	t.Run("two stores do not share state", func(t *testing.T) {
		t.Parallel()

		a, b := store.NewMemory(), store.NewMemory()
		require.NoError(t, a.SetState(ctx, "cursor", "42"))
		got, err := a.GetState(ctx, "cursor")
		require.NoError(t, err)
		assert.Equal(t, "42", got)
		_, err = b.GetState(ctx, "cursor")
		assert.ErrorIs(t, err, store.ErrNotFound)
	})

	t.Run("close is safe to repeat", func(t *testing.T) {
		t.Parallel()

		m := store.NewMemory()
		assert.NoError(t, m.Close())
		assert.NoError(t, m.Close())
	})
}
