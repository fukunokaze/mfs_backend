package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MFSUser holds the schema definition for the MFSUser entity.
type MFSUser struct {
	ent.Schema
}

// Fields of the MFSUser.
func (MFSUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("userId"),
		field.String("username"),
		field.String("password"),
		field.String("address").Optional(),
	}
}

// Edges of the MFSUser.
func (MFSUser) Edges() []ent.Edge {
	return nil
}
