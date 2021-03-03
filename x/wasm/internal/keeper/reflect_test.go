// nolint: staticcheck, unused, deadcode
package keeper

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"

	wasmTypes "github.com/CosmWasm/wasmvm/types"

	"github.com/line/lbm-sdk/x/coin"
	"github.com/line/lbm-sdk/x/wasm/internal/types"
)

// MaskInitMsg is {}

// MaskHandleMsg is used to encode handle messages
type MaskHandleMsg struct {
	Reflect *reflectPayload `json:"reflect_msg,omitempty"`
	Change  *ownerPayload   `json:"change_owner,omitempty"`
}

type ownerPayload struct {
	Owner sdk.Address `json:"owner"`
}

type reflectPayload struct {
	Msgs []wasmTypes.CosmosMsg `json:"msgs"`
}

// MaskQueryMsg is used to encode query messages
type MaskQueryMsg struct {
	Owner       *struct{}   `json:"owner,omitempty"`
	Capitalized *Text       `json:"capitalized,omitempty"`
	Chain       *ChainQuery `json:"chain,omitempty"`
}

type ChainQuery struct {
	Request *wasmTypes.QueryRequest `json:"request,omitempty"`
}

type Text struct {
	Text string `json:"text"`
}

type OwnerResponse struct {
	Owner string `json:"owner,omitempty"`
}

type ChainResponse struct {
	Data []byte `json:"data,omitempty"`
}

func buildMaskQuery(t *testing.T, query *MaskQueryMsg) []byte {
	bz, err := json.Marshal(query)
	require.NoError(t, err)
	return bz
}

func mustParse(t *testing.T, data []byte, res interface{}) {
	err := json.Unmarshal(data, res)
	require.NoError(t, err)
}

const MaskFeatures = "staking,mask"

func TestMaskReflectContractSend(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "wasm")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)
	ctx, keepers := CreateTestInput(t, false, tempDir, MaskFeatures, maskEncoders(MakeTestCodec()), nil)
	accKeeper, keeper := keepers.AccountKeeper, keepers.WasmKeeper

	deposit := sdk.NewCoins(sdk.NewInt64Coin("denom", 100000))
	creator := createFakeFundedAccount(ctx, accKeeper, deposit)
	_, _, bob := keyPubAddr()

	// upload mask code
	maskCode, err := ioutil.ReadFile("./testdata/reflect.wasm")
	require.NoError(t, err)
	maskID, err := keeper.Create(ctx, creator, maskCode, "", "", nil)
	require.NoError(t, err)
	require.Equal(t, uint64(1), maskID)

	// upload hackatom escrow code
	escrowCode, err := ioutil.ReadFile("./testdata/hackatom.wasm")
	require.NoError(t, err)
	escrowID, err := keeper.Create(ctx, creator, escrowCode, "", "", nil)
	require.NoError(t, err)
	require.Equal(t, uint64(2), escrowID)

	// creator instantiates a contract and gives it tokens
	maskStart := sdk.NewCoins(sdk.NewInt64Coin("denom", 40000))
	maskAddr, err := keeper.Instantiate(ctx, maskID, creator, nil, []byte("{}"), "mask contract 2", maskStart)
	require.NoError(t, err)
	require.NotEmpty(t, maskAddr)

	// now we set contract as verifier of an escrow
	initMsg := InitMsg{
		Verifier:    maskAddr,
		Beneficiary: bob,
	}
	initMsgBz, err := json.Marshal(initMsg)
	require.NoError(t, err)
	escrowStart := sdk.NewCoins(sdk.NewInt64Coin("denom", 25000))
	escrowAddr, err := keeper.Instantiate(ctx, escrowID, creator, nil, initMsgBz, "escrow contract 2", escrowStart)
	require.NoError(t, err)
	require.NotEmpty(t, escrowAddr)

	// let's make sure all balances make sense
	checkAccount(t, ctx, accKeeper, creator, sdk.NewCoins(sdk.NewInt64Coin("denom", 35000))) // 100k - 40k - 25k
	checkAccount(t, ctx, accKeeper, maskAddr, maskStart)
	checkAccount(t, ctx, accKeeper, escrowAddr, escrowStart)
	checkAccount(t, ctx, accKeeper, bob, nil)

	// now for the trick.... we reflect a message through the mask to call the escrow
	// we also send an additional 14k tokens there.
	// this should reduce the mask balance by 14k (to 26k)
	// this 14k is added to the escrow, then the entire balance is sent to bob (total: 39k)
	approveMsg := []byte(`{"release":{}}`)
	msgs := []wasmTypes.CosmosMsg{{
		Wasm: &wasmTypes.WasmMsg{
			Execute: &wasmTypes.ExecuteMsg{
				ContractAddr: escrowAddr.String(),
				Msg:          approveMsg,
				Send: []wasmTypes.Coin{{
					Denom:  "denom",
					Amount: "14000",
				}},
			},
		},
	}}
	reflectSend := MaskHandleMsg{
		Reflect: &reflectPayload{
			Msgs: msgs,
		},
	}
	reflectSendBz, err := json.Marshal(reflectSend)
	require.NoError(t, err)
	_, err = keeper.Execute(ctx, maskAddr, creator, reflectSendBz, nil)
	require.NoError(t, err)

	// did this work???
	checkAccount(t, ctx, accKeeper, creator, sdk.NewCoins(sdk.NewInt64Coin("denom", 35000)))  // same as before
	checkAccount(t, ctx, accKeeper, maskAddr, sdk.NewCoins(sdk.NewInt64Coin("denom", 26000))) // 40k - 14k (from send)
	checkAccount(t, ctx, accKeeper, escrowAddr, sdk.Coins{})                                  // emptied reserved
	checkAccount(t, ctx, accKeeper, bob, sdk.NewCoins(sdk.NewInt64Coin("denom", 39000)))      // all escrow of 25k + 14k
}

