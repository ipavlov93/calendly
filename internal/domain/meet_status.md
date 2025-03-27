## Meet status

### Existing statuses

- Created
- Canceled (Rejected, Failed)
- Scheduled
- InProgress
- Completed

### Created

Meet status is set to Created when meet is created (by creator participant).

### Canceled (Rejected, Failed)

Meet status is planned to be divided with the following statuses:
- Canceled (by creator)
- Rejected (by organizer)
- Failed (server failed to create meet link)

Meet could be canceled or rejected only before meet status changed to InProgress.

#### Failed

Potential reasons why meet has been failed. Fail reason should be saved.

- Server failed to create meet's link for the meet before it's planned to start.
- Max attempts number to create meet's link has been reached.

### Scheduled

Meet status is set to Scheduled after these steps:
- Meet's link has been created successfully.
- Meet is added to attendees calendar.
- letters are sent to creator and organizer (participants).
- notification letters are sent to other invited attendees (TODO: create them as participants)

### InProgress

Scheduled status transits to InProgress when From date-time has been reached.

### Completed

InProgress status transits to Completed when To date-time has been reached.
