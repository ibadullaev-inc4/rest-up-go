package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()

	config.DatabaseURL = databaseUrl
	store := New(config)
	if err := store.Open(); err != nil {
		t.Fatalf("failed to open store: %v", err)
	}
	return store, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := store.db.Exec(fmt.Sprintf("TRUNCATE %s", strings.Join(tables, ", "))); err != nil {
				t.Fatalf("failed to truncate tables: %v", err)
			}
			store.Close()
		}
	}
}
