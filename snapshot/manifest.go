package snapshot

import (
	"time"

	"github.com/kopia/kopia/repo"
)

// Manifest represents information about a single point-in-time filesystem snapshot.
type Manifest struct {
	Source SourceInfo `json:"source"`

	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`

	Handle       string        `json:"handle"`
	RootObjectID repo.ObjectID `json:"root"`
	HashCacheID  repo.ObjectID `json:"hashCache"`

	Stats Stats `json:"stats"`
}
