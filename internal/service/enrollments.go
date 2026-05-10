package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type EnrollmentService struct {
	q *repository.Queries
}

func NewEnrollmentService(pool *pgxpool.Pool) *EnrollmentService {
	return &EnrollmentService{q: repository.New(pool)}
}

type EnrollmentResponse struct {
	ID                  string   `json:"id"`
	UserID              string   `json:"userId"`
	CourseID            string   `json:"courseId"`
	CompletedLessonIDs  []string `json:"completedLessonIds"`
	PassedQuizIDs       []string `json:"passedQuizIds"`
	TimeSpentSeconds    int32    `json:"timeSpentSeconds"`
	EnrolledAt          string   `json:"enrolledAt"`
}

func (s *EnrollmentService) ListByUser(ctx context.Context, userID string) ([]EnrollmentResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	rows, err := s.q.ListEnrollmentsByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return []EnrollmentResponse{}, nil
	}
	eIDs := make([]pgtype.UUID, len(rows))
	for i, r := range rows {
		eIDs[i] = r.ID
	}
	progress, err := s.q.GetCompletedLessonIDsForEnrollments(ctx, eIDs)
	if err != nil {
		return nil, err
	}
	lessonMap := make(map[string][]string)
	for _, p := range progress {
		eid := uuidToString(p.EnrollmentID)
		lessonMap[eid] = append(lessonMap[eid], uuidToString(p.LessonID))
	}
	out := make([]EnrollmentResponse, len(rows))
	for i, r := range rows {
		out[i] = toEnrollmentResponse(r, lessonMap[uuidToString(r.ID)])
	}
	return out, nil
}

func (s *EnrollmentService) Get(ctx context.Context, userID, courseID string) (*EnrollmentResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	cid, err := parseUUID(courseID)
	if err != nil {
		return nil, err
	}
	e, err := s.q.GetEnrollment(ctx, repository.GetEnrollmentParams{UserID: uid, CourseID: cid})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	completedIDs, _ := s.q.GetCompletedLessonIDs(ctx, e.ID)
	resp := toEnrollmentResponse(e, uuidsToStrings(completedIDs))
	return &resp, nil
}

func (s *EnrollmentService) Enroll(ctx context.Context, userID, courseID string) (*EnrollmentResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	cid, err := parseUUID(courseID)
	if err != nil {
		return nil, err
	}
	e, err := s.q.InsertEnrollment(ctx, repository.InsertEnrollmentParams{UserID: uid, CourseID: cid})
	if err != nil {
		return nil, err
	}
	completedIDs, _ := s.q.GetCompletedLessonIDs(ctx, e.ID)
	resp := toEnrollmentResponse(e, uuidsToStrings(completedIDs))
	return &resp, nil
}

func (s *EnrollmentService) ToggleLessonComplete(ctx context.Context, userID, courseID, lessonID string) (*EnrollmentResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	cid, err := parseUUID(courseID)
	if err != nil {
		return nil, err
	}
	lid, err := parseUUID(lessonID)
	if err != nil {
		return nil, err
	}
	e, err := s.q.GetEnrollment(ctx, repository.GetEnrollmentParams{UserID: uid, CourseID: cid})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("not enrolled")
		}
		return nil, err
	}
	existing, _ := s.q.GetCompletedLessonIDs(ctx, e.ID)
	wasCompleted := uuidSliceContains(existing, lid)
	if wasCompleted {
		_ = s.q.DeleteLessonProgress(ctx, repository.DeleteLessonProgressParams{EnrollmentID: e.ID, LessonID: lid})
		_ = s.q.IncrementUserPoints(ctx, repository.IncrementUserPointsParams{ID: uid, Points: -10})
	} else {
		_ = s.q.InsertLessonProgress(ctx, repository.InsertLessonProgressParams{EnrollmentID: e.ID, LessonID: lid})
		_ = s.q.IncrementUserPoints(ctx, repository.IncrementUserPointsParams{ID: uid, Points: 10})
	}
	completedIDs, _ := s.q.GetCompletedLessonIDs(ctx, e.ID)
	resp := toEnrollmentResponse(e, uuidsToStrings(completedIDs))
	return &resp, nil
}

func (s *EnrollmentService) RecordTimeSpent(ctx context.Context, userID, courseID string, seconds int32) (*EnrollmentResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	cid, err := parseUUID(courseID)
	if err != nil {
		return nil, err
	}
	e, err := s.q.GetEnrollment(ctx, repository.GetEnrollmentParams{UserID: uid, CourseID: cid})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	updated, err := s.q.AddEnrollmentTimeSpent(ctx, repository.AddEnrollmentTimeSpentParams{
		ID:                e.ID,
		TimeSpentSeconds:  seconds,
	})
	if err != nil {
		return nil, err
	}
	completedIDs, _ := s.q.GetCompletedLessonIDs(ctx, e.ID)
	resp := toEnrollmentResponse(updated, uuidsToStrings(completedIDs))
	return &resp, nil
}

func (s *EnrollmentService) MarkQuizPassed(ctx context.Context, userID, courseID, quizID string) (*EnrollmentResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	cid, err := parseUUID(courseID)
	if err != nil {
		return nil, err
	}
	e, err := s.q.GetEnrollment(ctx, repository.GetEnrollmentParams{UserID: uid, CourseID: cid})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	already := false
	for _, id := range e.PassedQuizIds {
		if id == quizID {
			already = true
			break
		}
	}
	if !already {
		newIDs := append(e.PassedQuizIds, quizID)
		updated, err := s.q.UpdatePassedQuizIds(ctx, repository.UpdatePassedQuizIdsParams{
			ID:             e.ID,
			PassedQuizIds:  newIDs,
		})
		if err != nil {
			return nil, err
		}
		e = updated
	}
	completedIDs, _ := s.q.GetCompletedLessonIDs(ctx, e.ID)
	resp := toEnrollmentResponse(e, uuidsToStrings(completedIDs))
	return &resp, nil
}

func toEnrollmentResponse(e repository.Enrollment, completedLessonIDs []string) EnrollmentResponse {
	if completedLessonIDs == nil {
		completedLessonIDs = []string{}
	}
	passedQuizIDs := e.PassedQuizIds
	if passedQuizIDs == nil {
		passedQuizIDs = []string{}
	}
	return EnrollmentResponse{
		ID:                 uuidToString(e.ID),
		UserID:             uuidToString(e.UserID),
		CourseID:           uuidToString(e.CourseID),
		CompletedLessonIDs: completedLessonIDs,
		PassedQuizIDs:      passedQuizIDs,
		TimeSpentSeconds:   e.TimeSpentSeconds,
		EnrolledAt:         e.EnrolledAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func uuidsToStrings(ids []pgtype.UUID) []string {
	out := make([]string, len(ids))
	for i, id := range ids {
		out[i] = uuidToString(id)
	}
	return out
}

func uuidSliceContains(ids []pgtype.UUID, target pgtype.UUID) bool {
	for _, id := range ids {
		if id == target {
			return true
		}
	}
	return false
}