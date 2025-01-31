package types_test

import (
	"testing"

	"github.com/quicksilver-zone/quicksilver/x/interchainstaking/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestValidateParams(t *testing.T) {
	require.NoError(t, types.DefaultParams().Validate(), "default")
	require.NoError(t, types.NewParams(1, 1, sdk.NewDec(1), true).Validate(), "valid")
	require.Error(t, types.NewParams(0, 1, sdk.NewDec(1), true).Validate(), "0 deposit interval")
	require.Error(t, types.NewParams(1, 0, sdk.NewDec(1), true).Validate(), "0 valset interval")
	require.Error(t, types.NewParams(1, 1, sdk.NewDec(-1), true).Validate(), "negative commission rate")
}
