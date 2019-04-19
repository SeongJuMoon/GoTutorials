package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type router map[string]interface{}

type service struct {
	r  router
	db *sql.DB
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routerContext := s.r
	dbContext := s.db

	v, ok := routerContext[r.URL.Path]

	if ok == false {
		http.Error(w, "not found", http.StatusNotFound)
	} else {
		v.(func(*sql.DB, http.ResponseWriter, *http.Request))(dbContext, w, r)
	}
}

func welcome(dbService *sql.DB, response http.ResponseWriter, request *http.Request) {

	ctx, cancel := context.WithTimeout(request.Context(), 1*time.Second)
	defer cancel()

	err := dbService.PingContext(ctx)
	if err != nil {
		http.Error(response, fmt.Sprintf("db down: %v", err), http.StatusFailedDependency)
		return
	}

	response.WriteHeader(http.StatusOK)
}

func configuareRouter() router {
	r := make(router)
	route(r, "/", welcome)

	return r
}

func route(r router, urlPath string, m interface{}) {
	r[urlPath] = m
}

func connectRepository() *sql.DB {
	connStr := "postgres://web-user:1q2w3e4r@localhost/web-user?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return db
}

func main() {

	const port = ":10004"

	r := configuareRouter()

	db := connectRepository()

	service := &service{
		r,
		db,
	}

	http.ListenAndServe(port, service)
}
