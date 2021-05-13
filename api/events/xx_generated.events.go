// This file has been automatically generated. Don't edit it.

package events

func GetEventForType(name string) Event {
	switch name {
	case "Heartbeat":
		return &Heartbeat{}
	case "Exiting":
		return &Exiting{}
	case "ProfileChanged":
		return &ProfileChanged{}
	case "ProfileListChanged":
		return &ProfileListChanged{}
	case "RecordingStarting":
		return &RecordingStarting{}
	case "RecordingStarted":
		return &RecordingStarted{}
	case "RecordingStopping":
		return &RecordingStopping{}
	case "RecordingStopped":
		return &RecordingStopped{}
	case "ReplayStarting":
		return &ReplayStarting{}
	case "ReplayStarted":
		return &ReplayStarted{}
	case "ReplayStopping":
		return &ReplayStopping{}
	case "ReplayStopped":
		return &ReplayStopped{}
	case "SwitchScenes":
		return &SwitchScenes{}
	case "ScenesChanged":
		return &ScenesChanged{}
	case "SceneCollectionChanged":
		return &SceneCollectionChanged{}
	case "SceneCollectionListChanged":
		return &SceneCollectionListChanged{}
	case "SourceOrderChanged":
		return &SourceOrderChanged{}
	case "SceneItemAdded":
		return &SceneItemAdded{}
	case "SceneItemRemoved":
		return &SceneItemRemoved{}
	case "SceneItemVisibilityChanged":
		return &SceneItemVisibilityChanged{}
	case "StreamStarting":
		return &StreamStarting{}
	case "StreamStarted":
		return &StreamStarted{}
	case "StreamStopping":
		return &StreamStopping{}
	case "StreamStopped":
		return &StreamStopped{}
	case "StreamStatus":
		return &StreamStatus{}
	case "PreviewSceneChanged":
		return &PreviewSceneChanged{}
	case "StudioModeSwitched":
		return &StudioModeSwitched{}
	case "SwitchTransition":
		return &SwitchTransition{}
	case "TransitionListChanged":
		return &TransitionListChanged{}
	case "TransitionDurationChanged":
		return &TransitionDurationChanged{}
	case "TransitionBegin":
		return &TransitionBegin{}
	default:
		return nil
	}
}