package sqlite

import (
	"fmt"
	"testing"
	"time"

	"github.com/danmaxdanilov/zts.shared/pkg/logger"
	"gorm.io/gorm"
)

func TestSqlite(t *testing.T) {
	logger := logger.NewAppLogger(&logger.Config{})
	logger.InitLogger()
	cfg := &Config{
		DbPath: "file::memory:?cache=shared",
	}

	tests := []struct {
		input    int
		expected string
		wantErr  bool
	}{
		{input: 1, expected: "inserted records count 0", wantErr: true},
		{input: -1, expected: "inserted records count 1", wantErr: false},
	}

	for _, tt := range tests {
		t.Run("SUCCES. test sqlite db init and create record", func(t *testing.T) {
			context := NewDbContext(logger, *cfg)

			if tt.input == -1 {
				context.db.AutoMigrate(&fakeUser{})
			}

			user := fakeUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

			result := context.db.Create(&user)

			// logger.Debugf("inserted primary key is %v", user.ID)

			if (result.Error != nil) != tt.wantErr {
				t.Errorf("NewDbContext() error = %v, wantErr %v", result.Error, tt.wantErr)
				return
			}

			resultMessage := fmt.Sprintf("inserted records count %v", result.RowsAffected)
			if tt.expected != resultMessage {
				t.Errorf("NewDbContext() = %v, want %v", resultMessage, tt.expected)
			}
		})
	}

	// t.Run("FAIL. test sqlite db init and create record", func(t *testing.T) {
	// 	context := NewDbContext(logger, *cfg)

	// 	// context.db.AutoMigrate(&fakeUser{})

	// 	user := fakeUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// 	result := context.db.Create(&user)

	// 	logger.Debugf("inserted primary key is %v", user.ID)
	// 	if result.Error != nil {
	// 		logger.Fatalf("error - %v", result.Error)
	// 	}
	// 	logger.Debugf("inserted records count %v", result.RowsAffected)
	// })
}

type fakeUser struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}
