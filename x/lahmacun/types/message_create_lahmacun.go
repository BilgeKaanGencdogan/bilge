package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateLahmacun = "create_lahmacun"

var _ sdk.Msg = &MsgCreateLahmacun{}

func NewMsgCreateLahmacun(creator string, id string, heat string, dough string, meat string, tomato string, onion string) *MsgCreateLahmacun {
	return &MsgCreateLahmacun{
		Creator: creator,
		Id:      id,
		Heat:    heat,
		Dough:   dough,
		Meat:    meat,
		Tomato:  tomato,
		Onion:   onion,
	}
}

func (msg *MsgCreateLahmacun) TypeCreate() string {
	return TypeMsgCreateLahmacun
}

func (msg *MsgCreateLahmacun) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %s", err)
	}

	return nil
}
