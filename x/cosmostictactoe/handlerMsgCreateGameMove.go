package cosmostictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/keeper"
	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/types"
)

func getFieldIndex(moveRow uint8, moveColumn uint8) uint8 {
	return moveRow*3 + moveColumn
}

func getWinner(fields [9]string, moveRow uint8, moveColumn uint8, symbol string, player sdk.AccAddress) sdk.AccAddress {
	if (fields[getFieldIndex(moveRow, 0)] == symbol && fields[getFieldIndex(moveRow, 1)] == symbol && fields[getFieldIndex(moveRow, 2)] == symbol) ||
		(fields[getFieldIndex(0, moveColumn)] == symbol && fields[getFieldIndex(1, moveColumn)] == symbol && fields[getFieldIndex(2, moveColumn)] == symbol) ||
		(fields[getFieldIndex(0, 0)] == symbol && fields[getFieldIndex(1, 1)] == symbol && fields[getFieldIndex(2, 2)] == symbol) ||
		(fields[getFieldIndex(2, 0)] == symbol && fields[getFieldIndex(1, 1)] == symbol && fields[getFieldIndex(0, 2)] == symbol) {
		return player
	}

	return nil
}

func handleMsgCreateGameMove(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGameMove) (*sdk.Result, error) {
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only registered players can play")
	}

	if game.Turn == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game has already ended")
	}

	if !game.Turn.Equals(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Please wait for the opponent to play first")
	}

	fieldIndex := getFieldIndex(msg.Row, msg.Column)

	if game.Fields[fieldIndex] != types.GameFieldEmpty {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The same move has already been played")
	}

	var symbol string

	if game.Round%2 == 0 {
		symbol = types.GameFieldX
	} else {
		symbol = types.GameFieldO
	}

	game.Fields[fieldIndex] = symbol
	game.Round++

	if isInitiator {
		game.Turn = game.Challenger
	} else {
		game.Turn = game.Initiator
	}

	if winner := getWinner(game.Fields, msg.Row, msg.Column, symbol, msg.Creator); winner != nil {
		game.Winner = winner
		game.Turn = nil

		if _, err := k.CoinKeeper.AddCoins(ctx, game.Winner, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(game.Bet)))); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Failed awarding winner")
		}
	} else if game.Round == uint8(len(game.Fields)) {
		game.Turn = nil

		if _, err := k.CoinKeeper.AddCoins(ctx, game.Initiator, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(game.Bet/2)))); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Failed refunding initiator")
		}

		if _, err := k.CoinKeeper.AddCoins(ctx, game.Challenger, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(game.Bet/2)))); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Failed refunding challenger")
		}
	}

	game.UpdatedAt = ctx.BlockHeight()

	k.SetGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
