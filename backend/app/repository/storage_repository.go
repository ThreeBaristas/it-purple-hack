package repository

type MatricesMappingStorage interface {
	SegmentToMatrix(segmentId int64) (int64, bool)
	BaselineMatrix() int64
	SetUpStorage(req *SetUpStorageRequest) error
}

type InlineMappingStorage struct {
	baselineMatrix int64
	discounts      map[int64]int64
}

func (i *InlineMappingStorage) SegmentToMatrix(segmentId int64) (int64, bool) {
	if segmentId == 0 {
		return i.baselineMatrix, true
	}
	res, ok := i.discounts[segmentId]
	return res, ok
}

func (i *InlineMappingStorage) BaselineMatrix() int64 {
	return i.baselineMatrix
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
	i.baselineMatrix = req.BaselineMatrix
	i.discounts = make(map[int64]int64)
	for _, mapping := range req.Discounts {
		i.discounts[mapping.SegmentId] = mapping.MatrixId
	}
	return nil
}
