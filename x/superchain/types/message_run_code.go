package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRunCode = "run_code"

var _ sdk.Msg = &MsgRunCode{}

func NewMsgRunCode(creator string, id int32) *MsgRunCode {
	return &MsgRunCode{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRunCode) Route() string {
	return RouterKey
}

func (msg *MsgRunCode) Type() string {
	return TypeMsgRunCode
}

func (msg *MsgRunCode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRunCode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRunCode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
