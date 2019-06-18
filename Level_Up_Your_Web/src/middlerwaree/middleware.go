<<<<<<< HEAD:Level_Up_Your_Web/src/middlewaree/middlewaree.go
package middlewaree
=======
package middlerwaree
>>>>>>> a8c6af148516d9a83e29d78934d9f34a3a1501ae:Level_Up_Your_Web/src/middlerwaree/middleware.go

import "net/http"


type Middleware []http.Handler

//Adds a handler to the middleware
func (m *Middleware) Add(handler http.Handler){
	*m = append(*m, handler)
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// Wrap the supplied ResponseWriter
	mw := NewMiddlewareResponseWriter(w)

	// Loop through all of the registered handlers
	for _, handler := range m{
		// Call the handler with our MiddlewareResponseWriter
		handler.ServeHTTP(mw, r)

		// If there was a write, stop processing
		if mw.written {
			return
		}
	}
	// If no handlers wrote to the response, it’s a 404
	http.NotFound(w, r)
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	written bool
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	return &MiddlewareResponseWriter{
		ResponseWriter: w,
	}
}

func (w * MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}
