package product

import (
	"github.com/AlghazHernanda/dynamodb-go-crud-yt/internal/repository/adapter"
)

type Controller struct {
	repository adapter.Interface
}

// panggil interface yang telah dibuat  di atas ini
type Interface interface {
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entities []product.Product, err error)
	Create(entity *product.Product) (uuid.UUID, error)
	Update(id uuid.UUID, entity *product.Product) error
	Remove(id uuid.UUID) error
}

func NewController(repository adapter.Interface) Interface {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) ListOne(id uuid.UUID) (entity product.Product, err error) {
	entity.ID = id
	response, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())
	if err != nil {
		return entity, err
	}
	return product.ParseDynamoAttributeToStruct(response.Item)
}
