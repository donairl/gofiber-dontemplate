package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/donairl/gofiber-dontemplate/lib"
	"github.com/donairl/gofiber-dontemplate/lib/database"
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/donairl/gofiber-dontemplate/routers"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {

	// Initialize a session store
	sessConfig := session.Config{
		Expiration:     60 * time.Minute,        // Expire sessions after 30 minutes of inactivity
		KeyLookup:      "cookie:__Host-session", // Recommended to use the __Host- prefix when serving the app over TLS
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	}
	database.ConnectDb()
	// Migrate the User model
	database.Connection.AutoMigrate(&models.User{})
	lib.Store = session.New(sessConfig)

	app := routers.New()

	certFile := "./cert.pem"
	keyFile := "./key.pem"

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		fmt.Println("Self-signed certificate not found, generating...")
		if err := generateSelfSignedCert(certFile, keyFile); err != nil {
			panic(err)
		}
		fmt.Println("Self-signed certificate generated successfully")
		fmt.Println("You will need to accept the self-signed certificate in your browser")
	}

	// cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// if err != nil {
	// 	panic(err)
	// }

	// config := &tls.Config{Certificates: []tls.Certificate{cert}}

	// ln, err := tls.Listen("tcp", ":4443", config)
	// if err != nil {
	// 	panic(err)
	// }
	// app.Listener(ln)
	app.ListenTLS(":8443", certFile, keyFile)
	//log.Fatal(app.ListenTLS(":443", "./cert.pem", "./cert.key"))

}

// generateSelfSignedCert generates a self-signed certificate and key
// and saves them to the specified files
//
// This is only for testing purposes and should not be used in production
func generateSelfSignedCert(certFile string, keyFile string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	certOut, err := os.Create(certFile)
	if err != nil {
		return err
	}
	defer certOut.Close()

	_ = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer keyOut.Close()

	_ = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return nil
}
