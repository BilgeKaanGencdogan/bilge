package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateLahmacun = "update_lahmacun"

var _ sdk.Msg = &MsgUpdateLahmacun{}

func NewMsgUpdateLahmacun(creator string, id string, heat string, dough string, meat string, tomato string, onion string) *MsgUpdateLahmacun {
	return &MsgUpdateLahmacun{
		Creator: creator,
		Id:      id,
		Heat:    heat,
		Dough:   dough,
		Meat:    meat,
		Tomato:  tomato,
		Onion:   onion,
	}
}

func (msg *MsgCreateLahmacun) TypeUpdate() string {
	return TypeMsgCreateLahmacun
}
