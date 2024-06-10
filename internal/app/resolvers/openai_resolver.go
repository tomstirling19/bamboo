// openai_resolver.go
// Implements GraphQL resolvers for OpenAI queries. Handles GraphQL
// requests and integrates with the 'openai_service' to fetch responses.

package resolvers

import (
	"context"

	"bamboo/internal/app/services"

	"github.com/graph-gophers/graphql-go"
)

type OpenAIResolver struct {
	OpenAIService *services.OpenAIService
}

type Query struct {
	*OpenAIResolver
}

func (r *Query) GetOpenAIResponse(ctx context.Context, args struct{ Prompt string }) (string, error) {
	response, err := r.OpenAIService.GetResponse(args.Prompt)
	if err != nil {
		return "", err
	}
	return response, nil
}

func NewSchema(resolver *OpenAIResolver) *graphql.Schema {
	schema := `
		schema {
			query: Query
		}
		type Query {
			getOpenAIResponse(prompt: String!): String!
		}
	`
	return graphql.MustParseSchema(schema, &Query{OpenAIResolver: resolver})
}
