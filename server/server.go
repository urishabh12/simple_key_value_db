package server

import (
	"context"

	"github.com/urishabh12/simple_key_value_db/db"
	db_proto "github.com/urishabh12/simple_key_value_db/proto"
)

const (
	default_db = "default"
)

type DBServer struct {
	db *db.DB
}

func NewServer() (*DBServer, error) {
	d_db, err := db.NewDB(default_db)
	if err != nil {
		return nil, err
	}

	return &DBServer{
		db: d_db,
	}, nil
}

func (d *DBServer) Put(ctx context.Context, in *db_proto.PutRequest) (*db_proto.PutResponse, error) {
	err := d.db.Put(in.Key, in.Value)
	if err != nil {
		return nil, err
	}

	return &db_proto.PutResponse{
		Status: true,
	}, nil
}

func (d *DBServer) Get(ctx context.Context, in *db_proto.GetRequest) (*db_proto.GetResponse, error) {
	value, err := d.db.Get(in.Key)
	if err != nil {
		return nil, err
	}

	return &db_proto.GetResponse{
		Value: value,
	}, nil
}
