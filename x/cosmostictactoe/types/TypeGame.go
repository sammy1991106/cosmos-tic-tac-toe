package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	GameFieldEmpty = "[]"
	GameFieldX     = "X"
	GameFieldO     = "O"
)

type Game struct {
	ID         string         `json:"id" yaml:"id"`
	Initiator  sdk.AccAddress `json:"initiator" yaml:"initiator"`
	Challenger sdk.AccAddress `json:"challenger" yaml:"challenger"`
	Winner     sdk.AccAddress `json:"winner" yaml:"winner"`
	Turn       sdk.AccAddress `json:"turn" yaml:"turn"`
	Bet        int64          `json:"bet" yaml:"bet"`
	UpdatedAt  int64          `json:"updatedAt" yaml:"updatedAt"`
	Round      uint8          `json:"round" yaml:"round"`
	Fields     [9]string      `json:"fields" yaml:"fields"`
	Timeout    int64          `json:"timeout" yaml:"timeout"`
}
