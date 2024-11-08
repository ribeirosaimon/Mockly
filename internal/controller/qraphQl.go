package controller

import (
	"github.com/graphql-go/graphql"
)

func NewGraphQlDinamic() map[string]*graphql.Field {
	return map[string]*graphql.Field{
		"Teste": {
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "DEU BOM", nil
			},
		},
	}
}
