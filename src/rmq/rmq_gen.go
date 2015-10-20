package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/glycerine/msgp)
// DO NOT EDIT

import (
	"github.com/glycerine/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Payload) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Sub":
			var isz uint32
			isz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for isz > 0 {
				isz--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "A":
					z.Sub.A, err = dc.ReadString()
					if err != nil {
						return
					}
				case "B":
					z.Sub.B, err = dc.ReadInt()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "D":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.D) >= int(xsz) {
				z.D = z.D[:xsz]
			} else {
				z.D = make([]string, xsz)
			}
			for xvk := range z.D {
				z.D[xvk], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "E":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.E) >= int(xsz) {
				z.E = z.E[:xsz]
			} else {
				z.E = make([]int32, xsz)
			}
			for bzg := range z.E {
				z.E[bzg], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Payload) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Sub"
	// map header, size 2
	// write "A"
	err = en.Append(0x83, 0xa3, 0x53, 0x75, 0x62, 0x82, 0xa1, 0x41)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sub.A)
	if err != nil {
		return
	}
	// write "B"
	err = en.Append(0xa1, 0x42)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Sub.B)
	if err != nil {
		return
	}
	// write "D"
	err = en.Append(0xa1, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.D)))
	if err != nil {
		return
	}
	for xvk := range z.D {
		err = en.WriteString(z.D[xvk])
		if err != nil {
			return
		}
	}
	// write "E"
	err = en.Append(0xa1, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.E)))
	if err != nil {
		return
	}
	for bzg := range z.E {
		err = en.WriteInt32(z.E[bzg])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Payload) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Sub"
	// map header, size 2
	// string "A"
	o = append(o, 0x83, 0xa3, 0x53, 0x75, 0x62, 0x82, 0xa1, 0x41)
	o = msgp.AppendString(o, z.Sub.A)
	// string "B"
	o = append(o, 0xa1, 0x42)
	o = msgp.AppendInt(o, z.Sub.B)
	// string "D"
	o = append(o, 0xa1, 0x44)
	o = msgp.AppendArrayHeader(o, uint32(len(z.D)))
	for xvk := range z.D {
		o = msgp.AppendString(o, z.D[xvk])
	}
	// string "E"
	o = append(o, 0xa1, 0x45)
	o = msgp.AppendArrayHeader(o, uint32(len(z.E)))
	for bzg := range z.E {
		o = msgp.AppendInt32(o, z.E[bzg])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Payload) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Sub":
			var isz uint32
			isz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for isz > 0 {
				isz--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "A":
					z.Sub.A, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "B":
					z.Sub.B, bts, err = msgp.ReadIntBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "D":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.D) >= int(xsz) {
				z.D = z.D[:xsz]
			} else {
				z.D = make([]string, xsz)
			}
			for xvk := range z.D {
				z.D[xvk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "E":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.E) >= int(xsz) {
				z.E = z.E[:xsz]
			} else {
				z.E = make([]int32, xsz)
			}
			for bzg := range z.E {
				z.E[bzg], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Payload) Msgsize() (s int) {
	s = 1 + 4 + 1 + 2 + msgp.StringPrefixSize + len(z.Sub.A) + 2 + msgp.IntSize + 2 + msgp.ArrayHeaderSize
	for xvk := range z.D {
		s += msgp.StringPrefixSize + len(z.D[xvk])
	}
	s += 2 + msgp.ArrayHeaderSize + (len(z.E) * (msgp.Int32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Subload) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "A":
			z.A, err = dc.ReadString()
			if err != nil {
				return
			}
		case "B":
			z.B, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Subload) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "A"
	err = en.Append(0x82, 0xa1, 0x41)
	if err != nil {
		return err
	}
	err = en.WriteString(z.A)
	if err != nil {
		return
	}
	// write "B"
	err = en.Append(0xa1, 0x42)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.B)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Subload) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "A"
	o = append(o, 0x82, 0xa1, 0x41)
	o = msgp.AppendString(o, z.A)
	// string "B"
	o = append(o, 0xa1, 0x42)
	o = msgp.AppendInt(o, z.B)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Subload) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "A":
			z.A, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "B":
			z.B, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z Subload) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.A) + 2 + msgp.IntSize
	return
}
