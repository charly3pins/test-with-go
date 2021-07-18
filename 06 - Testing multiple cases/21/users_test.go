package psql

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestUserStore(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=jon sslmode=disable dbname=test_user_store")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err = %s", err))
	}
	defer db.Close()
	us := &UserStore{
		sql: db,
	}
	t.Run("Find", testUserStore_Find(us))
	t.Run("Create", testUserStore_Find(us))
	t.Run("Delete", testUserStore_Find(us))
	t.Run("Subscribe", testUserStore_Find(us))
	// teardown
}

func testUserStore_Find(us *UserStore) func(t *testing.T) {
	return func(t *testing.T) {
		jon := &User{
			Name:  "Jon Calhoun",
			Email: "jon@calhoun.io",
		}
		err := us.Create(jon)
		if err != nil {
			t.Errorf("us.Create() err = %s", err)
		}
		defer func() {
			err := us.Delete(jon.ID)
			if err != nil {
				t.Errorf("us.Delete() err = %s", err)
			}
		}()

		tests := []struct {
			name    string
			id      int
			want    *User
			wantErr error
		}{
			{"Found", jon.ID, jon, nil},
			{"Not Found", -1, nil, ErrNotFound},
		}
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				got, err := us.Find(tc.id)
				if err != tc.wantErr {
					t.Errorf("us.Find() err = %s", err)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("us.Find() = %+v, want %+v", got, tc.want)
				}
			})
		}
	}
}
