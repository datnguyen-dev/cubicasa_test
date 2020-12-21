package postgresdb

import (
	"fmt"

	"datnguyen.cubicasa.test/common"
	"datnguyen.cubicasa.test/store"
	"github.com/jinzhu/gorm"
)

type hub struct {
	db *gorm.DB
}

//AddHub - Add hub into databases
func (h *hub) AddHub(hub *store.Hub) (string, error) {
	if hub.HubID == "" {
		hub.HubID = common.GenUID()
	} else {
		var hu *store.Hub
		h.db.First(&hu, "hub_id = ?", hub.HubID)
		if hu != nil {
			return "", fmt.Errorf("Duplicate")
		}
	}
	h.db.Create(&hub)
	return hub.HubID, nil
}

//UpdateHub - Update existed Hub
func (h *hub) UpdateHub(hub *store.Hub) (bool, error) {
	if hub.HubID == "" {
		return false, fmt.Errorf("NotFound")
	}
	var hu store.Hub
	h.db.First(&hu, "hub_id = ?", hub.HubID)
	if &hu == nil {
		return false, fmt.Errorf("NotFound")
	}

	h.db.Model(hu).Updates(hub)
	return true, nil
}

//DeleteHub - Delete existed Hub
func (h *hub) DeleteHub(id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("NotFound")
	}
	var hu store.Hub
	h.db.First(&hu, "hub_id = ?", id)
	if &hu == nil {
		return false, fmt.Errorf("NotFound")
	}
	h.db.Delete(hu, "hub_id = ?", id)
	return true, nil
}
