package tcp

import "encoding/binary"

/*
  TCP POC-service-profile Option:

         Kind: 10  Length: 3 bytes

                                       1 bit        1 bit    6 bits
             +----------+----------+------------+----------+--------+
             |  Kind=10 | Length=3 | Start_flag | End_flag | Filler |
             +----------+----------+------------+----------+--------+

*/
type optionPartialOrderServiceProfile []byte

func (o optionPartialOrderServiceProfile) Length() byte { return o[0] }
func (o optionPartialOrderServiceProfile) PartialOrderServiceProfile() uint16 {
	return binary.BigEndian.Uint16(o[1:])
	//
}
func (o optionPartialOrderServiceProfile) NextOption() []byte { return o[3:] }

func (o optionPartialOrderServiceProfile) Process(s *Socket) error {
	return nil
}
