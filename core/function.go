package core

import (
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
		switch {
		case pageSize > 50:
			pageSize = 50
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ResponseData(count int, r *http.Request) (int, bool, bool) {
	prev, next := false, false
	var page, pageSize int
	if count != 0 {
		var err error
		if r.URL.Query().Get("page") == "" {
			page = 1
		} else {
			page, err = strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				log.Println(err)
			}
		}
		if r.URL.Query().Get("page_size") == "" {
			pageSize = 10
		} else {
			pageSize, err = strconv.Atoi(r.URL.Query().Get("page_size"))
			if pageSize > 50 {
				pageSize = 50
			}
			if err != nil {
				log.Println(err)
			}
		}

		if (page * pageSize) < count {
			next = true
		}

		if page > 1 {
			prev = true
		}
	}
	return page, prev, next
}
