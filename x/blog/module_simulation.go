package blog

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/deep2essence/planet/testutil/sample"
	blogsimulation "github.com/deep2essence/planet/x/blog/simulation"
	"github.com/deep2essence/planet/x/blog/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = blogsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePost int = 100

	opWeightMsgUpdatePost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePost int = 100

	opWeightMsgDeletePost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePost int = 100

	opWeightMsgCreateSentPost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSentPost int = 100

	opWeightMsgUpdateSentPost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSentPost int = 100

	opWeightMsgDeleteSentPost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSentPost int = 100

	opWeightMsgCreateTimedoutPost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTimedoutPost int = 100

	opWeightMsgUpdateTimedoutPost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTimedoutPost int = 100

	opWeightMsgDeleteTimedoutPost = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTimedoutPost int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blogGenesis := types.GenesisState{
		PostList: []types.Post{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PostCount: 2,
		SentPostList: []types.SentPost{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		SentPostCount: 2,
		TimedoutPostList: []types.TimedoutPost{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TimedoutPostCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blogGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		blogsimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePost, &weightMsgUpdatePost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePost = defaultWeightMsgUpdatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePost,
		blogsimulation.SimulateMsgUpdatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = defaultWeightMsgDeletePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePost,
		blogsimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSentPost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSentPost, &weightMsgCreateSentPost, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSentPost = defaultWeightMsgCreateSentPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSentPost,
		blogsimulation.SimulateMsgCreateSentPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSentPost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSentPost, &weightMsgUpdateSentPost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSentPost = defaultWeightMsgUpdateSentPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSentPost,
		blogsimulation.SimulateMsgUpdateSentPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSentPost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSentPost, &weightMsgDeleteSentPost, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSentPost = defaultWeightMsgDeleteSentPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSentPost,
		blogsimulation.SimulateMsgDeleteSentPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTimedoutPost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTimedoutPost, &weightMsgCreateTimedoutPost, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTimedoutPost = defaultWeightMsgCreateTimedoutPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTimedoutPost,
		blogsimulation.SimulateMsgCreateTimedoutPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTimedoutPost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTimedoutPost, &weightMsgUpdateTimedoutPost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTimedoutPost = defaultWeightMsgUpdateTimedoutPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTimedoutPost,
		blogsimulation.SimulateMsgUpdateTimedoutPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTimedoutPost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTimedoutPost, &weightMsgDeleteTimedoutPost, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTimedoutPost = defaultWeightMsgDeleteTimedoutPost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTimedoutPost,
		blogsimulation.SimulateMsgDeleteTimedoutPost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
