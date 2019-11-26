package entity

import (
	"database/sql"
	"fmt"
	"strconv"
)


type ResQuery struct {
	Id sql.NullInt32 `db:"id"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (r *ResQuery)String()string {
	str := ""
	if r.Id.Valid {str += "Id: " +  strconv.Itoa(int(r.Id.Int32)) + "\n" } else {str +="Id: null\n"}
	if r.UpdatedAt.Valid {str += "UpdatedAt: " +  r.UpdatedAt.Time.String() + "\n"} else {str +="UpdatedAt: null\n"}
	return  str
}

func (r *ResQuery)XML(url string) *string{
	if !r.Id.Valid {
		return nil
	}

	if !r.UpdatedAt.Valid {
		return nil
	}

	update := r.UpdatedAt.Time.Format("2006-01-02")

	str := fmt.Sprintf("<url><loc>%s/%d</loc>", url, r.Id.Int32)
	str += fmt.Sprintf("<lastmod>%s</lastmod>", update)
	str +="<changefreq>hourly</changefreq></url>"
	return &str
}