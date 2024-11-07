package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/giftee/cqrs-example-go/application/handler/graphql/generated"
	"github.com/giftee/cqrs-example-go/application/handler/graphql/resolver"
	"github.com/giftee/cqrs-example-go/application/infrastructure/database"
	"github.com/giftee/cqrs-example-go/application/usecase/command"
)

func main() {
	conn := database.Connect("cqrs_example", "root", "password", "db", 3306, 2, 5)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	promotionRepository := database.PromotionRepository{Conn: conn}
	couponRepository := database.CouponRepository{Conn: conn}

	submitPromotionCommand := command.SubmitPromotionCommand{Repository: promotionRepository}
	publishPromotionCommand := command.PublishPromotionCommand{Repository: promotionRepository}
	applyPromotionCommand := command.ApplyPromotionCommand{Repository: promotionRepository}
	grantCouponCommand := command.GrantCouponCommand{Repository: couponRepository}
	invalidateCouponCommand := command.InvalidateCouponCommand{Repository: couponRepository}

	customerQueryService := database.CustomerQueryService{Conn: conn}
	promotionQueryService := database.PromotionQueryService{Conn: conn}
	couponQueryService := database.CouponQueryService{Conn: conn}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		SubmitPromotionCommand:  submitPromotionCommand,
		PublishPromotionCommand: publishPromotionCommand,
		ApplyPromotionCommand:   applyPromotionCommand,
		GrantCouponCommand:      grantCouponCommand,
		InvalidateCouponCommand: invalidateCouponCommand,
		CustomerQueryService:    customerQueryService,
		CouponQueryService:      couponQueryService,
		PromotionQueryService:   promotionQueryService,
	}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	http.Handle("/", srv)

	log.Printf("connect to http://localhost:3000/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
