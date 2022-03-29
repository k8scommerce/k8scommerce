package eventkey

import (
	"encoding/json"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

// catalog

/////////////////
// CatalogImageUploadedKey
/////////////////

func (s *CatalogImageUploadedKey) Marshal(obj *models.Asset) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CatalogImageUploadedKey) Unmarshal(data []byte) (obj *models.Asset, err error) {
	obj = &models.Asset{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CatalogImageUploadedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CatalogProductAddedKey
/////////////////

func (s *CatalogProductAddedKey) Marshal(obj *catalog.Product) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CatalogProductAddedKey) Unmarshal(data []byte) (obj *catalog.Product, err error) {
	obj = &catalog.Product{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CatalogProductAddedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CatalogProductUpdatedKey
/////////////////

func (s *CatalogProductUpdatedKey) Marshal(obj *catalog.Product) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CatalogProductUpdatedKey) Unmarshal(data []byte) (obj *catalog.Product, err error) {
	obj = &catalog.Product{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CatalogProductUpdatedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CatalogProductDeletedKey
/////////////////

func (s *CatalogProductDeletedKey) Marshal(obj *catalog.Product) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CatalogProductDeletedKey) Unmarshal(data []byte) (obj *catalog.Product, err error) {
	obj = &catalog.Product{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CatalogProductDeletedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CatalogCategoryAddedKey
/////////////////

func (s *CatalogCategoryAddedKey) Marshal(obj *catalog.Category) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CatalogCategoryAddedKey) Unmarshal(data []byte) (obj *catalog.Category, err error) {
	obj = &catalog.Category{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CatalogCategoryAddedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CatalogCategoryUpdatedKey
/////////////////

func (s *CatalogCategoryUpdatedKey) Marshal(obj *catalog.Category) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CatalogCategoryUpdatedKey) Unmarshal(data []byte) (obj *catalog.Category, err error) {
	obj = &catalog.Category{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CatalogCategoryUpdatedKey) AsKey() EventKey {
	return EventKey(*s)
}

/////////////////
// CatalogCategoryDeletedKey
/////////////////

func (s *CatalogCategoryDeletedKey) Marshal(obj *catalog.Category) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *CatalogCategoryDeletedKey) Unmarshal(data []byte) (obj *catalog.Category, err error) {
	obj = &catalog.Category{}
	if err = json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *CatalogCategoryDeletedKey) AsKey() EventKey {
	return EventKey(*s)
}
