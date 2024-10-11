package domain

type Information struct {
	Member Member
}

func (i *Information) AddMember(member Member) {
	// add member
	i.Member = member
}
