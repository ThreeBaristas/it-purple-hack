package repository

type MatricesMappingStorage interface {
	SegmentToMatrix(segmentId int64) (int64, bool)
	BaselineMatrix() int64
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
	segmentId int64
	matrixId  int64
}

type SetUpStorageRequest struct {
	baselineMatrix int64
	discounts      []DiscountMappingDTO
}

func (i *InlineMappingStorage) SetUpStorage(req *SetUpStorageRequest) error {
	i.baselineMatrix = req.baselineMatrix
	i.discounts = make(map[int64]int64)
	for _, mapping := range req.discounts {
		i.discounts[mapping.segmentId] = mapping.matrixId
	}
	return nil
}
