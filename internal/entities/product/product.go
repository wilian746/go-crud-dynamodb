package product

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/wilian746/go-crud-dynamodb/internal/entities"
	"time"
)

type Product struct {
	entities.Base
	Name string `json:"name"`
}

func InterfaceToModel(data interface{}) (instance *Product, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Product) GetFilterId() map[string]interface{} {
	return map[string]interface{}{"_id": p.ID.String()}
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":       p.ID.String(),
		"name":      p.Name,
		"createdAt": p.CreatedAt.Format(entities.GetTimeFormat()),
		"updatedAt": p.UpdatedAt.Format(entities.GetTimeFormat()),
	}
}

func ParseDynamoAtributeToStruct(response map[string]*dynamodb.AttributeValue) (p Product, err error) {
	if response == nil || (response != nil && len(response) == 0) {
		return p, errors.New("Item not found")
	}
	for key, value := range response {
		if key == "_id" {
			p.ID, err = uuid.Parse(*value.S)
			if p.ID == uuid.Nil {
				err = errors.New("Item not found")
			}
		}
		if key == "name" {
			p.Name = *value.S
		}
		if key == "createdAt" {
			p.CreatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "updatedAt" {
			p.UpdatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if err != nil {
			return p, err
		}
	}

	return p, nil
}
