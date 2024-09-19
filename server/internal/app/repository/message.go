package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type IMessageRepository interface {
	AddMessage(ctx context.Context, timestamp int64, message string) error
	GetMessagesByTimeRange(ctx context.Context, currentTimestamp, timeRange int64) (map[string]int, error)
}

type MessageRepository struct {
	rc *redis.Client
}

func NewMessageRepository(rc *redis.Client) IMessageRepository {
	return &MessageRepository{rc: rc}
}

const sortedSetKey = "recent_timestamps"

func (mr *MessageRepository) AddMessage(ctx context.Context, timestamp int64, message string) error {
	hashKey := fmt.Sprintf("timestamp:%d", timestamp)

	if err := mr.rc.HIncrBy(ctx, hashKey, message, 1).Err(); err != nil {
		return err
	}

	if err := mr.rc.ZAdd(ctx, sortedSetKey, redis.Z{
		Score:  float64(timestamp),
		Member: hashKey,
	}).Err(); err != nil {
		return err
	}

	if err := mr.rc.Expire(ctx, hashKey, 1*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func (mr *MessageRepository) GetMessagesByTimeRange(ctx context.Context, currentTimestamp, timeRange int64) (map[string]int, error) {
	timestamps, err := mr.rc.ZRangeByScore(ctx, sortedSetKey, &redis.ZRangeBy{
		Min: strconv.FormatInt(currentTimestamp-timeRange, 10),
		Max: strconv.FormatInt(currentTimestamp, 10),
	}).Result()
	if err != nil {
		return nil, err
	}
	aggregatedData := make(map[string]int)
	for _, timestamp := range timestamps {
		messages, err := mr.rc.HGetAll(ctx, timestamp).Result()
		if err != nil {
			return nil, err
		}
		for message, count := range messages {
			count, err := strconv.Atoi(count)
			if err != nil {
				return nil, err
			}
			aggregatedData[message] += count
		}
	}
	return aggregatedData, nil
}
