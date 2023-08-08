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
		field.String("hostname"),
		field.String("domainname"),
		field.String("user"),
		field.Bool("attach_stdin"),
		field.Bool("attach_stdout"),
		field.Bool("attach_stderr"),
		field.Bool("tty"),
		field.Bool("open_stdin"),
		field.Bool("stdin_once"),
		field.Strings("env"),
		field.Strings("cmd"),
		field.String("image"),
		field.String("labels"),
		field.String("volumes"),
		field.String("working_dir"),
		field.String("entrypoint"),
		field.Bool("network_disabled"),
		field.String("mac_address"),
		field.String("expose_ports"),
		field.String("stop_signal"),
		field.String("stop_timeout"),
		field.String("host_config"),
		field.String("network_config"),
	}
}

// Edges of the Container.
func (Container) Edges() []ent.Edge {
	return nil
}
