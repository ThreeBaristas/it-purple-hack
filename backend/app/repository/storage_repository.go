package repository

type MatricesMappingStorage interface {
	SegmentToMatrix(segmentId int64) (int64, bool)
	BaselineMatrix() int64
}

type InlineMappingStorage struct {
	baselineMatrix int64
	mapping        map[int64]int64
}

func (i *InlineMappingStorage) SegmentToMatrix(segmentId int64) (int64, bool) {
	res, ok := i.mapping[segmentId]
	return res, ok
}

func (i *InlineMappingStorage) BaselineMatrix() int64 {
	return i.baselineMatrix
}
