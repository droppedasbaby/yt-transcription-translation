package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Text("video_url").NotEmpty().Unique(),
		field.Text("video_id").NotEmpty().Unique(),
		field.Text("file_location").NotEmpty().Optional(),
		field.Enum("status").
			NamedValues(
				"Downloading", "downloading",
				"Downloaded", "downloaded",
				"Transcribed", "transcribed",
				"Translated", "translated",
			),
		field.UUID("run_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}
