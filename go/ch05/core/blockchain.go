package core

// *
type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

// *
func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)

	err := bc.addBlockWithoutValidation(genesis)
	return bc, err
}

// *
func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

// *
func (bc *Blockchain) AddBlock(b *Block) error {
	return nil
}

// *
func (bc *Blockchain) HasBlock(height uint32) bool {
	return height < bc.Height()
}

// *
// 블록 최대 높이는 저장된 마지막 헤더의 인덱스
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers)) - 1
}

// *
func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	return bc.store.Put(b)
}
