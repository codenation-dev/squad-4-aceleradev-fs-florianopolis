package postgres

import (
	"testing"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
)

var fakeDB struct {
	// connector:    c,
	// openerCh:     make(chan struct{}, connectionRequestQueueSize),
	// resetterCh:   make(chan *driverConn, 50),
	// lastPut:      make(map[*driverConn]string),
	// connRequests: make(map[uint64]chan connRequest),
	// stop:         cancel,
}

// func Test(){
// 	sql.Open("postgres", )
// }

var dbName = "test_uati"
var tableName = "test_users"

func TestConnect(t *testing.T) {
	fakeDB, err := Connect(dbName)
	if err != nil {
		t.Error(err)
	}

	if fakeDB.Ping() != nil {
		t.Error(err)
	}
}

func TestCreateTableUsers(t *testing.T) {
	// fakeDB, _ := Connect(dbName)
	s := NewStorage(dbName)
	err := s.createUsersTable(tableName)
	if err != nil {
		t.Error(err)
	}
	defer s.dropTable(tableName)
}

var pipa = entity.User{"pipa@email.com", "42"}

// func populateDB(dbName string) {
// 	var err error
// 	assertErr := func(err error){
// 		if err != nil{
// 			log.Fatal(err)
// 		}
// 	}
// 	s := NewStorage(dbName)
// 	err s.createUsersTable(tableName)
// 	s.db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
// 	s.CreateUser(pipa)
// }
