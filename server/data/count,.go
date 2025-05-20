package data

import (
	"context"
	"fmt"
)

func (s *Store) IncrementCount(ctx context.Context) (int, error) {
	const query = `
		UPDATE global_stats 
		SET total_count = total_count + 1
		RETURNING total_count`

	var newCount int
	if err := s.db.QueryRow(ctx, query).Scan(&newCount); err != nil {
		return 0, fmt.Errorf("failed to increment count: %w", err)
	}

	return newCount, nil
}