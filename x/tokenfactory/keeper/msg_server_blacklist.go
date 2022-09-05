package keeper

import (
	"context"

	"noble/x/tokenfactory/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Blacklist(goCtx context.Context, msg *types.MsgBlacklist) (*types.MsgBlacklistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	blacklister, found := k.GetBlacklister(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUserNotFound, "blacklister isn't set")
	}

	if blacklister.Address != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "you are not the blacklister")
	}

	blacklisted := types.Blacklisted{
		Address: msg.Address,
	}

	k.SetBlacklisted(ctx, blacklisted)

	return &types.MsgBlacklistResponse{}, nil
}
