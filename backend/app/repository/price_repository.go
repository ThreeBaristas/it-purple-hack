package repository

import (
	"database/sql"

	"github.com/lib/pq"
)

type GetPriceRequest struct {
	LocationId []int64
	CategoryId []int64
	Matrices   []int64
}

type GetPriceResponse struct {
	LocationId int64
	CategoryId int64
	Price      int64
	MatrixId   int64
}

type PriceRepository interface {
	GetPricesBatch(nodes *GetPriceRequest) ([]GetPriceResponse, error)
}

type PostgresPriceRepository struct {
	db *sql.DB
}

func NewPostgresPriceRepository(db *sql.DB) PriceRepository {
	return &PostgresPriceRepository{
		db: db,
	}
}

func (r *PostgresPriceRepository) GetPricesBatch(req *GetPriceRequest) ([]GetPriceResponse, error) {
	rows, err := r.db.Query("SELECT location_id, category_id, price, matrix_id FROM prices WHERE location_id = any($1) AND category_id = any($2) AND (matrix_id = any($3) OR matrix_id = 0) ORDER BY matrix_id DESC", pq.Array(req.LocationId), pq.Array(req.CategoryId), pq.Array(req.Matrices))
	if err != nil {
		return nil, err
	}

	var ans []GetPriceResponse
	for rows.Next() {
		cur := GetPriceResponse{}
		var matrixId sql.Null[int64]
		err := rows.Scan(&cur.LocationId, &cur.CategoryId, &cur.Price, &matrixId)
		if err != nil {
			return nil, err
		}

		if matrixId.Valid {
			cur.MatrixId = matrixId.V
		} else {
			// cur.MatrixId = 0
		}

		ans = append(ans, cur)
	}
	return ans, nil
}
