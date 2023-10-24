package user

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

//func (s *Service) UpsertUserFunds(ctx context.Context) ([]Fund, error) {
//	funds := s.repo.UpsertUserFunds(ctx, searchTerm)
//	if err != nil {
//		return nil, err
//	}
//	return funds, nil
//}
