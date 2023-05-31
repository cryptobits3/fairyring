package types

import (
	peptypes "fairyring/x/pep/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// PepKeeper defines the expected interface needed to get and set active and queued public keys
type PepKeeper interface {
	SetActivePubKey(ctx sdk.Context, activePubKey peptypes.ActivePubKey)
	SetQueuedPubKey(ctx sdk.Context, queuedPubKey peptypes.QueuedPubKey)
	GetActivePubKey(ctx sdk.Context) (val peptypes.ActivePubKey, found bool)
	GetQueuedPubKey(ctx sdk.Context) (val peptypes.QueuedPubKey, found bool)
	DeleteActivePubKey(ctx sdk.Context)
	DeleteQueuedPubKey(ctx sdk.Context)
}