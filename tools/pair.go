package tools

type Pair struct {
	key   uint
	value interface{}
}

func NewPair(key uint, value interface{}) *Pair {
	return &Pair{
		key,
		value,
	}
}

func (pair *Pair) CompareKey(pair2 *Pair) bool {
	return pair.key == pair2.key
}

func (pair *Pair) Value() interface{} {
	return pair.value
}

func NewMap(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}
