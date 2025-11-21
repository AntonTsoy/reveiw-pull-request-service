package repository

type SubsRepo struct {
	db *sqlx.DB
}

func NewSubsRepo(db *sqlx.DB) *SubsRepo {
	return &SubsRepo{db: db}
}
