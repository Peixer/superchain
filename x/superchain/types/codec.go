package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePythonCode{}, "superchain/CreatePythonCode", nil)
	cdc.RegisterConcrete(&MsgUpdatePythonCode{}, "superchain/UpdatePythonCode", nil)
	cdc.RegisterConcrete(&MsgDeletePythonCode{}, "superchain/DeletePythonCode", nil)
	cdc.RegisterConcrete(&MsgRunCode{}, "superchain/RunCode", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePythonCode{},
		&MsgUpdatePythonCode{},
		&MsgDeletePythonCode{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRunCode{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
