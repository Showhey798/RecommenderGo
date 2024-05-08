package bigtable

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigtable"
	"recommender.package/internal/domain/entity"
)

const (
	ColumnFamily = "movie_info"
)

type MovieRepository struct {
	table *bigtable.Table
}

func (repo *MovieRepository) Get(ctx context.Context, id string) (movie entity.Movie, err error) {
	row, err := repo.table.ReadRow(
		ctx,
		id,
		bigtable.RowFilter(bigtable.FamilyFilter(ColumnFamily)),
	)
	if err != nil {
		return entity.Movie{}, fmt.Errorf("bigtable: could not read row: %v", err)
	}

	if len(row) == 0 {
		return entity.Movie{}, fmt.Errorf("bigtable: row not found")
	}
	err = json.Unmarshal(row[ColumnFamily][0].Value, &movie)
	if err != nil {
		return entity.Movie{}, fmt.Errorf("bigtable: could not unmarshal movie: %v", err)
	}
	return movie, nil
}
