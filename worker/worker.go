package worker

import (
	"avito/banner"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Worker struct {
	db          *sql.DB
	redisClient *redis.Client
	actualTime  time.Duration
}

// можно поменять
func makeKeyForRedis(featureID, tagID int) string {
	return fmt.Sprint(10000*featureID + tagID)
}

func (w *Worker) GetBanner(ctx context.Context, featureID, tagID int, needMoreActual bool) (banner.Banner, error) {
	if !needMoreActual {
		ban, err := w.redisClient.Get(ctx, makeKeyForRedis(featureID, tagID)).Result()
		// какая - то неожиданная ошибка
		if err != nil && !errors.Is(err, redis.Nil) {
			return banner.Banner{}, err
		}
		if err == nil {
			var b banner.Banner
			err = json.Unmarshal([]byte(ban), &b)
			if err != nil {
				return banner.Banner{}, err
			}
			return b, nil
		}
	}
	// если нужны свежие данные или в кеше нет нужной записи

	var data []byte
	var isActive bool
	err := w.db.QueryRowContext(ctx, fmt.Sprintf("SELECT feature_id, tag_id, data, isActive FROM test WHERE feature_id = $%d AND tag_id = $%d", featureID, tagID), 1).Scan(featureID, tagID, &data, &IsActive)
	if err != nil || isActive == false {
		if errors.Is(err, sql.ErrNoRows) {
			return banner.Banner{}, err
		}
		return banner.Banner{}, err
	}
	ban, err := w.SetBanner(ctx, featureID, tagID, data, isActive)
	if err != nil {
		return ban, err
	}
	return ban, nil
}

func (w *Worker) SetBanner(ctx context.Context, featureID, tagID int, data []byte, isActive bool) (banner.Banner, error) {
	ban := banner.Banner{Id: makeIdForBanner(featureID, []int{tagID}), Data: data, FeatureID: featureID, TagsIDs: []int{tagID}, IsActive: isActive}
	err := w.redisClient.Set(ctx, makeKeyForRedis(featureID, tagID), ban, w.actualTime).Err()
	if err != nil {
		return ban, err
	}
	return ban, nil
}

func makeIdForBanner(featureId int, tagsIDs []int) int {
	return 0
}
