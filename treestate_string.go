// Code generated by "stringer -type=TreeState"; DO NOT EDIT.

package gitstatus

import "strconv"

const _TreeState_name = "DefaultRebasingAMAMRebaseMergingCherryPickingRevertingBisecting"

var _TreeState_index = [...]uint8{0, 7, 15, 17, 25, 32, 45, 54, 63}

func (i TreeState) String() string {
	if i < 0 || i >= TreeState(len(_TreeState_index)-1) {
		return "TreeState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TreeState_name[_TreeState_index[i]:_TreeState_index[i+1]]
}
