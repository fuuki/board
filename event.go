package board

type EventType string

const (
	EventTypePhaseChange EventType = "phase_change"
)

type Event struct {
	EventType   EventType    `json:"event_type"`
	PhaseChange *PhaseChange `json:"phase_change,omitempty"`
}

type PhaseChange struct {
	PhaseCount int `json:"phase_count"`
}

type eventChan struct {
	ch chan *Event
}

// newEventChan returns a new event channel.
func newEventChan() (<-chan *Event, *eventChan) {
	ch := make(chan *Event)
	return ch, &eventChan{
		ch: ch,
	}
}

// sendPhaseChange sends a phase change event.
func (ec *eventChan) sendPhaseChange(phaseCount int) {
	ec.ch <- &Event{
		EventType: EventTypePhaseChange,
		PhaseChange: &PhaseChange{
			PhaseCount: phaseCount,
		},
	}
}

// close closes the channel.
func (ec *eventChan) close() {
	close(ec.ch)
}
