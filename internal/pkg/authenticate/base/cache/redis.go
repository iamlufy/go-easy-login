package cache

//TODO redis
var cache map[string]string = make(map[string]string)

func Add(key, value string, activeSeconds int) {
	cache[key] = value
}

func Get(key string) string {
	return cache[key]
}
