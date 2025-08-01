package infrastracture

import (
	"context"
	"database/sql"

	"github.com/tf63/go-tx-sample/internal/context-pattern/db"
	"github.com/tf63/go-tx-sample/internal/context-pattern/db/xcontext"
)

type BaseRepository struct {
	_db          *sql.DB
	txContextKey any
}

func NewBaseRepository(db *sql.DB) BaseRepository {
	return BaseRepository{
		_db:          db,
		txContextKey: xcontext.Key(),
	}
}

// Repository内で呼ぶ用
func (b *BaseRepository) DB(ctx context.Context) db.DB {
	// ctxにトランザクションが含まれている場合はそれを使用し、そうでなければ通常のDBを返す
	if tx, ok := xcontext.GetTx(ctx); ok {
		return tx
	}
	return b._db
}
