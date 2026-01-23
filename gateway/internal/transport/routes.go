package transport

import (
	"net/http"
	"strings"

	"gateway/internal/middleware"
	"gateway/internal/proxy"

	"github.com/gin-gonic/gin"
)

type Config struct {
	JWTSecret             string
	UserServiceURL        string
	VenueServiceURL       string
	ReservationServiceURL string
	PaymentServiceURL     string
}

func RegisterRoutes(r *gin.Engine, cfg Config) error {
	userUpstream, err := proxy.NewUpstream(cfg.UserServiceURL, "/api", nil)
	if err != nil {
		return err
	}
	venueUpstream, err := proxy.NewUpstream(cfg.VenueServiceURL, "/api", nil)
	if err != nil {
		return err
	}
	reservationUpstream, err := proxy.NewUpstream(cfg.ReservationServiceURL, "/api", nil)
	if err != nil {
		return err
	}
	paymentUpstream, err := proxy.NewUpstream(cfg.PaymentServiceURL, "/api", nil)
	if err != nil {
		return err
	}

	api := r.Group("/api")
	api.Use(middleware.AuthUnless(cfg.JWTSecret, isPublicRequest))

	api.Any("/auth/*path", gin.WrapH(http.HandlerFunc(userUpstream.ServeHTTP)))
	api.Any("/users/*path", gin.WrapH(http.HandlerFunc(userUpstream.ServeHTTP)))

	// Venue routes - сначала специфичные для reservation, потом общие
	api.Any("/venues", gin.WrapH(http.HandlerFunc(venueUpstream.ServeHTTP)))
	
	// Создаем специальный handler для venue маршрутов, который определяет upstream по пути
	venueHandler := func(c *gin.Context) {
		path := c.Request.URL.Path
		// Если путь заканчивается на /availability или /bookings, используем reservation service
		if strings.HasSuffix(path, "/availability") || strings.HasSuffix(path, "/bookings") {
			reservationUpstream.ServeHTTP(c.Writer, c.Request)
		} else {
			// Иначе используем venue service
			venueUpstream.ServeHTTP(c.Writer, c.Request)
		}
	}
	api.Any("/venues/*path", venueHandler)
	api.Any("/venue-types", gin.WrapH(http.HandlerFunc(venueUpstream.ServeHTTP)))
	api.Any("/venue-types/*path", gin.WrapH(http.HandlerFunc(venueUpstream.ServeHTTP)))

	// Bookings routes - используем handler для определения upstream
	aggregator := NewAggregator(cfg)
	
	bookingsHandler := func(c *gin.Context) {
		path := c.Request.URL.Path
		
		// Если это GET запрос к /summary, обрабатываем через aggregator
		if c.Request.Method == http.MethodGet && strings.HasSuffix(path, "/summary") {
			aggregator.GetBookingSummary(c)
			return
		}
		
		// Если путь заканчивается на /payment, используем payment service
		if strings.HasSuffix(path, "/payment") {
			paymentUpstream.ServeHTTP(c.Writer, c.Request)
		} else {
			// Иначе используем reservation service
			reservationUpstream.ServeHTTP(c.Writer, c.Request)
		}
	}
	api.Any("/bookings", gin.WrapH(http.HandlerFunc(reservationUpstream.ServeHTTP)))
	api.Any("/bookings/*path", bookingsHandler)

	api.Any("/payments", gin.WrapH(http.HandlerFunc(paymentUpstream.ServeHTTP)))
	api.Any("/payments/*path", gin.WrapH(http.HandlerFunc(paymentUpstream.ServeHTTP)))

	return nil
}

func isPublicRequest(r *http.Request) bool {
	path := r.URL.Path
	method := r.Method

	if strings.HasPrefix(path, "/api/auth/") {
		return true
	}

	if method != http.MethodGet {
		return false
	}

	if path == "/api/venue-types" || strings.HasPrefix(path, "/api/venue-types/") {
		return true
	}

	if path == "/api/venues" || strings.HasPrefix(path, "/api/venues/") {
		if strings.HasSuffix(path, "/bookings") {
			return false
		}
		return true
	}

	return false
}

 
