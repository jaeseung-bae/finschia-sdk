package types

import (
	"strings"
	"testing"
	"time"

	"github.com/line/lbm-sdk/codec"
	"github.com/line/lbm-sdk/codec/types"
	codectypes "github.com/line/lbm-sdk/codec/types"
	sdk "github.com/line/lbm-sdk/types"
	govtypes "github.com/line/lbm-sdk/x/gov/types"
	"github.com/line/ostracon/libs/rand"
	wasmvmtypes "github.com/line/wasmvm/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContractInfoValidateBasic(t *testing.T) {
	specs := map[string]struct {
		srcMutator func(*ContractInfo)
		expError   bool
	}{
		"all good": {srcMutator: func(_ *ContractInfo) {}},
		"code id empty": {
			srcMutator: func(c *ContractInfo) { c.CodeID = 0 },
			expError:   true,
		},
		"creator empty": {
			srcMutator: func(c *ContractInfo) { c.Creator = "" },
			expError:   true,
		},
		"creator not an address": {
			srcMutator: func(c *ContractInfo) { c.Creator = "invalid address" },
			expError:   true,
		},
		"admin empty": {
			srcMutator: func(c *ContractInfo) { c.Admin = "" },
			expError:   false,
		},
		"admin not an address": {
			srcMutator: func(c *ContractInfo) { c.Admin = "invalid address" },
			expError:   true,
		},
		"label empty": {
			srcMutator: func(c *ContractInfo) { c.Label = "" },
			expError:   true,
		},
		"label exceeds limit": {
			srcMutator: func(c *ContractInfo) { c.Label = strings.Repeat("a", MaxLabelSize+1) },
			expError:   true,
		},
		"invalid extension": {
			srcMutator: func(c *ContractInfo) {
				// any protobuf type with ValidateBasic method
				any, err := codectypes.NewAnyWithValue(&govtypes.TextProposal{})
				require.NoError(t, err)
				c.Extension = any
			},
			expError: true,
		},
		"not validatable extension": {
			srcMutator: func(c *ContractInfo) {
				// any protobuf type with ValidateBasic method
				any, err := codectypes.NewAnyWithValue(&govtypes.Proposal{})
				require.NoError(t, err)
				c.Extension = any
			},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			state := ContractInfoFixture(spec.srcMutator)
			got := state.ValidateBasic()
			if spec.expError {
				require.Error(t, got)
				return
			}
			require.NoError(t, got)
		})
	}
}

func TestCodeInfoValidateBasic(t *testing.T) {
	specs := map[string]struct {
		srcMutator func(*CodeInfo)
		expError   bool
	}{
		"all good": {srcMutator: func(_ *CodeInfo) {}},
		"code hash empty": {
			srcMutator: func(c *CodeInfo) { c.CodeHash = []byte{} },
			expError:   true,
		},
		"code hash nil": {
			srcMutator: func(c *CodeInfo) { c.CodeHash = nil },
			expError:   true,
		},
		"creator empty": {
			srcMutator: func(c *CodeInfo) { c.Creator = "" },
			expError:   true,
		},
		"creator not an address": {
			srcMutator: func(c *CodeInfo) { c.Creator = "invalid address" },
			expError:   true,
		},
		"source empty": {
			srcMutator: func(c *CodeInfo) { c.Source = "" },
		},
		"source not an url": {
			srcMutator: func(c *CodeInfo) { c.Source = "invalid" },
			expError:   true,
		},
		"source not an absolute url": {
			srcMutator: func(c *CodeInfo) { c.Source = "../bar.txt" },
			expError:   true,
		},
		"source not https schema url": {
			srcMutator: func(c *CodeInfo) { c.Source = "http://example.com" },
			expError:   true,
		},
		"builder tag exceeds limit": {
			srcMutator: func(c *CodeInfo) { c.Builder = strings.Repeat("a", MaxBuildTagSize+1) },
			expError:   true,
		},
		"builder tag does not match pattern": {
			srcMutator: func(c *CodeInfo) { c.Builder = "invalid" },
			expError:   true,
		},
		"Instantiate config invalid": {
			srcMutator: func(c *CodeInfo) { c.InstantiateConfig = AccessConfig{} },
			expError:   true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			state := CodeInfoFixture(spec.srcMutator)
			got := state.ValidateBasic()
			if spec.expError {
				require.Error(t, got)
				return
			}
			require.NoError(t, got)
		})
	}
}

func TestContractInfoSetExtension(t *testing.T) {
	anyTime := time.Now().UTC()
	aNestedProtobufExt := func() ContractInfoExtension {
		// using gov proposal here as a random protobuf types as it contains an Any type inside for nested unpacking
		myExtension, err := govtypes.NewProposal(&govtypes.TextProposal{Title: "bar"}, 1, anyTime, anyTime)
		require.NoError(t, err)
		myExtension.TotalDeposit = nil
		return &myExtension
	}

	specs := map[string]struct {
		src    ContractInfoExtension
		expErr bool
		expNil bool
	}{
		"all good with any proto extension": {
			src: aNestedProtobufExt(),
		},
		"nil allowed": {
			src:    nil,
			expNil: true,
		},
		"validated and accepted": {
			src: &govtypes.TextProposal{Title: "bar", Description: "set"},
		},
		"validated and rejected": {
			src:    &govtypes.TextProposal{Title: "bar"},
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			var c ContractInfo
			gotErr := c.SetExtension(spec.src)
			if spec.expErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			if spec.expNil {
				return
			}
			require.NotNil(t, c.Extension)
			assert.NotNil(t, c.Extension.GetCachedValue())
		})
	}
}

