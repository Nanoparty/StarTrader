package menus

import (
	"time"
	"startrader/globals"
)

// StartMissionTimers should be called once at game start
func StartMissionTimers() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			updateActiveMissionTimers()
		}
	}()
}

func updateActiveMissionTimers() {
	for _, ship := range globals.Company.Ships {
		mission := ship.AssignedMission
		if mission != nil && mission.Status == "In Progress" {
			if mission.Minutes > 0 || mission.Seconds > 0 {
				if mission.Seconds == 0 {
					if mission.Minutes > 0 {
						mission.Minutes--
						mission.Seconds = 59
					}
				} else {
					mission.Seconds--
				}
				if mission.Minutes == 0 && mission.Seconds == 0 {
					mission.Status = "Complete"
				}
			}
		}
	}
}
