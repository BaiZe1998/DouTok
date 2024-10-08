package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func (a *Adapter) AddFavorite(ctx context.Context, id, userId int64, target v1.FavoriteTarget, favoriteType v1.FavoriteType) error {
	req := &v1.AddFavoriteRequest{
		Id:     id,
		UserId: userId,
		Target: target,
		Type:   favoriteType,
	}

	resp, err := a.favorite.AddFavorite(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) RemoveFavorite(ctx context.Context, id, userId int64, target v1.FavoriteTarget, favoriteType v1.FavoriteType) error {
	req := &v1.RemoveFavoriteRequest{
		Id:     id,
		UserId: userId,
		Target: target,
		Type:   favoriteType,
	}

	resp, err := a.favorite.RemoveFavorite(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) IsUserFavoriteVideo(ctx context.Context, userId int64, videoIdList []int64) (map[int64]bool, error) {
	var items []*v1.IsFavoriteRequestItem
	for _, id := range videoIdList {
		items = append(items, &v1.IsFavoriteRequestItem{
			BizId:  id,
			UserId: userId,
		})
	}

	req := &v1.IsFavoriteRequest{
		Items: items,
	}
	resp, err := a.favorite.IsFavorite(ctx, req)
	return respcheck.CheckT[map[int64]bool, *v1.Metadata](
		resp, err,
		func() map[int64]bool {
			result := make(map[int64]bool)
			if len(resp.Result) == 0 {
				return result
			}

			for _, item := range resp.Result {
				result[item.BizId] = item.IsFavorite
			}
			return result
		},
	)
}
