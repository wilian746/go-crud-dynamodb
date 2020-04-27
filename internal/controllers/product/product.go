package product

import (
	"github.com/google/uuid"
	"github.com/wilian746/go-crud-dynamodb/internal/entities/product"
	"github.com/wilian746/go-crud-dynamodb/internal/repository/adapter"
	"time"
)

type Controller struct {
	repository adapter.Interface
}

type Interface interface {
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entities []product.Product, err error)
	Create(entity *product.Product) (uuid.UUID, error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error
}

func NewController(repository adapter.Interface) Interface {
	return &Controller{repository: repository}
}

func (c *Controller) ListOne(id uuid.UUID) (entity product.Product, err error) {
	response, err := c.repository.Find(map[string]interface{}{"_id": id}, entity.TableName())
	if err != nil {
		return entity, err
	}

	return entity, c.repository.ParseDynamoAtributeToStruct(&entity, response.Item)
}

func (c *Controller) ListAll() (entities []product.Product, err error) {
	var entity product.Product
	response, err := c.repository.Find(map[string]interface{}{}, entity.TableName())
	if err != nil {
		return entities, err
	}

	return entities, c.repository.ParseDynamoAtributeToStruct(&entities, response.Item)
}

func (c *Controller) Create(entity *product.Product) (uuid.UUID, error) {
	entity.CreatedAt = time.Now()
	response, err := c.repository.Create(entity, entity.TableName())
	if err != nil {
		return entity.ID, err
	}

	err = c.repository.ParseDynamoAtributeToStruct(entity, response.Attributes)
	return entity.ID, err
}

func (c *Controller) Update(id uuid.UUID, entity *product.Product) error {
	entity.UpdatedAt = time.Now()

	_, err := c.repository.Update(map[string]interface{}{"_id": id}, &entity, entity.TableName())

	return err
}

func (c *Controller) Remove(id uuid.UUID) error {
	var entity product.Product

	_, err := c.repository.Delete(map[string]interface{}{"_id": id}, entity.TableName())

	return err
}
