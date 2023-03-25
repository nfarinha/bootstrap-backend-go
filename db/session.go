package db

type IDBSession interface {
	NewSession() IDBSession

	Find(dest any, query any) (any, error)
	Get(ptr any) error
	Insert(ptr any) error
	//InnerJoin(ptr Tabler) IDBSession
	Joins(pWhere any) IDBSession
	//Select(query any, args ...any) IDBSession
	Update(ptr any) error
	//Native() any
}

type IDB interface {
	IDBSession
	Sync(model any) error
}
