package vacancy

import (
	"strings"

	"github.com/predator272/app-go/internal/helper"
)

type Model struct {
	ID          uint64
	Position    string
	Experience  uint64
	Description string
}

func (this *Model) Validate() bool {
	this.Position = strings.Join(strings.Fields(strings.TrimSpace(this.Position)), " ")
	this.Experience = helper.Clamp(this.Experience, 0, 100)
	this.Description = strings.Join(strings.Fields(strings.TrimSpace(this.Description)), " ")
	return (len(this.Position) > 0 && this.Experience >= 0 && len(this.Description) > 0)
}
