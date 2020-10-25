package cosmostictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/keeper"
	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/types"
)

func handleMsgChallengeGameTimeout(ctx sdk.Context, k keeper.Keeper, msg types.MsgChallengeGameTimeout) (*sdk.Result, error) {
	if !k.GameExists(ctx, msg.ID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game doesn't exist")
	}

	game, err := k.GetGame(ctx, msg.ID)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Failed getting game")
	}

	if game.Challenger == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game has not started yet")
	}

	isInitiator := game.Initiator.Equals(msg.Creator)
	isChallenger := game.Challenger.Equals(msg.Creator)

	if !isInitiator && !isChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only registered players can challenge for game timeout")
	}

	if msg.Creator.Equals(game.Turn) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Play for the current turn cannot challenge for game timeout")
	}

	blockHeight := ctx.BlockHeight()

	if blockHeight < game.UpdatedAt+game.Timeout {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The game has not timed out yet")
	}

	game.Winner = msg.Creator
	game.Turn = nil
	game.UpdatedAt = blockHeight

	if _, err := k.CoinKeeper.AddCoins(ctx, game.Winner, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(game.Bet)))); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Failed awarding winner")
	}

	k.SetGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
