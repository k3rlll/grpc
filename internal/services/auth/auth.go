package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"go.mod/internal/domain/models"
	"go.mod/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash string,
	) (uid int64, err error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New creates new Auth service
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:         log,
		usrSaver:    userSaver,
		usrProvider: userProvider,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

// Login authenticates user and returns JWT token
func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appID int,
) (string, error) {

	const op = "auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)
	log.Info("login attempt")

	user, err := a.usrProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", slog.String("error", err.Error()))
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
	}

}

func (a *Auth) ReisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {

	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)
	log.Info("registering new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to hash password", slog.String("error", err.Error()))
		return 0, err
	}

	id, err := a.usrSaver.SaveUser(ctx, email, string(passHash))
	if err != nil {
		log.Error("failed to save user", slog.String("error", err.Error()))
		return 0, err
	}
	log.Info("user registered", slog.Int64("user_id", id))
	return id, nil

}

func (a *Auth) isAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("implement me")
}
