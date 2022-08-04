package object

type (
	StatusID    = int64

	Status struct {
		ID StatusID `json:"id"`

		Account *Account `json:"account" db:"account"`

		AccountId int64 `json:"-" db:"account_id"`

		Content *string `json:"content"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)