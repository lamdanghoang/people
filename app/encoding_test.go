package app_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/lamdanghoang/people/gen/go/people"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

type Parent interface {
	proto.Message
}

func NewTestInterfaceRegistry() types.InterfaceRegistry {
	registry := types.NewInterfaceRegistry()
	registry.RegisterInterface("Parent", (*Parent)(nil))
	registry.RegisterImplementations(
		(*Parent)(nil),
		&people.SaleStaff{},
	)
	return registry
}

func TestMarshalStaff(t *testing.T) {
	registry := types.NewInterfaceRegistry()
	registry.RegisterImplementations(
		(*Parent)(nil),
		&people.SaleStaff{},
	)
	cdc := codec.NewProtoCodec(registry)

	sale := &people.SaleStaff{Id: 123, Name: "Kitty", Email: "123@123", ProductQuantity: 25}

	bz, err := cdc.MarshalInterface(sale)
	require.NoError(t, err)
	fmt.Println("Marshal: ", bz)

	// var staff Parent

	// // empty registry should fail
	// err = cdc.UnmarshalInterface(bz, &staff)
	// fmt.Println("Error empty registry: ", err)
	// require.Error(t, err)

	// // wrong type registration should fail
	// registry.RegisterImplementations((*people.Staff)(nil), &people.OfficeStaff{})
	// err = cdc.UnmarshalInterface(bz, &staff)
	// fmt.Println("Error wrong type registration: ", err)
	// require.Error(t, err)

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
