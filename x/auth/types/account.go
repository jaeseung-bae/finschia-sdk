package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/line/ostracon/crypto"
	"gopkg.in/yaml.v2"

	"github.com/line/lfb-sdk/codec"
	codectypes "github.com/line/lfb-sdk/codec/types"
	"github.com/line/lfb-sdk/crypto/keys/ed25519"
	"github.com/line/lfb-sdk/crypto/keys/multisig"
	"github.com/line/lfb-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/line/lfb-sdk/crypto/types"
	sdk "github.com/line/lfb-sdk/types"
)

var (
	_ AccountI                           = (*BaseAccount)(nil)
	_ GenesisAccount                     = (*BaseAccount)(nil)
	_ codectypes.UnpackInterfacesMessage = (*BaseAccount)(nil)
	_ GenesisAccount                     = (*ModuleAccount)(nil)
	_ ModuleAccountI                     = (*ModuleAccount)(nil)

	BaseAccountSig   = []byte("bacc")
	ModuleAccountSig = []byte("macc")
)

// NewBaseAccount creates a new BaseAccount object
//nolint:interfacer
func NewBaseAccount(address sdk.AccAddress, pubKey cryptotypes.PubKey, sequence uint64) *BaseAccount {
	acc := &BaseAccount{
		Address:  address,
		Sequence: sequence,
	}

	err := acc.SetPubKey(pubKey)
	if err != nil {
		panic(err)
	}

	return acc
}

// ProtoBaseAccount - a prototype function for BaseAccount
func ProtoBaseAccount() AccountI {
	return &BaseAccount{}
}

// NewBaseAccountWithAddress - returns a new base account with a given address
func NewBaseAccountWithAddress(addr sdk.AccAddress) *BaseAccount {
	return &BaseAccount{
		Address: addr,
	}
}

// GetAddress - Implements sdk.AccountI.
func (acc BaseAccount) GetAddress() sdk.AccAddress {
	return sdk.AccAddress(acc.Address)
}

// SetAddress - Implements sdk.AccountI.
func (acc *BaseAccount) SetAddress(addr sdk.AccAddress) error {
	if len(acc.Address) != 0 {
		return errors.New("cannot override BaseAccount address")
	}

	acc.Address = addr.Bytes()
	return nil
}

// GetPubKey - Implements sdk.AccountI.
func (acc BaseAccount) GetPubKey() cryptotypes.PubKey {
	if acc.Ed25519PubKey != nil {
		return acc.Ed25519PubKey
	} else if acc.Secp256K1PubKey != nil {
		return acc.Secp256K1PubKey
	} else if acc.MultisigPubKey != nil {
		return acc.MultisigPubKey
	}
	return nil
}

// SetPubKey - Implements sdk.AccountI.
func (acc *BaseAccount) SetPubKey(pubKey cryptotypes.PubKey) error {
	if pubKey == nil {
		acc.Ed25519PubKey, acc.Secp256K1PubKey, acc.MultisigPubKey = nil, nil, nil
	} else if pk, ok := pubKey.(*ed25519.PubKey); ok {
		acc.Ed25519PubKey, acc.Secp256K1PubKey, acc.MultisigPubKey = pk, nil, nil
	} else if pk, ok := pubKey.(*secp256k1.PubKey); ok {
		acc.Ed25519PubKey, acc.Secp256K1PubKey, acc.MultisigPubKey = nil, pk, nil
	} else if pk, ok := pubKey.(*multisig.LegacyAminoPubKey); ok {
		acc.Ed25519PubKey, acc.Secp256K1PubKey, acc.MultisigPubKey = nil, nil, pk
	} else {
		return fmt.Errorf("invalid pubkey")
	}
	return nil
}

// GetSequence - Implements sdk.AccountI.
func (acc BaseAccount) GetSequence() uint64 {
	return acc.Sequence
}

// SetSequence - Implements sdk.AccountI.
func (acc *BaseAccount) SetSequence(seq uint64) error {
	acc.Sequence = seq
	return nil
}