func TestContractInfoMarshalUnmarshal(t *testing.T) {
	var myAddr = sdk.BytesToAccAddress(rand.Bytes(sdk.BytesAddrLen))
	var myOtherAddr = sdk.BytesToAccAddress(rand.Bytes(sdk.BytesAddrLen))
	var anyPos = AbsoluteTxPosition{BlockHeight: 1, TxIndex: 2}

	anyTime := time.Now().UTC()
	// using gov proposal here as a random protobuf types as it contains an Any type inside for nested unpacking
	myExtension, err := govtypes.NewProposal(&govtypes.TextProposal{Title: "bar"}, 1, anyTime, anyTime)
	require.NoError(t, err)
	myExtension.TotalDeposit = nil

	src := NewContractInfo(1, myAddr, myOtherAddr, "bar", &anyPos, ContractStatusActive)
	err = src.SetExtension(&myExtension)
	require.NoError(t, err)

	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	RegisterInterfaces(interfaceRegistry)
	// register proposal as extension type
	interfaceRegistry.RegisterImplementations(
		(*ContractInfoExtension)(nil),
		&govtypes.Proposal{},
	)
	// register gov types for nested Anys
	govtypes.RegisterInterfaces(interfaceRegistry)

	// when encode
	bz, err := marshaler.MarshalBinaryBare(&src)
	require.NoError(t, err)
	// and decode
	var dest ContractInfo
	err = marshaler.UnmarshalBinaryBare(bz, &dest)
	// then
	require.NoError(t, err)
	assert.Equal(t, src, dest)
	// and sanity check nested any
	var destExt govtypes.Proposal
	require.NoError(t, dest.ReadExtension(&destExt))
	assert.Equal(t, destExt.GetTitle(), "bar")
}

func TestContractInfoReadExtension(t *testing.T) {
	anyTime := time.Now().UTC()
	myExtension, err := govtypes.NewProposal(&govtypes.TextProposal{Title: "foo"}, 1, anyTime, anyTime)
	require.NoError(t, err)
	type TestExtensionAsStruct struct {
		ContractInfoExtension
	}

	specs := map[string]struct {
		setup  func(*ContractInfo)
		param  func() ContractInfoExtension
		expVal ContractInfoExtension
		expErr bool
	}{
		"all good": {
			setup: func(i *ContractInfo) {
				i.SetExtension(&myExtension)
			},
			param: func() ContractInfoExtension {
				return &govtypes.Proposal{}
			},
			expVal: &myExtension,
		},
		"no extension set": {
			setup: func(i *ContractInfo) {
			},
			param: func() ContractInfoExtension {
				return &govtypes.Proposal{}
			},
			expVal: &govtypes.Proposal{},
		},
		"nil argument value": {
			setup: func(i *ContractInfo) {
				i.SetExtension(&myExtension)
			},
			param: func() ContractInfoExtension {
				return nil
			},
			expErr: true,
		},
		"non matching types": {
			setup: func(i *ContractInfo) {
				i.SetExtension(&myExtension)
			},
			param: func() ContractInfoExtension {
				return &govtypes.TextProposal{}
			},
			expErr: true,
		},
		"not a pointer argument": {
			setup: func(i *ContractInfo) {
			},
			param: func() ContractInfoExtension {
				return TestExtensionAsStruct{}
			},
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			var c ContractInfo
			spec.setup(&c)
			// when

			gotValue := spec.param()
			gotErr := c.ReadExtension(gotValue)

			// then
			if spec.expErr {
				require.Error(t, gotErr)
				return
			}
			require.NoError(t, gotErr)
			assert.Equal(t, spec.expVal, gotValue)
		})
	}
}

func TestNewEnv(t *testing.T) {
	myTime := time.Unix(0, 1619700924259075000)
	t.Logf("++ unix: %d", myTime.UnixNano())
	var myContractAddr = sdk.BytesToAccAddress(rand.Bytes(sdk.BytesAddrLen))
	specs := map[string]struct {
		srcCtx sdk.Context
		exp    wasmvmtypes.Env
	}{
		"all good": {
			srcCtx: sdk.Context{}.WithBlockHeight(1).WithBlockTime(myTime).WithChainID("testing"),
			exp: wasmvmtypes.Env{
				Block: wasmvmtypes.BlockInfo{
					Height:  1,
					Time:    1619700924259075000,
					ChainID: "testing",
				},
				Contract: wasmvmtypes.ContractInfo{
					Address: myContractAddr.String(),
				},
			},
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, spec.exp, NewEnv(spec.srcCtx, myContractAddr))
		})
	}

}