func TestMaskReflectCustomMsg(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "wasm")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)
	ctx, keepers := CreateTestInput(t, false, tempDir, MaskFeatures, maskEncoders(MakeTestCodec()), maskPlugins())
	accKeeper, keeper := keepers.AccountKeeper, keepers.WasmKeeper

	deposit := sdk.NewCoins(sdk.NewInt64Coin("denom", 100000))
	creator := createFakeFundedAccount(ctx, accKeeper, deposit)
	bob := createFakeFundedAccount(ctx, accKeeper, deposit)
	_, _, fred := keyPubAddr()

	// upload code
	maskCode, err := ioutil.ReadFile("./testdata/reflect.wasm")
	require.NoError(t, err)
	codeID, err := keeper.Create(ctx, creator, maskCode, "", "", nil)
	require.NoError(t, err)
	require.Equal(t, uint64(1), codeID)

	// creator instantiates a contract and gives it tokens
	contractStart := sdk.NewCoins(sdk.NewInt64Coin("denom", 40000))
	contractAddr, err := keeper.Instantiate(ctx, codeID, creator, nil, []byte("{}"), "mask contract 1", contractStart)
	require.NoError(t, err)
	require.NotEmpty(t, contractAddr)

	// set owner to bob
	transfer := MaskHandleMsg{
		Change: &ownerPayload{
			Owner: bob,
		},
	}
	transferBz, err := json.Marshal(transfer)
	require.NoError(t, err)
	_, err = keeper.Execute(ctx, contractAddr, creator, transferBz, nil)
	require.NoError(t, err)

	// check some account values
	checkAccount(t, ctx, accKeeper, contractAddr, contractStart)
	checkAccount(t, ctx, accKeeper, bob, deposit)
	checkAccount(t, ctx, accKeeper, fred, nil)

	// bob can send contract's tokens to fred (using SendMsg)
	msgs := []wasmTypes.CosmosMsg{{
		Bank: &wasmTypes.BankMsg{
			Send: &wasmTypes.SendMsg{
				FromAddress: contractAddr.String(),
				ToAddress:   fred.String(),
				Amount: []wasmTypes.Coin{{
					Denom:  "denom",
					Amount: "15000",
				}},
			},
		},
	}}
	reflectSend := MaskHandleMsg{
		Reflect: &reflectPayload{
			Msgs: msgs,
		},
	}
	reflectSendBz, err := json.Marshal(reflectSend)
	require.NoError(t, err)
	_, err = keeper.Execute(ctx, contractAddr, bob, reflectSendBz, nil)
	require.NoError(t, err)

	// fred got coins
	checkAccount(t, ctx, accKeeper, fred, sdk.NewCoins(sdk.NewInt64Coin("denom", 15000)))
	// contract lost them
	checkAccount(t, ctx, accKeeper, contractAddr, sdk.NewCoins(sdk.NewInt64Coin("denom", 25000)))
	checkAccount(t, ctx, accKeeper, bob, deposit)

	// construct an opaque message
	var sdkSendMsg sdk.Msg = &coin.MsgSend{
		From:   contractAddr,
		To:     fred,
		Amount: sdk.NewCoins(sdk.NewInt64Coin("denom", 23000)),
	}
	opaque, err := toMaskRawMsg(keeper.cdc, sdkSendMsg)
	require.NoError(t, err)
	reflectOpaque := MaskHandleMsg{
		Reflect: &reflectPayload{
			Msgs: []wasmTypes.CosmosMsg{opaque},
		},
	}
	reflectOpaqueBz, err := json.Marshal(reflectOpaque)
	require.NoError(t, err)

	_, err = keeper.Execute(ctx, contractAddr, bob, reflectOpaqueBz, nil)
	require.NoError(t, err)

	// fred got more coins
	checkAccount(t, ctx, accKeeper, fred, sdk.NewCoins(sdk.NewInt64Coin("denom", 38000)))
	// contract lost them
	checkAccount(t, ctx, accKeeper, contractAddr, sdk.NewCoins(sdk.NewInt64Coin("denom", 2000)))
	checkAccount(t, ctx, accKeeper, bob, deposit)
}

