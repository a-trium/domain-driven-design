package product

import (
	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi"
	productapi "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/product"
	"github.com/go-openapi/runtime/middleware"
	dto "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagmodel"
)

type ProductHandler interface {
	Configure(handlerRegistry *swagapi.GatewayAPI)
	FindAllProducts(itemCountPerPage int, currentPageOffset int) (int, []*Product, e.Exception)
}

type productHandlerImpl struct {
	productRepository Repository
}

func NewProductHandler(repo Repository) ProductHandler {
	return &productHandlerImpl{
		productRepository: repo,
	}
}

func (h *productHandlerImpl) Configure(registry *swagapi.GatewayAPI) () {
	registry.ProductFindAllProductHandler = productapi.FindAllProductHandlerFunc(
		func(params productapi.FindAllProductParams) middleware.Responder {
			currentPageOffset := int(*params.CurrentPageOffset)
			itemCountPerPage := int(*params.ItemCountPerPage)
			totalCount, productList, ex := h.FindAllProducts(currentPageOffset, itemCountPerPage)
			totalItemCount := int64(totalCount)

			if ex != nil {
				return productapi.NewFindAllProductDefault(ex.StatusCode()).WithPayload(ex.ToSwaggerError())
			}

			rows := make([]*dto.Product, 0)
			for i := range productList {
				product := productList[i]
				dto := product.convertToDTO()
				rows = append(rows, dto)
			}

			response := dto.FindAllProductOKBody{
				Pagination: &dto.Pagination{
					CurrentPageOffset: params.CurrentPageOffset,
					ItemCountPerPage: params.ItemCountPerPage,
					TotalItemCount: &totalItemCount,
				},

				Rows: rows,
			}

			return productapi.NewFindAllProductOK().WithPayload(&response)
		})
}

func (h *productHandlerImpl) FindAllProducts(currentPageOffset int, itemCountPerPage int) (int, []*Product, e.Exception) {
	if itemCountPerPage <= 0 || itemCountPerPage > 20 {
		itemCountPerPage = 20
	}
	if currentPageOffset < 0 {
		currentPageOffset = 0
	}

	totalCount, productList, ex := h.productRepository.FindAllProducts(itemCountPerPage, currentPageOffset)
	if ex != nil {
		return 0, nil, ex
	}

	return totalCount, productList, nil
}
