package keys

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Bococoin/core/client/flags"
	"github.com/Bococoin/core/crypto/keys"
)

func TestCommands(t *testing.T) {
	rootCommands := Commands()
	assert.NotNil(t, rootCommands)

	// Commands are registered
	assert.Equal(t, 11, len(rootCommands.Commands()))
}

func TestMain(m *testing.M) {
	viper.Set(flags.FlagKeyringBackend, keys.BackendTest)
	os.Exit(m.Run())
}