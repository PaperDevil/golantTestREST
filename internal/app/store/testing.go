package store

import (
	"fmt"
	"strings"
	"testing"
)

// Функция имитирующая работу серверного API, создающая конфиги и экземпляр структуры Store
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			_, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
