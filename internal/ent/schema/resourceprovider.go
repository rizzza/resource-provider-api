package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/vektah/gqlparser/v2/ast"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// ResourceProvider holds the schema definition for the ResourceProvider entity.
type ResourceProvider struct {
	ent.Schema
}

// Mixin of the ResourceProvider type
func (ResourceProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ResourceProvider.
func (ResourceProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the resource provider.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ResourceProviderPrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the resource provider.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("description").
			Optional().
			Comment("The description of the resource provider.").
			Annotations(
				entgql.OrderField("DESCRIPTION"),
			),
		field.String("owner_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the owner for this resource provider.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
				entx.EventsHookAdditionalSubject("owner"),
			),
	}
}

// Edges of the ResourceProvider.
func (ResourceProvider) Edges() []ent.Edge {
	return nil
}

// Indexes of the ResourceProvider.
func (ResourceProvider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

// Annotations of the ResourceProvider
func (ResourceProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents a Resource Provider on the graph."),
		prefixIDDirective(ResourceProviderPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new ResourceProvider."),
			entgql.MutationUpdate().Description("Update an existing ResourceProvider."),
		),
		entx.EventsHookSubjectName("resource-provider"),
	}
}

func prefixIDDirective(prefix string) entgql.Annotation {
	var args []*ast.Argument
	if prefix != "" {
		args = append(args, &ast.Argument{
			Name: "prefix",
			Value: &ast.Value{
				Raw:  prefix,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.Directives(entgql.NewDirective("prefixedID", args...))
}
