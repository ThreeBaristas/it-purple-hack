package repository

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
  GetPricesBatch(nodes []*GetPriceRequest) ([]*GetPriceResponse, error)
}
