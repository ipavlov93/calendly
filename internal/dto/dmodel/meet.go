package dmodel

import (
	"database/sql"
	"event-calendar/internal/domain"
	"time"
)

type Meet struct {
	ID          int64             `db:"id"`
	Title       string            `db:"title"`
	Status      domain.MeetStatus `db:"status"`
	From        time.Time         `db:"from"`
	To          time.Time         `db:"to"`
	OrganizerID int64             `db:"organizer_id"`
	CreatedBy   int64             `db:"created_by"`
	Description string            `db:"description"`
	Link        sql.NullString    `db:"link"`
	CreatedAt   time.Time         `db:"created_at"`
}

// NewMeet init meet with given fields.
// Set status SCHEDULED.
func NewMeet(
	title string,
	from, to time.Time,
	description string,
	creatorID int64,
	organizerID int64,
) Meet {
	return newMeet(
		title,
		domain.CREATED,
		from, to,
		description,
		creatorID,
		organizerID,
	)
}

// NewScheduledMeet init meet with given fields.
// Set status CREATED.
func NewScheduledMeet(
	title string,
	from, to time.Time,
	description string,
	creatorID int64,
	organizerID int64,
) Meet {
	return newMeet(
		title,
		domain.SCHEDULED,
		from, to,
		description,
		creatorID,
		organizerID,
	)
}

// newMeet init meet with given fields.
func newMeet(
	title string,
	status domain.MeetStatus,
	from, to time.Time,
	description string,
	creatorID int64,
	organizerID int64,
) Meet {
	meet := Meet{
		Title:       title,
		Status:      status,
		From:        from,
		To:          to,
		Description: description,
		OrganizerID: organizerID,
		CreatedBy:   creatorID,
	}
	return meet
}
