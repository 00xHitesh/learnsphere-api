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

type CourseService struct {
	q    *repository.Queries
	pool *pgxpool.Pool
}

func NewCourseService(pool *pgxpool.Pool) *CourseService {
	return &CourseService{q: repository.New(pool), pool: pool}
}

type LessonResponse struct {
	ID          string  `json:"id"`
	CourseID    string  `json:"courseId"`
	Title       string  `json:"title"`
	Duration    string  `json:"duration"`
	VideoURL    *string `json:"videoUrl"`
	IsPreview   bool    `json:"isPreview"`
	Description *string `json:"description"`
	SortOrder   int16   `json:"sortOrder"`
}

type ReviewResponse struct {
	ID           string  `json:"id"`
	CourseID     string  `json:"courseId"`
	UserID       *string `json:"userId"`
	ReviewerName string  `json:"reviewerName"`
	Rating       int16   `json:"rating"`
	Comment      string  `json:"comment"`
	ReviewedAt   string  `json:"reviewedAt"`
}

type QuizOptionResponse struct {
	ID        string `json:"id"`
	OptionKey string `json:"optionKey"`
	Text      string `json:"text"`
}

type QuizQuestionResponse struct {
	ID                string               `json:"id"`
	Prompt            string               `json:"prompt"`
	CorrectOptionKeys []string             `json:"correctOptionKeys"`
	IsMultiSelect     bool                 `json:"isMultiSelect"`
	Explanation       *string              `json:"explanation"`
	Options           []QuizOptionResponse `json:"options"`
}

type QuizResponse struct {
	ID           string                 `json:"id"`
	CourseID     string                 `json:"courseId"`
	Title        string                 `json:"title"`
	Description  *string                `json:"description"`
	PassingScore int16                  `json:"passingScore"`
	Questions    []QuizQuestionResponse `json:"questions"`
}

type CourseResponse struct {
	ID            string          `json:"id"`
	Title         string          `json:"title"`
	Slug          string          `json:"slug"`
	Description   string          `json:"description"`
	Summary       string          `json:"summary"`
	Category      string          `json:"category"`
	Level         string          `json:"level"`
	ListPrice     float64         `json:"listPrice"`
	SalePrice     *float64        `json:"salePrice"`
	Rating        float64         `json:"rating"`
	StudentCount  int32           `json:"studentCount"`
	Instructor    string          `json:"instructor"`
	InstructorBio string          `json:"instructorBio"`
	Cover         string          `json:"cover"`
	Status        string          `json:"status"`
	Tags          []string        `json:"tags"`
	Lessons       []LessonResponse `json:"lessons,omitempty"`
	Reviews       []ReviewResponse `json:"reviews,omitempty"`
	Quiz          *QuizResponse   `json:"quiz,omitempty"`
	CreatedAt     string          `json:"createdAt"`
	UpdatedAt     string          `json:"updatedAt"`
}

type UpsertCourseParams struct {
	ID            string
	Title         string
	Slug          string
	Description   string
	Summary       string
	Category      string
	Level         string
	ListPrice     float64
	SalePrice     *float64
	Instructor    string
	InstructorBio string
	Cover         string
	Status        string
	Tags          []string
}

func (s *CourseService) ListPublished(ctx context.Context) ([]CourseResponse, error) {
	courses, err := s.q.ListPublishedCourses(ctx)
	if err != nil {
		return nil, err
	}
	return s.attachTagsToList(ctx, courses)
}

func (s *CourseService) ListAll(ctx context.Context) ([]CourseResponse, error) {
	courses, err := s.q.ListAllCourses(ctx)
	if err != nil {
		return nil, err
	}
	return s.attachTagsToList(ctx, courses)
}

