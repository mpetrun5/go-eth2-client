// Code generated by fastssz. DO NOT EDIT.
// Hash: 53f7a78b79c00bbf864316c122774dbb8e8aae2e6c5245947a816bca4690d094
// Version: 0.1.3
package deneb

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlindedBlockContents object
func (b *BlindedBlockContents) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlindedBlockContents object to a target array
func (b *BlindedBlockContents) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'BlindedBlock'
	dst = ssz.WriteOffset(dst, offset)
	if b.BlindedBlock == nil {
		b.BlindedBlock = new(BlindedBeaconBlock)
	}
	offset += b.BlindedBlock.SizeSSZ()

	// Offset (1) 'BlindedBlobSidecars'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.BlindedBlobSidecars) * 216

	// Field (0) 'BlindedBlock'
	if dst, err = b.BlindedBlock.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'BlindedBlobSidecars'
	if size := len(b.BlindedBlobSidecars); size > 4 {
		err = ssz.ErrListTooBigFn("BlindedBlockContents.BlindedBlobSidecars", size, 4)
		return
	}
	for ii := 0; ii < len(b.BlindedBlobSidecars); ii++ {
		if dst, err = b.BlindedBlobSidecars[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BlindedBlockContents object
func (b *BlindedBlockContents) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'BlindedBlock'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'BlindedBlobSidecars'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'BlindedBlock'
	{
		buf = tail[o0:o1]
		if b.BlindedBlock == nil {
			b.BlindedBlock = new(BlindedBeaconBlock)
		}
		if err = b.BlindedBlock.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'BlindedBlobSidecars'
	{
		buf = tail[o1:]
		num, err := ssz.DivideInt2(len(buf), 216, 4)
		if err != nil {
			return err
		}
		b.BlindedBlobSidecars = make([]*BlindedBlobSidecar, num)
		for ii := 0; ii < num; ii++ {
			if b.BlindedBlobSidecars[ii] == nil {
				b.BlindedBlobSidecars[ii] = new(BlindedBlobSidecar)
			}
			if err = b.BlindedBlobSidecars[ii].UnmarshalSSZ(buf[ii*216 : (ii+1)*216]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlindedBlockContents object
func (b *BlindedBlockContents) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'BlindedBlock'
	if b.BlindedBlock == nil {
		b.BlindedBlock = new(BlindedBeaconBlock)
	}
	size += b.BlindedBlock.SizeSSZ()

	// Field (1) 'BlindedBlobSidecars'
	size += len(b.BlindedBlobSidecars) * 216

	return
}

// HashTreeRoot ssz hashes the BlindedBlockContents object
func (b *BlindedBlockContents) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlindedBlockContents object with a hasher
func (b *BlindedBlockContents) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'BlindedBlock'
	if err = b.BlindedBlock.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'BlindedBlobSidecars'
	{
		subIndx := hh.Index()
		num := uint64(len(b.BlindedBlobSidecars))
		if num > 4 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.BlindedBlobSidecars {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlindedBlockContents object
func (b *BlindedBlockContents) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
