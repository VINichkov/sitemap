package models

import "sitemap/models/entity"

type ObjectRepository interface {
	Count() (int, error)
	ObjectsForSitemap(int, int)(*[]entity.ResQuery, error)
}