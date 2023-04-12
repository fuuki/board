package resource

import (
	"fmt"

	"github.com/fuuki/board/player"
)

type Resource struct {
	point int
}

func NewResource() *Resource {
	return &Resource{}
}

func (r *Resource) AddPoint(point int) {
	r.point += point
}

func (r *Resource) Point() int {
	return r.point
}

type ResourceProfile struct {
	resources map[player.Player]*Resource
}

func NewResourceProfile() *ResourceProfile {
	return &ResourceProfile{
		resources: make(map[player.Player]*Resource),
	}
}

func (rp *ResourceProfile) AddResource(player player.Player, resource *Resource) {
	rp.resources[player] = resource
}

func (rp *ResourceProfile) Player(player player.Player) *Resource {
	return rp.resources[player]
}

// Show print all resources
func (rp *ResourceProfile) Show() string {
	s := ""
	for p, r := range rp.resources {
		s += fmt.Sprintf("Player %d: %d pt(s)\n", p, r.point)
	}
	return s
}

func (rp *ResourceProfile) Equal(rp2 *ResourceProfile) bool {
	for p, r := range rp.resources {
		if r.point != rp2.resources[p].point {
			return false
		}
	}
	return true
}
