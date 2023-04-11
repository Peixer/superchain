package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePythonCode = "create_python_code"
	TypeMsgUpdatePythonCode = "update_python_code"
	TypeMsgDeletePythonCode = "delete_python_code"
)

var _ sdk.Msg = &MsgCreatePythonCode{}

func NewMsgCreatePythonCode(creator string, uRI string) *MsgCreatePythonCode {
	return &MsgCreatePythonCode{
		Creator: creator,
		URI:     uRI,
	}
}

func (msg *MsgCreatePythonCode) Route() string {
	return RouterKey
}

func (msg *MsgCreatePythonCode) Type() string {
	return TypeMsgCreatePythonCode
}

func (msg *MsgCreatePythonCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePythonCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePythonCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePythonCode{}

func NewMsgUpdatePythonCode(creator string, id uint64, uRI string) *MsgUpdatePythonCode {
	return &MsgUpdatePythonCode{
		Id:      id,
		Creator: creator,
		URI:     uRI,
	}
}

func (msg *MsgUpdatePythonCode) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePythonCode) Type() string {
	return TypeMsgUpdatePythonCode
}

func (msg *MsgUpdatePythonCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePythonCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePythonCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePythonCode{}

func NewMsgDeletePythonCode(creator string, id uint64) *MsgDeletePythonCode {
	return &MsgDeletePythonCode{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePythonCode) Route() string {
	return RouterKey
}

func (msg *MsgDeletePythonCode) Type() string {
	return TypeMsgDeletePythonCode
}

func (msg *MsgDeletePythonCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePythonCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePythonCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
