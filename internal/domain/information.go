package domain

const (
	// ServiceMEMBER is a constant for member service
	ServiceMEMBER = "member"
)

type Information struct {
	Member []*Member
}

func (i *Information) AddMember(member *Member) {
	// add member
	i.Member = append(i.Member, member)
}
