package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

// RegisterCodec registers the necessary types and interfaces for the module.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateLahmacun{}, "lahmacun/CreateLahmacun", nil)
	cdc.RegisterConcrete(&MsgUpdateLahmacun{}, "lahmacun/UpdateLahmacun", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "lahmacun/UpdateParams", nil)
	cdc.RegisterConcrete(&QueryLahmacunRequest{}, "lahmacun/GetLahmacun", nil)
}

// RegisterInterfaces registers the module's interface types
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// Register the implementations of the sdk.Msg interface
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgCreateLahmacun{},
		&MsgUpdateLahmacun{},
		&QueryLahmacunRequest{},
	)
	// Register the message service description
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
