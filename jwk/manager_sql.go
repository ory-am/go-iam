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

package jwk

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/square/go-jose.v2"

	"github.com/ory/hydra/pkg"
	"github.com/ory/x/sqlcon"
)

type SQLManager struct {
	DB     *sqlx.DB
	Cipher *AEAD
}

func NewSQLManager(db *sqlx.DB, key []byte) *SQLManager {
	return &SQLManager{DB: db, Cipher: &AEAD{Key: key}}
}

var sharedMigrations = map[string]*migrate.Migration{
	"1": {
		Id: "1",
		Up: []string{
			`CREATE TABLE IF NOT EXISTS hydra_jwk (
	sid     varchar(255) NOT NULL,
	kid 	varchar(255) NOT NULL,
	version int NOT NULL DEFAULT 0,
	keydata text NOT NULL,
	PRIMARY KEY (sid, kid)
)`,
		},
		Down: []string{
			"DROP TABLE hydra_jwk",
		},
	},
	"2": {
		Id: "2",
		Up: []string{
			`ALTER TABLE hydra_jwk ADD created_at TIMESTAMP NOT NULL DEFAULT NOW()`,
		},
		Down: []string{
			`ALTER TABLE hydra_jwk DROP COLUMN created_at`,
		},
	},
	// See https://github.com/ory/hydra/issues/921
	"3": {
		Id: "3",
		Up: []string{
			`DELETE FROM hydra_jwk WHERE sid='hydra.openid.id-token'`,
		},
		Down: []string{},
	},
}

var Migrations = map[string]*migrate.MemoryMigrationSource{
	"mysql": {
		Migrations: []*migrate.Migration{
			sharedMigrations["1"],
			sharedMigrations["2"],
			sharedMigrations["3"],
			{
				Id: "4",
				Up: []string{
					`ALTER TABLE hydra_jwk DROP PRIMARY KEY`,
					`CREATE UNIQUE INDEX hydra_jwk_idx_id_uq ON hydra_jwk (sid, kid)`,
					`ALTER TABLE hydra_jwk ADD pk INT UNSIGNED AUTO_INCREMENT PRIMARY KEY`,
				},
				Down: []string{
					`ALTER TABLE hydra_jwk DROP COLUMN pk`,
					`ALTER TABLE hydra_jwk DROP INDEX hydra_jwk_idx_id_uq`,
					`ALTER TABLE hydra_jwk ADD PRIMARY KEY (sid, kid)`,
				},
			},
		},
	},
	"postgres": {
		Migrations: []*migrate.Migration{
			sharedMigrations["1"],
			sharedMigrations["2"],
			sharedMigrations["3"],
			{
				Id: "4",
				Up: []string{
					`ALTER TABLE hydra_jwk DROP CONSTRAINT hydra_jwk_pkey`,
					`ALTER TABLE hydra_jwk ADD pk SERIAL`,
					`ALTER TABLE hydra_jwk ADD PRIMARY KEY (pk)`,
					`CREATE UNIQUE INDEX hydra_jwk_idx_id_uq ON hydra_jwk (sid, kid)`,
				},
				Down: []string{
					`ALTER TABLE hydra_jwk DROP CONSTRAINT hydra_jwk_pkey`,
					`ALTER TABLE hydra_jwk DROP COLUMN pk`,
					`DROP INDEX hydra_jwk_idx_id_uq`,
					`ALTER TABLE hydra_jwk ADD PRIMARY KEY (sid, kid)`,
				},
			},
		},
	},
}

type sqlData struct {
	PK        int    	`db:"pk"`
	Set       string    `db:"sid"`
	KID       string    `db:"kid"`
	Version   int       `db:"version"`
	CreatedAt time.Time `db:"created_at"`
	Key       string    `db:"keydata"`
}

func (m *SQLManager) CreateSchemas() (int, error) {
	database := m.DB.DriverName()
	switch database {
	case "pgx", "pq":
		database = "postgres"
	}

	migrate.SetTable("hydra_jwk_migration")
	n, err := migrate.Exec(m.DB.DB, m.DB.DriverName(), Migrations[database], migrate.Up)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not migrate sql schema, applied %d Migrations", n)
	}
	return n, nil
}

