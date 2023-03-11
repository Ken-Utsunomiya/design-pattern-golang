package composite

import (
	"fmt"

	"github.com/projects/design-pattern-golang/Structure/Composite/component"
)

type Folder struct {
	Components []component.Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c component.Component) {
	f.Components = append(f.Components, c)
}
