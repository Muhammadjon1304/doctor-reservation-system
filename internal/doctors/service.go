package doctors

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SearchDoctors(query string) ([]Doctor, error) {
	// Validate input
	if query == "" {
		return s.repo.Search("")
	}

	return s.repo.Search(query)
}

func (s *Service) GetDoctorDetails(id int) (*Doctor, error) {
	return s.repo.GetDoctorByID(id)
}
