package main

import (
	"fmt"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)

	})
}


func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}


// secureHeaders es un middleware que agrega encabezados de seguridad a las respuestas HTTP,
// protegiendo contra ataques como inyección de contenido, clickjacking y robo de información.
//
// Los encabezados establecidos son:
// - Content-Security-Policy: Define fuentes de contenido seguras.
// - Referrer-Policy: Controla la información enviada en el encabezado Referer.
// - X-Content-Type-Options: Previene que el navegador adivine el tipo de contenido.
// - X-Frame-Options: Evita que la página se cargue en un iframe de otros sitios.
// - X-XSS-Protection: Desactiva la protección XSS en los navegadores.
//
// Después de configurar estos encabezados, la solicitud se pasa al siguiente handler.


// logRequest es un middleware que registra información sobre cada solicitud HTTP,
// incluyendo la dirección remota del cliente, el protocolo, el método y la URI solicitada.
// Después de registrar los detalles, pasa la solicitud al siguiente handler en la cadena.
// 
// INFO    2025/01/08 10:28:09 [::1]:59838 - HTTP/1.1 GET /


// recoverPanic es un middleware que maneja pánicos en las solicitudes HTTP.
// Si ocurre un pánico durante la ejecución de la solicitud, lo captura,
// cierra la conexión y responde con un error 500 utilizando el método serverError.
// Después de manejar cualquier pánico, pasa la solicitud al siguiente handler.
