package hash

type SimpleAddHashHelper struct {
	step int64
	size int64
}

func NewSimpleAddHashHelper(step int64, size int64) SimpleAddHashHelper {
	return SimpleAddHashHelper{
		step: step,
		size: size,
	}
}

func (sa SimpleAddHashHelper) Name() string {
	return "SimpleAddHashHelper"
}

func (sa SimpleAddHashHelper) Hash(data int64) int64 {
	return (data + sa.step) % sa.size
}
