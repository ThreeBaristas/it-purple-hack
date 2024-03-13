package repository

type MatricesMappingStorage interface {
	SegmentToMatrix(segmentId int64) (int64, bool)
	BaselineMatrix() int64
	SetUpStorage(req *SetUpStorageRequest) error
	GetStorage() (*SetUpStorageRequest, error)
  GetSegmentByMatrix(matrix int64) (int64, bool)
}

type InlineMappingStorage struct {
	baselineMatrix int64
	segmentToMatrix      map[int64]int64
  // Между сегментами и матрицами существует биекция,
  // а для биекции, как известно, существует обратная ей
  matrixToSegment   map[int64]int64
}

func DefaultInlineMappingStorage() MatricesMappingStorage {
	return &InlineMappingStorage{
		baselineMatrix: 0,
		segmentToMatrix:      make(map[int64]int64),
    matrixToSegment: make(map[int64]int64),
	}
}

func (i *InlineMappingStorage) SegmentToMatrix(segmentId int64) (int64, bool) {
	if segmentId == 0 {
		return i.baselineMatrix, true
	}
	res, ok := i.segmentToMatrix[segmentId]
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
	i.segmentToMatrix = make(map[int64]int64)
  
  // Design convention: segment 0 is baseline matrix
  i.segmentToMatrix[0] = req.BaselineMatrix
  i.matrixToSegment[req.BaselineMatrix] = 0;

	for _, mapping := range req.Discounts {
		i.segmentToMatrix[mapping.SegmentId] = mapping.MatrixId
    i.matrixToSegment[mapping.MatrixId] = mapping.SegmentId
	}
	return nil
}

func (i *InlineMappingStorage) GetStorage() (*SetUpStorageRequest, error) {
	resp := SetUpStorageRequest{}
	resp.Discounts = []DiscountMappingDTO{}
	resp.BaselineMatrix = i.baselineMatrix
	for segment, matrix := range i.segmentToMatrix {
		resp.Discounts = append(resp.Discounts, DiscountMappingDTO{SegmentId: segment, MatrixId: matrix})
	}
	return &resp, nil
}

func (i *InlineMappingStorage) GetSegmentByMatrix(matrix int64) (int64, bool) {
  val, ok := i.matrixToSegment[matrix]
  return val, ok
}
