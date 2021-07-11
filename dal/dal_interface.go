package dal

import "github.com/touchps/hackernews/types"

type Dal interface {
	SaveLink(link types.Link) (int64, error)
	GetLinks() ([]types.Link, error)
}
