/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: HttpStatus
 Date: 7/28/22 6:28 AM
*/
package common

type HttpStatus struct {
	Code int
	Msg  string
	Data interface{}
}

var (
	CONTINUE                        = &HttpStatus{Code: 100, Msg: "Continue"}
	SWITCHING_PROTOCOLS             = &HttpStatus{Code: 101, Msg: "Switching Protocols"}
	PROCESSING                      = &HttpStatus{Code: 102, Msg: "Processing"}
	CHECKPOINT                      = &HttpStatus{Code: 103, Msg: "Checkpoint"}
	OK                              = &HttpStatus{Code: 200, Msg: "OK"}
	CREATED                         = &HttpStatus{Code: 201, Msg: "Created"}
	ACCEPTED                        = &HttpStatus{Code: 202, Msg: "Accepted"}
	NON_AUTHORITATIVE_INFORMATION   = &HttpStatus{Code: 203, Msg: "Non-Authoritative Information"}
	NO_CONTENT                      = &HttpStatus{Code: 204, Msg: "No Content"}
	RESET_CONTENT                   = &HttpStatus{Code: 205, Msg: "Reset Content"}
	PARTIAL_CONTENT                 = &HttpStatus{Code: 206, Msg: "Partial Content"}
	MULTI_STATUS                    = &HttpStatus{Code: 207, Msg: "Multi-Status"}
	ALREADY_REPORTED                = &HttpStatus{Code: 208, Msg: "Already Reported"}
	IM_USED                         = &HttpStatus{Code: 226, Msg: "IM Used"}
	MULTIPLE_CHOICES                = &HttpStatus{Code: 300, Msg: "Multiple Choices"}
	MOVED_PERMANENTLY               = &HttpStatus{Code: 301, Msg: "Moved Permanently"}
	FOUND                           = &HttpStatus{Code: 302, Msg: "Found"}
	SEE_OTHER                       = &HttpStatus{Code: 303, Msg: "See Other"}
	NOT_MODIFIED                    = &HttpStatus{Code: 304, Msg: "Not Modified"}
	USE_PROXY                       = &HttpStatus{Code: 305, Msg: "Use Proxy"}
	TEMPORARY_REDIRECT              = &HttpStatus{Code: 307, Msg: "Temporary Redirect"}
	PERMANENT_REDIRECT              = &HttpStatus{Code: 308, Msg: "Permanent Redirect"}
	BAD_REQUEST                     = &HttpStatus{Code: 400, Msg: "Bad Request"}
	UNAUTHORIZED                    = &HttpStatus{Code: 401, Msg: "Unauthorized"}
	PAYMENT_REQUIRED                = &HttpStatus{Code: 402, Msg: "Payment Required"}
	FORBIDDEN                       = &HttpStatus{Code: 403, Msg: "Forbidden"}
	NOT_FOUND                       = &HttpStatus{Code: 404, Msg: "Not Found"}
	METHOD_NOT_ALLOWED              = &HttpStatus{Code: 405, Msg: "Method Not Allowed"}
	NOT_ACCEPTABLE                  = &HttpStatus{Code: 406, Msg: "Not Acceptable"}
	PROXY_AUTHENTICATION_REQUIRED   = &HttpStatus{Code: 407, Msg: "Proxy Authentication Required"}
	REQUEST_TIMEOUT                 = &HttpStatus{Code: 408, Msg: "Request Timeout"}
	CONFLICT                        = &HttpStatus{Code: 409, Msg: "Conflict"}
	GONE                            = &HttpStatus{Code: 410, Msg: "Gone"}
	LENGTH_REQUIRED                 = &HttpStatus{Code: 411, Msg: "Length Required"}
	PRECONDITION_FAILED             = &HttpStatus{Code: 412, Msg: "Precondition Failed"}
	PAYLOAD_TOO_LARGE               = &HttpStatus{Code: 413, Msg: "Payload Too Large"}
	URI_TOO_LONG                    = &HttpStatus{Code: 414, Msg: "URI Too Long"}
	UNSUPPORTED_MEDIA_TYPE          = &HttpStatus{Code: 415, Msg: "Unsupported Media Type"}
	REQUESTED_RANGE_NOT_SATISFIABLE = &HttpStatus{Code: 416, Msg: "Requested range not satisfiable"}
	EXPECTATION_FAILED              = &HttpStatus{Code: 417, Msg: "Expectation Failed"}
	I_AM_A_TEAPOT                   = &HttpStatus{Code: 418, Msg: "I'm a teapot"}
	INSUFFICIENT_SPACE_ON_RESOURCE  = &HttpStatus{Code: 419, Msg: "Insufficient Space On Resource"}
	METHOD_FAILURE                  = &HttpStatus{Code: 420, Msg: "Method Failure"}
	DESTINATION_LOCKED              = &HttpStatus{Code: 421, Msg: "Destination Locked"}
	UNPROCESSABLE_ENTITY            = &HttpStatus{Code: 422, Msg: "Unprocessable Entity"}
	LOCKED                          = &HttpStatus{Code: 423, Msg: "Locked"}
	FAILED_DEPENDENCY               = &HttpStatus{Code: 424, Msg: "Failed Dependency"}
	TOO_EARLY                       = &HttpStatus{Code: 425, Msg: "Too Early"}
	UPGRADE_REQUIRED                = &HttpStatus{Code: 426, Msg: "Upgrade Required"}
	PRECONDITION_REQUIRED           = &HttpStatus{Code: 428, Msg: "Precondition Required"}
	TOO_MANY_REQUESTS               = &HttpStatus{Code: 429, Msg: "Too Many Requests"}
	REQUEST_HEADER_FIELDS_TOO_LARGE = &HttpStatus{Code: 431, Msg: "Request Header Fields Too Large"}
	UNAVAILABLE_FOR_LEGAL_REASONS   = &HttpStatus{Code: 451, Msg: "Unavailable For Legal Reasons"}
	INTERNAL_SERVER_ERROR           = &HttpStatus{Code: 500, Msg: "Internal Server Error"}
	NOT_IMPLEMENTED                 = &HttpStatus{Code: 501, Msg: "Not Implemented"}
	BAD_GATEWAY                     = &HttpStatus{Code: 502, Msg: "Bad Gateway"}
	SERVICE_UNAVAILABLE             = &HttpStatus{Code: 503, Msg: "Service Unavailable"}
	GATEWAY_TIMEOUT                 = &HttpStatus{Code: 504, Msg: "Gateway Timeout"}
	HTTP_VERSION_NOT_SUPPORTED      = &HttpStatus{Code: 505, Msg: "HTTP Version not supported"}
	VARIANT_ALSO_NEGOTIATES         = &HttpStatus{Code: 506, Msg: "Variant Also Negotiates"}
	INSUFFICIENT_STORAGE            = &HttpStatus{Code: 507, Msg: "Insufficient Storage"}
	LOOP_DETECTED                   = &HttpStatus{Code: 508, Msg: "Loop Detected"}
	BANDWIDTH_LIMIT_EXCEEDED        = &HttpStatus{Code: 509, Msg: "Bandwidth Limit Exceeded"}
	NOT_EXTENDED                    = &HttpStatus{Code: 510, Msg: "Not Extended"}
	NETWORK_AUTHENTICATION_REQUIRED = &HttpStatus{Code: 511, Msg: "Network Authentication Required"}
)
