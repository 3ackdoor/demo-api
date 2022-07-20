package request

type SortDirection string

const (
	ASC  SortDirection = "asc"
	DESC SortDirection = "desc"
)

var (
	DefaultSortColumn    string = "created_at"
	DefaultSortDirection string = "asc"
)

type BasePaginationRequest struct {
	Page          int           `form:"page,default=1"`
	Limit         int           `form:"limit,default=10"`
	SortColumn    string        `form:"sort-col,default=updated_at"`
	SortDirection SortDirection `form:"sort-dir,default=desc"`
}
