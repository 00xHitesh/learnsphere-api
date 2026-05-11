package handler

import (
	"context"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

type authService interface {
	Login(ctx context.Context, email, password string) (*service.UserResponse, *service.TokenPair, error)
	Register(ctx context.Context, name, email, password string) (*service.UserResponse, *service.TokenPair, error)
	GetUser(ctx context.Context, userID string) (*service.UserResponse, error)
	UpdateInterests(ctx context.Context, userID string, interests []string) (*service.UserResponse, error)
	UpdateLocale(ctx context.Context, userID, regionCode, currencyCode string) (*service.UserResponse, error)
	RequestReset(ctx context.Context, email string) (string, error)
	ResetPassword(ctx context.Context, token, password string) error
	RefreshToken(ctx context.Context, refreshToken string) (*service.TokenPair, error)
}

type courseService interface {
	ListPublished(ctx context.Context) ([]service.CourseResponse, error)
	ListAll(ctx context.Context) ([]service.CourseResponse, error)
	GetBySlug(ctx context.Context, slug string) (*service.CourseResponse, error)
	Upsert(ctx context.Context, p service.UpsertCourseParams) (*service.CourseResponse, error)
	Delete(ctx context.Context, id string) error
	SetStatus(ctx context.Context, id, status string) (*service.CourseResponse, error)
	UpsertLesson(ctx context.Context, courseID string, p service.UpsertLessonParams) (*service.LessonResponse, error)
	DeleteLesson(ctx context.Context, courseID, lessonID string) error
	UpsertQuiz(ctx context.Context, courseID string, p service.UpsertQuizParams) (*service.QuizResponse, error)
	DeleteQuiz(ctx context.Context, courseID string) error
}

type enrollmentService interface {
	ListByUser(ctx context.Context, userID string) ([]service.EnrollmentResponse, error)
	Get(ctx context.Context, userID, courseID string) (*service.EnrollmentResponse, error)
	Enroll(ctx context.Context, userID, courseID string) (*service.EnrollmentResponse, error)
	ToggleLessonComplete(ctx context.Context, userID, courseID, lessonID string) (*service.EnrollmentResponse, error)
	RecordTimeSpent(ctx context.Context, userID, courseID string, seconds int32) (*service.EnrollmentResponse, error)
	MarkQuizPassed(ctx context.Context, userID, courseID, quizID string) (*service.EnrollmentResponse, error)
}

type userService interface {
	List(ctx context.Context) ([]service.UserResponse, error)
	Delete(ctx context.Context, id string) error
}

type transactionService interface {
	ListByUser(ctx context.Context, userID string) ([]service.TransactionResponse, error)
	ListAll(ctx context.Context) ([]service.TransactionResponse, error)
	Purchase(ctx context.Context, userID, courseID string) (*service.TransactionResponse, *service.EnrollmentResponse, error)
}

type wishlistService interface {
	ListByUser(ctx context.Context, userID string) ([]service.WishlistItemResponse, error)
	Toggle(ctx context.Context, userID, courseID string) (bool, error)
}

type chatService interface {
	List(ctx context.Context, userID string, courseID *string) ([]service.ChatMessageResponse, error)
	Append(ctx context.Context, userID string, courseID *string, role, content string) (*service.ChatMessageResponse, error)
	Reset(ctx context.Context, userID string, courseID *string) error
}

type statsService interface {
	Overview(ctx context.Context) (*service.OverviewResponse, error)
	RevenueByPeriod(ctx context.Context, period string) ([]service.PeriodPoint, error)
	ActiveUsersByPeriod(ctx context.Context, period string) ([]service.PeriodPoint, error)
}

type featureService interface {
	List(ctx context.Context) ([]service.FeatureFlagResponse, error)
	Toggle(ctx context.Context, name string) ([]service.FeatureFlagResponse, error)
}

type liveSessionService interface {
	ListPublished(ctx context.Context) ([]service.LiveSessionResponse, error)
	ListAll(ctx context.Context) ([]service.LiveSessionResponse, error)
	Upsert(ctx context.Context, p service.UpsertLiveSessionParams) (*service.LiveSessionResponse, error)
	Delete(ctx context.Context, id string) error
	SetStatus(ctx context.Context, id, status string) (*service.LiveSessionResponse, error)
}

type localeService interface {
	GetCurrencies(ctx context.Context) ([]service.CurrencyResponse, error)
	GetRegions(ctx context.Context) ([]service.RegionResponse, error)
}