func TestMaskReflectCustomQuery(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "wasm")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)
	ctx, keepers := CreateTestInput(t, false, tempDir, MaskFeatures, maskEncoders(MakeTestCodec()), maskPlugins())
	accKeeper, keeper := keepers.AccountKeeper, keepers.WasmKeeper

	deposit := sdk.NewCoins(sdk.NewInt64Coin("denom", 100000))
	creator := createFakeFundedAccount(ctx, accKeeper, deposit)

	// upload code
	maskCode, err := ioutil.ReadFile("./testdata/reflect.wasm")
	require.NoError(t, err)
	codeID, err := keeper.Create(ctx, creator, maskCode, "", "", nil)
	require.NoError(t, err)
	require.Equal(t, uint64(1), codeID)

	// creator instantiates a contract and gives it tokens
	contractStart := sdk.NewCoins(sdk.NewInt64Coin("denom", 40000))
	contractAddr, err := keeper.Instantiate(ctx, codeID, creator, nil, []byte("{}"), "mask contract 1", contractStart)
	require.NoError(t, err)
	require.NotEmpty(t, contractAddr)

	// let's perform a normal query of state
	ownerQuery := MaskQueryMsg{
		Owner: &struct{}{},
	}
	ownerQueryBz, err := json.Marshal(ownerQuery)
	require.NoError(t, err)
	ownerRes, err := keeper.QuerySmart(ctx, contractAddr, ownerQueryBz)
	require.NoError(t, err)
	var res OwnerResponse
	err = json.Unmarshal(ownerRes, &res)
	require.NoError(t, err)
	assert.Equal(t, res.Owner, creator.String())

	// and now making use of the custom querier callbacks
	customQuery := MaskQueryMsg{
		Capitalized: &Text{
			Text: "all Caps noW",
		},
	}
	customQueryBz, err := json.Marshal(customQuery)
	require.NoError(t, err)
	custom, err := keeper.QuerySmart(ctx, contractAddr, customQueryBz)
	require.NoError(t, err)
	var resp capitalizedResponse
	err = json.Unmarshal(custom, &resp)
	require.NoError(t, err)
	assert.Equal(t, resp.Text, "ALL CAPS NOW")
}

type maskState struct {
	Owner []byte `json:"owner"`
}

