package jwk

import (
	"testing"

	"github.com/square/go-jose"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"fmt"
)

func TestGenerator(t *testing.T) {
	for k, c := range []struct {
		g     KeyGenerator
		check func(*jose.JsonWebKeySet)
	}{
		{
			g: &RS256Generator{},
			check: func(ks *jose.JsonWebKeySet) {
				assert.Len(t, ks, 2)
			},
		},
		{
			g: &ECDSA521Generator{},
			check: func(ks *jose.JsonWebKeySet) {
				assert.Len(t, ks, 2)
			},
		},
		{
			g: &ECDSA256Generator{},
			check: func(ks *jose.JsonWebKeySet) {
				assert.Len(t, ks, 2)
			},
		},
		{
			g: &HS256Generator{
				Length: 32,
			},
			check: func(ks *jose.JsonWebKeySet) {
				assert.Len(t, ks, 1)
			},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			keys, err := c.g.Generate("foo")
			require.NoError(t, err)
			if err != nil {
				c.check(keys)
			}
		})
	}
}
