package types

const (
	// ModuleName defines the module name
	ModuleName = "superchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_superchain"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PythonCodeKey      = "PythonCode/value/"
	PythonCodeCountKey = "PythonCode/count/"
)

const (
	GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
	GameCreatedEventCreator   = "creator"          // Subsidiary information
	GameCreatedEventGameIndex = "game-index"       // What game is relevant
	GameCreatedEventBlack     = "black"            // Is it relevant to me?
	GameCreatedEventRed       = "red"              // Is it relevant to me?
)
