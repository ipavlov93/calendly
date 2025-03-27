package domain

import "fmt"

type MeetStatus string

const (
	CREATED  MeetStatus = "Created"
	CANCELED MeetStatus = "Canceled"
	//DEACTIVATED MeetStatus = "Deactivated"

	SCHEDULED   MeetStatus = "Scheduled"
	IN_PROGRESS MeetStatus = "InProgress"
	COMPLETED   MeetStatus = "Completed"
)

var stateName = map[MeetStatus]string{
	CREATED:     "Created",
	CANCELED:    "Canceled",
	SCHEDULED:   "Scheduled",
	IN_PROGRESS: "InProgress",
	COMPLETED:   "Completed",
}

func (m MeetStatus) String() string {
	s, ok := stateName[m]
	if !ok {
		return fmt.Sprintf("MeetStatus(%s)", string(m))
	}
	return s
}

func (m MeetStatus) transition(meetStatus MeetStatus) (MeetStatus, error) {
	_, ok := stateName[m]
	if !ok {
		panic(fmt.Errorf("unknown meet status: %s", meetStatus))
	}

	switch m {
	case CREATED:
		if meetStatus == CANCELED || meetStatus == SCHEDULED {
			return meetStatus, nil
		}
	case SCHEDULED:
		if meetStatus == CANCELED || meetStatus == IN_PROGRESS {
			return meetStatus, nil
		}
	case IN_PROGRESS:
		if meetStatus == COMPLETED {
			return meetStatus, nil
		}
	default:
		return m, fmt.Errorf("can't set meet status %s to %s", m, meetStatus)
	}
	return meetStatus, nil
}
