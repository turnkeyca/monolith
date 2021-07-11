package shorturl

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response shortUrlErrorResponse
//lint:ignore U1000 for docs
type shortUrlErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A short url
// swagger:response shortUrlResponse
//lint:ignore U1000 for docs
type shortUrlResponseWrapper struct {
	// A short url
	// in: body
	Body ShortUrlDto
}

// swagger:parameters getShortUrl
//lint:ignore U1000 for docs
type shortUrlParamsWrapper struct {
	// The url to be converted
	// in: query
	// required: true
	Url string `json:"url"`
}
