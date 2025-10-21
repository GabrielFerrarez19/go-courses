package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type repository struct {
	rdb *redis.Client
}

type Repository interface {
	GetFullURL(ctx context.Context, code string) (string, error)
	SaveShortnedURL(ctx context.Context, _url string) (string, error)
}

func NewRepository(rdb *redis.Client) Repository {
	return repository{rdb}
}

func (r repository) SaveShortnedURL(ctx context.Context, _url string) (string, error) {
	var code string
	for range 5 {
		code = genCode()
		err := r.rdb.HGet(ctx, "encurtador", code).Err()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}
			return "", fmt.Errorf("failed to get code from encurtador hashmap: %w", err)
		}

	}

	if err := r.rdb.HSet(ctx, "encurtador", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code in encurtador hashmap: %w", err)
	}
	return code, nil
}

func (r repository) GetFullURL(ctx context.Context, code string) (string, error) {
	fullURL, err := r.rdb.HGet(ctx, "encurtador", code).Result()
	fmt.Println(fullURL)
	if err != nil {
		return "", fmt.Errorf("failed to get code from encurtador hashmap: %w", err)
	}
	return fullURL, nil
}
