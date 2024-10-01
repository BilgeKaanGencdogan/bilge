package types

const (
	// ModuleName defines the module name
	ModuleName = "bilge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bilge"
)

var (
	ParamsKey = []byte("p_bilge")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
