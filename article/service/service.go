package service

import (
	"context"

	"github.com/tkr53/go-gprc-graphql/article/pb"
	"github.com/tkr53/go-gprc-graphql/article/repository"
)

type Service interface {
	CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error)
	ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error)
	UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error)
	ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{r}
}

func (s *service) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	input := req.GetArticleInput()

	id, err := s.repository.InsertArticle(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error) {
	id := req.GetId()

	a, err := s.repository.SelectArticleById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.ReadArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  a.Author,
			Title:   a.Title,
			Content: a.Content,
		},
	}, nil
}

func (s *service) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	id := req.GetId()
	input := req.GetArticleInput()

	if err := s.repository.UpdateArticle(ctx, id, input); err != nil {
		return nil, err
	}

	return &pb.UpdateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	id := req.GetId()

	if err := s.repository.DeleteArticle(ctx, id); err != nil {
		return nil, err
	}

	return &pb.DeleteArticleResponse{Id: id}, nil
}

func (s *service) ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error {
	rows, err := s.repository.SelectAllArticles()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var a pb.Article
		err := rows.Scan(&a.Id, &a.Author, &a.Title, &a.Content)
		if err != nil {
			return err
		}
		stream.Send(&pb.ListArticleResponse{Article: &a})
	}

	return nil
}
