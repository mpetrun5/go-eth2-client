// Code generated by fastssz. DO NOT EDIT.
// Hash: cb2f1c63717d30324163608065e759ce532e3cf10373eb26de9e3dc4f0553100
// Version: 0.1.3
package deneb

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedBlindedBlockContents object
func (s *SignedBlindedBlockContents) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedBlindedBlockContents object to a target array
func (s *SignedBlindedBlockContents) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'SignedBlindedBlock'
	dst = ssz.WriteOffset(dst, offset)
	if s.SignedBlindedBlock == nil {
		s.SignedBlindedBlock = new(SignedBlindedBeaconBlock)
	}
	offset += s.SignedBlindedBlock.SizeSSZ()

	// Offset (1) 'SignedBlindedBlobSidecars'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.SignedBlindedBlobSidecars) * 312

	// Field (0) 'SignedBlindedBlock'
	if dst, err = s.SignedBlindedBlock.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'SignedBlindedBlobSidecars'
	if size := len(s.SignedBlindedBlobSidecars); size > 6 {
		err = ssz.ErrListTooBigFn("SignedBlindedBlockContents.SignedBlindedBlobSidecars", size, 6)
		return
	}
	for ii := 0; ii < len(s.SignedBlindedBlobSidecars); ii++ {
		if dst, err = s.SignedBlindedBlobSidecars[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedBlindedBlockContents object
func (s *SignedBlindedBlockContents) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'SignedBlindedBlock'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'SignedBlindedBlobSidecars'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'SignedBlindedBlock'
	{
		buf = tail[o0:o1]
		if s.SignedBlindedBlock == nil {
			s.SignedBlindedBlock = new(SignedBlindedBeaconBlock)
		}
		if err = s.SignedBlindedBlock.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'SignedBlindedBlobSidecars'
	{
		buf = tail[o1:]
		num, err := ssz.DivideInt2(len(buf), 312, 6)
		if err != nil {
			return err
		}
		s.SignedBlindedBlobSidecars = make([]*SignedBlindedBlobSidecar, num)
		for ii := 0; ii < num; ii++ {
			if s.SignedBlindedBlobSidecars[ii] == nil {
				s.SignedBlindedBlobSidecars[ii] = new(SignedBlindedBlobSidecar)
			}
			if err = s.SignedBlindedBlobSidecars[ii].UnmarshalSSZ(buf[ii*312 : (ii+1)*312]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBlindedBlockContents object
func (s *SignedBlindedBlockContents) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'SignedBlindedBlock'
	if s.SignedBlindedBlock == nil {
		s.SignedBlindedBlock = new(SignedBlindedBeaconBlock)
	}
	size += s.SignedBlindedBlock.SizeSSZ()

	// Field (1) 'SignedBlindedBlobSidecars'
	size += len(s.SignedBlindedBlobSidecars) * 312

	return
}

// HashTreeRoot ssz hashes the SignedBlindedBlockContents object
func (s *SignedBlindedBlockContents) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedBlindedBlockContents object with a hasher
func (s *SignedBlindedBlockContents) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SignedBlindedBlock'
	if err = s.SignedBlindedBlock.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'SignedBlindedBlobSidecars'
	{
		subIndx := hh.Index()
		num := uint64(len(s.SignedBlindedBlobSidecars))
		if num > 6 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.SignedBlindedBlobSidecars {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 6)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedBlindedBlockContents object
func (s *SignedBlindedBlockContents) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
