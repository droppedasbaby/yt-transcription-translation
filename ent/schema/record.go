package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/GrewalAS/yt-transcription-translation/types"
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
			NamedValues(string(types.RecordStatusDownloading), string(types.RecordStatusDownloaded)),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}
