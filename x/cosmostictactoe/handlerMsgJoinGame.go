package cosmostictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/crypto/sha3"

	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/keeper"
	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/types"
)

func getGameTurn(blockHash []byte, initiator sdk.AccAddress, challenger sdk.AccAddress) sdk.AccAddress {
	bytes := append(blockHash[:], []byte(initiator.String() + challenger.String())[:]...)

	if (sha3.Sum256(bytes)[0]>>7)&1 == 0 {
		return initiator
	}

	return challenger
}

func handleMsgJoinGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgJoinGame) (*sdk.Result, error) {
	if !k.GameExists(ctx, msg.ID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game doesn't exist")
	}

	game, err := k.GetGame(ctx, msg.ID)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Failed getting game")
	}

	if game.Challenger != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Game has already started")
	}

	challenger := msg.Creator

	if challenger.Equals(game.Initiator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Player cannot play against itself")
	}

	if _, err := k.CoinKeeper.SubtractCoins(ctx, challenger, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(game.Bet)))); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Challenger doesn't have enough balance to match bet")
	}

	game.Challenger = challenger
	game.Turn = getGameTurn(ctx.BlockHeader().DataHash, game.Initiator, game.Challenger)
	game.Bet *= 2
	game.UpdatedAt = ctx.BlockHeight()

	k.SetGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
