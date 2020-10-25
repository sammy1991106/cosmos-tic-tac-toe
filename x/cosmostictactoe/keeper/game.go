package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/sammy1991106/cosmos-tic-tac-toe/x/cosmostictactoe/types"
)

// CreateGame creates a game
func (k Keeper) CreateGame(ctx sdk.Context, game types.Game) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + game.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	store.Set(key, value)
}

// GetGame returns the game information
func (k Keeper) GetGame(ctx sdk.Context, key string) (types.Game, error) {
	store := ctx.KVStore(k.storeKey)
	var game types.Game
	byteKey := []byte(types.GamePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &game)
	if err != nil {
		return game, err
	}
	return game, nil
}

// SetGame sets a game
func (k Keeper) SetGame(ctx sdk.Context, game types.Game) {
	gameKey := game.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	key := []byte(types.GamePrefix + gameKey)
	store.Set(key, bz)
}

//
// Functions used by querier
//

func listGame(ctx sdk.Context, k Keeper) ([]byte, error) {
	var gameList []types.Game
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var game types.Game
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &game)
		gameList = append(gameList, game)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, gameList)
	return res, nil
}

func getGame(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, game)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetGameOwner(ctx sdk.Context, key string) sdk.AccAddress {
	game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil
	}
	return game.Initiator
}

// Check if the key exists in the store
func (k Keeper) GameExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.GamePrefix + key))
}
