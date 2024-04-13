package app_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	v1 "github.com/lamdanghoang/people/gen/go/people/v1"
	"github.com/stretchr/testify/require"
)

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	registry.RegisterInterface("Staff", (*v1.Staff)(nil))
	registry.RegisterImplementations(
		(*v1.Staff)(nil),
		&v1.OfficeStaff{},
		&v1.SaleStaff{},
	)
	return registry
}

func TestMarshalAny(t *testing.T) {
	registry := types.NewInterfaceRegistry()
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&v1.SaleStaff{},
	)
	cdc := codec.NewProtoCodec(registry)

	sale := &v1.SaleStaff{Name: "Kitty"}
	bz, err := cdc.MarshalInterface(sale)
	require.NoError(t, err)

	var staff v1.Staff

	// empty registry should fail
	err = cdc.UnmarshalInterface(bz, &staff)
	require.Error(t, err)

	// wrong type registration should fail
	registry.RegisterImplementations((*v1.Staff)(nil), &v1.OfficeStaff{})
	err = cdc.UnmarshalInterface(bz, &staff)
	require.Error(t, err)

	// should pass
	registry = NewTestInterfaceRegistry()
	cdc = codec.NewProtoCodec(registry)
	err = cdc.UnmarshalInterface(bz, &staff)
	require.NoError(t, err)
	require.Equal(t, sale, staff)

	// nil should fail
	_ = NewTestInterfaceRegistry()
	err = cdc.UnmarshalInterface(bz, nil)
	require.Error(t, err)
}
