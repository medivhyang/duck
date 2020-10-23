package snowflake

import "sync"

var instance *Node
var once = sync.Once{}

func Generate() ID {
	once.Do(func() {
		var err error
		if instance, err = NewNode(0); err != nil {
			panic(err)
		}
	})
	return instance.Generate()
}
