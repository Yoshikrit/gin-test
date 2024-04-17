package repositories_test

import (
	"testing"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gin-test/models"
	"gin-test/repositories"
	"gin-test/utils/errs"
)

const (
	StubDBConnectionErrMsg = "an error '%s' was not expected when opening a stub database connection"
	GormDBErrMsg           = "an error '%s' was not expected when opening a gorm database"
)


func TestCreate(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(StubDBConnectionErrMsg, err)
	}
	defer mockDb.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb}), &gorm.Config{})
	if err != nil {
		t.Fatalf(GormDBErrMsg, err)
	}

	prodTypeCreateMock := &models.ProductTypeCreate{
		Id:   1,
		Name: "A",
	}

	t.Run("test case : fail exist producttype", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)
		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A")

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
			WithArgs("A", 1).
        	WillReturnRows(rows)

		_, err := repo.Create(prodTypeCreateMock)

		expectedRes := errs.NewConflictError("Product Type with the same name already exists")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})

	t.Run("test case : create pass", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)
		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A")

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "producttype"`).
			WithArgs("A", 1).
			WillReturnRows(rows)
		mock.ExpectCommit()

		result, err := repo.Create(prodTypeCreateMock)

		expectedRes := &models.ProductTypeEntity{Id: 1, Name: "A"}
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, result)
	})

	t.Run("test case : create fail create", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "producttype"`).
        	WithArgs("A", 1).
        	WillReturnError(errs.NewUnexpectedError(""))
    	mock.ExpectRollback()

		_, err := repo.Create(prodTypeCreateMock)

		expectedRes := errs.NewUnexpectedError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})
}

func TestGetAll(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(StubDBConnectionErrMsg, err)
	}
	defer mockDb.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb}), &gorm.Config{})
	if err != nil {
		t.Fatalf(GormDBErrMsg, err)
	}

	entityRes := []models.ProductTypeEntity{
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "B",
		},
	}
	t.Run("test case : get all pass", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A").AddRow(2, "B")
		mock.ExpectQuery(`SELECT \* FROM "producttype"`).
			WillReturnRows(rows)

		result, err := repo.GetAll()

		expectedRes := entityRes
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, result)
	})
	t.Run("test case : get all fail", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectQuery(`SELECT \* FROM "producttype"`).
			WillReturnError(errs.NewUnexpectedError(""))

		_, err := repo.GetAll()

		expectedRes := errs.NewUnexpectedError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})
}

func TestGetById(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(StubDBConnectionErrMsg, err)
	}
	defer mockDb.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb}), &gorm.Config{})
	if err != nil {
		t.Fatalf(GormDBErrMsg, err)
	}

	entityRes := &models.ProductTypeEntity{
		Id:   1,
		Name: "A",
	}
	t.Run("test case : get pass", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)
		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A")

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
			WillReturnRows(rows)

		result, err := repo.GetById(1)

		expectedRes := entityRes
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, result)
	})
	t.Run("test case : get fail gorm not found", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.GetById(1)

		expectedRes := errs.NewNotFoundError("record not found")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})

	t.Run("test case : get fail get id", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
			WillReturnError(errs.NewUnexpectedError(""))

		_, err := repo.GetById(1)

		expectedRes := errs.NewUnexpectedError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})
}

func TestUpdate(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(StubDBConnectionErrMsg, err)
	}
	defer mockDb.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb}), &gorm.Config{})
	if err != nil {
		t.Fatalf(GormDBErrMsg, err)
	}

	prodTypeUpdateMock := &models.ProductTypeUpdate{
		Name: "A",
	}

	t.Run("test case : fail id not found", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE `).
        	WillReturnError(errs.NewNotFoundError(""))

		_, err := repo.Update(1, prodTypeUpdateMock)

		expectedRes := errs.NewNotFoundError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})

	t.Run("test case : update pass", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)
		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A")

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
        	WillReturnRows(rows)
		
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WithArgs("A", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		result, err := repo.Update(1, prodTypeUpdateMock)

		expectedRes := &models.ProductTypeEntity{Id: 1, Name: "A"}
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, result)
	})
	
	t.Run("test case : update fail update", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)
		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A")

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
        	WillReturnRows(rows)
		
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE`).
			WithArgs("A", 1).
			WillReturnError(errs.NewUnexpectedError(""))
		mock.ExpectRollback()

		_, err := repo.Update(1, prodTypeUpdateMock)

		expectedRes := errs.NewUnexpectedError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})
}

func TestDelete(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(StubDBConnectionErrMsg, err)
	}
	defer mockDb.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb}), &gorm.Config{})
	if err != nil {
		t.Fatalf(GormDBErrMsg, err)
	}

	t.Run("test case : delete fail no id", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectQuery(`SELECT \* FROM "producttype"`).
			WillReturnError(errs.NewNotFoundError(""))

		err := repo.DeleteById(1)

		expectedRes := errs.NewNotFoundError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})

	t.Run("test case : delete pass", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)
		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A")

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
        	WillReturnRows(rows)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE").
    		WithArgs(1).
    		WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.DeleteById(1)

		assert.NoError(t, err)
	})
	t.Run("test case : delete fail cant delete", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)
		rows := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(1, "A")

		mock.ExpectQuery(`SELECT \* FROM "producttype" WHERE`).
        	WillReturnRows(rows)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE").
    		WithArgs(1).
    		WillReturnError(errs.NewUnexpectedError(""))
		mock.ExpectRollback()

		err := repo.DeleteById(1)

		expectedRes := errs.NewUnexpectedError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})
}

func TestGetCount(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(StubDBConnectionErrMsg, err)
	}
	defer mockDb.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb}), &gorm.Config{})
	if err != nil {
		t.Fatalf(GormDBErrMsg, err)
	}

	t.Run("test case : get count pass", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "producttype"`)).
      		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		result, err := repo.GetCount()

		expectedRes := int64(1)
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, result)
	})
	t.Run("test case : get count fail", func(t *testing.T) {
		repo := repositories.NewProductTypeRepositoryDB(db)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "producttype"`)).
			WillReturnError(errs.NewUnexpectedError(""))

		_, err := repo.GetCount()

		expectedRes := errs.NewUnexpectedError("")
		assert.Error(t, err)
		assert.Equal(t, expectedRes, err)
	})
}