// Validate checks for errors on the account fields
func (acc BaseAccount) Validate() error {
	if acc.Address == nil || acc.GetPubKey() == nil {
		return nil
	}

	if !bytes.Equal(acc.GetPubKey().Address().Bytes(), acc.Address) {
		return errors.New("account address and pubkey address do not match")
	}

	return nil
}

func (acc BaseAccount) String() string {
	out, _ := acc.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of an account.
func (acc BaseAccount) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &acc)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (acc BaseAccount) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if acc.MultisigPubKey != nil {
		return codectypes.UnpackInterfaces(acc.MultisigPubKey, unpacker)
	}
	return nil
}

func (acc *BaseAccount) MarshalX() ([]byte, error) {
	bz, err := acc.Marshal()
	if err != nil {
		return nil, err
	}
	t := BaseAccountSig
	b := make([]byte, len(t)+len(bz))
	copy(b, t)
	copy(b[len(t):], bz)
	return b, nil
}

// NewModuleAddress creates an AccAddress from the hash of the module's name
func NewModuleAddress(name string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(name)))
}

// NewEmptyModuleAccount creates a empty ModuleAccount from a string
func NewEmptyModuleAccount(name string, permissions ...string) *ModuleAccount {
	moduleAddress := NewModuleAddress(name)
	baseAcc := NewBaseAccountWithAddress(moduleAddress)

	if err := validatePermissions(permissions...); err != nil {
		panic(err)
	}

	return &ModuleAccount{
		BaseAccount: baseAcc,
		Name:        name,
		Permissions: permissions,
	}
}

// NewModuleAccount creates a new ModuleAccount instance
func NewModuleAccount(ba *BaseAccount, name string, permissions ...string) *ModuleAccount {
	if err := validatePermissions(permissions...); err != nil {
		panic(err)
	}

	return &ModuleAccount{
		BaseAccount: ba,
		Name:        name,
		Permissions: permissions,
	}
}

// HasPermission returns whether or not the module account has permission.
func (ma ModuleAccount) HasPermission(permission string) bool {
	for _, perm := range ma.Permissions {
		if perm == permission {
			return true
		}
	}
	return false
}

// GetName returns the the name of the holder's module
func (ma ModuleAccount) GetName() string {
	return ma.Name
}

// GetPermissions returns permissions granted to the module account
func (ma ModuleAccount) GetPermissions() []string {
	return ma.Permissions
}

// SetPubKey - Implements AccountI
func (ma ModuleAccount) SetPubKey(pubKey cryptotypes.PubKey) error {
	return fmt.Errorf("not supported for module accounts")
}

// SetSequence - Implements AccountI
func (ma ModuleAccount) SetSequence(seq uint64) error {
	return fmt.Errorf("not supported for module accounts")
}

// Validate checks for errors on the account fields
func (ma ModuleAccount) Validate() error {
	if strings.TrimSpace(ma.Name) == "" {
		return errors.New("module account name cannot be blank")
	}

	if !bytes.Equal(ma.Address, sdk.AccAddress(crypto.AddressHash([]byte(ma.Name))).Bytes()) {
		return fmt.Errorf("address %s cannot be derived from the module name '%s'", sdk.AccAddress(ma.Address).String(), ma.Name)
	}

	return ma.BaseAccount.Validate()
}

func (ma *ModuleAccount) MarshalX() ([]byte, error) {
	bz, err := ma.Marshal()
	if err != nil {
		return nil, err
	}
	t := ModuleAccountSig
	b := make([]byte, len(t)+len(bz))
	copy(b, t)
	copy(b[len(t):], bz)
	return b, nil
}

type moduleAccountPretty struct {
	Address     sdk.AccAddress `json:"address" yaml:"address"`
	PubKey      string         `json:"public_key" yaml:"public_key"`
	Sequence    uint64         `json:"sequence" yaml:"sequence"`
	Name        string         `json:"name" yaml:"name"`
	Permissions []string       `json:"permissions" yaml:"permissions"`
}

