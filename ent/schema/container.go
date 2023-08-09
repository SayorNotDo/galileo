package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"regexp"
)

// Container holds the schema definition for the Container entity.
type Container struct {
	ent.Schema
}

func (Container) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "container",
		},
	}
}

// Fields of the Container.
func (Container) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Match(regexp.MustCompile("/?[a-zA-Z0-9][a-zA-Z0-9_.-]+")),
		field.String("hostname").NotEmpty(),
		field.String("domainname").NotEmpty(),
		field.String("user").Optional(),
		field.Strings("env").Optional(),
		field.Strings("cmd").Optional(),
		field.String("image").NotEmpty(),
		field.Strings("labels").Optional(),
		field.Strings("volumes").Optional(),
		field.String("working_dir").Optional(),
		field.Strings("entrypoint").Optional(),
		field.String("mac_address").NotEmpty(),
		field.Strings("expose_ports").Optional(),
		field.String("compose_file_url").Optional(),
		field.String("dockerfile_url").Optional(),
	}
}

// Edges of the Container.
func (Container) Edges() []ent.Edge {
	return nil
}
