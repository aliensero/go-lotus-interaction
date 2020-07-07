package types

import (
	"github.com/filecoin-project/specs-actors/actors/builtin"
)

const FilecoinPrecision = uint64(1_000_000_000_000_000_000)
const TotalFilecoin = uint64(2_000_000_000)
const BlockGasLimit = 100_000_000_000

type Version uint32

var BlocksPerEpoch = uint64(builtin.ExpectedLeadersPerEpoch)
