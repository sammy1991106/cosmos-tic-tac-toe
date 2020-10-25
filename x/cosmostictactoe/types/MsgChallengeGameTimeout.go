package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgChallengeGameTimeout{}

type MsgChallengeGameTimeout struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgChallengeGameTimeout(creator sdk.AccAddress, id string) MsgChallengeGameTimeout {
	return MsgChallengeGameTimeout{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgChallengeGameTimeout) Route() string {
	return RouterKey
}

func (msg MsgChallengeGameTimeout) Type() string {
	return "CreateGame"
}

func (msg MsgChallengeGameTimeout) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgChallengeGameTimeout) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgChallengeGameTimeout) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Creator can't be empty")
	}

	return nil
}
