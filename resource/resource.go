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
func (rp *ResourceProfile) Show() {
	for p, r := range rp.resources {
		fmt.Printf("Player %d: %d\n", p, r.point)
	}
}

func (rp *ResourceProfile) Equal(rp2 *ResourceProfile) bool {
	for p, r := range rp.resources {
		if r.point != rp2.resources[p].point {
			return false
		}
	}
	return true
}
