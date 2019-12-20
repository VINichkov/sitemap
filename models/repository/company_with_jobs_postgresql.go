package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"sitemap/models/entity"
	"time"
)

type DbCompanyWithJobsRepo struct {
	Conn *sqlx.DB
}


func NewSQLCompanyWithJobsRepo(Conn *sqlx.DB) *DbCompanyWithJobsRepo{
	return &DbCompanyWithJobsRepo{
		Conn: Conn,
	}
}

func (l *DbCompanyWithJobsRepo) Count()(int, error){
	t := time.Now()
	query := "select count(distinct(jobs.company_id)) from jobs"
	var count int
	err := l.Conn.Get(&count, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (l *DbCompanyWithJobsRepo)ObjectsForSitemap(page int, limit int)(*[]entity.ResQuery,error){
	t := time.Now()
	offset:= (page - 1) *limit
	query := fmt.Sprintf("select company_id as id ,max(updated_at) as updated_at from jobs group by company_id limit %d  OFFSET %d", limit, offset )
	result := &[]entity.ResQuery{}
	err := l.Conn.Select(result, query)
	log.Debug().Msg(fmt.Sprintf("query: %s , %d ms",query, time.Now().Sub(t).Milliseconds()))
	if err != nil {
		return nil, err
	}
	return result, nil
}


