package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTimedoutPost = "create_timedout_post"
	TypeMsgUpdateTimedoutPost = "update_timedout_post"
	TypeMsgDeleteTimedoutPost = "delete_timedout_post"
)

var _ sdk.Msg = &MsgCreateTimedoutPost{}

func NewMsgCreateTimedoutPost(creator string, title string, chain string) *MsgCreateTimedoutPost {
	return &MsgCreateTimedoutPost{
		Creator: creator,
		Title:   title,
		Chain:   chain,
	}
}

func (msg *MsgCreateTimedoutPost) Route() string {
	return RouterKey
}

func (msg *MsgCreateTimedoutPost) Type() string {
	return TypeMsgCreateTimedoutPost
}

func (msg *MsgCreateTimedoutPost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTimedoutPost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTimedoutPost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTimedoutPost{}

func NewMsgUpdateTimedoutPost(creator string, id uint64, title string, chain string) *MsgUpdateTimedoutPost {
	return &MsgUpdateTimedoutPost{
		Id:      id,
		Creator: creator,
		Title:   title,
		Chain:   chain,
	}
}

func (msg *MsgUpdateTimedoutPost) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTimedoutPost) Type() string {
	return TypeMsgUpdateTimedoutPost
}

func (msg *MsgUpdateTimedoutPost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTimedoutPost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTimedoutPost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTimedoutPost{}

func NewMsgDeleteTimedoutPost(creator string, id uint64) *MsgDeleteTimedoutPost {
	return &MsgDeleteTimedoutPost{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTimedoutPost) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTimedoutPost) Type() string {
	return TypeMsgDeleteTimedoutPost
}

func (msg *MsgDeleteTimedoutPost) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTimedoutPost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTimedoutPost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
