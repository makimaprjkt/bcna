package keeper

import (
	"context"

	"github.com/BitCannaGlobal/bcna/x/bcna/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBitcannaid(goCtx context.Context, msg *types.MsgCreateBitcannaid) (*types.MsgCreateBitcannaidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if a BitCannaID with the same Bcnaid already exists
	if k.HasBitcannaidWithBcnaid(ctx, msg.Bcnaid) {
		return nil, types.ErrDuplicateBitcannaid.Wrapf("BitCannaID with Bcnaid %s already exists", msg.Bcnaid)
	}
	var bitcannaid = types.Bitcannaid{
		Creator: msg.Creator,
		Bcnaid:  msg.Bcnaid,
		Address: msg.Address,
	}

	id := k.AppendBitcannaid(
		ctx,
		bitcannaid,
	)

	return &types.MsgCreateBitcannaidResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBitcannaid(goCtx context.Context, msg *types.MsgUpdateBitcannaid) (*types.MsgUpdateBitcannaidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if a BitCannaID with the same Bcnaid already exists
	if k.HasBitcannaidWithBcnaid(ctx, msg.Bcnaid) {
		return nil, types.ErrDuplicateBitcannaid.Wrapf("BitCannaID with Bcnaid %s already exists", msg.Bcnaid)
	}
	var bitcannaid = types.Bitcannaid{
		Creator: msg.Creator,
		Id:      msg.Id,
		Bcnaid:  msg.Bcnaid,
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetBitcannaid(ctx, msg.Id)
	if !found {
		return nil, types.ErrKeyNotFound.Wrapf("key doesn't exist: %d", msg.Id)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, types.ErrUnauthorized.Wrapf("Unauthorized: %s,", msg.Creator)
	}

	k.SetBitcannaid(ctx, bitcannaid)

	return &types.MsgUpdateBitcannaidResponse{}, nil
}

func (k msgServer) DeleteBitcannaid(goCtx context.Context, msg *types.MsgDeleteBitcannaid) (*types.MsgDeleteBitcannaidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetBitcannaid(ctx, msg.Id)
	if !found {
		return nil, types.ErrKeyNotFound.Wrapf("key doesn't exist: %d", msg.Id)
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, types.ErrUnauthorized.Wrapf("Unauthorized: %s,", msg.Creator)
	}

	k.RemoveBitcannaid(ctx, msg.Id)

	return &types.MsgDeleteBitcannaidResponse{}, nil
}
