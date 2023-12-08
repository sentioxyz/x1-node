package jsonrpc

import (
	"context"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/jackc/pgx/v4"
)

// DBTxManager allows to do scopped DB txs
type DBTxManager struct{}

// DBTxScopedFn function to do scopped DB txs
type DBTxScopedFn func(ctx context.Context, dbTx pgx.Tx) (interface{}, types.Error)

// DBTxer interface to begin DB txs
type DBTxer interface {
	BeginStateTransaction(ctx context.Context) (pgx.Tx, error)
}

// NewDbTxScope function to initiate DB scopped txs
func (f *DBTxManager) NewDbTxScope(db DBTxer, scopedFn DBTxScopedFn) (interface{}, types.Error) {
	ctx := context.Background()
	dbTx, err := db.BeginStateTransaction(ctx)
	if err != nil {
		return RPCErrorResponse(types.DefaultErrorCode, "failed to connect to the state", err)
	}

	v, rpcErr := scopedFn(ctx, dbTx)
	if rpcErr != nil {
		if txErr := dbTx.Rollback(context.Background()); txErr != nil {
			return RPCErrorResponse(types.DefaultErrorCode, "failed to rollback db transaction", txErr)
		}
		return v, rpcErr
	}

	if txErr := dbTx.Commit(context.Background()); txErr != nil {
		return RPCErrorResponse(types.DefaultErrorCode, "failed to commit db transaction", txErr)
	}
	return v, rpcErr
}
func (f *DBTxManager) NewDbTxScopeSCF(db DBTxer, scopedFn DBTxScopedFn, inputLog bool) (interface{}, types.Error) {
	ts := time.Now()
	ctx := context.Background()
	dbTx, err := db.BeginStateTransaction(ctx)
	if err != nil {
		return RPCErrorResponse(types.DefaultErrorCode, "failed to connect to the state", err)
	}
	if inputLog {
		log.Infof("SCF NewDbTxScope BeginStateTransaction=%d", time.Now().Sub(ts).Milliseconds())
	}

	v, rpcErr := scopedFn(ctx, dbTx)
	if inputLog {
		log.Infof("SCF NewDbTxScope scopedFn=%d", time.Now().Sub(ts).Milliseconds())
	}

	if rpcErr != nil {
		if txErr := dbTx.Rollback(context.Background()); txErr != nil {
			return RPCErrorResponse(types.DefaultErrorCode, "failed to rollback db transaction", txErr)
		}
		return v, rpcErr
	}
	if inputLog {
		log.Infof("SCF NewDbTxScope before commit =%d", time.Now().Sub(ts).Milliseconds())
	}

	if txErr := dbTx.Commit(context.Background()); txErr != nil {
		return RPCErrorResponse(types.DefaultErrorCode, "failed to commit db transaction", txErr)
	}
	if inputLog {
		log.Infof("SCF NewDbTxScope end commit =%d", time.Now().Sub(ts).Milliseconds())
	}

	return v, rpcErr
}
