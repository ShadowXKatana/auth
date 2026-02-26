package graphql

import (
	"time"

	"github.com/gin-gonic/gin"
	graphqlLib "github.com/graphql-go/graphql"
	graphqlHandler "github.com/graphql-go/handler"
)

type Handler struct {
	httpHandler *graphqlHandler.Handler
}

func NewHandler() (*Handler, error) {
	h := &Handler{}

	schema, err := h.buildSchema()
	if err != nil {
		return nil, err
	}

	h.httpHandler = graphqlHandler.New(&graphqlHandler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return h, nil
}

func (h *Handler) Serve(c *gin.Context) {
	h.httpHandler.ContextHandler(c.Request.Context(), c.Writer, c.Request)
}

func (h *Handler) buildSchema() (graphqlLib.Schema, error) {
	query := graphqlLib.NewObject(graphqlLib.ObjectConfig{
		Name: "Query",
		Fields: graphqlLib.Fields{
			"ping": &graphqlLib.Field{
				Type: graphqlLib.NewNonNull(graphqlLib.String),
				Resolve: func(p graphqlLib.ResolveParams) (interface{}, error) {
					return "pong", nil
				},
			},
			"serverTime": &graphqlLib.Field{
				Type: graphqlLib.NewNonNull(graphqlLib.String),
				Resolve: func(p graphqlLib.ResolveParams) (interface{}, error) {
					return time.Now().UTC().Format(time.RFC3339), nil
				},
			},
		},
	})

	return graphqlLib.NewSchema(graphqlLib.SchemaConfig{Query: query})
}
