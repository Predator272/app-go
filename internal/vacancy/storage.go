package vacancy

import (
	"github.com/jmoiron/sqlx"

	"github.com/predator272/app-go/internal/helper"
)

type Storage struct {
	db *sqlx.DB
}

func (this *Storage) QueryOne(query string, args ...any) *Model {
	model := Model{}
	if helper.ErrorArg(this.db.Preparex(query)).Get(&model, args...) != nil {
		return nil
	}
	return &model
}

func (this *Storage) QueryAll(query string, args ...any) []Model {
	model := []Model{}
	if helper.ErrorArg(this.db.Preparex(query)).Select(&model, args...) != nil {
		return nil
	}
	return model
}

func (this *Storage) Create() {
	helper.ErrorArg(this.db.Exec("CREATE TABLE IF NOT EXISTS vacancy(id INTEGER PRIMARY KEY, position TEXT, experience UNSIGNED INTEGER, description TEXT)"))
}

func (this *Storage) Save(model Model) *Model {
	return this.QueryOne("INSERT INTO vacancy(position, experience, description) VALUES(?, ?, ?) RETURNING *", model.Position, model.Experience, model.Description)
}

func (this *Storage) UpdateById(model Model) *Model {
	return this.QueryOne("UPDATE vacancy SET position = ?, experience = ?, description = ? WHERE id = ? RETURNING *", model.Position, model.Experience, model.Description, model.ID)
}

func (this *Storage) DeleteById(model Model) *Model {
	return this.QueryOne("DELETE FROM vacancy WHERE id = ? RETURNING *", model.ID)
}

func (this *Storage) FindById(model Model) *Model {
	return this.QueryOne("SELECT * FROM vacancy WHERE id = ?", model.ID)
}

func (this *Storage) All() []Model {
	return this.QueryAll("SELECT * FROM vacancy")
}
