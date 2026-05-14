package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/trjade/wallet-transfer-service/internal/domain"
	"github.com/trjade/wallet-transfer-service/internal/repository"
)

type LedgerRepository struct {
	db *pgxpool.Pool
}

func NewLedgerRepository(db *pgxpool.Pool) repository.LedgerRepository {
	return &LedgerRepository{db: db}
}

func (r *LedgerRepository) getExecutor(ctx context.Context) QueryExecutor {
	tx := extractTx(ctx)
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *LedgerRepository) CreateEntries(ctx context.Context, entries []domain.LedgerEntry) error {
	query := `
		INSERT INTO ledger_entries (
		id, 
		transfer_id, 
		wallet_id, 
		amount, 
		entry_type, 
		created_at
		)	
		VALUES ($1, $2, $3, $4, $5, NOW())
	`

	batch := &pgxpool.Batch{}

	for _, entry := range entries {
		batch.Queue(query, 
			entry.ID, 
			entry.TransferID, 
			entry.WalletID, 
			entry.Amount, 
			entry.EntryType
		)
	}
	be := r.getExecutor(ctx).SendBatch(ctx, batch)
	defer be.Close()

	for range entries {
		if _, err := be.Exec(); err != nil {
			return err
		}
	}

	return nil
}
