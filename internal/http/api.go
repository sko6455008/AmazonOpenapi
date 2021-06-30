package api

import (
	"amazonOpenApi/internal/http/gen"
	"amazonOpenApi/internal/http/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Api struct {
	amazon *usecase.Amazon
}

func NewApi(db *gorm.DB) *Api {
	return &Api{amazon: usecase.NewAmazon(db)}
}

var _ gen.ServerInterface = (*Api)(nil)

func (p *Api) AddAmazon(ctx echo.Context) error {
	return p.amazon.AddAmazon(ctx)
}

func (p *Api) GetAmazon(ctx echo.Context, asin string) error {
	return p.amazon.GetAmazon(ctx, asin)
}

func (p *Api) PutAmazon(ctx echo.Context, asin string) error {
	return p.amazon.PutAmazon(ctx, asin)
}

func (p *Api) PatchAmazon(ctx echo.Context, asin string) error {
	return p.amazon.PatchAmazon(ctx, asin)
}

func (p *Api) ActiveAmazon(ctx echo.Context, asin string) error {
	return p.amazon.ActiveAmazon(ctx, asin)
}

func (p *Api) InactiveAmazon(ctx echo.Context, asin string) error {
	return p.amazon.InactiveAmazon(ctx, asin)
}

func (p *Api) DeleteAmazon(ctx echo.Context, asin string) error {
	return p.amazon.DeleteAmazon(ctx, asin)
}
