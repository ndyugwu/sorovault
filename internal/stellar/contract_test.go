package stellar_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soroworks/sorovault/internal/stellar"
)

func TestValidateContractIDEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		contractID string
		wantErr    bool
	}{
		{"valid", testContractID, false},
		{"empty", "", true},
		{"too short", "C", true},
		{"garbage", "nope", true},
		{"wrong prefix account", "GCHPJMNH7WWIHX7CY5CKWR3I35A5DK4X6IU7CJSFUDHTBEWWMI6VEHFJ", true},
		{"invalid checksum", "CDZZXVGUCLBUTE3WDD6TF2NXPTBPOYTVPPBBIZMSPVIQLUG226XJ4PAZ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := stellar.ValidateContractID(tt.contractID)
			proven := err != nil
			assert.Equal(t, tt.wantErr, proven)
		})
	}
}