func TestMaskReflectWasmQueries(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "wasm")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)
	ctx, keepers := CreateTestInput(t, false, tempDir, MaskFeatures, maskEncoders(MakeTestCodec()), nil)
	accKeeper, keeper := keepers.AccountKeeper, keepers.WasmKeeper

	deposit := sdk.NewCoins(sdk.NewInt64Coin("denom", 100000))
	creator := createFakeFundedAccount(ctx, accKeeper, deposit)

	// upload mask code
	maskCode, err := ioutil.ReadFile("./testdata/reflect.wasm")
	require.NoError(t, err)
	maskID, err := keeper.Create(ctx, creator, maskCode, "", "", nil)
	require.NoError(t, err)
	require.Equal(t, uint64(1), maskID)

	// creator instantiates a contract and gives it tokens
	maskStart := sdk.NewCoins(sdk.NewInt64Coin("denom", 40000))
	maskAddr, err := keeper.Instantiate(ctx, maskID, creator, nil, []byte("{}"), "mask contract 2", maskStart)
	require.NoError(t, err)
	require.NotEmpty(t, maskAddr)

	// for control, let's make some queries directly on the mask
	ownerQuery := buildMaskQuery(t, &MaskQueryMsg{Owner: &struct{}{}})
	res, err := keeper.QuerySmart(ctx, maskAddr, ownerQuery)
	require.NoError(t, err)
	var ownerRes OwnerResponse
	mustParse(t, res, &ownerRes)
	require.Equal(t, ownerRes.Owner, creator.String())

	// and a raw query: cosmwasm_storage::Singleton uses 2 byte big-endian length-prefixed to store data
	configKey := append([]byte{0, 6}, []byte("config")...)
	raw := keeper.QueryRaw(ctx, maskAddr, configKey)
	var stateRes maskState
	mustParse(t, raw, &stateRes)
	require.Equal(t, stateRes.Owner, []byte(creator))

	// now, let's reflect a smart query into the x/wasm handlers and see if we get the same result
	reflectOwnerQuery := MaskQueryMsg{Chain: &ChainQuery{Request: &wasmTypes.QueryRequest{Wasm: &wasmTypes.WasmQuery{
		Smart: &wasmTypes.SmartQuery{
			ContractAddr: maskAddr.String(),
			Msg:          ownerQuery,
		},
	}}}}
	reflectOwnerBin := buildMaskQuery(t, &reflectOwnerQuery)
	res, err = keeper.QuerySmart(ctx, maskAddr, reflectOwnerBin)
	require.NoError(t, err)
	// first we pull out the data from chain response, before parsing the original response
	var reflectRes ChainResponse
	mustParse(t, res, &reflectRes)
	var reflectOwnerRes OwnerResponse
	mustParse(t, reflectRes.Data, &reflectOwnerRes)
	require.Equal(t, reflectOwnerRes.Owner, creator.String())

	// and with queryRaw
	reflectStateQuery := MaskQueryMsg{Chain: &ChainQuery{Request: &wasmTypes.QueryRequest{Wasm: &wasmTypes.WasmQuery{
		Raw: &wasmTypes.RawQuery{
			ContractAddr: maskAddr.String(),
			Key:          configKey,
		},
	}}}}
	reflectStateBin := buildMaskQuery(t, &reflectStateQuery)
	res, err = keeper.QuerySmart(ctx, maskAddr, reflectStateBin)
	require.NoError(t, err)
	// first we pull out the data from chain response, before parsing the original response
	var reflectRawRes ChainResponse
	mustParse(t, res, &reflectRawRes)
	// now, with the raw data, we can parse it into state
	var reflectStateRes maskState
	mustParse(t, reflectRawRes.Data, &reflectStateRes)
	require.Equal(t, reflectStateRes.Owner, []byte(creator))
}

func TestWasmRawQueryWithNil(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "wasm")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)
	ctx, keepers := CreateTestInput(t, false, tempDir, MaskFeatures, maskEncoders(MakeTestCodec()), nil)
	accKeeper, keeper := keepers.AccountKeeper, keepers.WasmKeeper

	deposit := sdk.NewCoins(sdk.NewInt64Coin("denom", 100000))
	creator := createFakeFundedAccount(ctx, accKeeper, deposit)

	// upload mask code
	maskCode, err := ioutil.ReadFile("./testdata/reflect.wasm")
	require.NoError(t, err)
	maskID, err := keeper.Create(ctx, creator, maskCode, "", "", nil)
	require.NoError(t, err)
	require.Equal(t, uint64(1), maskID)

	// creator instantiates a contract and gives it tokens
	maskStart := sdk.NewCoins(sdk.NewInt64Coin("denom", 40000))
	maskAddr, err := keeper.Instantiate(ctx, maskID, creator, nil, []byte("{}"), "mask contract 2", maskStart)
	require.NoError(t, err)
	require.NotEmpty(t, maskAddr)

	// control: query directly
	missingKey := []byte{0, 1, 2, 3, 4}
	raw := keeper.QueryRaw(ctx, maskAddr, missingKey)
	require.Nil(t, raw)

	// and with queryRaw
	reflectQuery := MaskQueryMsg{Chain: &ChainQuery{Request: &wasmTypes.QueryRequest{Wasm: &wasmTypes.WasmQuery{
		Raw: &wasmTypes.RawQuery{
			ContractAddr: maskAddr.String(),
			Key:          missingKey,
		},
	}}}}
	reflectStateBin := buildMaskQuery(t, &reflectQuery)
	res, err := keeper.QuerySmart(ctx, maskAddr, reflectStateBin)
	require.NoError(t, err)

	// first we pull out the data from chain response, before parsing the original response
	var reflectRawRes ChainResponse
	mustParse(t, res, &reflectRawRes)
	// and make sure there is no data
	require.Empty(t, reflectRawRes.Data)
	// we get an empty byte slice not nil (if anyone care in go-land)
	require.Equal(t, []byte{}, reflectRawRes.Data)
}

