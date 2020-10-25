package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreateGame{}, "cosmostictactoe/CreateGame", nil)
	cdc.RegisterConcrete(MsgJoinGame{}, "cosmostictactoe/JoinGame", nil)
	cdc.RegisterConcrete(MsgCreateGameMove{}, "cosmostictactoe/CreateGameMove", nil)
	cdc.RegisterConcrete(MsgChallengeGameTimeout{}, "cosmostictactoe/ChallengeGameTimeout", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