func (ma ModuleAccount) String() string {
	out, _ := ma.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of a ModuleAccount.
func (ma ModuleAccount) MarshalYAML() (interface{}, error) {
	bs, err := yaml.Marshal(moduleAccountPretty{
		Address:     ma.Address,
		PubKey:      "",
		Sequence:    ma.Sequence,
		Name:        ma.Name,
		Permissions: ma.Permissions,
	})

	if err != nil {
		return nil, err
	}

	return string(bs), nil
}

// MarshalJSON returns the JSON representation of a ModuleAccount.
func (ma ModuleAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(moduleAccountPretty{
		Address:     ma.Address,
		PubKey:      "",
		Sequence:    ma.Sequence,
		Name:        ma.Name,
		Permissions: ma.Permissions,
	})
}

// UnmarshalJSON unmarshals raw JSON bytes into a ModuleAccount.
func (ma *ModuleAccount) UnmarshalJSON(bz []byte) error {
	var alias moduleAccountPretty
	if err := json.Unmarshal(bz, &alias); err != nil {
		return err
	}

	ma.BaseAccount = NewBaseAccount(alias.Address, nil, alias.Sequence)
	ma.Name = alias.Name
	ma.Permissions = alias.Permissions

	return nil
}

// AccountI is an interface used to store coins at a given address within state.
// It presumes a notion of sequence numbers for replay protection,
// a notion of sig block height for replay protection for previously pruned accounts,
// and a pubkey for authentication purposes.
//
// Many complex conditions can be used in the concrete struct which implements AccountI.
type AccountI interface {
	proto.Message

	GetAddress() sdk.AccAddress
	SetAddress(sdk.AccAddress) error // errors if already set.

	GetPubKey() cryptotypes.PubKey // can return nil.
	SetPubKey(cryptotypes.PubKey) error

	GetSequence() uint64
	SetSequence(uint64) error

	// Ensure that account implements stringer
	String() string

	MarshalX() ([]byte, error)
}

func MarshalAccountX(cdc codec.BinaryMarshaler, acc AccountI) ([]byte, error) {
	if bacc, ok := acc.(*BaseAccount); ok && bacc.MultisigPubKey == nil {
		return acc.MarshalX()
	} else if macc, ok := acc.(*ModuleAccount); ok && macc.MultisigPubKey == nil {
		return acc.MarshalX()
	} else {
		return cdc.MarshalInterface(acc)
	}
}

func UnmarshalAccountX(cdc codec.BinaryMarshaler, bz []byte) (AccountI, error) {
	sigLen := len(BaseAccountSig)
	if len(bz) < sigLen {
		return nil, fmt.Errorf("invalid data")
	}
	if bytes.Equal(bz[:sigLen], BaseAccountSig) {
		acc := &BaseAccount{}
		if err := acc.Unmarshal(bz[sigLen:]); err != nil {
			return nil, err
		}
		return acc, nil
	} else if bytes.Equal(bz[:sigLen], ModuleAccountSig) {
		acc := &ModuleAccount{}
		if err := acc.Unmarshal(bz[sigLen:]); err != nil {
			return nil, err
		}
		return acc, nil
	} else {
		var acc AccountI
		return acc, cdc.UnmarshalInterface(bz, &acc)
	}
}

// ModuleAccountI defines an account interface for modules that hold tokens in
// an escrow.
type ModuleAccountI interface {
	AccountI

	GetName() string
	GetPermissions() []string
	HasPermission(string) bool
}

// GenesisAccounts defines a slice of GenesisAccount objects
type GenesisAccounts []GenesisAccount

// Contains returns true if the given address exists in a slice of GenesisAccount
// objects.
func (ga GenesisAccounts) Contains(addr sdk.Address) bool {
	for _, acc := range ga {
		if acc.GetAddress().Equals(addr) {
			return true
		}
	}

	return false
}

// GenesisAccount defines a genesis account that embeds an AccountI with validation capabilities.
type GenesisAccount interface {
	AccountI

	Validate() error
}
