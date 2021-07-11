package hash

type LinearHashHelper struct {
	weight int64
	bias   int64
	size   int64
}

/**
create a w * x + b hash helper
*/
func NewLinearHashHelper(weight int64, bias int64, size int64) LinearHashHelper {
	return LinearHashHelper{
		weight: weight,
		bias:   bias,
		size:   size,
	}
}

func (sa LinearHashHelper) Name() string {
	return "LinearHashHelper"
}

func (sa LinearHashHelper) Hash(data int64) int64 {
	return (sa.weight*data + sa.bias) % sa.size
}
