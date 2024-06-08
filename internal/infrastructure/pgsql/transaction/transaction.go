package transaction

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type TxManager interface {
	ReadCommitted(ctx context.Context, fn func(context.Context) error) error
}

type Manager struct {
	db *sql.DB
}

func New(db *sql.DB) *Manager {
	return &Manager{db: db}
}

func (m Manager) ReadCommitted(ctx context.Context, fn func(context.Context) error) error {
	return m.performTx(ctx, fn, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
}

func (m Manager) performTx(ctx context.Context, fn func(context.Context) error, opts *sql.TxOptions) error {
	// Если это вложенная транзакция, пропускаем инициацию новой транзакции и выполняем обработчик.
	if tx := ExtractTx(ctx); tx != nil {
		return fn(ctx)
	}

	// Стартуем новую транзакцию.
	tx, err := m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	// Настраиваем функцию отсрочки для отката или коммита транзакции.
	defer func() {
		// восстанавливаемся после паники
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		// откатываем транзакцию, если произошла ошибка
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}
			return
		}

		// если ошибок не было, коммитим транзакцию
		if err = tx.Commit(); err != nil {
			err = errors.Wrap(err, "tx commit failed")
		}
	}()

	// Выполните код внутри транзакции.
	// Если функция терпит неудачу, возвращаем ошибку, и функция отсрочки выполняет откат
	// или в противном случае транзакция коммитится.
	if err = fn(InjectTx(ctx, tx)); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}
