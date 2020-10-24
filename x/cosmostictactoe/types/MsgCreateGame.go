package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateGame{}

type MsgCreateGame struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Bet     int64          `json:"bet" yaml:"bet"`
}

func NewMsgCreateGame(creator sdk.AccAddress, bet int64) MsgCreateGame {
	return MsgCreateGame{
		ID:      uuid.New().String(),
		Creator: creator,
		Bet:     bet,
	}
}

func (msg MsgCreateGame) Route() string {
	return RouterKey
}

func (msg MsgCreateGame) Type() string {
	return "CreateGame"
}

func (msg MsgCreateGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGame) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Creator can't be empty")
	}

	if msg.Bet <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Bet has to be a positive integer")
	}

	return nil
}
