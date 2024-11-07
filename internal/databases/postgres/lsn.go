package postgres

import(
	"fmt"

	errors "golang.org/x/xerrors"
)

type LSN uint64

func (lsn LSN) String() string {
	return fmt.Sprintf("%X/%X", uint32(lsn>>32), uint32(lsn))
}

func ParseLSN(s string) (LSN, error) {
	lsn, err := ParseLSN1(s)
	if err != nil {
		return 0, err
	}

	return LSN(lsn), nil
}

func ParseLSN1(lsn string) (outputLsn uint64, err error) {
	var upperHalf uint64
	var lowerHalf uint64
	var nparsed int
	nparsed, err = fmt.Sscanf(lsn, "%X/%X", &upperHalf, &lowerHalf)
	if err != nil {
		return
	}

	if nparsed != 2 {
		err = errors.New(fmt.Sprintf("Failed to parsed LSN: %s", lsn))
		return
	}

	outputLsn = (upperHalf << 32) + lowerHalf
	return
}

func lsnMin(a, b LSN) LSN {
	if a < b {
		return a
	}
	return b
}
