package hash

type HashHelper interface {
	Name() string
	Hash(data int64) int64
}
