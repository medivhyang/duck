package snowflake

import "sync"

var instance *Node
var once = sync.Once{}

func getDefault() *Node {
	once.Do(func() {
		var err error
		if instance, err = NewNode(0); err != nil {
			panic(err)
		}
	})
	return instance
}

func Generate() ID {
	return getDefault().Generate()
}
