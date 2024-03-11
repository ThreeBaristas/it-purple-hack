package repository

import (
	"database/sql"

	"github.com/lib/pq"
)

type GetPriceRequest struct {
  locationId int64
  categoryId int64
  segments   []int64
}

type GetPriceResponse struct {
  locationId int64
  categoryId int64
  price      int64
  matrix_id  int64
}

type PriceRepository interface {
  GetPricesBatch(nodes []*GetPriceRequest) ([]GetPriceResponse, error)
}

type PostgresPriceRepository struct {
  db *sql.DB
}

func NewPostgresPriceRepository(db *sql.DB) PostgresPriceRepository {
  return PostgresPriceRepository{
    db: db,
  }
}

func (r *PostgresPriceRepository) GetPricesBatch(nodes []*GetPriceRequest) ([]GetPriceResponse, error) {
  var locations []int64;
  var categories []int64;
  for _, node := range nodes {
    locations = append(locations, node.locationId)
    categories = append(categories, node.categoryId)
  }
  rows, err := r.db.Query("SELECT location_id, category_id, price, matrix_id FROM prices WHERE location_id IN $1 AND category_id IN $2", pq.Array(locations), pq.Array(categories))
  if(err != nil) {
    return nil, err
  }

  var ans []GetPriceResponse;
  for rows.Next() {
    cur := GetPriceResponse{}
    err := rows.Scan(&cur.locationId, &cur.categoryId, &cur.price, &cur.matrix_id);
    if(err != nil) {
      return nil, err
    }
    ans = append(ans, cur)
  }
  return ans, nil
}
