package main
import (
    "fmt"
	"errors"
	"strconv"
    "github.com/tecbot/gorocksdb"
)

type Rock struct {
	db *gorocksdb.DB
	ro *gorocksdb.ReadOptions
	wo *gorocksdb.WriteOptions
}

func (r *Rock) Init(id int) (Rock, error) {
    bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
    bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
    opts := gorocksdb.NewDefaultOptions()
    opts.SetBlockBasedTableFactory(bbto)
    opts.SetCreateIfMissing(true)
    db, err := gorocksdb.OpenDb(opts, "/tmp/db" + strconv.FormatInt(int64(id),10) )
    if err != nil {
		fmt.Println("open db error", err)
		return *r, err
	}

	r.ro = gorocksdb.NewDefaultReadOptions()
	r.wo = gorocksdb.NewDefaultWriteOptions()
	r.db = db

	return *r, nil
}

func (r *Rock) Close() {
	r.ro.Destroy()
	r.wo.Destroy()
	r.db.Close()
}

func (r *Rock) Put(key, value []byte) error {
	err := r.db.Put(r.wo, key, value)
	if err != nil {
		fmt.Println("put error", err)
		return err
	}
	return nil
}

func (r *Rock) Get(key []byte) (value []byte, err error) {
	v, err := r.db.Get(r.ro, key)
	if err != nil {
		fmt.Println("get error", err)
		return nil, err
	}

	defer v.Free()
	if v.Exists() {
		fmt.Println("Got value:", string(v.Data()))
		return v.Data(), nil
	} 
	fmt.Println("Not found key: %s", string(key))
	return nil, errors.New("not found")
}

func (r *Rock) Delete(key []byte) error {
	err := r.db.Delete(r.wo, key)
	if err != nil {
		fmt.Println("put error", err)
		return err
	}
	return nil
}
