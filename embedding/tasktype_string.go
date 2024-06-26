// Code generated by "stringer -type=TaskType"; DO NOT EDIT.

package embedding

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TaskTypeGeneral-0]
	_ = x[TaskTypeSearchQuery-1]
	_ = x[TaskTypeSearchDocument-2]
	_ = x[TaskTypeSemanticSimilarity-3]
	_ = x[TaskTypeClassification-4]
	_ = x[TaskTypeClustering-5]
	_ = x[TaskTypeFactVerification-6]
	_ = x[TaskTypeQA-7]
}

const _TaskType_name = "TaskTypeGeneralTaskTypeSearchQueryTaskTypeSearchDocumentTaskTypeSemanticSimilarityTaskTypeClassificationTaskTypeClusteringTaskTypeFactVerificationTaskTypeQA"

var _TaskType_index = [...]uint8{0, 15, 34, 56, 82, 104, 122, 146, 156}

func (i TaskType) String() string {
	if i >= TaskType(len(_TaskType_index)-1) {
		return "TaskType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TaskType_name[_TaskType_index[i]:_TaskType_index[i+1]]
}
