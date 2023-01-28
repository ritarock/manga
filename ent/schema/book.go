package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("isbn").Immutable().Optional(),
		field.String("title").Immutable().Optional(),
		field.String("publisher").Immutable().Optional(),
		field.String("pubdate").Immutable().Optional(),
		field.String("cover").Immutable().Optional(),
		field.String("author").Immutable().Optional(),
		field.String("subject_code").Immutable().Optional(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
