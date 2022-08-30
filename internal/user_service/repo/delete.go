package repo

import "context"

func (u *UserRepository) Delete(ctx context.Context, uid int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.defaultTimeOut)
	defer cancel()
	err := u.storage.Delete(ctx, uid)
	if err != nil {
		return err
	}

	err = u.cache.Reset(ctx)
	if err != nil {
		u.logger.Println(err)
	}

	return nil
}
