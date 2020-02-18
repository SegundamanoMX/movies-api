package handler

import (
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
	"net/http"
	"reflect"
	"testing"
)


func Test_createSearchMoviesHandler_(t *testing.T) {
	type args struct {
		s movies.MovieSearcher
	}
	var tests []struct {
		name string
		args args
		want func(http.ResponseWriter, *http.Request)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSearchMoviesHandler(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSearchMoviesHandler()")
			}
		})
	}
}

func Test_createSearchMoviesSortedHandler(t *testing.T) {
	type args struct {
		s movies.MovieSearcher
	}
	var tests  []struct {
		name string
		args args
		want func(http.ResponseWriter, *http.Request)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSearchMoviesSortedHandler(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSearchMoviesSortedHandler")
			}
		})
	}
}