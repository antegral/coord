// Code generated by "stringer -type=SegmentType -linecomment"; DO NOT EDIT.

package llm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SegmentTypeUnknown-0]
	_ = x[SegmentTypeText-1]
	_ = x[SegmentTypeInlineData-2]
	_ = x[SegmentTypeFunctionCall-3]
	_ = x[SegmentTypeFunctionResponse-4]
}

const _SegmentType_name = "unknowntextinline_datafunction_callfunction_response"

var _SegmentType_index = [...]uint8{0, 7, 11, 22, 35, 52}

func (i SegmentType) String() string {
	if i >= SegmentType(len(_SegmentType_index)-1) {
		return "SegmentType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SegmentType_name[_SegmentType_index[i]:_SegmentType_index[i+1]]
}
