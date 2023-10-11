package types

import "github.com/okx/zkevm-node/aggregator/prover"

// FinalProofInputs struct
type FinalProofInputs struct {
	FinalProof       *prover.FinalProof
	NewLocalExitRoot []byte
	NewStateRoot     []byte
}
