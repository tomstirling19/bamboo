package services

import (
	"bamboo/internal/app/utils"
	"bamboo/internal/config"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type GraphQLService struct {
    Config *config.GraphQLConfig
}

func NewGraphQLService(config *config.GraphQLConfig) *GraphQLService {
    return &GraphQLService{Config: config}
}

func (s *GraphQLService) LoadSchema(resolver interface{}) (*graphql.Schema, error) {
    schemaContent, err := utils.ReadFile(s.Config.SchemaPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read schema file: %w", err)
    }

    parsedSchema, err := graphql.ParseSchema(schemaContent, resolver)
    if err != nil {
        return nil, fmt.Errorf("failed to parse schema: %w", err)
    }

    return parsedSchema, nil
}
