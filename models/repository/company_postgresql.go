package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"sitemap/models/entity"
	"time"
)

func NewSQLCompanyRepo(Conn *sqlx.DB) *DbCompanyRepo {
	return &DbCompanyRepo{
		Conn: Conn,
	}
}

type DbCompanyRepo struct {
	Conn *sqlx.DB
}

func (l *DbCompanyRepo) Count()(int, error){
	t := time.Now()
	query := "SELECT COUNT(1) FROM COMPANIES where (companies.description is not null) or (companies.site is not null) "
	var count int
	err := l.Conn.Get(&count, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (l *DbCompanyRepo)ObjectsForSitemap(page int, limit int)(*[]entity.ResQuery,error){
	t := time.Now()
	offset:= (page - 1) *limit
	query := fmt.Sprintf("select id, updated_at from companies where (companies.description is not null) or (companies.site is not null) order by updated_at limit %d OFFSET %d", limit, offset )
	result := &[]entity.ResQuery{}
	err := l.Conn.Select(result, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return nil, err
	}
	return result, nil
}
