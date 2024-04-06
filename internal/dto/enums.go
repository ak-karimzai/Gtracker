package dto

// CompletionStatus
// @Description A pre-defined choices for goal status
type CompletionStatus string

const (
	Progress  CompletionStatus = "Progress"
	Completed CompletionStatus = "Completed"
	Skipped   CompletionStatus = "Skipped"
)

func (cp CompletionStatus) Validate() bool {
	return cp == Progress || cp == Completed || cp == Skipped
}

// Frequency
// @Description A pre-defined choices for task frequency
type Frequency string

const (
	Daily   Frequency = "Daily"
	Weekly  Frequency = "Weekly"
	Monthly Frequency = "Monthly"
)

func (freq Frequency) Validate() bool {
	return freq == Daily || freq == Weekly || freq == Monthly
}
