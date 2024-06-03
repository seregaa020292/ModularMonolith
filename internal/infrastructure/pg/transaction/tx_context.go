package transaction

import (
	"context"
	"database/sql"
)

type txKey struct{}

// InjectTx внедряет транзакцию в контекст
func InjectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// ExtractTx извлекает транзакцию из контекста
func ExtractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}
