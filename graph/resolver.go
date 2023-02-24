package graph

import "github.com/ritarock/manga/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EntClient *ent.Client
}
