package hash

type SquareHashHelper struct {
	size int64
}

func NewSquareHashHelper(size int64) SquareHashHelper {
	return SquareHashHelper{
		size: size,
	}
}

func (sa SquareHashHelper) Name() string {
	return "SquareHashHelper"
}

func (sa SquareHashHelper) Hash(data int64) int64 {
	return (data * data) % sa.size
}
