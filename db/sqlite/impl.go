package sqlite

import (
	"reflect"

	"github.com/nfarinha/bootstrap-backend-go/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dbSession struct {
	gorm *gorm.DB
}

type dbConnection struct {
	dbSession
}

func getStructName(beans any) (res string) {
	t := reflect.TypeOf(beans)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func New(path string) (db db.IDB, err error) {
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	db = &dbConnection{dbSession{conn}}
	return db, err
}

func (c *dbSession) NewSession() db.IDBSession {
	return &dbSession{c.gorm.Session(&gorm.Session{})}
}

func (d *dbSession) Find(dest any, query any) (any, error) {
	res := d.gorm.Find(&dest, query)
	return dest, res.Error
}

func (d *dbSession) Get(ptr any) error {
	return d.gorm.Take(ptr, ptr).Error
}

func (d *dbSession) InnerJoin(ptr any) db.IDBSession {
	d.gorm = d.gorm.InnerJoins(getStructName(ptr), d.gorm.Where(ptr))
	return d
}

func (d *dbSession) Insert(ptr any) error {
	return d.gorm.Create(ptr).Error
}

func (d *dbSession) Joins(ptr any) db.IDBSession {
	d.gorm = d.gorm.Joins(getStructName(ptr), d.gorm.Where(ptr))
	return d
}

func (d *dbSession) Update(ptr any) error {
	return d.gorm.Model(ptr).Updates(ptr).Error
}

func (c *dbConnection) Sync(model any) error {
	return c.gorm.AutoMigrate(model)
}
