package repo

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manc88/demo/internal/user_service/mocks"
)

type repoFixture struct {
	repo        *UserRepository
	mockCache   *mocks.MockICache
	mockStorage *mocks.MockIStorage
	ctrl        *gomock.Controller
	ErrStorage  error
	ErrCache    error
}

func setUP(t *testing.T) *repoFixture {
	var fix repoFixture
	fix.ctrl = gomock.NewController(t)
	fix.mockCache = mocks.NewMockICache(fix.ctrl)
	fix.mockStorage = mocks.NewMockIStorage(fix.ctrl)
	logger := log.New(os.Stdout, "[REPO UNIT TESTING] ", log.Ldate|log.Ltime|log.Lshortfile)
	fix.repo = &UserRepository{
		cache:   fix.mockCache,
		storage: fix.mockStorage,
		logger:  logger,
	}
	fix.ErrStorage = errors.New("storage error")
	fix.ErrCache = errors.New("cache error")

	return &fix

}

func (f *repoFixture) tearDown() {
	f.ctrl.Finish()
}
