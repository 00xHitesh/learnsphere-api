package handler

import (
	"context"
	"errors"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

var errMock = errors.New("mock error")

// mockAuthService

type mockAuthService struct {
	loginFn           func(ctx context.Context, email, password string) (*service.UserResponse, *service.TokenPair, error)
	registerFn        func(ctx context.Context, name, email, password string) (*service.UserResponse, *service.TokenPair, error)
	getUserFn         func(ctx context.Context, userID string) (*service.UserResponse, error)
	updateInterestsFn func(ctx context.Context, userID string, interests []string) (*service.UserResponse, error)
	updateLocaleFn    func(ctx context.Context, userID, regionCode, currencyCode string) (*service.UserResponse, error)
	requestResetFn    func(ctx context.Context, email string) (string, error)
	resetPasswordFn   func(ctx context.Context, token, password string) error
}

func (m *mockAuthService) Login(ctx context.Context, email, password string) (*service.UserResponse, *service.TokenPair, error) {
	if m.loginFn != nil {
		return m.loginFn(ctx, email, password)
	}
	return &service.UserResponse{ID: "u1", Email: email}, &service.TokenPair{AccessToken: "at", RefreshToken: "rt"}, nil
}

func (m *mockAuthService) Register(ctx context.Context, name, email, password string) (*service.UserResponse, *service.TokenPair, error) {
	if m.registerFn != nil {
		return m.registerFn(ctx, name, email, password)
	}
	return &service.UserResponse{ID: "u1", Name: name, Email: email}, &service.TokenPair{AccessToken: "at", RefreshToken: "rt"}, nil
}

func (m *mockAuthService) GetUser(ctx context.Context, userID string) (*service.UserResponse, error) {
	if m.getUserFn != nil {
		return m.getUserFn(ctx, userID)
	}
	return &service.UserResponse{ID: userID}, nil
}

func (m *mockAuthService) UpdateInterests(ctx context.Context, userID string, interests []string) (*service.UserResponse, error) {
	if m.updateInterestsFn != nil {
		return m.updateInterestsFn(ctx, userID, interests)
	}
	return &service.UserResponse{ID: userID, Interests: interests}, nil
}

func (m *mockAuthService) UpdateLocale(ctx context.Context, userID, regionCode, currencyCode string) (*service.UserResponse, error) {
	if m.updateLocaleFn != nil {
		return m.updateLocaleFn(ctx, userID, regionCode, currencyCode)
	}
	return &service.UserResponse{ID: userID, RegionCode: regionCode, CurrencyCode: currencyCode}, nil
}

func (m *mockAuthService) RequestReset(ctx context.Context, email string) (string, error) {
	if m.requestResetFn != nil {
		return m.requestResetFn(ctx, email)
	}
	return "reset-token", nil
}

func (m *mockAuthService) ResetPassword(ctx context.Context, token, password string) error {
	if m.resetPasswordFn != nil {
		return m.resetPasswordFn(ctx, token, password)
	}
	return nil
}

// mockCourseService

type mockCourseService struct {
	listPublishedFn func(ctx context.Context) ([]service.CourseResponse, error)
	listAllFn       func(ctx context.Context) ([]service.CourseResponse, error)
	getBySlugFn     func(ctx context.Context, slug string) (*service.CourseResponse, error)
	upsertFn        func(ctx context.Context, p service.UpsertCourseParams) (*service.CourseResponse, error)
	deleteFn        func(ctx context.Context, id string) error
	setStatusFn     func(ctx context.Context, id, status string) (*service.CourseResponse, error)
}

func (m *mockCourseService) ListPublished(ctx context.Context) ([]service.CourseResponse, error) {
	if m.listPublishedFn != nil {
		return m.listPublishedFn(ctx)
	}
	return []service.CourseResponse{{ID: "c1", Title: "Course 1"}}, nil
}

func (m *mockCourseService) ListAll(ctx context.Context) ([]service.CourseResponse, error) {
	if m.listAllFn != nil {
		return m.listAllFn(ctx)
	}
	return []service.CourseResponse{{ID: "c1", Title: "Course 1"}}, nil
}

func (m *mockCourseService) GetBySlug(ctx context.Context, slug string) (*service.CourseResponse, error) {
	if m.getBySlugFn != nil {
		return m.getBySlugFn(ctx, slug)
	}
	return &service.CourseResponse{ID: "c1", Slug: slug}, nil
}

func (m *mockCourseService) Upsert(ctx context.Context, p service.UpsertCourseParams) (*service.CourseResponse, error) {
	if m.upsertFn != nil {
		return m.upsertFn(ctx, p)
	}
	return &service.CourseResponse{ID: "c1", Title: p.Title, Slug: p.Slug}, nil
}

func (m *mockCourseService) Delete(ctx context.Context, id string) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}
	return nil
}

