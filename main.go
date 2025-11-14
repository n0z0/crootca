package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"PT. Mencari Cinta Sejati CA"},
			Country:       []string{"ID"},
			Province:      []string{"Jakarta"},
			Locality:      []string{"Jakarta"},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour * 10), // 10 years
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		// Self-signed certificate - this makes it a root CA
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}

	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	// Save private key
	keyOut, err := os.OpenFile("root-ca.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open root-ca.key for writing: %v", err)
	}
	defer keyOut.Close()

	keyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(keyOut, keyPEM); err != nil {
		log.Fatalf("Failed to write private key: %v", err)
	}

	// Save certificate
	certOut, err := os.OpenFile("root-ca.crt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open root-ca.crt for writing: %v", err)
	}
	defer certOut.Close()

	certPEM := &pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	if err := pem.Encode(certOut, certPEM); err != nil {
		log.Fatalf("Failed to write certificate: %v", err)
	}

	fmt.Println("✅ Root CA created successfully!")
	fmt.Println("📁 Files created:")
	fmt.Println("   - root-ca.key (private key)")
	fmt.Println("   - root-ca.crt (certificate)")
	fmt.Println("")
	fmt.Println("🔧 Next steps:")
	fmt.Println("   1. Install root-ca.crt to Windows Certificate Store")
	fmt.Println("   2. Use these files in your proxy configuration")
	fmt.Println("   3. Browsers will trust all certificates signed by this CA")
	fmt.Println("")
	fmt.Println("📋 Installation command for Windows:")
	fmt.Println("   certutil -addstore -f \"ROOT\" root-ca.crt")
}
