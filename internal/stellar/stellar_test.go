package stellar_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soroworks/sorovault/internal/stellar"
)

func TestNetworkNameEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		passphrase string
		want       string
	}{
		{"public", stellar.NetworkPublic, "public"},
		{"testnet", stellar.NetworkTestnet, "testnet"},
		{"futurenet", stellar.NetworkFuturenet, "futurenet"},
		{"empty", "", "unknown"},
		{"custom passphrase", "custom-network-passphrase", "custom-network-passphrase"},
		{"whitespace", "   ", "   "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, stellar.NetworkName(tt.passphrase))
		})
	}
}