func (m *mockCourseService) SetStatus(ctx context.Context, id, status string) (*service.CourseResponse, error) {
	if m.setStatusFn != nil {
		return m.setStatusFn(ctx, id, status)
	}
	return &service.CourseResponse{ID: id, Status: status}, nil
}

// mockEnrollmentService

type mockEnrollmentService struct {
	listByUserFn          func(ctx context.Context, userID string) ([]service.EnrollmentResponse, error)
	getFn                 func(ctx context.Context, userID, courseID string) (*service.EnrollmentResponse, error)
	enrollFn              func(ctx context.Context, userID, courseID string) (*service.EnrollmentResponse, error)
	toggleLessonFn        func(ctx context.Context, userID, courseID, lessonID string) (*service.EnrollmentResponse, error)
	recordTimeSpentFn     func(ctx context.Context, userID, courseID string, seconds int32) (*service.EnrollmentResponse, error)
	markQuizPassedFn      func(ctx context.Context, userID, courseID, quizID string) (*service.EnrollmentResponse, error)
}

func (m *mockEnrollmentService) ListByUser(ctx context.Context, userID string) ([]service.EnrollmentResponse, error) {
	if m.listByUserFn != nil {
		return m.listByUserFn(ctx, userID)
	}
	return []service.EnrollmentResponse{{ID: "e1", UserID: userID}}, nil
}

func (m *mockEnrollmentService) Get(ctx context.Context, userID, courseID string) (*service.EnrollmentResponse, error) {
	if m.getFn != nil {
		return m.getFn(ctx, userID, courseID)
	}
	return &service.EnrollmentResponse{ID: "e1", UserID: userID, CourseID: courseID}, nil
}

func (m *mockEnrollmentService) Enroll(ctx context.Context, userID, courseID string) (*service.EnrollmentResponse, error) {
	if m.enrollFn != nil {
		return m.enrollFn(ctx, userID, courseID)
	}
	return &service.EnrollmentResponse{ID: "e1", UserID: userID, CourseID: courseID}, nil
}

func (m *mockEnrollmentService) ToggleLessonComplete(ctx context.Context, userID, courseID, lessonID string) (*service.EnrollmentResponse, error) {
	if m.toggleLessonFn != nil {
		return m.toggleLessonFn(ctx, userID, courseID, lessonID)
	}
	return &service.EnrollmentResponse{ID: "e1"}, nil
}

func (m *mockEnrollmentService) RecordTimeSpent(ctx context.Context, userID, courseID string, seconds int32) (*service.EnrollmentResponse, error) {
	if m.recordTimeSpentFn != nil {
		return m.recordTimeSpentFn(ctx, userID, courseID, seconds)
	}
	return &service.EnrollmentResponse{ID: "e1"}, nil
}

func (m *mockEnrollmentService) MarkQuizPassed(ctx context.Context, userID, courseID, quizID string) (*service.EnrollmentResponse, error) {
	if m.markQuizPassedFn != nil {
		return m.markQuizPassedFn(ctx, userID, courseID, quizID)
	}
	return &service.EnrollmentResponse{ID: "e1"}, nil
}

// mockUserService

type mockUserService struct {
	listFn   func(ctx context.Context) ([]service.UserResponse, error)
	deleteFn func(ctx context.Context, id string) error
}

func (m *mockUserService) List(ctx context.Context) ([]service.UserResponse, error) {
	if m.listFn != nil {
		return m.listFn(ctx)
	}
	return []service.UserResponse{{ID: "u1"}}, nil
}

