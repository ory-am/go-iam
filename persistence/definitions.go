// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package persistence

import (
	"context"

	"github.com/gobuffalo/pop/v6"

	"github.com/ory/hydra/v2/client"
	"github.com/ory/hydra/v2/consent"
	"github.com/ory/hydra/v2/jwk"
	"github.com/ory/hydra/v2/oauth2/trust"
	"github.com/ory/hydra/v2/x"
	"github.com/ory/x/popx"
)

type (
	Persister interface {
		consent.Manager
		client.Manager
		x.FositeStorer
		jwk.Manager
		trust.GrantManager

		MigrationStatus(ctx context.Context) (popx.MigrationStatuses, error)
		MigrateDown(context.Context, int) error
		MigrateUp(context.Context) error
		PrepareMigration(context.Context) error
		Connection(context.Context) *pop.Connection
		Ping() error
	}
	Provider interface {
		Persister() Persister
	}
)
