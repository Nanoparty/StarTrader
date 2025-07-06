package utils

import (
	"startrader/globals"
	"time"
)

// StartContractTimers should be called once at game start
func StartContractTimers() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			updateActiveContractTimers()
		}
	}()
}

func updateActiveContractTimers() {
	for _, ship := range globals.Company.Ships {
		contract := ship.AssignedContract
		if contract != nil && contract.Status == "In Progress" {
			if contract.Minutes > 0 || contract.Seconds > 0 {
				if contract.Seconds == 0 {
					if contract.Minutes > 0 {
						contract.Minutes--
						contract.Seconds = 59
					}
				} else {
					contract.Seconds--
				}
				if contract.Minutes == 0 && contract.Seconds == 0 {
					contract.Status = "Complete"
				}
			}
		}
	}
}
