package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"sitemap/models/entity"
	"time"
)

func NewSQLdbJobRepo(Conn *sqlx.DB) *DbJobRepo {
	return &DbJobRepo{
		Conn: Conn,
	}
}

type DbJobRepo struct {
	Conn *sqlx.DB
}


func (l *DbJobRepo) Count()(int, error){
	t := time.Now()
	query := "SELECT COUNT(1) FROM JOBS"
	var count int
	err := l.Conn.Get(&count, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (l *DbJobRepo)ObjectsForSitemap(page int, limit int)(*[]entity.ResQuery,error){
	t := time.Now()
	offset:= (page - 1) *limit
	query := fmt.Sprintf("select id, updated_at from jobs limit %d OFFSET %d", limit, offset)
	result := &[]entity.ResQuery{}
	err := l.Conn.Select(result, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return nil, err
	}
	return result, nil
}
