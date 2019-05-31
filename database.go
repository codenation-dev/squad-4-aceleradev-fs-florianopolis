package main

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

var dbPath = "db/uati.db"

var sal = "SALARY"
var bank = "BANKFUNC"
var public = "PUBLIC"
var users = "USERS"
var warn = "WARNINGS"

type accountHolder struct {
	Name         string `json:"name"`
	IsBankFunc   string `json:"isBankFunc"`   // 1 = true, 0 = false
	IsPublicFunc string `json:"isPublicFunc"` // 1 = true, 0 = false
	Salary       string `json:"salary"`
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type warning struct {
	Datetime    time.Time `json:"datetime"`
	FromAccount string    `json:"fromAccount"`
	ToUser      string    `json:"toUser"`
	Message     string    `json:"message"`
}

func OpenRead() (*bolt.DB, error) {
	return openDB(true)
}

func OpenWrite() (*bolt.DB, error) {
	return openDB(false)
}

func openDB(readOnly bool) (*bolt.DB, error) {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{ReadOnly: readOnly, Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not read db file: %v", err)
	}
	return db, err
}

func setupDB() error {
	db, err := OpenWrite()
	if err != nil {
		return fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(sal))
		if err != nil {
			return fmt.Errorf("could not create salary bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(bank))
		if err != nil {
			return fmt.Errorf("could not create bank func bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(public))
		if err != nil {
			return fmt.Errorf("could not create public func bucket: %v", err)
		}

		_, err = root.CreateBucketIfNotExists([]byte(users))
		if err != nil {
			return fmt.Errorf("could not create users bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(warn))
		if err != nil {
			return fmt.Errorf("could not create warnings bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB configurado.")
	return nil
}

func addUser(u user) error {
	db, err := OpenWrite()
	if err != nil {
		return fmt.Errorf("could not open DB: %v", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(users)).Put(
			[]byte(u.Email),
			[]byte(u.Password),
		)
		if err != nil {
			return fmt.Errorf("Could not insert user: %v", err)
		}
		return nil
	})
	fmt.Printf("\n%v cadastrado com sucesso.", u.Email)
	return nil
}

func addAccountHolder(ah accountHolder) error {
	db, err := OpenWrite()
	if err != nil {
		return fmt.Errorf("could not open DB: %v", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(sal)).Put(
			[]byte(ah.Name),
			[]byte(ah.Salary),
		)
		if err != nil {
			return fmt.Errorf("could not insert salary: %v", err)
		}
		err = tx.Bucket([]byte("DB")).Bucket([]byte(bank)).Put(
			[]byte(ah.Name),
			[]byte(ah.IsBankFunc),
		)
		if err != nil {
			return fmt.Errorf("could not insert bank func: %v", err)
		}
		err = tx.Bucket([]byte("DB")).Bucket([]byte(public)).Put(
			[]byte(ah.Name),
			[]byte(ah.IsPublicFunc),
		)
		if err != nil {
			return fmt.Errorf("could not insert public func: %v", err)
		}
		return nil
	})
	fmt.Printf("%v cadastrado com sucesso.", ah.Name)
	return nil
}

func readUsers() error {
	db, err := OpenRead()
	if err != nil {
		return err
	}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(users))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	return nil
}

func readAccountHolders() error {
	db, err := OpenRead()
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(sal))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not read account holders bucket: %v", err)
	}
	return nil
}
