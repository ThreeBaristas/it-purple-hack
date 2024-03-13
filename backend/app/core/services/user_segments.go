package services

import "math/rand"

func NewGetUserSegmentsService() *GetUserSegmentsService {
	return &GetUserSegmentsService{}
}

type GetUserSegmentsService struct {
}

func randomArray() []int64 {
  return []int64{
    int64(rand.Intn(200) + 1),
    int64(rand.Intn(200) + 1),
    int64(rand.Intn(200) + 1),
    int64(rand.Intn(200) + 1),
    int64(rand.Intn(200) + 1),
  }
}

var db = map[int64][]int64{
	2100: randomArray(),
	2200: randomArray(),
	2300: randomArray(),
	2400: randomArray(),
	2500: randomArray(),
	2600: randomArray(),
	2700: randomArray(),
	2800: randomArray(),
	2900: randomArray(),
	3000: randomArray(),
	3100: randomArray(),
	3200: randomArray(),
	3300: randomArray(),
	3400: randomArray(),
	3500: randomArray(),
	3600: randomArray(),
	3700: randomArray(),
	3800: randomArray(),
	3900: randomArray(),
	4000: randomArray(),
	4100: randomArray(),
	4200: randomArray(),
}

func (g *GetUserSegmentsService) GetSegments(userId int64) ([]int64, error) {
	val, ok := db[userId]
	if !ok {
		return nil, nil
	}
	return val, nil
}
