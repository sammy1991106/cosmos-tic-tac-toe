package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGameMove{}

type MsgCreateGameMove struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Row     uint8          `json:"row" yaml:"row"`
	Column  uint8          `json:"column" yaml:"column"`
}

func NewMsgCreateGameMove(creator sdk.AccAddress, id string, row uint8, y uint8) MsgCreateGameMove {
	return MsgCreateGameMove{
		ID:      id,
		Creator: creator,
		Row:     row,
		Column:  y,
	}
}

func (msg MsgCreateGameMove) Route() string {
	return RouterKey
}

func (msg MsgCreateGameMove) Type() string {
	return "CreateGameMove"
}

func (msg MsgCreateGameMove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGameMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGameMove) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Creator can't be empty")
	}

	if msg.Row >= 3 || msg.Column >= 3 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid move")
	}

	return nil
}
