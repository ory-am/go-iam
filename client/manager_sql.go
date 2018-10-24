/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/square/go-jose.v2"

	"github.com/ory/fosite"
	"github.com/ory/go-convenience/stringsx"
	"github.com/ory/x/sqlcon"
)

var sharedMigrations = []*migrate.Migration{
	{
		Id: "1",
		Up: []string{`CREATE TABLE IF NOT EXISTS hydra_client (
	id      	varchar(255) NOT NULL PRIMARY KEY,
	client_name  	text NOT NULL,
	client_secret  	text NOT NULL,
	redirect_uris  	text NOT NULL,
	grant_types  	text NOT NULL,
	response_types  text NOT NULL,
	scope  			text NOT NULL,
	owner  			text NOT NULL,
	policy_uri  	text NOT NULL,
	tos_uri  		text NOT NULL,
	client_uri  	text NOT NULL,
	logo_uri  		text NOT NULL,
	contacts  		text NOT NULL,
	public  		boolean NOT NULL
)`},
		Down: []string{
			"DROP TABLE hydra_client",
		},
	},
	{
		Id: "2",
		Up: []string{
			`ALTER TABLE hydra_client ADD client_secret_expires_at INTEGER NOT NULL DEFAULT 0`,
		},
		Down: []string{
			`ALTER TABLE hydra_client DROP COLUMN client_secret_expires_at`,
		},
	},
	{
		Id: "3",
		Up: []string{
			`ALTER TABLE hydra_client ADD sector_identifier_uri TEXT`,
			`ALTER TABLE hydra_client ADD jwks TEXT`,
			`ALTER TABLE hydra_client ADD jwks_uri TEXT`,
			`ALTER TABLE hydra_client ADD request_uris TEXT`,
			`ALTER TABLE hydra_client ADD token_endpoint_auth_method VARCHAR(25) NOT NULL DEFAULT ''`,
			`ALTER TABLE hydra_client ADD request_object_signing_alg  VARCHAR(10) NOT NULL DEFAULT ''`,
			`ALTER TABLE hydra_client ADD userinfo_signed_response_alg VARCHAR(10) NOT NULL DEFAULT ''`,
		},
		Down: []string{
			`ALTER TABLE hydra_client DROP COLUMN sector_identifier_uri`,
			`ALTER TABLE hydra_client DROP COLUMN jwks`,
			`ALTER TABLE hydra_client DROP COLUMN jwks_uri`,
			`ALTER TABLE hydra_client DROP COLUMN token_endpoint_auth_method`,
			`ALTER TABLE hydra_client DROP COLUMN request_uris`,
			`ALTER TABLE hydra_client DROP COLUMN request_object_signing_alg`,
			`ALTER TABLE hydra_client DROP COLUMN userinfo_signed_response_alg`,
		},
	},
	{
		Id: "5",
		Up: []string{
			`UPDATE hydra_client SET token_endpoint_auth_method='none' WHERE public=TRUE`,
			`ALTER TABLE hydra_client DROP COLUMN public`,
		},
		Down: []string{
			`ALTER TABLE hydra_client ADD public BOOLEAN NOT NULL DEFAULT FALSE`,
			`UPDATE hydra_client SET public=TRUE WHERE token_endpoint_auth_method='none'`,
		},
	},
	{
		Id: "6",
		Up: []string{
			`ALTER TABLE hydra_client ADD subject_type VARCHAR(15) NOT NULL DEFAULT ''`,
		},
		Down: []string{
			`ALTER TABLE hydra_client DROP COLUMN subject_type`,
		},
	},
	{
		Id: "7",
		Up: []string{
			`ALTER TABLE hydra_client ADD allowed_cors_origins TEXT`,
		},
		Down: []string{
			`ALTER TABLE hydra_client DROP COLUMN allowed_cors_origins`,
		},
	},
}

