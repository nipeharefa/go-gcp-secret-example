// Sample quickstart is a basic program that uses Secret Manager.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func main() {

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	nameBuilder := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", os.Getenv("GCP_PROJECT_ID"), os.Getenv("GCP_SECRET_NAME"))
	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: nameBuilder,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	// // Print the secret payload.
	// // WARNING: Do not print the secret in a production environment - this
	// // snippet is showing how to access the secret material.
	log.Printf("Plaintext: %s", result.Payload.Data)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