func (m *mockUserService) Delete(ctx context.Context, id string) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}
	return nil
}

// mockTransactionService

type mockTransactionService struct {
	listByUserFn func(ctx context.Context, userID string) ([]service.TransactionResponse, error)
	listAllFn    func(ctx context.Context) ([]service.TransactionResponse, error)
	purchaseFn   func(ctx context.Context, userID, courseID string) (*service.TransactionResponse, *service.EnrollmentResponse, error)
}

func (m *mockTransactionService) ListByUser(ctx context.Context, userID string) ([]service.TransactionResponse, error) {
	if m.listByUserFn != nil {
		return m.listByUserFn(ctx, userID)
	}
	return []service.TransactionResponse{{ID: "t1"}}, nil
}

func (m *mockTransactionService) ListAll(ctx context.Context) ([]service.TransactionResponse, error) {
	if m.listAllFn != nil {
		return m.listAllFn(ctx)
	}
	return []service.TransactionResponse{{ID: "t1"}}, nil
}

func (m *mockTransactionService) Purchase(ctx context.Context, userID, courseID string) (*service.TransactionResponse, *service.EnrollmentResponse, error) {
	if m.purchaseFn != nil {
		return m.purchaseFn(ctx, userID, courseID)
	}
	return &service.TransactionResponse{ID: "t1"}, &service.EnrollmentResponse{ID: "e1"}, nil
}

// mockWishlistService

type mockWishlistService struct {
	listByUserFn func(ctx context.Context, userID string) ([]service.WishlistItemResponse, error)
	toggleFn     func(ctx context.Context, userID, courseID string) (bool, error)
}

func (m *mockWishlistService) ListByUser(ctx context.Context, userID string) ([]service.WishlistItemResponse, error) {
	if m.listByUserFn != nil {
		return m.listByUserFn(ctx, userID)
	}
	return []service.WishlistItemResponse{{UserID: userID}}, nil
}

func (m *mockWishlistService) Toggle(ctx context.Context, userID, courseID string) (bool, error) {
	if m.toggleFn != nil {
		return m.toggleFn(ctx, userID, courseID)
	}
	return true, nil
}

// mockChatService

type mockChatService struct {
	listFn   func(ctx context.Context, userID string, courseID *string) ([]service.ChatMessageResponse, error)
	appendFn func(ctx context.Context, userID string, courseID *string, role, content string) (*service.ChatMessageResponse, error)
	resetFn  func(ctx context.Context, userID string, courseID *string) error
}

func (m *mockChatService) List(ctx context.Context, userID string, courseID *string) ([]service.ChatMessageResponse, error) {
	if m.listFn != nil {
		return m.listFn(ctx, userID, courseID)
	}
	return []service.ChatMessageResponse{{ID: "msg1", Role: "user", Content: "hello"}}, nil
}

func (m *mockChatService) Append(ctx context.Context, userID string, courseID *string, role, content string) (*service.ChatMessageResponse, error) {
	if m.appendFn != nil {
		return m.appendFn(ctx, userID, courseID, role, content)
	}
	return &service.ChatMessageResponse{ID: "msg1", Role: role, Content: content}, nil
}

func (m *mockChatService) Reset(ctx context.Context, userID string, courseID *string) error {
	if m.resetFn != nil {
		return m.resetFn(ctx, userID, courseID)
	}
	return nil
}

// mockStatsService

type mockStatsService struct {
	overviewFn          func(ctx context.Context) (*service.OverviewResponse, error)
	revenueByPeriodFn   func(ctx context.Context, period string) ([]service.PeriodPoint, error)
	activeUsersByPeriodFn func(ctx context.Context, period string) ([]service.PeriodPoint, error)
}

func (m *mockStatsService) Overview(ctx context.Context) (*service.OverviewResponse, error) {
	if m.overviewFn != nil {
		return m.overviewFn(ctx)
	}
	return &service.OverviewResponse{TotalUsers: 10, TotalCourses: 5}, nil
}

