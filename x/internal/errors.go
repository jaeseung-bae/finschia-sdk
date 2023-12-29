package internal

import "cosmossdk.io/errors"

const (
	linkCodespace = "link"
)

// NO additional errors allowed into this codespace
var (
	ErrInvalidPermission = errors.Register(linkCodespace, 2, "invalid permission")
	ErrInvalidDenom      = errors.Register(linkCodespace, 3, "invalid denom")
)

const contractCodespace = "contract"

var (
	ErrInvalidContractID = errors.Register(contractCodespace, 2, "invalid contractID")
	ErrContractNotExist  = errors.Register(contractCodespace, 3, "contract does not exist")
)
