package shite

import "math/rand"

// const charset = "abcdefghijklmnopqrstuvwxyz" +
//
//	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
type length struct {
	less int
	more int
}

func (l length) Rand() int {
	delta := rand.Intn(l.more - l.less)
	return l.less + delta
}

type charSet struct {
	set     []rune
	lengths length
}

var (
	SmallLetter = []rune("abcdefghijklmnopqrstuvwxyz")
	BigLetter   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Number      = []rune("012345678")
)

func StringWithCharset(cSet *charSet) []rune {
	var l = cSet.lengths.less
	if cSet.lengths.more != 0 {
		l = cSet.lengths.Rand()
	}
	b := make([]rune, l)
	set := cSet.set
	for i := range b {
		b[i] = set[rand.Intn(len(set))]
	}
	return b
}
func (s *Shite) String() string {
	res := make([]rune, 0)
	for _, charset := range s.containers {
		res = append(res, StringWithCharset(charset)...)
	}
	return string(res)
}

type Shite struct {
	containers []*charSet
}

func NewShite() *Shite {
	return &Shite{
		containers: make([]*charSet, 0),
	}
}
func NewItem(list []rune, lens ...int) *charSet {
	cl := make([]rune, len(list))
	for i, c := range list {
		cl[i] = c
	}
	var less, more int
	if len(lens) == 1 {
		less = lens[0]
	} else if len(lens) == 2 {
		less, more = lens[0], lens[1]
	} else {
		panic("err")
	}
	return &charSet{
		set: cl,
		lengths: length{
			less: less,
			more: more,
		},
	}
}
func (s *Shite) Add(set *charSet) {
	s.containers = append(s.containers, set)
}
