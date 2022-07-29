package gakujo

type AssignmentKind string

const (
	AssignmentKindReport   = AssignmentKind("レポート")
	AssignmentKindMinitest = AssignmentKind("小テスト")
)

func NewAssignmentKind(s string) AssignmentKind {
	switch s {
	case "レポート":
		return AssignmentKindReport
	case "小テスト":
		return AssignmentKindMinitest
	default:
		return AssignmentKind("")
	}
}

type Semester string

const (
	SemesterFirst  = Semester("前期")
	SemesterSecond = Semester("後期")
)

func NewSemester(s string) Semester {
	switch s {
	case "前期":
		return SemesterFirst
	case "後期":
		return SemesterSecond
	default:
		return Semester("")
	}
}

type AssignmentStatus string

const (
	AssignmentStatusOpen      = "受付中"
	AssignmentStatusClose     = "締切"
	AssignmentStatusUndefined = "未定義" // undefined!
)

func NewAssignmentStatus(s string) AssignmentStatus {
	switch s {
	case "受付中":
		return AssignmentStatusOpen
	case "締切":
		return AssignmentStatusClose
	default:
		return AssignmentStatusUndefined
	}
}