func (m *SQLManager) AddKey(ctx context.Context, set string, key *jose.JSONWebKey) error {
	out, err := json.Marshal(key)
	if err != nil {
		return errors.WithStack(err)
	}

	encrypted, err := m.Cipher.Encrypt(out)
	if err != nil {
		return errors.WithStack(err)
	}

	if _, err = m.DB.NamedExecContext(ctx, `INSERT INTO hydra_jwk (sid, kid, version, keydata) VALUES (:sid, :kid, :version, :keydata)`, &sqlData{
		Set:     set,
		KID:     key.KeyID,
		Version: 0,
		Key:     encrypted,
	}); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (m *SQLManager) AddKeySet(ctx context.Context, set string, keys *jose.JSONWebKeySet) error {
	tx, err := m.DB.BeginTxx(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := m.addKeySet(ctx, tx, m.Cipher, set, keys); err != nil {
		if re := tx.Rollback(); re != nil {
			return errors.Wrap(err, re.Error())
		}
		return sqlcon.HandleError(err)
	}

	if err := tx.Commit(); err != nil {
		if re := tx.Rollback(); re != nil {
			return errors.Wrap(err, re.Error())
		}
		return sqlcon.HandleError(err)
	}
	return nil
}

func (m *SQLManager) addKeySet(ctx context.Context, tx *sqlx.Tx, cipher *AEAD, set string, keys *jose.JSONWebKeySet) error {
	for _, key := range keys.Keys {
		out, err := json.Marshal(key)
		if err != nil {
			return errors.WithStack(err)
		}

		encrypted, err := cipher.Encrypt(out)
		if err != nil {
			return errors.WithStack(err)
		}

		if _, err = tx.NamedExecContext(ctx, `INSERT INTO hydra_jwk (sid, kid, version, keydata) VALUES (:sid, :kid, :version, :keydata)`, &sqlData{
			Set:     set,
			KID:     key.KeyID,
			Version: 0,
			Key:     encrypted,
		}); err != nil {
			return sqlcon.HandleError(err)
		}
	}

	return nil
}

func (m *SQLManager) GetKey(ctx context.Context, set, kid string) (*jose.JSONWebKeySet, error) {
	var d sqlData
	if err := m.DB.GetContext(ctx, &d, m.DB.Rebind("SELECT * FROM hydra_jwk WHERE sid=? AND kid=? ORDER BY created_at DESC"), set, kid); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	key, err := m.Cipher.Decrypt(d.Key)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var c jose.JSONWebKey
	if err := json.Unmarshal(key, &c); err != nil {
		return nil, errors.WithStack(err)
	}

	return &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{c},
	}, nil
}

func (m *SQLManager) GetKeySet(ctx context.Context, set string) (*jose.JSONWebKeySet, error) {
	var ds []sqlData
	if err := m.DB.SelectContext(ctx, &ds, m.DB.Rebind("SELECT * FROM hydra_jwk WHERE sid=? ORDER BY created_at DESC"), set); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	if len(ds) == 0 {
		return nil, errors.Wrap(pkg.ErrNotFound, "")
	}

	keys := &jose.JSONWebKeySet{Keys: []jose.JSONWebKey{}}
	for _, d := range ds {
		key, err := m.Cipher.Decrypt(d.Key)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		var c jose.JSONWebKey
		if err := json.Unmarshal(key, &c); err != nil {
			return nil, errors.WithStack(err)
		}
		keys.Keys = append(keys.Keys, c)
	}

	if len(keys.Keys) == 0 {
		return nil, errors.WithStack(pkg.ErrNotFound)
	}

	return keys, nil
}

func (m *SQLManager) DeleteKey(ctx context.Context, set, kid string) error {
	if _, err := m.DB.ExecContext(ctx, m.DB.Rebind(`DELETE FROM hydra_jwk WHERE sid=? AND kid=?`), set, kid); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (m *SQLManager) DeleteKeySet(ctx context.Context, set string) error {
	tx, err := m.DB.BeginTxx(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := m.deleteKeySet(ctx, tx, set); err != nil {
		if re := tx.Rollback(); re != nil {
			return errors.Wrap(err, re.Error())
		}
		return sqlcon.HandleError(err)
	}

	if err := tx.Commit(); err != nil {
		if re := tx.Rollback(); re != nil {
			return errors.Wrap(err, re.Error())
		}
		return sqlcon.HandleError(err)
	}
	return nil
}

func (m *SQLManager) deleteKeySet(ctx context.Context, tx *sqlx.Tx, set string) error {
	if _, err := tx.ExecContext(ctx, m.DB.Rebind(`DELETE FROM hydra_jwk WHERE sid=?`), set); err != nil {
		return sqlcon.HandleError(err)
	}
	return nil
}

func (m *SQLManager) RotateKeys(ctx context.Context, new *AEAD) error {
	sids := make([]string, 0)
	if err := m.DB.Select(&sids, "SELECT sid FROM hydra_jwk GROUP BY sid"); err != nil {
		return sqlcon.HandleError(err)
	}

	sets := make([]jose.JSONWebKeySet, 0)
	for _, sid := range sids {
		set, err := m.GetKeySet(ctx, sid)
		if err != nil {
			return errors.WithStack(err)
		}
		sets = append(sets, *set)
	}

	tx, err := m.DB.Beginx()
	if err != nil {
		return errors.WithStack(err)
	}

	for k, set := range sets {
		if err := m.deleteKeySet(ctx, tx, sids[k]); err != nil {
			if re := tx.Rollback(); re != nil {
				return errors.Wrap(err, re.Error())
			}
			return sqlcon.HandleError(err)
		}

		if err := m.addKeySet(ctx, tx, new, sids[k], &set); err != nil {
			if re := tx.Rollback(); re != nil {
				return errors.Wrap(err, re.Error())
			}
			return sqlcon.HandleError(err)
		}
	}

	if err := tx.Commit(); err != nil {
		if re := tx.Rollback(); re != nil {
			return errors.Wrap(err, re.Error())
		}
		return sqlcon.HandleError(err)
	}
	return nil
}
