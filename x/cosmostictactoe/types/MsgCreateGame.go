package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateGame{}

const (
	MinGameTimeout = 10
)

type MsgCreateGame struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Bet     int64          `json:"bet" yaml:"bet"`
	Timeout int64          `json:"timeout" yaml:"timeout"`
}

func NewMsgCreateGame(creator sdk.AccAddress, bet int64, timeout int64) MsgCreateGame {
	return MsgCreateGame{
		ID:      uuid.New().String(),
		Creator: creator,
		Bet:     bet,
		Timeout: timeout,
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

	if msg.Timeout < MinGameTimeout {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Game timeout has to be greater than or equal to %d", MinGameTimeout))
	}

	return nil
}
