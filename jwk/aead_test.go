package jwk

import (
	"crypto/rand"
	"io"
	"testing"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RandomBytes returns n random bytes by reading from crypto/rand.Reader
func randomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return []byte{}, errors.WithStack(err)
	}
	return bytes, nil
}

func TestAEAD(t *testing.T) {
	key, err := randomBytes(32)
	require.NoError(t, err)

	a := &AEAD{
		Key: key,
	}

	for i := 0; i < 100; i++ {
		plain := []byte(uuid.New())
		ct, err := a.Encrypt(plain)
		require.NoError(t, err)

		res, err := a.Decrypt(ct)
		require.NoError(t, err)
		assert.Equal(t, plain, res)
	}
}