func (s *CourseService) GetBySlug(ctx context.Context, slug string) (*CourseResponse, error) {
	course, err := s.q.GetCourseBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	resp, err := s.buildCourseResponse(ctx, course, true)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *CourseService) Upsert(ctx context.Context, p UpsertCourseParams) (*CourseResponse, error) {
	var course repository.Course
	var err error

	level := repository.CourseLevel(p.Level)
	status := repository.ContentStatus(p.Status)
	listPrice := floatToNumeric(p.ListPrice)
	salePrice := optFloatToNumeric(p.SalePrice)

	if p.ID == "" {
		course, err = s.q.InsertCourse(ctx, repository.InsertCourseParams{
			Title: p.Title, Slug: p.Slug, Description: p.Description,
			Summary: p.Summary, Category: p.Category, Level: level,
			ListPrice: listPrice, SalePrice: salePrice,
			Instructor: p.Instructor, InstructorBio: p.InstructorBio,
			Cover: p.Cover, Status: status,
		})
	} else {
		uid, parseErr := parseUUID(p.ID)
		if parseErr != nil {
			return nil, parseErr
		}
		course, err = s.q.UpdateCourse(ctx, repository.UpdateCourseParams{
			ID: uid, Title: p.Title, Slug: p.Slug, Description: p.Description,
			Summary: p.Summary, Category: p.Category, Level: level,
			ListPrice: listPrice, SalePrice: salePrice,
			Instructor: p.Instructor, InstructorBio: p.InstructorBio,
			Cover: p.Cover, Status: status,
		})
		if err != nil && errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("course not found")
		}
	}
	if err != nil {
		return nil, err
	}

	if err := s.syncTags(ctx, course.ID, p.Tags); err != nil {
		return nil, err
	}

	resp, err := s.buildCourseResponse(ctx, course, false)
	if err != nil {
		return nil, err
	}
	resp.Tags = p.Tags
	return &resp, nil
}

func (s *CourseService) Delete(ctx context.Context, id string) error {
	uid, err := parseUUID(id)
	if err != nil {
		return err
	}
	return s.q.DeleteCourse(ctx, uid)
}

