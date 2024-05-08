package appcontext

import (
	"context"
	"errors"
)

type key int

const (
	loggedInUserIDKey key = iota
	loggerPtr
	TxProcessKey
)

func SetLoggedInUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, loggedInUserIDKey, userID)
}

func GetLoggedinUserID(ctx context.Context) (string, error) {
	value := ctx.Value(loggedInUserIDKey)
	uid, ok := value.(string)
	if !ok || uid == "" {
		return "", errors.New("not authorized")
	}
	return uid, nil
}
