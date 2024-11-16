package article

type Author struct {
	Id   int
	Name string
}

type Article struct {
	Id      int
	Title   string
	Content string
	Author  Author
}

type Storage interface {
	Add(a *Article) (int, error)
}

type Service struct {
	Storage Storage
}

func (s *Service) Add(a *Article) error {
	id, err := s.Storage.Add(a)
	if err != nil {
		return err
	}
	a.Id = id
    return nil
}

func (s *Service) Update(a *Article) error {
	return nil
}

func (s *Service) Delete(id int) error {
	return nil
}

func (s *Service) Get(id int) error {
	return nil
}

func (s *Service) List() error {
	return nil
}
