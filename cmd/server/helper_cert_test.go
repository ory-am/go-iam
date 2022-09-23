package server_test

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
	"gopkg.in/square/go-jose.v2"

	"github.com/ory/x/configx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tlsx"

	"github.com/ory/hydra/cmd/server"
	"github.com/ory/hydra/driver"
	"github.com/ory/hydra/driver/config"
	"github.com/ory/hydra/internal/testhelpers"
	"github.com/ory/hydra/jwk"
)

func TestGetOrCreateTLSCertificate(t *testing.T) {
	certPath, keyPath, cert, priv := testhelpers.GenerateTLSCertificateFilesForTests(t)
	logger := logrusx.New("", "")
	logger.Logger.ExitFunc = func(code int) { t.Fatalf("Logger called os.Exit(%v)", code) }
	hook := test.NewLocal(logger.Logger)
	cfg := config.MustNew(
		context.Background(),
		logger,
		configx.WithValues(map[string]interface{}{
			"dsn":                 config.DSNMemory,
			"serve.tls.enabled":   true,
			"serve.tls.cert.path": certPath,
			"serve.tls.key.path":  keyPath,
		}),
	)
	d, err := driver.NewRegistryWithoutInit(cfg, logger)
	require.NoError(t, err)
	getCert := server.GetOrCreateTLSCertificate(context.Background(), d, config.AdminInterface, nil)
	require.NotNil(t, getCert)
	tlsCert, err := getCert(nil)
	require.NoError(t, err)
	require.NotNil(t, tlsCert)
	if tlsCert.Leaf == nil {
		tlsCert.Leaf, err = x509.ParseCertificate(tlsCert.Certificate[0])
		require.NoError(t, err)
	}
	require.True(t, tlsCert.Leaf.Equal(cert))
	require.True(t, priv.Equal(tlsCert.PrivateKey))

	// generate new cert+key
	newCertPath, newKeyPath, newCert, newPriv := testhelpers.GenerateTLSCertificateFilesForTests(t)
	require.False(t, cert.Equal(newCert))
	require.False(t, priv.Equal(newPriv))
	require.NotEqual(t, certPath, newCertPath)
	require.NotEqual(t, keyPath, newKeyPath)

	// move them into place
	require.NoError(t, os.Rename(newKeyPath, keyPath))
	require.NoError(t, os.Rename(newCertPath, certPath))

	// give it some time and check we're reloaded
	time.Sleep(150 * time.Millisecond)
	require.Nil(t, hook.LastEntry())

	// request another certificate: it should be the new one
	tlsCert, err = getCert(nil)
	require.NoError(t, err)
	if tlsCert.Leaf == nil {
		tlsCert.Leaf, err = x509.ParseCertificate(tlsCert.Certificate[0])
		require.NoError(t, err)
	}
	require.True(t, tlsCert.Leaf.Equal(newCert))
	require.True(t, newPriv.Equal(tlsCert.PrivateKey))

	require.NoError(t, os.WriteFile(certPath, []byte{'j', 'u', 'n', 'k'}, 0))

	timeout := time.After(500 * time.Millisecond)
	for {
		if hook.LastEntry() != nil {
			break
		}
		select {
		case <-timeout:
			require.FailNow(t, "expected error log entry")
		default:
		}
	}
	require.Contains(t, hook.LastEntry().Message, "Failed to reload TLS certificates. Using the previously loaded certificates.")
}

func TestCreateSelfSignedCertificate(t *testing.T) {
	keys, err := jwk.GenerateJWK(context.Background(), jose.RS256, uuid.New().String(), "sig")
	require.NoError(t, err)

	private := keys.Keys[0]
	cert, err := tlsx.CreateSelfSignedCertificate(private.Key)
	require.NoError(t, err)
	server.AttachCertificate(&private, cert)

	var actual jose.JSONWebKeySet
	var b bytes.Buffer
	require.NoError(t, json.NewEncoder(&b).Encode(keys))
	require.NoError(t, json.NewDecoder(&b).Decode(&actual))
}
