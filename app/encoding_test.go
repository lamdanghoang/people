package app_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/lamdanghoang/people/gen/go/people/v1"
	"github.com/stretchr/testify/require"
)

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	registry.RegisterInterface("Staff", (*people.Staff)(nil))
	registry.RegisterImplementations(
		(*people.Staff)(nil),
		&people.OfficeStaff{},
		&people.SaleStaff{},
	)
	return registry
}

func TestMarshalStaff(t *testing.T) {
	registry := types.NewInterfaceRegistry()
	registry.RegisterImplementations(
		(*tx.TxExtensionOptionI)(nil),
		&people.SaleStaff{},
	)
	cdc := codec.NewProtoCodec(registry)

	sale := &people.SaleStaff{Id: 123, Name: "Kitty", Email: "kitty@gmail.com", ProductQuantity: 25}
	fmt.Println(sale.Print_Something())
	bz, err := cdc.MarshalInterface(sale)
	require.NoError(t, err)
	fmt.Println("Marshal: ", bz)

	var staff people.Staff

	// empty registry should fail
	err = cdc.UnmarshalInterface(bz, &staff)
	fmt.Println("Error empty registry: ", err)
	require.Error(t, err)

	// wrong type registration should fail
	registry.RegisterImplementations((*people.Staff)(nil), &people.OfficeStaff{})
	err = cdc.UnmarshalInterface(bz, &staff)
	fmt.Println("Error wrong type registration: ", err)
	require.Error(t, err)

	// should pass
	// registry = NewTestInterfaceRegistry()
	// cdc = codec.NewProtoCodec(registry)
	// err = cdc.UnmarshalInterface(bz, &staff)
	// fmt.Println("Error: ", err)
	// require.NoError(t, err)
	// require.Equal(t, sale, staff)

	// // nil should fail
	// _ = NewTestInterfaceRegistry()
	// err = cdc.UnmarshalInterface(bz, nil)
	// require.Error(t, err)
}