func checkAccount(t *testing.T, ctx sdk.Context, accKeeper auth.AccountKeeper, addr sdk.AccAddress, expected sdk.Coins) {
	acct := accKeeper.GetAccount(ctx, addr)
	if expected == nil {
		assert.Nil(t, acct)
	} else {
		assert.NotNil(t, acct)
		if expected.Empty() {
			// there is confusion between nil and empty slice... let's just treat them the same
			assert.True(t, acct.GetCoins().Empty())
		} else {
			assert.Equal(t, acct.GetCoins(), expected)
		}
	}
}

/**** Code to support custom messages *****/

type maskCustomMsg struct {
	Debug string `json:"debug,omitempty"`
	Raw   []byte `json:"raw,omitempty"`
}

// toMaskRawMsg encodes an sdk msg using amino json encoding.
// Then wraps it as an opaque message
func toMaskRawMsg(cdc *codec.Codec, msg sdk.Msg) (wasmTypes.CosmosMsg, error) {
	rawBz, err := cdc.MarshalJSON(msg)
	if err != nil {
		return wasmTypes.CosmosMsg{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	customMsg, err := json.Marshal(maskCustomMsg{
		Raw: rawBz,
	})
	res := wasmTypes.CosmosMsg{
		Custom: customMsg,
	}
	return res, nil
}

// maskEncoders needs to be registered in test setup to handle custom message callbacks
func maskEncoders(cdc *codec.Codec) *MessageEncoders {
	return &MessageEncoders{
		Custom: fromMaskRawMsg(cdc),
	}
}

// fromMaskRawMsg decodes msg.Data to an sdk.Msg using amino json encoding.
// this needs to be registered on the Encoders
func fromMaskRawMsg(cdc *codec.Codec) CustomEncoder {
	return func(_sender sdk.AccAddress, msg json.RawMessage, router types.Router) ([]sdk.Msg, error) {
		var custom maskCustomMsg
		err := json.Unmarshal(msg, &custom)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
		}
		if custom.Raw != nil {
			var sdkMsg sdk.Msg
			err := cdc.UnmarshalJSON(custom.Raw, &sdkMsg)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
			}
			return []sdk.Msg{sdkMsg}, nil
		}
		if custom.Debug != "" {
			return nil, sdkerrors.Wrapf(types.ErrInvalidMsg, "Custom Debug: %s", custom.Debug)
		}
		return nil, sdkerrors.Wrap(types.ErrInvalidMsg, "Unknown Custom message variant")
	}
}

type maskCustomQuery struct {
	Ping        *struct{} `json:"ping,omitempty"`
	Capitalized *Text     `json:"capitalized,omitempty"`
}

// this is from the go code back to the contract (capitalized or ping)
type customQueryResponse struct {
	Msg string `json:"msg"`
}

// these are the return values from contract -> go depending on type of query
type ownerResponse struct {
	Owner string `json:"owner"`
}

type capitalizedResponse struct {
	Text string `json:"text"`
}

type chainResponse struct {
	Data []byte `json:"data"`
}

// maskPlugins needs to be registered in test setup to handle custom query callbacks
func maskPlugins() *QueryPlugins {
	return &QueryPlugins{
		Custom: performCustomQuery,
	}
}

func performCustomQuery(_ sdk.Context, request json.RawMessage) ([]byte, error) {
	var custom maskCustomQuery
	err := json.Unmarshal(request, &custom)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	if custom.Capitalized != nil {
		msg := strings.ToUpper(custom.Capitalized.Text)
		return json.Marshal(customQueryResponse{Msg: msg})
	}
	if custom.Ping != nil {
		return json.Marshal(customQueryResponse{Msg: "pong"})
	}
	return nil, sdkerrors.Wrap(types.ErrInvalidMsg, "Unknown Custom query variant")
}
