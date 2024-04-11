package redisadapter

import "fmt"

type Key string

func (k Key) Add(key string) Key {
	return k + ":" + Key(key)
}

func (k Key) FormatAndAdd(a ...any) Key {
	return k.Add(fmt.Sprintf(string(k), a...))
}
