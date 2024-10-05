package types

const (
	// ModuleName defines the module name
	ModuleName = "lahmacun"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lahmacun"
)

var (
	ParamsKey = []byte("p_lahmacun")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
