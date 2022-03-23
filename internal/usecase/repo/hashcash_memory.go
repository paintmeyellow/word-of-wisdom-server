package repo

type HashcashRepo struct {
	Hashes map[string]struct{}
}

func NewHashcashRepo() *HashcashRepo {
	return &HashcashRepo{
		Hashes: make(map[string]struct{}),
	}
}

func (r *HashcashRepo) Store(hash string) {
	if _, ok := r.Hashes[hash]; !ok {
		r.Hashes[hash] = struct{}{}
	}
}

func (r *HashcashRepo) Exists(hash string) bool {
	_, ok := r.Hashes[hash]
	return ok
}
