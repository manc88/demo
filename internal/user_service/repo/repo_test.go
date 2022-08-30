package repo

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manc88/demo/internal/models"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {

	t.Run("positive", func(t *testing.T) {

		f := setUP(t)
		defer f.tearDown()

		expectedUser := &models.User{
			UID:     1,
			Name:    "test_name",
			Email:   "mail@mail.m",
			Age:     11,
			Deleted: false,
		}

		f.mockStorage.EXPECT().Create(gomock.Any(), expectedUser).Return(int64(1), nil).Times(1)
		f.mockCache.EXPECT().Reset(gomock.Any()).Times(1)
		id, err := f.repo.Create(context.Background(), expectedUser)
		require.NoError(t, err)
		require.Equal(t, id, int64(1))

	})

	t.Run("error", func(t *testing.T) {

		f := setUP(t)
		defer f.tearDown()

		expectedUser := &models.User{
			UID:     1,
			Name:    "test_name",
			Email:   "mail@mail.m",
			Age:     11,
			Deleted: false,
		}

		f.mockStorage.EXPECT().Create(gomock.Any(), expectedUser).Return(int64(0), f.ErrStorage).Times(1)
		id, err := f.repo.Create(context.Background(), expectedUser)

		require.ErrorIs(t, f.ErrStorage, err)
		require.Equal(t, id, int64(0))

	})

}

func TestGetAll(t *testing.T) {

	t.Run("positive full path", func(t *testing.T) {

		expectedUsers := []*models.User{
			{
				UID:     1,
				Name:    "test_name",
				Email:   "mail@mail.m",
				Age:     11,
				Deleted: false},
			{
				UID:     2,
				Name:    "test_name2",
				Email:   "mail2@mail.m",
				Age:     22,
				Deleted: false},
		}

		f := setUP(t)
		defer f.tearDown()
		f.mockStorage.EXPECT().GetAll(gomock.Any()).Return(expectedUsers, nil).Times(1)
		f.mockCache.EXPECT().Load(gomock.Any()).Return(nil, f.ErrCache).Times(1)
		f.mockCache.EXPECT().Store(gomock.Any(), gomock.Any())
		res, err := f.repo.GetAll(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedUsers, res)

	})

	t.Run("positive from cache", func(t *testing.T) {

		expectedUsers := []*models.User{
			{
				UID:     1,
				Name:    "test_name",
				Email:   "mail@mail.m",
				Age:     11,
				Deleted: false},
			{
				UID:     2,
				Name:    "test_name2",
				Email:   "mail2@mail.m",
				Age:     22,
				Deleted: false},
		}

		f := setUP(t)
		defer f.tearDown()
		f.mockCache.EXPECT().Load(gomock.Any()).Return(expectedUsers, nil).Times(1)
		res, err := f.repo.GetAll(context.Background())

		require.NoError(t, err)
		require.Equal(t, expectedUsers, res)

	})

	t.Run("error storage", func(t *testing.T) {

		f := setUP(t)
		defer f.tearDown()
		f.mockStorage.EXPECT().GetAll(gomock.Any()).Return(nil, f.ErrStorage).Times(1)
		f.mockCache.EXPECT().Load(gomock.Any()).Return(nil, f.ErrCache).Times(1)

		_, err := f.repo.GetAll(context.Background())

		require.ErrorIs(t, f.ErrStorage, err)

	})

}

func TestDelete(t *testing.T) {

	t.Run("positive", func(t *testing.T) {

		f := setUP(t)
		defer f.tearDown()
		f.mockStorage.EXPECT().Delete(gomock.Any(), int64(1)).Return(nil).Times(1)
		f.mockCache.EXPECT().Reset(gomock.Any()).Times(1)
		err := f.repo.Delete(context.Background(), int64(1))
		require.NoError(t, err)

	})

	t.Run("error", func(t *testing.T) {

		f := setUP(t)
		defer f.tearDown()
		f.mockStorage.EXPECT().Delete(gomock.Any(), int64(1)).Return(f.ErrStorage).Times(1)
		err := f.repo.Delete(context.Background(), int64(1))
		require.ErrorIs(t, f.ErrStorage, err)

	})

}
