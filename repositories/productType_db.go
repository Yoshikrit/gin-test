package repositories

import (
    "gorm.io/gorm"

	"gin-test/utils/errs"
	"gin-test/models"
)

type productTypeRepositoryDB struct {
	db *gorm.DB
}

func NewProductTypeRepositoryDB(db *gorm.DB) ProductTypeRepository {
	return &productTypeRepositoryDB{db: db}
}


func (r *productTypeRepositoryDB) Create(prodTypeReq *models.ProductTypeCreate) (*models.ProductTypeEntity, error) {
	var existProdTypeEntity models.ProductTypeEntity
	if err := r.db.Where("prodtype_name = ?", prodTypeReq.Name).First(&existProdTypeEntity).Error; err == nil {
		return nil, errs.NewConflictError("Product Type with the same name already exists")
	}

	var prodTypeEntity models.ProductTypeEntity
	prodTypeEntity.Id = prodTypeReq.Id
	prodTypeEntity.Name = prodTypeReq.Name

	if err := r.db.Create(&prodTypeEntity).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &prodTypeEntity, nil
}


func (r *productTypeRepositoryDB) GetAll() ([]models.ProductTypeEntity, error) {
	var prodTypeFromDB []models.ProductTypeEntity
	err := r.db.Find(&prodTypeFromDB).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return prodTypeFromDB, nil
}

func (r *productTypeRepositoryDB) GetById(id int) (*models.ProductTypeEntity, error) {
	var prodTypeFromDB models.ProductTypeEntity
	err := r.db.First(&prodTypeFromDB, id).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &prodTypeFromDB, nil
}

func (r *productTypeRepositoryDB) Update(id int, updateProdType *models.ProductTypeUpdate) (*models.ProductTypeEntity, error) {
	prodTypeFromDB, err := r.GetById(id)
	if err != nil {
		return nil, errs.NewNotFoundError(err.Error())
	}

	prodTypeFromDB.Name = updateProdType.Name

	if err := r.db.Save(prodTypeFromDB).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return prodTypeFromDB, nil
}

func (r *productTypeRepositoryDB) DeleteById(id int) error {
	_, err := r.GetById(id)
	if err != nil {
		return errs.NewNotFoundError(err.Error())
	}

	if err := r.db.Delete(&models.ProductTypeEntity{}, id).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func (r *productTypeRepositoryDB) GetCount() (int64, error) {
	var count int64
	err := r.db.Table("producttype").Count(&count).Error
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	return count, nil
}






