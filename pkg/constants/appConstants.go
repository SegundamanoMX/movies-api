package constants

var (
	//Properties keys
	URL_PROVIDER_INFO_MOVIES             = "url.provider.info.movies"
	API_KEY_FOR_URL_PROVIDER_INFO_MOVIES = "api.key.for.url.provider.info.movies"
	WEB_SERVICE_PORT                     = "web.service.port"

	//Request and response for info movies provider
	PARAM_MOVIE_TITLE_TO_SEARCH = "s"
	PARAM_API_KEY               = "apikey"
	PARAM_TYPE                  = "type"
	PARAM_PAGE                  = "page"
	MOVIE_TYPE                  = "movie"
	QUESTION_MARK               = "?"
	TRUE_AS_STR                 = "True"

	//Request and response for endpoints
	WS_PARAM_MOVIE_TITLE_TO_SEARCH       = "q"
	WS_PARAM_MOVIE_TITLE_TO_SEARCH_ERROR = "query parameter 'q' with the title of movie doesn't exist"
	ERROR                                = "error"
	RESULTS                              = "result"

	//Routes of endpoints
	ROUTE_MOVIES        = "/movies"
	ROUTE_MOVIES_SORTED = "/movies-sorted"

	//Commons
	EMPTY                         = ""
	CONTENT_TYPE                  = "Content-Type"
	APPLICATION_JSON_CONTENT_TYPE = "application/json"
)
