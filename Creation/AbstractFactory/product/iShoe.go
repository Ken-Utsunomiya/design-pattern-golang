package product

// abstract product

type IShoe interface {
	SetLogo(Logo string)
	SetSize(Size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	Logo string
	Size int
}

func (s *Shoe) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.Logo
}

func (s *Shoe) SetSize(size int) {
	s.Size = size
}

func (s *Shoe) GetSize() int {
	return s.Size
}
