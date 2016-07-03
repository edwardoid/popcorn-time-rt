package www

type Context struct {
	URL string
	Parameters map[string]string
	Headers map[string]string
	Cookies map[string]string
	Data map[string]string
	Method string
}