package repository

import (
	"context"
	"database/sql"
)

type MatricesMappingStorage interface {
	SegmentToMatrix(segmentId int64) (int64, bool)
	BaselineMatrix() int64
	SetUpStorage(req *SetUpStorageRequest) error
	GetStorage() (*SetUpStorageRequest, error)
	GetSegmentByMatrix(matrix int64) (int64, bool)
}

type InlineMappingStorage struct {
	db *sql.DB
}

func DefaultInlineMappingStorage(db *sql.DB) MatricesMappingStorage {
	return &InlineMappingStorage{
		db: db,
	}
}

func (i *InlineMappingStorage) SegmentToMatrix(segmentId int64) (int64, bool) {
	row := i.db.QueryRow("SELECT matrix_id FROM storage_mapping WHERE segment_id = $1", segmentId)
	var matrixId int64
	err := row.Scan(&matrixId)
	return matrixId, err == nil
}

func (i *InlineMappingStorage) BaselineMatrix() int64 {
	data, _ := i.SegmentToMatrix(0)
	return data
}

type DiscountMappingDTO struct {
	SegmentId int64 `json:"segment_id"`
	MatrixId  int64 `json:"matrix_id"`
}

type SetUpStorageRequest struct {
	BaselineMatrix int64                `json:"baseline_matrix_id"`
	Discounts      []DiscountMappingDTO `json:"discounts"`
}

func (i *InlineMappingStorage) SetUpStorage(req *SetUpStorageRequest) error {
	tx, err := i.db.BeginTx(context.TODO(), nil)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM storage_mapping")
	if err != nil {
		return err
	}
	setSegmentMatrix := func(segment int64, matrix int64) error {
		_, err = tx.Exec("INSERT INTO storage_mapping(segment_id, matrix_id) VALUES ($1, $2) ON CONFLICT (segment_id) DO UPDATE SET matrix_id = $2", segment, matrix)
		return err
	}
	err = setSegmentMatrix(0, req.BaselineMatrix)
	if err != nil {
		return tx.Rollback()
	}
	for _, dto := range req.Discounts {
		if setSegmentMatrix(dto.SegmentId, dto.MatrixId) != nil {
			return tx.Rollback()
		}
	}
	return tx.Commit()
}

func (i *InlineMappingStorage) GetStorage() (*SetUpStorageRequest, error) {
	resp := SetUpStorageRequest{}
	resp.Discounts = []DiscountMappingDTO{}

	rows, err := i.db.Query("SELECT segment_id, matrix_id FROM storage_mapping")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cur := DiscountMappingDTO{}
		rows.Scan(&cur.SegmentId, &cur.MatrixId)
		if cur.SegmentId == 0 {
			resp.BaselineMatrix = cur.MatrixId
		}
		resp.Discounts = append(resp.Discounts, cur)
	}

	return &resp, nil
}

func (i *InlineMappingStorage) GetSegmentByMatrix(matrix int64) (int64, bool) {
	row := i.db.QueryRow("SELECT segment_id FROM storage_mapping WHERE matrix_id = $1", matrix)
	var segment int64
	err := row.Scan(&segment)
	return segment, err == nil
}