var Migrations = map[string]*migrate.MemoryMigrationSource{
	"mysql": {Migrations: []*migrate.Migration{
		sharedMigrations[0],
		sharedMigrations[1],
		sharedMigrations[2],
		{
			Id: "4",
			Up: []string{
				`UPDATE hydra_client SET sector_identifier_uri='', jwks='', jwks_uri='', request_uris=''`,
				`ALTER TABLE hydra_client MODIFY sector_identifier_uri TEXT NOT NULL`,
				`ALTER TABLE hydra_client MODIFY jwks TEXT NOT NULL`,
				`ALTER TABLE hydra_client MODIFY jwks_uri TEXT NOT NULL`,
				`ALTER TABLE hydra_client MODIFY request_uris TEXT NOT NULL`,
			},
			Down: []string{
				`ALTER TABLE hydra_client MODIFY sector_identifier_uri TEXT`,
				`ALTER TABLE hydra_client MODIFY jwks TEXT`,
				`ALTER TABLE hydra_client MODIFY jwks_uri TEXT`,
				`ALTER TABLE hydra_client MODIFY request_uris TEXT`,
			},
		},
		sharedMigrations[3],
		sharedMigrations[4],
		sharedMigrations[5],
		{
			Id: "8",
			Up: []string{
				`UPDATE hydra_client SET allowed_cors_origins=''`,
				`ALTER TABLE hydra_client MODIFY allowed_cors_origins TEXT NOT NULL`,
			},
			Down: []string{
				`ALTER TABLE hydra_client MODIFY allowed_cors_origins TEXT`,
			},
		},
		{
			Id: "9",
			Up: []string{
				`ALTER TABLE hydra_client DROP PRIMARY KEY`,
				`CREATE UNIQUE INDEX hydra_client_idx_id_uq ON hydra_client (id)`,
				`ALTER TABLE hydra_client ADD pk INT UNSIGNED AUTO_INCREMENT PRIMARY KEY`,
			},
			Down: []string{
				`ALTER TABLE hydra_client DROP COLUMN pk`,
				`ALTER TABLE hydra_client DROP INDEX hydra_client_idx_id_uq`,
				`ALTER TABLE hydra_client ADD PRIMARY KEY (id)`,
			},
		},
	}},
	"postgres": {Migrations: []*migrate.Migration{
		sharedMigrations[0],
		sharedMigrations[1],
		sharedMigrations[2],
		{
			Id: "4",
			Up: []string{
				`UPDATE hydra_client SET sector_identifier_uri='', jwks='', jwks_uri='', request_uris=''`,
				`ALTER TABLE hydra_client ALTER COLUMN sector_identifier_uri SET NOT NULL`,
				`ALTER TABLE hydra_client ALTER COLUMN jwks SET NOT NULL`,
				`ALTER TABLE hydra_client ALTER COLUMN jwks_uri SET NOT NULL`,
				`ALTER TABLE hydra_client ALTER COLUMN request_uris SET NOT NULL`,
			},
			Down: []string{
				`ALTER TABLE hydra_client ALTER COLUMN sector_identifier_uri DROP NOT NULL`,
				`ALTER TABLE hydra_client ALTER COLUMN jwks DROP NOT NULL`,
				`ALTER TABLE hydra_client ALTER COLUMN jwks_uri DROP NOT NULL`,
				`ALTER TABLE hydra_client ALTER COLUMN request_uris DROP NOT NULL`,
			},
		},
		sharedMigrations[3],
		sharedMigrations[4],
		sharedMigrations[5],
		{
			Id: "8",
			Up: []string{
				`UPDATE hydra_client SET allowed_cors_origins=''`,
				`ALTER TABLE hydra_client ALTER COLUMN allowed_cors_origins SET NOT NULL`,
			},
			Down: []string{
				`ALTER TABLE hydra_client ALTER COLUMN allowed_cors_origins DROP NOT NULL`,
			},
		},
		{
			Id: "9",
			Up: []string{
				`ALTER TABLE hydra_client DROP CONSTRAINT hydra_client_pkey`,
				`ALTER TABLE hydra_client ADD pk SERIAL`,
				`ALTER TABLE hydra_client ADD PRIMARY KEY (pk)`,
				`CREATE UNIQUE INDEX hydra_client_idx_id_uq ON hydra_client (id)`,
			},
			Down: []string{
				`ALTER TABLE hydra_client DROP CONSTRAINT hydra_client_pkey`,
				`ALTER TABLE hydra_client DROP COLUMN pk`,
				`DROP INDEX hydra_client_idx_id_uq`,
				`ALTER TABLE hydra_client ADD PRIMARY KEY (id)`,
			},
		},
	}},
}

type SQLManager struct {
	Hasher fosite.Hasher
	DB     *sqlx.DB
}

