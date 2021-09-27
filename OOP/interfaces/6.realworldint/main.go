package main

import "net/http"

// we wish to parse the body of an HTTP reuqest into some object data.

type Entity interface {
	UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v Entity) error {
	return v.UnmarshalHTTP(r)
}

// func (u *User) UnmarshalHTTP(r *http.Request) error {

// }
