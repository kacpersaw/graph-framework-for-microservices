// Code generated by "stringer -type=GNSPluginState"; DO NOT EDIT.

package nexus

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GNSPluginState_UNKNOWN-0]
	_ = x[GNSPluginState_IDLE-1]
	_ = x[GNSPluginState_PROCESSING-2]
	_ = x[GNSPluginState_SYNCING-3]
	_ = x[GNSPluginState_SYNCED-4]
	_ = x[GNSPluginState_ERROR-5]
}

const _GNSPluginState_name = "GNSPluginState_UNKNOWNGNSPluginState_IDLEGNSPluginState_PROCESSINGGNSPluginState_SYNCINGGNSPluginState_SYNCEDGNSPluginState_ERROR"

var _GNSPluginState_index = [...]uint8{0, 22, 41, 66, 88, 109, 129}

func (i GNSPluginState) String() string {
	if i < 0 || i >= GNSPluginState(len(_GNSPluginState_index)-1) {
		return "GNSPluginState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GNSPluginState_name[_GNSPluginState_index[i]:_GNSPluginState_index[i+1]]
}
