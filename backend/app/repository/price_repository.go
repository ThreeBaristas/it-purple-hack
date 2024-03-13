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
	SetPrice(locationId int64, categoryId int64, matrixId int64, price int64) (*GetPriceResponse, error)
	DeletePrice(locationId int64, categoryId int64, matrixId int64) (bool, error)
	GetRules(pageSize int32, page int64) ([]GetPriceResponse, int, error)
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
			cur.MatrixId = 0
		}

		ans = append(ans, cur)
	}
	return ans, nil
}

func (r *PostgresPriceRepository) SetPrice(locationId int64, categoryId int64, matrixid int64, price int64) (*GetPriceResponse, error) {
	_, err := r.db.Exec("INSERT INTO prices (location_id, category_id, matrix_id, price) VALUES ($1, $2, $3, $4) ON CONFLICT (category_id, location_id, matrix_id) DO UPDATE SET price = $4", locationId, categoryId, matrixid, price)
	if err != nil {
		return nil, err
	}
	return &GetPriceResponse{
		LocationId: locationId,
		CategoryId: categoryId,
		MatrixId:   matrixid,
		Price:      price,
	}, nil
}

func (r *PostgresPriceRepository) DeletePrice(locationId int64, categoryId int64, matrixId int64) (bool, error) {
	res, err := r.db.Exec("DELETE FROM prices WHERE location_id = $1 AND category_id = $2 AND matrix_id = $3", locationId, categoryId, matrixId)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

func (r *PostgresPriceRepository) GetRules(pageSize int32, page int64) ([]GetPriceResponse, int, error) {
	rows, err := r.db.Query("SELECT location_id, category_id, price, matrix_id, count(*) OVER() AS full_count FROM prices ORDER BY (location_id, category_id, matrix_id) LIMIT $1 OFFSET $2", pageSize, page*int64(pageSize))
	if err != nil {
		return nil, 0, err
	}

	var ans []GetPriceResponse
	var count int
	for rows.Next() {
		cur := GetPriceResponse{}
		var matrixId sql.Null[int64]
		err := rows.Scan(&cur.LocationId, &cur.CategoryId, &cur.Price, &matrixId, &count)
		if err != nil {
			return nil, 0, err
		}

		if matrixId.Valid {
			cur.MatrixId = matrixId.V
		} else {
			cur.MatrixId = 0
		}

		ans = append(ans, cur)
	}
	return ans, (count + int(pageSize) - 1) / int(pageSize), nil
}
