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
	"crypto/rand"
	"io"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/square/go-jose.v2"
)

func RandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return []byte{}, errors.WithStack(err)
	}
	return bytes, nil
}

func TestHelperManagerKey(m Manager, keys *jose.JSONWebKeySet, suffix string) func(t *testing.T) {
	pub := keys.Key("public:" + suffix)
	priv := keys.Key("private:" + suffix)

	return func(t *testing.T) {
		t.Parallel()
		_, err := m.GetKey("faz", "baz")
		assert.NotNil(t, err)

		err = m.AddKey("faz", First(priv))
		assert.Nil(t, err)

		got, err := m.GetKey("faz", "private:"+suffix)
		assert.Nil(t, err)
		assert.Equal(t, priv, got.Keys)

		err = m.AddKey("faz", First(pub))
		assert.Nil(t, err)

		got, err = m.GetKey("faz", "private:"+suffix)
		assert.Nil(t, err)
		assert.Equal(t, priv, got.Keys)

		got, err = m.GetKey("faz", "public:"+suffix)
		assert.Nil(t, err)
		assert.Equal(t, pub, got.Keys)

		First(pub).KeyID = "new-key-id:" + suffix
		err = m.AddKey("faz", First(pub))
		assert.Nil(t, err)

		_, err = m.GetKey("faz", "new-key-id:"+suffix)
		assert.Nil(t, err)

		keys, err = m.GetKeySet("faz")
		assert.Nil(t, err)
		assert.EqualValues(t, "new-key-id:"+suffix, First(keys.Keys).KeyID)

		err = m.DeleteKey("faz", "public:"+suffix)
		assert.Nil(t, err)

		_, err = m.GetKey("faz", "public:"+suffix)
		assert.NotNil(t, err)
	}
}

func TestHelperManagerKeySet(m Manager, keys *jose.JSONWebKeySet, suffix string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		_, err := m.GetKeySet("foo")
		require.Error(t, err)

		err = m.AddKeySet("bar", keys)
		assert.Nil(t, err)

		got, err := m.GetKeySet("bar")
		assert.Nil(t, err)
		assert.Equal(t, keys.Key("public:"+suffix), got.Key("public:"+suffix))
		assert.Equal(t, keys.Key("private:"+suffix), got.Key("private:"+suffix))

		err = m.DeleteKeySet("bar")
		assert.Nil(t, err)

		_, err = m.GetKeySet("bar")
		assert.NotNil(t, err)
	}
}