type sqlData struct {
	PK                            int    `db:"pk"`
	ID                            string `db:"id"`
	Name                          string `db:"client_name"`
	Secret                        string `db:"client_secret"`
	RedirectURIs                  string `db:"redirect_uris"`
	GrantTypes                    string `db:"grant_types"`
	ResponseTypes                 string `db:"response_types"`
	Scope                         string `db:"scope"`
	Owner                         string `db:"owner"`
	PolicyURI                     string `db:"policy_uri"`
	TermsOfServiceURI             string `db:"tos_uri"`
	ClientURI                     string `db:"client_uri"`
	LogoURI                       string `db:"logo_uri"`
	Contacts                      string `db:"contacts"`
	SecretExpiresAt               int    `db:"client_secret_expires_at"`
	SectorIdentifierURI           string `db:"sector_identifier_uri"`
	JSONWebKeysURI                string `db:"jwks_uri"`
	JSONWebKeys                   string `db:"jwks"`
	TokenEndpointAuthMethod       string `db:"token_endpoint_auth_method"`
	RequestURIs                   string `db:"request_uris"`
	SubjectType                   string `db:"subject_type"`
	RequestObjectSigningAlgorithm string `db:"request_object_signing_alg"`
	UserinfoSignedResponseAlg     string `db:"userinfo_signed_response_alg"`
	AllowedCORSOrigins            string `db:"allowed_cors_origins"`
}

var sqlParams = []string{
	"id",
	"client_name",
	"client_secret",
	"redirect_uris",
	"grant_types",
	"response_types",
	"scope",
	"owner",
	"policy_uri",
	"tos_uri",
	"client_uri",
	"subject_type",
	"logo_uri",
	"contacts",
	"client_secret_expires_at",
	"sector_identifier_uri",
	"jwks",
	"jwks_uri",
	"token_endpoint_auth_method",
	"request_uris",
	"request_object_signing_alg",
	"userinfo_signed_response_alg",
	"allowed_cors_origins",
}

func sqlDataFromClient(d *Client) (*sqlData, error) {
	jwks := ""

	if d.JSONWebKeys != nil {
		out, err := json.Marshal(d.JSONWebKeys)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		jwks = string(out)
	}

	return &sqlData{
		ID:                            d.GetID(),
		Name:                          d.Name,
		Secret:                        d.Secret,
		RedirectURIs:                  strings.Join(d.RedirectURIs, "|"),
		GrantTypes:                    strings.Join(d.GrantTypes, "|"),
		ResponseTypes:                 strings.Join(d.ResponseTypes, "|"),
		Scope:                         d.Scope,
		Owner:                         d.Owner,
		PolicyURI:                     d.PolicyURI,
		TermsOfServiceURI:             d.TermsOfServiceURI,
		ClientURI:                     d.ClientURI,
		LogoURI:                       d.LogoURI,
		Contacts:                      strings.Join(d.Contacts, "|"),
		SecretExpiresAt:               d.SecretExpiresAt,
		SectorIdentifierURI:           d.SectorIdentifierURI,
		JSONWebKeysURI:                d.JSONWebKeysURI,
		JSONWebKeys:                   jwks,
		TokenEndpointAuthMethod:       d.TokenEndpointAuthMethod,
		RequestObjectSigningAlgorithm: d.RequestObjectSigningAlgorithm,
		RequestURIs:                   strings.Join(d.RequestURIs, "|"),
		UserinfoSignedResponseAlg:     d.UserinfoSignedResponseAlg,
		SubjectType:                   d.SubjectType,
		AllowedCORSOrigins:            strings.Join(d.AllowedCORSOrigins, "|"),
	}, nil
}

