package cosmostictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/keeper"
	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/types"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	if _, err := k.CoinKeeper.SubtractCoins(ctx, msg.Creator, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(msg.Bet)))); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Creator doesn't have enough balance to place bet")
	}

	fields := [9]string{}
	for i := range fields {
		fields[i] = types.GameFieldEmpty
	}

	game := types.Game{
		ID:         msg.ID,
		Initiator:  msg.Creator,
		Challenger: nil,
		Winner:     nil,
		Turn:       nil,
		Bet:        msg.Bet,
		UpdatedAt:  ctx.BlockHeight(),
		Round:      0,
		Fields:     fields,
	}
	k.CreateGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