func (m *mockStatsService) RevenueByPeriod(ctx context.Context, period string) ([]service.PeriodPoint, error) {
	if m.revenueByPeriodFn != nil {
		return m.revenueByPeriodFn(ctx, period)
	}
	return []service.PeriodPoint{{Label: "2024-01-01", Value: 100}}, nil
}

func (m *mockStatsService) ActiveUsersByPeriod(ctx context.Context, period string) ([]service.PeriodPoint, error) {
	if m.activeUsersByPeriodFn != nil {
		return m.activeUsersByPeriodFn(ctx, period)
	}
	return []service.PeriodPoint{{Label: "2024-01-01", Value: 5}}, nil
}

// mockFeatureService

type mockFeatureService struct {
	listFn   func(ctx context.Context) ([]service.FeatureFlagResponse, error)
	toggleFn func(ctx context.Context, name string) ([]service.FeatureFlagResponse, error)
}

func (m *mockFeatureService) List(ctx context.Context) ([]service.FeatureFlagResponse, error) {
	if m.listFn != nil {
		return m.listFn(ctx)
	}
	return []service.FeatureFlagResponse{{ID: "f1", Name: "dark-mode", Enabled: true}}, nil
}

func (m *mockFeatureService) Toggle(ctx context.Context, name string) ([]service.FeatureFlagResponse, error) {
	if m.toggleFn != nil {
		return m.toggleFn(ctx, name)
	}
	return []service.FeatureFlagResponse{{ID: "f1", Name: name, Enabled: false}}, nil
}

// mockLiveSessionService

type mockLiveSessionService struct {
	listPublishedFn func(ctx context.Context) ([]service.LiveSessionResponse, error)
	listAllFn       func(ctx context.Context) ([]service.LiveSessionResponse, error)
	upsertFn        func(ctx context.Context, p service.UpsertLiveSessionParams) (*service.LiveSessionResponse, error)
	deleteFn        func(ctx context.Context, id string) error
	setStatusFn     func(ctx context.Context, id, status string) (*service.LiveSessionResponse, error)
}

func (m *mockLiveSessionService) ListPublished(ctx context.Context) ([]service.LiveSessionResponse, error) {
	if m.listPublishedFn != nil {
		return m.listPublishedFn(ctx)
	}
	return []service.LiveSessionResponse{{ID: "s1", Title: "Session 1"}}, nil
}

func (m *mockLiveSessionService) ListAll(ctx context.Context) ([]service.LiveSessionResponse, error) {
	if m.listAllFn != nil {
		return m.listAllFn(ctx)
	}
	return []service.LiveSessionResponse{{ID: "s1", Title: "Session 1"}}, nil
}

func (m *mockLiveSessionService) Upsert(ctx context.Context, p service.UpsertLiveSessionParams) (*service.LiveSessionResponse, error) {
	if m.upsertFn != nil {
		return m.upsertFn(ctx, p)
	}
	return &service.LiveSessionResponse{ID: "s1", Title: p.Title}, nil
}

func (m *mockLiveSessionService) Delete(ctx context.Context, id string) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}
	return nil
}

func (m *mockLiveSessionService) SetStatus(ctx context.Context, id, status string) (*service.LiveSessionResponse, error) {
	if m.setStatusFn != nil {
		return m.setStatusFn(ctx, id, status)
	}
	return &service.LiveSessionResponse{ID: id, Status: status}, nil
}

// mockLocaleService

type mockLocaleService struct {
	getCurrenciesFn func(ctx context.Context) ([]service.CurrencyResponse, error)
	getRegionsFn    func(ctx context.Context) ([]service.RegionResponse, error)
}

func (m *mockLocaleService) GetCurrencies(ctx context.Context) ([]service.CurrencyResponse, error) {
	if m.getCurrenciesFn != nil {
		return m.getCurrenciesFn(ctx)
	}
	return []service.CurrencyResponse{{Code: "USD", Symbol: "$", Name: "US Dollar"}}, nil
}

func (m *mockLocaleService) GetRegions(ctx context.Context) ([]service.RegionResponse, error) {
	if m.getRegionsFn != nil {
		return m.getRegionsFn(ctx)
	}
	return []service.RegionResponse{{Code: "US", Name: "United States", DefaultCurrencyCode: "USD"}}, nil
}
