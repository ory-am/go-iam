package persistence

import (
	"context"
	"io"

	"github.com/ory/hydra/client"
	"github.com/ory/hydra/consent"
	"github.com/ory/hydra/jwk"
	"github.com/ory/hydra/x"

	"github.com/gobuffalo/pop/v5"
)

type (
	Persister interface {
		consent.Manager
		client.Manager
		x.FositeStorer
		jwk.Manager

		MigrationStatus(context.Context, io.Writer) error
		MigrateDown(context.Context, int) error
		MigrateUp(context.Context) error
		PrepareMigration(context.Context) error
		Connection(context.Context) *pop.Connection
	}
	Provider interface {
		Persister() Persister
	}
)
