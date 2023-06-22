// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql

import (
	"context"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"

	"github.com/ory/x/errorsx"
	"github.com/ory/x/otelx"

	"github.com/ory/fosite"
	"github.com/ory/hydra/v2/client"
	"github.com/ory/x/sqlcon"
)

func (p *Persister) GetConcreteClient(ctx context.Context, id string) (c *client.Client, err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetConcreteClient")
	defer otelx.End(span, &err)

	var cl client.Client
	if err := p.QueryWithNetwork(ctx).Where("id = ?", id).First(&cl); err != nil {
		return nil, sqlcon.HandleError(err)
	}
	return &cl, nil
}

func (p *Persister) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	return p.GetConcreteClient(ctx, id)
}

func (p *Persister) UpdateClient(ctx context.Context, cl *client.Client) (err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.UpdateClient")
	defer otelx.End(span, &err)

	return p.transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		o, err := p.GetConcreteClient(ctx, cl.GetID())
		if err != nil {
			return err
		}

		if cl.Secret == "" {
			cl.Secret = string(o.GetHashedSecret())
		} else {
			h, err := p.r.ClientHasher().Hash(ctx, []byte(cl.Secret))
			if err != nil {
				return errorsx.WithStack(err)
			}
			cl.Secret = string(h)
		}
		// set the internal primary key
		cl.ID = o.ID

		// Set the legacy client ID
		cl.LegacyClientID = o.LegacyClientID

		if err = cl.BeforeSave(c); err != nil {
			return sqlcon.HandleError(err)
		}

		count, err := p.UpdateWithNetwork(ctx, cl)
		if err != nil {
			return sqlcon.HandleError(err)
		} else if count == 0 {
			return sqlcon.HandleError(sqlcon.ErrNoRows)
		}
		return sqlcon.HandleError(err)
	})
}

func (p *Persister) Authenticate(ctx context.Context, id string, secret []byte) (_ *client.Client, err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.Authenticate")
	defer otelx.End(span, &err)

	c, err := p.GetConcreteClient(ctx, id)
	if err != nil {
		return nil, errorsx.WithStack(err)
	}

	if err := p.r.ClientHasher().Compare(ctx, c.GetHashedSecret(), secret); err != nil {
		return nil, errorsx.WithStack(err)
	}

	return c, nil
}

func (p *Persister) CreateClient(ctx context.Context, c *client.Client) (err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.CreateClient")
	defer otelx.End(span, &err)

	h, err := p.r.ClientHasher().Hash(ctx, []byte(c.Secret))
	if err != nil {
		return err
	}

	c.Secret = string(h)
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV4())
	}
	if c.LegacyClientID == "" {
		c.LegacyClientID = c.ID.String()
	}
	return sqlcon.HandleError(p.CreateWithNetwork(ctx, c))
}

func (p *Persister) DeleteClient(ctx context.Context, id string) (err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.DeleteClient")
	defer otelx.End(span, &err)

	_, err = p.GetConcreteClient(ctx, id)
	if err != nil {
		return err
	}

	return sqlcon.HandleError(p.QueryWithNetwork(ctx).Where("id = ?", id).Delete(&client.Client{}))
}

func (p *Persister) GetClients(ctx context.Context, filters client.Filter) (_ []client.Client, err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetClients")
	defer otelx.End(span, &err)

	cs := make([]client.Client, 0)

	query := p.QueryWithNetwork(ctx).
		Paginate(filters.Offset/filters.Limit+1, filters.Limit).
		Order("pk")

	if filters.Name != "" {
		query.Where("client_name = ?", filters.Name)
	}
	if filters.Owner != "" {
		query.Where("owner = ?", filters.Owner)
	}

	if err := query.All(&cs); err != nil {
		return nil, sqlcon.HandleError(err)
	}
	return cs, nil
}

func (p *Persister) CountClients(ctx context.Context) (n int, err error) {
	ctx, span := p.r.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.CountClients")
	defer otelx.End(span, &err)

	n, err = p.QueryWithNetwork(ctx).Count(&client.Client{})
	return n, sqlcon.HandleError(err)
}
