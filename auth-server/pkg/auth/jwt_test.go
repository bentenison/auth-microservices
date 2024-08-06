package auth

import (
	"testing"
	"time"
)

func TestTokenGeneration(t *testing.T) {
	// config.InitConfig()
	svc := NewAuthService(&AuthConf{
		ExpirationDuration: time.Duration(2 * time.Hour),
		SecretKey:          []byte("mysecret"),
	})
	// for i := 0; i < b.N; i++ {
	token, err := svc.CreateToken("123456789")
	if err != nil {
		t.Errorf("error creating token %v", err)
	}
	_ = token
	// }
}
func BenchmarkTokenGeneration(b *testing.B) {
	svc := NewAuthService(&AuthConf{
		ExpirationDuration: time.Duration(2 * time.Hour),
		SecretKey:          []byte("mysecret"),
	})
	for i := 0; i < b.N; i++ {
		token, err := svc.CreateToken("123456789")
		if err != nil {
			b.Errorf("error creating token %v", err)
		}
		_ = token
	}
}
