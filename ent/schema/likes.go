package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Likes holds the schema definition for the Likes entity.
type Likes struct {
	ent.Schema
}

func (Likes) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "post_id"),
	}
}

// Fields of the Likes.
func (Likes) Fields() []ent.Field {
	return []ent.Field{
		field.Time("liked_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("post_id"),
	}
}

// Edges of the Likes.
func (Likes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("tweet", Post.Type).
			Unique().
			Required().
			Field("post_id"),
	}
}
