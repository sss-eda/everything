package serial

type ErrStationRecording struct{}

func (err ErrStationRecording) Error() string {
	return "cannot execute this command while the station is recording"
}
