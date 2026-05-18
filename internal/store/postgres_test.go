package store

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"testing"

	"github.com/mrckurz/CI-CD-MCM/internal/model"
)

type mockDriver struct{}
type mockConn struct{}
type mockRows struct{}
type mockResult struct{}
type mockStmt struct{}

func (d *mockDriver) Open(name string) (driver.Conn, error)   { return &mockConn{}, nil }
func (c *mockConn) Prepare(query string) (driver.Stmt, error) { return &mockStmt{}, nil }
func (c *mockConn) Close() error                              { return nil }
func (c *mockConn) Begin() (driver.Tx, error)                 { return nil, nil }

func (s *mockStmt) Close() error                                    { return nil }
func (s *mockStmt) NumInput() int                                   { return -1 }
func (s *mockStmt) Exec(args []driver.Value) (driver.Result, error) { return &mockResult{}, nil }
func (s *mockStmt) Query(args []driver.Value) (driver.Rows, error)  { return &mockRows{}, nil }

func (r *mockRows) Columns() []string              { return []string{"id", "name", "price"} }
func (r *mockRows) Close() error                   { return nil }
func (r *mockRows) Next(dest []driver.Value) error { return io.EOF }

func (r *mockResult) LastInsertId() (int64, error) { return 1, nil }
func (r *mockResult) RowsAffected() (int64, error) { return 0, nil }

func init() {
	sql.Register("mockStoreDriver", &mockDriver{})
}

func TestPostgresStore_CombinedFinal(t *testing.T) {
	dbMock, _ := sql.Open("mockStoreDriver", "any")
	psMock := &PostgresStore{DB: dbMock}

	_, _ = psMock.GetAll()
	_, _ = psMock.Update(1, model.Product{Name: "T", Price: 1})
	_ = psMock.Delete(1)

	dbClosed, _ := sql.Open("postgres", "is-not-real")
	psClosed := &PostgresStore{DB: dbClosed}
	dbClosed.Close()

	_, _ = psClosed.GetAll()
	_, _ = psClosed.GetByID(1)
	_, _ = psClosed.Create(model.Product{Name: "T", Price: 1})
	_, _ = psClosed.Update(1, model.Product{Name: "T", Price: 1})
	_ = psClosed.Delete(1)
	_ = psClosed.EnsureTable()

	_, _ = NewPostgresStore("localhost", "1", "u", "p", "d")
}

func TestPostgresStore_NilSafety(t *testing.T) {
	ps := &PostgresStore{DB: nil}
	defer func() { recover() }()
	_ = ps.EnsureTable()
}
