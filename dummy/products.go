package dummyjson

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type ProductResponse struct {
	Products []Product `json:"products"`
	Total    uint      `json:"total"`
	Skip     uint      `json:"skip"`
	Limit    uint      `json:"limit"`
}

type Product struct {
	Id                   int       `json:"id"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	Category             string    `json:"category"`
	Price                float64   `json:"price"`
	DiscountPercentage   float64   `json:"discountPercentage"`
	Rating               float64   `json:"rating"`
	Stock                uint      `json:"stock"`
	Tags                 []string  `json:"tags"`
	Brand                string    `json:"brand"`
	Sku                  string    `json:"sku"`
	Weight               float64   `json:"weight"`
	Dimensions           Dimension `json:"dimensions"`
	WarrantyInfo         string    `json:"warrantyInformation"`
	ShippingInfo         string    `json:"shippingInformation"`
	AvailabilityStatus   string    `json:"availabilityStatus"`
	Reviews              []Review  `json:"reviews"`
	ReturnPolicy         string    `json:"returnPolicy"`
	MinimumOrderQuantity uint      `json:"minimumOrderQuantity"`
	Meta                 Meta      `json:"meta"`
	Thumbnail            string    `json:"thumbnail"`
	Images               []string  `json:"images"`
	IsDeleted            bool      `json:"isDeleted,omitempty"`
	DeletedOn            time.Time `json:"deletedOn,omitempty"`
}

type Meta struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Barcode   string    `json:"barcode"`
	QrCode    string    `json:"qrCode"`
}

type Review struct {
	Rating        uint8     `json:"rating"`
	Comment       string    `json:"comment"`
	Date          time.Time `json:"date"`
	ReviewerName  string    `json:"reviewerName"`
	ReviewerEmail string    `json:"reviewerEmail"`
}

type Dimension struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
}

type DummyClient struct {
	client *resty.Client
}

type DummyError struct {
	Message string
}

func (de DummyError) Error() string {
	return fmt.Sprintf("error from the server: %v", de.Message)
}

func NewDummyClient(url string) *DummyClient {
	return &DummyClient{client: resty.New().SetBaseURL(url)}
}

func (dc *DummyClient) GetProducts() ([]Product, error) {
	skip, limit := 0, 30
	var prodRes ProductResponse
	req := dc.client.R().SetResult(&prodRes)
	res, err := req.Get("/products")
	if res.IsError() {
		return nil, DummyError{Message: string(res.Body())}
	}
	if err != nil {
		return nil, err
	}
	var products []Product

	for limit == 30 {
		res, err = req.SetQueryParams(map[string]string{
			"limit": strconv.Itoa(limit),
			"skip":  strconv.Itoa(skip),
		}).Get("/products")
		if res.IsError() {
			return nil, DummyError{Message: string(res.Body())}
		}
		if err != nil {
			return nil, err
		}
		skip += int(prodRes.Limit)
		limit = int(prodRes.Limit)
		products = append(products, prodRes.Products...)
	}
	return products, nil
}

func (dc *DummyClient) GetProduct(id int) (Product, error) {
	var product Product
	res, err := dc.client.R().SetResult(&product).Get(fmt.Sprintf("/products/%d", id))
	if res.IsError() {
		return Product{}, DummyError{Message: string(res.Body())}
	}
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (dc *DummyClient) UploadProduct(prod Product) (Product, error) {
	var created Product
	res, err := dc.client.R().SetBody(prod).SetResult(&created).Post("/products/add")
	if res.IsError() {
		return Product{}, DummyError{Message: string(res.Body())}
	}
	if err != nil {
		return Product{}, err
	}
	return created, nil
}

func (dc *DummyClient) UpdateProduct(id int, prod Product) (Product, error) {
	var updated Product
	res, err := dc.client.R().SetBody(prod).SetResult(&updated).Patch(fmt.Sprintf("/products/%d", id))
	if res.IsError() {
		return Product{}, DummyError{Message: string(res.Body())}
	}
	if err != nil {
		return Product{}, err
	}
	return updated, nil
}

func (dc *DummyClient) DeleteProduct(id int) (Product, error) {
	var deleted Product
	res, err := dc.client.R().SetResult(&deleted).Delete(fmt.Sprintf("/products/%d", id))
	if res.IsError() {
		return Product{}, DummyError{Message: string(res.Body())}
	}
	if err != nil {
		return Product{}, err
	}
	return deleted, nil
}
