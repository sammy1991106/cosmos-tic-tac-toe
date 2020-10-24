package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgJoinGame{}

type MsgJoinGame struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgJoinGame(creator sdk.AccAddress, id string) MsgJoinGame {
	return MsgJoinGame{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgJoinGame) Route() string {
	return RouterKey
}

func (msg MsgJoinGame) Type() string {
	return "JoinGame"
}

func (msg MsgJoinGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgJoinGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgJoinGame) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Creator can't be empty")
	}
	return nil
}
