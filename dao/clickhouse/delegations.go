package clickhouse

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/everstake/cosmoscan-api/dmodels"
)

func (db DB) CreateDelegations(delegations []dmodels.Delegation) error {
	if len(delegations) == 0 {
		return nil
	}
	q := squirrel.Insert(dmodels.DelegationsTable).Columns("dlg_id", "dlg_tx_hash", "dlg_delegator", "dlg_validator", "dlg_amount", "dlg_created_at")
	for _, delegation := range delegations {
		if delegation.ID == "" {
			return fmt.Errorf("field ID can not be empty")
		}
		if delegation.TxHash == "" {
			return fmt.Errorf("field TxHash can not be empty")
		}
		if delegation.Delegator == "" {
			return fmt.Errorf("field Delegator can not be empty")
		}
		if delegation.Validator == "" {
			return fmt.Errorf("field Validator can not be empty")
		}
		if delegation.CreatedAt.IsZero() {
			return fmt.Errorf("field CreatedAt can not be zero")
		}
		q = q.Values(delegation.ID, delegation.TxHash, delegation.Delegator, delegation.Validator, delegation.Amount, delegation.CreatedAt)
	}
	return db.Insert(q)
}