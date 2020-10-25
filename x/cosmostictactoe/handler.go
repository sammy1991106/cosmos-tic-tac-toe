package cosmostictactoe

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/keeper"
	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		if err := msg.ValidateBasic(); err != nil {
			return nil, err
		}

		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgCreateGame:
			return handleMsgCreateGame(ctx, k, msg)
		case types.MsgJoinGame:
			return handleMsgJoinGame(ctx, k, msg)
		case types.MsgCreateGameMove:
			return handleMsgCreateGameMove(ctx, k, msg)
		case types.MsgChallengeGameTimeout:
			return handleMsgChallengeGameTimeout(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
