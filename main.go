package main

import (
	"log"
	"net/http"

	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/constants"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/handler"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/movies"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/utils"
)

//Main method used to config and init the service
func main() {

	resource := utils.ResourceManager{}

	movieSearcher := &movies.APIMovieSearcher{
		APIKey: resource.GetProperty(constants.API_KEY_FOR_URL_PROVIDER_INFO_MOVIES),
		URL:    resource.GetProperty(constants.URL_PROVIDER_INFO_MOVIES),
	}

	handler := handler.ConfigRouter(movieSearcher)
	log.Fatal(http.ListenAndServe(":"+resource.GetProperty(constants.WEB_SERVICE_PORT), handler))
}
