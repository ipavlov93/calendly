package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"event-calendar/internal/domain/claims"
	firebaseauth "event-calendar/internal/service/authentication"
	auth "event-calendar/internal/service/authorization"
)

const (
	bearerPrefix        = "Bearer "
	authorizationHeader = "Authorization"

	firebaseClaimsKey = "firebase-claims"
	userClaimsKey     = "user-claims"

	loggerPrefix = "api-middleware"
)

type AuthMiddleware struct {
	firebaseAuthService firebaseauth.FirebaseAuthService
	logger              *log.Logger
	providerKeySetURLs  []string // auth provider key set URLs
}

// NewAuthMiddleware set default logger. Use WithOption() to set custom logger.
func NewAuthMiddleware(
	service firebaseauth.FirebaseAuthService,
	providerKeySetURLs []string,
) AuthMiddleware {
	return AuthMiddleware{
		firebaseAuthService: service,
		providerKeySetURLs:  providerKeySetURLs,
		logger:              log.New(os.Stdout, loggerPrefix, log.LstdFlags|log.Lshortfile),
	}
}

func (m AuthMiddleware) WithOption(logger *log.Logger) {
	if logger != nil {
		m.logger = logger
	}
}

func (m AuthMiddleware) RequireValidIDToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token, err := retrieveBearerToken(r)
		if err != nil {
			http.Error(rw,
				fmt.Sprintf("%s: %v", http.StatusText(http.StatusBadRequest), err),
				http.StatusBadRequest)
			return
		}

		idToken, err := m.firebaseAuthService.VerifyIDToken(token)
		if err != nil {
			http.Error(rw,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		parsedClaims, err := parseIDTokenClaims(idToken.Claims)
		if err != nil {
			m.logger.Printf("parseIDTokenClaims(): parse ID token claims error %s", err)
			http.Error(rw,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		// set claims to context
		ctx := context.WithValue(r.Context(), firebaseClaimsKey, parsedClaims)

		// proceed with the request handling
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (m AuthMiddleware) RequireValidAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		accessToken, err := retrieveBearerToken(r)
		if err != nil {
			http.Error(rw,
				fmt.Sprintf("%s: %v", http.StatusText(http.StatusUnauthorized), err),
				http.StatusUnauthorized)
			return
		}

		// Initialize JWK Set client with your JWK Set URLs
		jwks, err := auth.InitializeJWKSetClient(m.providerKeySetURLs)
		if err != nil {
			m.logger.Printf("InitializeJWKSetClient(): %s", err)
			http.Error(rw,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		// Verify the access token
		userClaims, err := auth.VerifyAccessToken(jwks, accessToken)
		if err != nil {
			http.Error(rw,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		r.WithContext(context.WithValue(r.Context(), userClaimsKey, userClaims))

		// Proceed with the request handling
		next.ServeHTTP(rw, r)
	})
}

func retrieveBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get(authorizationHeader)
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", fmt.Errorf("missing or invalid authorization header")
	}
	return strings.TrimPrefix(authHeader, bearerPrefix), nil
}

func parseIDTokenClaims(claimsMap map[string]any) (parsedClaims *claims.FirebaseAuthClaims, err error) {
	claimsJSON, err := json.Marshal(claimsMap)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal claims: %w", err)
	}

	parsedClaims = &claims.FirebaseAuthClaims{}
	err = json.Unmarshal(claimsJSON, parsedClaims)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal claims: %w", err)
	}
	return parsedClaims, nil
}
