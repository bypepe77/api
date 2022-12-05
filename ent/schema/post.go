package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("author_id").
			Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.String("content").
			NotEmpty(),
		field.JSON("hashtags", []string{}).
			Optional(),
		field.JSON("links", []string{}).
			Optional(),
		field.Bool("shared").
			Default(false),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).
			Field("author_id").
			Unique(),
		edge.From("liked_users", User.Type).
			Ref("liked_posts").
			Through("likes", Likes.Type),
	}
}
