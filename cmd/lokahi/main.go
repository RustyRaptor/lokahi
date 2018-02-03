package main

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/Xe/ln"
	"github.com/Xe/lokahi/internal/lokahiserver"
	"github.com/Xe/lokahi/rpc/lokahi"
	"github.com/caarlos0/env"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/heroku/x/scrub"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	NoPass      bool   `env:"NO_PASS" envDefault:"false"`
	Port        string `env:"PORT" envDefault:"5000"`
}

func (c config) F() ln.F {
	result := ln.F{
		"env_NO_PASS": c.NoPass,
		"env_PORT":    c.Port,
	}

	u, err := url.Parse(c.DatabaseURL)
	if err != nil {
		result["env_DATABASE_URL_err"] = err
	} else {
		result["env_DATABASE_URL"] = scrub.URL(u)
	}

	return result
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = ln.WithF(ctx, ln.F{"in": "main"})

	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		ln.FatalErr(ctx, err)
	}

	db, err := gorm.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		ln.FatalErr(ctx, err)
	}

	cks := &lokahiserver.Checks{
		DB: db,
	}
	mux := http.NewServeMux()

	mux.Handle(lokahi.ChecksPathPrefix, lokahi.NewChecksServer(cks, makeLnHooks()))

	ln.Log(ctx, ln.F{"port": os.Getenv("PORT")}, ln.Action("Listening on http"))
	ln.FatalErr(ctx, http.ListenAndServe(":"+cfg.Port, metaInfo(mux)), ln.Action("http server stopped for some reason"))
}

func metaInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		f := ln.F{
			"remote_ip":       host,
			"x_forwarded_for": r.Header.Get("X-Forwarded-For"),
		}
		ctx := ln.WithF(r.Context(), f)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}