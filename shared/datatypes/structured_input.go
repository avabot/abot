package dt

import "time"

// StructuredInput is generated by Abot and sent to plugins as a helper tool.
// Additional fields should be added, covering Times, Places, etc. to make
// plugin development even easier. Note that right now People is unused.
type StructuredInput struct {
	Commands []string
	Objects  []string
	Intents  []string
	People   []Person
	Times    []time.Time

	// TODO
	// Places []string
}

// SIT is a Structured Input Type. It corresponds to either a Command or an
// Object with additional Structured Input Types to be added later.
type SIT int
