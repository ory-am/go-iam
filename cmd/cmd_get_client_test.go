package cmd_test

import (
	"context"
	"encoding/json"
	"github.com/ory/hydra/client"
	"github.com/ory/hydra/cmd"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/snapshotx"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"testing"
)

func TestGetClient(t *testing.T) {
	ctx := context.Background()
	c := cmd.NewGetClientsCmd(new(cobra.Command))
	reg := setup(t, c)

	expected := &client.Client{}
	require.NoError(t, reg.ClientManager().CreateClient(ctx, expected))

	t.Run("case=gets client", func(t *testing.T) {
		actual := gjson.Parse(cmdx.ExecNoErr(t, c, expected.ID.String()))
		assert.NotEmpty(t, actual.Get("client_id").String())
		assert.Empty(t, actual.Get("client_secret").String())

		expected, err := reg.ClientManager().GetClient(ctx, actual.Get("client_id").String())
		require.NoError(t, err)

		assert.Equal(t, expected.GetID(), actual.Get("client_id").String())
		snapshotx.SnapshotT(t, json.RawMessage(actual.Raw), snapshotExcludedClientFields...)
	})

	t.Run("case=gets multiple clients", func(t *testing.T) {
		actual := gjson.Parse(cmdx.ExecNoErr(t, c, expected.ID.String(), expected.ID.String()))
		snapshotx.SnapshotT(t, json.RawMessage(actual.Raw), snapshotExcludedClientFields...)
	})
}
