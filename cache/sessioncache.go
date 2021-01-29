package cache

import "github.com/quasoft/memstore"

var SessionStore = memstore.NewMemStore(
	[]byte("123123123"),
)
