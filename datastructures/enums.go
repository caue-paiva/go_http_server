package datastructs

type HttpMethod string
const (
	GET    HttpMethod = "GET"
	PUT    HttpMethod = "PUT"
	POST   HttpMethod = "POST"
	HEAD   HttpMethod = "HEAD"
	DELETE HttpMethod = "DELETE"
)

type HttpError string
const (
	BadRequest          HttpError = "400 Bad Request"
	Unauthorized        HttpError = "401 Unauthorized"
	Forbidden           HttpError = "403 Forbidden"
	NotFound            HttpError = "404 Not Found"
	InternalServerError HttpError = "500 Internal Server Error"
)

type ContentType string
const (
	TextPlain       ContentType = "text/plain"
	TextHTML        ContentType = "text/html"
	TextCSS         ContentType = "text/css"
	TextJavaScript  ContentType = "text/javascript"
	ApplicationJSON ContentType = "application/json"
	ApplicationXML  ContentType = "application/xml"
	ApplicationForm ContentType = "application/x-www-form-urlencoded"
	MultipartForm   ContentType = "multipart/form-data"
	ImageJPEG       ContentType = "image/jpeg"
	ImagePNG        ContentType = "image/png"
	ImageGIF        ContentType = "image/gif"
	ImageSVG        ContentType = "image/svg+xml"
)