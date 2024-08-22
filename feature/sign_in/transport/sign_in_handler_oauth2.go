package transport

import (
    "net/http"                // Provides HTTP client and server implementations
    "github.com/gin-gonic/gin" // Gin web framework for handling HTTP requests in Go
    "github.com/markbates/goth" // Goth package for handling multiple authentication providers
    "github.com/markbates/goth/gothic" // Gothic provides simple authentication flow using Goth
    "github.com/markbates/goth/providers/facebook" // Facebook provider for Goth
    "github.com/markbates/goth/providers/google"   // Google provider for Goth
)

// init function is called automatically when the package is initialized.
// It sets up the authentication providers with their respective credentials and callback URLs.
func init() {
    goth.UseProviders(
        google.New("GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "http://localhost:8080/auth/google/callback"), // Google provider configuration
        facebook.New("FACEBOOK_APP_ID", "FACEBOOK_APP_SECRET", "http://localhost:8080/auth/facebook/callback"), // Facebook provider configuration
    )
}

// RegisterAuthRoutes registers the authentication routes with the provided Gin router.
func RegisterAuthRoutes(r *gin.Engine) {
    r.GET("/auth/:provider", authHandler)           // Route to initiate authentication with the chosen provider (e.g., Google, Facebook)
    r.GET("/auth/:provider/callback", authCallback) // Route to handle the callback from the provider after authentication
}

// authHandler initiates the authentication process by redirecting the user to the provider's login page.
func authHandler(c *gin.Context) {
    gothic.BeginAuthHandler(c.Writer, c.Request) // Begins the authentication process with the selected provider
}

// authCallback handles the callback from the authentication provider.
// If authentication is successful, it returns the user information as a JSON response.
func authCallback(c *gin.Context) {
    user, err := gothic.CompleteUserAuth(c.Writer, c.Request) // Completes the authentication and retrieves the user information
    if err != nil {
        c.Redirect(http.StatusTemporaryRedirect, "/") // If an error occurs, redirect the user back to the homepage
        return
    }
    // Handle the user information here (e.g., create a session or generate a JWT token for the user)
    c.JSON(http.StatusOK, gin.H{"user": user}) // Respond with the authenticated user information as JSON
}