func (s *CourseService) SetStatus(ctx context.Context, id, status string) (*CourseResponse, error) {
	uid, err := parseUUID(id)
	if err != nil {
		return nil, err
	}
	course, err := s.q.UpdateCourseStatus(ctx, repository.UpdateCourseStatusParams{
		ID:     uid,
		Status: repository.ContentStatus(status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	resp, err := s.buildCourseResponse(ctx, course, false)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *CourseService) syncTags(ctx context.Context, courseID pgtype.UUID, tagNames []string) error {
	if err := s.q.DeleteCourseTagsByCourse(ctx, courseID); err != nil {
		return err
	}
	for _, name := range tagNames {
		tag, err := s.q.UpsertTag(ctx, name)
		if err != nil {
			return err
		}
		_ = s.q.InsertCourseTag(ctx, repository.InsertCourseTagParams{
			CourseID: courseID,
			TagID:    tag.ID,
		})
	}
	return nil
}

func (s *CourseService) attachTagsToList(ctx context.Context, courses []repository.Course) ([]CourseResponse, error) {
	if len(courses) == 0 {
		return []CourseResponse{}, nil
	}
	ids := make([]pgtype.UUID, len(courses))
	for i, c := range courses {
		ids[i] = c.ID
	}
	tagRows, err := s.q.GetTagNamesForCourseIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	tagMap := make(map[string][]string)
	for _, row := range tagRows {
		cid := uuidToString(row.CourseID)
		tagMap[cid] = append(tagMap[cid], row.Name)
	}
	result := make([]CourseResponse, len(courses))
	for i, c := range courses {
		resp := toCourseResponse(c)
		tags := tagMap[uuidToString(c.ID)]
		if tags == nil {
			tags = []string{}
		}
		resp.Tags = tags
		result[i] = resp
	}
	return result, nil
}

func (s *CourseService) buildCourseResponse(ctx context.Context, c repository.Course, withDetails bool) (CourseResponse, error) {
	resp := toCourseResponse(c)
	tags, err := s.q.GetTagNamesByCourseID(ctx, c.ID)
	if err != nil {
		return resp, err
	}
	if tags == nil {
		tags = []string{}
	}
	resp.Tags = tags

	if !withDetails {
		return resp, nil
	}

	lessons, _ := s.q.ListLessonsByCourse(ctx, c.ID)
	resp.Lessons = toLessonResponses(lessons)

	reviews, _ := s.q.ListReviewsByCourse(ctx, c.ID)
	resp.Reviews = toReviewResponses(reviews)

	quiz, err := s.q.GetQuizByCourse(ctx, c.ID)
	if err == nil {
		questions, _ := s.q.ListQuizQuestions(ctx, quiz.ID)
		qIDs := make([]pgtype.UUID, len(questions))
		for i, q := range questions {
			qIDs[i] = q.ID
		}
		var opts []repository.QuizOption
		if len(qIDs) > 0 {
			opts, _ = s.q.ListQuizOptions(ctx, qIDs)
		}
		resp.Quiz = toQuizResponse(quiz, questions, opts)
	}

	return resp, nil
}

func toCourseResponse(c repository.Course) CourseResponse {
	return CourseResponse{
		ID:            uuidToString(c.ID),
		Title:         c.Title,
		Slug:          c.Slug,
		Description:   c.Description,
		Summary:       c.Summary,
		Category:      c.Category,
		Level:         string(c.Level),
		ListPrice:     numericToFloat(c.ListPrice),
		SalePrice:     nullNumericToPtr(c.SalePrice),
		Rating:        numericToFloat(c.Rating),
		StudentCount:  c.StudentCount,
		Instructor:    c.Instructor,
		InstructorBio: c.InstructorBio,
		Cover:         c.Cover,
		Status:        string(c.Status),
		Tags:          []string{},
		CreatedAt:     c.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:     c.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func toLessonResponses(lessons []repository.Lesson) []LessonResponse {
	out := make([]LessonResponse, len(lessons))
	for i, l := range lessons {
		out[i] = LessonResponse{
			ID:          uuidToString(l.ID),
			CourseID:    uuidToString(l.CourseID),
			Title:       l.Title,
			Duration:    l.Duration,
			VideoURL:    nullTextToPtr(l.VideoUrl),
			IsPreview:   l.IsPreview,
			Description: nullTextToPtr(l.Description),
			SortOrder:   l.SortOrder,
		}
	}
	return out
}

func toReviewResponses(reviews []repository.Review) []ReviewResponse {
	out := make([]ReviewResponse, len(reviews))
	for i, r := range reviews {
		var uid *string
		if r.UserID.Valid {
			s := uuidToString(r.UserID)
			uid = &s
		}
		out[i] = ReviewResponse{
			ID:           uuidToString(r.ID),
			CourseID:     uuidToString(r.CourseID),
			UserID:       uid,
			ReviewerName: r.ReviewerName,
			Rating:       r.Rating,
			Comment:      r.Comment,
			ReviewedAt:   r.ReviewedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		}
	}
	return out
}

func toQuizResponse(quiz repository.Quiz, questions []repository.QuizQuestion, opts []repository.QuizOption) *QuizResponse {
	optMap := make(map[string][]QuizOptionResponse)
	for _, o := range opts {
		qid := uuidToString(o.QuestionID)
		optMap[qid] = append(optMap[qid], QuizOptionResponse{
			ID:        uuidToString(o.ID),
			OptionKey: o.OptionKey,
			Text:      o.Text,
		})
	}
	qs := make([]QuizQuestionResponse, len(questions))
	for i, q := range questions {
		qid := uuidToString(q.ID)
		qs[i] = QuizQuestionResponse{
			ID:                qid,
			Prompt:            q.Prompt,
			CorrectOptionKeys: q.CorrectOptionKeys,
			IsMultiSelect:     q.IsMultiSelect,
			Explanation:       nullTextToPtr(q.Explanation),
			Options:           optMap[qid],
		}
		if qs[i].Options == nil {
			qs[i].Options = []QuizOptionResponse{}
		}
	}
	return &QuizResponse{
		ID:           uuidToString(quiz.ID),
		CourseID:     uuidToString(quiz.CourseID),
		Title:        quiz.Title,
		Description:  nullTextToPtr(quiz.Description),
		PassingScore: quiz.PassingScore,
		Questions:    qs,
	}
}

func floatToNumeric(f float64) pgtype.Numeric {
	var n pgtype.Numeric
	_ = n.Scan(fmt.Sprintf("%.2f", f))
	return n
}

func optFloatToNumeric(f *float64) pgtype.Numeric {
	if f == nil {
		return pgtype.Numeric{}
	}
	return floatToNumeric(*f)
}