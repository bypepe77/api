package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.String("username").
			Unique().
			NotEmpty().
			MaxLen(25),
		field.String("password").
			NotEmpty().
			Default("unknown"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Bool("active").
			Default(false),
		field.Int("follows_count").
			Positive().
			Default(1),
		field.Int("following_count").
			Positive().
			Default(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			From("followers"),
		edge.To("liked_posts", Post.Type).
			Through("likes", Likes.Type),
	}
}
