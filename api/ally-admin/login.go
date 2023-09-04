package handler

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
	ally "github.com/tubstrr/ally"
	"github.com/tubstrr/ally/environment"
)

func AllyAdminLogin(w http.ResponseWriter, r *http.Request) {
	host := environment.Get_environment_variable("ALLY_REDIS_HOST", "redis")
	port := environment.Get_environment_variable("ALLY_REDIS_PORT", "6379")
	db_name := environment.Get_environment_variable("ALLY_REDIS_DB_NAME", "")
	user := environment.Get_environment_variable("ALLY_REDIS_USERNAME", "default")
	password := environment.Get_environment_variable("ALLY_REDIS_PASSWORD", "")

	redis_url := "redis://" + user + ":" + password + "@" + host + ":" + port
	if (db_name != "") {
		redis_url += "/" + db_name
	}

	// Load client cert
	cert, err := tls.LoadX509KeyPair("redis_user.crt", "redis_user_private.key")
	if err != nil {
			log.Fatal(err)
	}

	// Load CA cert
	caCert, err := os.ReadFile("redis_ca.pem")
	if err != nil {
			log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Username: user,
		Password: password, // no password set
		DB:       0,  // use default DB
		TLSConfig: &tls.Config{
        MinVersion:   tls.VersionTLS12,
        Certificates: []tls.Certificate{cert},
        RootCAs:      caCertPool,
    },
	})

	ctx := context.Background()
	
	val, err := client.Get(ctx, "test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	
	ally.AdminLogin(w, r)
}