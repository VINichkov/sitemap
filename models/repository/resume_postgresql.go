package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"sitemap/models/entity"
	"time"
)

func NewSQLdbResumeRepo(Conn *sqlx.DB) *DbResumeRepo {
	return &DbResumeRepo{
		Conn: Conn,
	}
}

type DbResumeRepo struct {
	Conn *sqlx.DB
}

func (l *DbResumeRepo)Count() (int, error){
	t := time.Now()
	query := "SELECT COUNT(1) FROM RESUMES"
	var id int
	err := l.Conn.Get(&id, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (l *DbResumeRepo)ObjectsForSitemap(page int, limit int)(*[]entity.ResQuery,error){
	t := time.Now()
	offset:= (page - 1) *limit
	query := fmt.Sprintf( "select id, updated_at from resumes limit %d OFFSET %d", limit, offset )
	result := &[]entity.ResQuery{}
	err := l.Conn.Select(result, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return nil, err
	}
	return result, nil
}
