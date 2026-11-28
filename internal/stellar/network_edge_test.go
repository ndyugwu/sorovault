package stellar_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/sorovault/internal/stellar"
)

func TestNetworkNameEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{"matching is case sensitive", strings.ToLower(stellar.NetworkTestnet), strings.ToLower(stellar.NetworkTestnet)},
		{"surrounding whitespace is not trimmed", " " + stellar.NetworkPublic, " " + stellar.NetworkPublic},
		{"whitespace alone is passed through", " ", " "},
		{"a known passphrase with a suffix is passed through", stellar.NetworkFuturenet + "!", stellar.NetworkFuturenet + "!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, stellar.NetworkName(tt.in))
		})
	}
}

func TestValidateContractIDEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
	}{
		{"lowercase", strings.ToLower(testContractID)},
		{"leading space", " " + testContractID},
		{"trailing newline", testContractID + "\n"},
		{"truncated", testContractID[:len(testContractID)-1]},
		{"extra character", testContractID + "A"},
		{"corrupted checksum", testContractID[:len(testContractID)-1] + "B"},
		{"muxed account address", "MA7QYNF7SOWQ3GLR2BGMZEHXAVIRZA4KVWLTJJFC7MGXUA74P7UJVAAAAAAAAAAAAAJLK"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := stellar.ValidateContractID(tt.in)
			require.Error(t, err)
			assert.Contains(t, err.Error(), "is not a valid contract address")
		})
	}
}