func (d *sqlData) ToClient() (*Client, error) {
	c := &Client{
		ClientID:                      d.ID,
		Name:                          d.Name,
		Secret:                        d.Secret,
		RedirectURIs:                  stringsx.Splitx(d.RedirectURIs, "|"),
		GrantTypes:                    stringsx.Splitx(d.GrantTypes, "|"),
		ResponseTypes:                 stringsx.Splitx(d.ResponseTypes, "|"),
		Scope:                         d.Scope,
		Owner:                         d.Owner,
		PolicyURI:                     d.PolicyURI,
		TermsOfServiceURI:             d.TermsOfServiceURI,
		ClientURI:                     d.ClientURI,
		LogoURI:                       d.LogoURI,
		Contacts:                      stringsx.Splitx(d.Contacts, "|"),
		SecretExpiresAt:               d.SecretExpiresAt,
		SectorIdentifierURI:           d.SectorIdentifierURI,
		JSONWebKeysURI:                d.JSONWebKeysURI,
		TokenEndpointAuthMethod:       d.TokenEndpointAuthMethod,
		RequestObjectSigningAlgorithm: d.RequestObjectSigningAlgorithm,
		RequestURIs:                   stringsx.Splitx(d.RequestURIs, "|"),
		UserinfoSignedResponseAlg:     d.UserinfoSignedResponseAlg,
		SubjectType:                   d.SubjectType,
		AllowedCORSOrigins:            stringsx.Splitx(d.AllowedCORSOrigins, "|"),
	}

	if d.JSONWebKeys != "" {
		c.JSONWebKeys = new(jose.JSONWebKeySet)
		if err := json.Unmarshal([]byte(d.JSONWebKeys), &c.JSONWebKeys); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return c, nil
}

func (m *SQLManager) CreateSchemas() (int, error) {
	database := m.DB.DriverName()
	switch database {
	case "pgx", "pq":
		database = "postgres"
	}

	migrate.SetTable("hydra_client_migration")
	n, err := migrate.Exec(m.DB.DB, m.DB.DriverName(), Migrations[database], migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not migrate sql schema, applied %d Migrations", n)
	}
	return n, nil
}

func (m *SQLManager) GetConcreteClient(ctx context.Context, id string) (*Client, error) {
	var d sqlData
	if err := m.DB.GetContext(ctx, &d, m.DB.Rebind("SELECT * FROM hydra_client WHERE id=?"), id); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	return d.ToClient()
}

func (m *SQLManager) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	return m.GetConcreteClient(ctx, id)
}

func (m *SQLManager) UpdateClient(ctx context.Context, c *Client) error {
	o, err := m.GetClient(ctx, c.GetID())
	if err != nil {
		return errors.WithStack(err)
	}

	if c.Secret == "" {
		c.Secret = string(o.GetHashedSecret())
	} else {
		h, err := m.Hasher.Hash(ctx, []byte(c.Secret))
		if err != nil {
			return errors.WithStack(err)
		}
		c.Secret = string(h)
	}

	s, err := sqlDataFromClient(c)
	if err != nil {
		return errors.WithStack(err)
	}

	var query []string
	for _, param := range sqlParams {
		query = append(query, fmt.Sprintf("%s=:%s", param, param))
	}

	if _, err := m.DB.NamedExecContext(ctx, fmt.Sprintf(`UPDATE hydra_client SET %s WHERE id=:id`, strings.Join(query, ", ")), s); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (m *SQLManager) Authenticate(ctx context.Context, id string, secret []byte) (*Client, error) {
	c, err := m.GetConcreteClient(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := m.Hasher.Compare(ctx, c.GetHashedSecret(), secret); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}

func (m *SQLManager) CreateClient(ctx context.Context, c *Client) error {
	h, err := m.Hasher.Hash(ctx, []byte(c.Secret))
	if err != nil {
		return errors.WithStack(err)
	}
	c.Secret = string(h)

	data, err := sqlDataFromClient(c)
	if err != nil {
		return errors.WithStack(err)
	}

	if _, err := m.DB.NamedExecContext(ctx, fmt.Sprintf(
		"INSERT INTO hydra_client (%s) VALUES (%s)",
		strings.Join(sqlParams, ", "),
		":"+strings.Join(sqlParams, ", :"),
	), data); err != nil {
		return sqlcon.HandleError(err)
	}

	return nil
}

func (m *SQLManager) DeleteClient(ctx context.Context, id string) error {
	if _, err := m.DB.ExecContext(ctx, m.DB.Rebind(`DELETE FROM hydra_client WHERE id=?`), id); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (m *SQLManager) GetClients(ctx context.Context, limit, offset int) (clients map[string]Client, err error) {
	d := make([]sqlData, 0)
	clients = make(map[string]Client)

	if err := m.DB.SelectContext(ctx, &d, m.DB.Rebind("SELECT * FROM hydra_client ORDER BY id LIMIT ? OFFSET ?"), limit, offset); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	for _, k := range d {
		c, err := k.ToClient()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		clients[k.ID] = *c
	}
	return clients, nil
}
