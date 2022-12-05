package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Follows holds the schema definition for the Follows entity.
type Follows struct {
	ent.Schema
}

// Fields of the Follows.
func (Follows) Fields() []ent.Field {
	return []ent.Field{
		field.Int("followedby").
			Positive(),
		field.Int("follower").
			Positive(),
	}
}

// Edges of the Follows.
func (Follows) Edges() []ent.Edge {
	return nil
}
