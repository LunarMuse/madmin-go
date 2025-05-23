package madmin

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"time"

	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BucketScanInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	var zb0001Mask uint8 /* 1 bits */
	_ = zb0001Mask
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "pool":
			z.Pool, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Pool")
				return
			}
		case "set":
			z.Set, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Set")
				return
			}
		case "cycle":
			z.Cycle, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Cycle")
				return
			}
		case "ongoing":
			z.Ongoing, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Ongoing")
				return
			}
		case "last_update":
			z.LastUpdate, err = dc.ReadTimeUTC()
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "last_started":
			z.LastStarted, err = dc.ReadTimeUTC()
			if err != nil {
				err = msgp.WrapError(err, "LastStarted")
				return
			}
		case "last_error":
			z.LastError, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "LastError")
				return
			}
		case "completed":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Completed")
				return
			}
			if cap(z.Completed) >= int(zb0002) {
				z.Completed = (z.Completed)[:zb0002]
			} else {
				z.Completed = make([]time.Time, zb0002)
			}
			for za0001 := range z.Completed {
				z.Completed[za0001], err = dc.ReadTimeUTC()
				if err != nil {
					err = msgp.WrapError(err, "Completed", za0001)
					return
				}
			}
			zb0001Mask |= 0x1
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	// Clear omitted fields.
	if (zb0001Mask & 0x1) == 0 {
		z.Completed = nil
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z *BucketScanInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// check for omitted fields
	zb0001Len := uint32(8)
	var zb0001Mask uint8 /* 8 bits */
	_ = zb0001Mask
	if z.Completed == nil {
		zb0001Len--
		zb0001Mask |= 0x80
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}

	// skip if no fields are to be emitted
	if zb0001Len != 0 {
		// write "pool"
		err = en.Append(0xa4, 0x70, 0x6f, 0x6f, 0x6c)
		if err != nil {
			return
		}
		err = en.WriteInt(z.Pool)
		if err != nil {
			err = msgp.WrapError(err, "Pool")
			return
		}
		// write "set"
		err = en.Append(0xa3, 0x73, 0x65, 0x74)
		if err != nil {
			return
		}
		err = en.WriteInt(z.Set)
		if err != nil {
			err = msgp.WrapError(err, "Set")
			return
		}
		// write "cycle"
		err = en.Append(0xa5, 0x63, 0x79, 0x63, 0x6c, 0x65)
		if err != nil {
			return
		}
		err = en.WriteUint64(z.Cycle)
		if err != nil {
			err = msgp.WrapError(err, "Cycle")
			return
		}
		// write "ongoing"
		err = en.Append(0xa7, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67)
		if err != nil {
			return
		}
		err = en.WriteBool(z.Ongoing)
		if err != nil {
			err = msgp.WrapError(err, "Ongoing")
			return
		}
		// write "last_update"
		err = en.Append(0xab, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65)
		if err != nil {
			return
		}
		err = en.WriteTime(z.LastUpdate)
		if err != nil {
			err = msgp.WrapError(err, "LastUpdate")
			return
		}
		// write "last_started"
		err = en.Append(0xac, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64)
		if err != nil {
			return
		}
		err = en.WriteTime(z.LastStarted)
		if err != nil {
			err = msgp.WrapError(err, "LastStarted")
			return
		}
		// write "last_error"
		err = en.Append(0xaa, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72)
		if err != nil {
			return
		}
		err = en.WriteString(z.LastError)
		if err != nil {
			err = msgp.WrapError(err, "LastError")
			return
		}
		if (zb0001Mask & 0x80) == 0 { // if not omitted
			// write "completed"
			err = en.Append(0xa9, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64)
			if err != nil {
				return
			}
			err = en.WriteArrayHeader(uint32(len(z.Completed)))
			if err != nil {
				err = msgp.WrapError(err, "Completed")
				return
			}
			for za0001 := range z.Completed {
				err = en.WriteTime(z.Completed[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Completed", za0001)
					return
				}
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BucketScanInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// check for omitted fields
	zb0001Len := uint32(8)
	var zb0001Mask uint8 /* 8 bits */
	_ = zb0001Mask
	if z.Completed == nil {
		zb0001Len--
		zb0001Mask |= 0x80
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))

	// skip if no fields are to be emitted
	if zb0001Len != 0 {
		// string "pool"
		o = append(o, 0xa4, 0x70, 0x6f, 0x6f, 0x6c)
		o = msgp.AppendInt(o, z.Pool)
		// string "set"
		o = append(o, 0xa3, 0x73, 0x65, 0x74)
		o = msgp.AppendInt(o, z.Set)
		// string "cycle"
		o = append(o, 0xa5, 0x63, 0x79, 0x63, 0x6c, 0x65)
		o = msgp.AppendUint64(o, z.Cycle)
		// string "ongoing"
		o = append(o, 0xa7, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67)
		o = msgp.AppendBool(o, z.Ongoing)
		// string "last_update"
		o = append(o, 0xab, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65)
		o = msgp.AppendTime(o, z.LastUpdate)
		// string "last_started"
		o = append(o, 0xac, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64)
		o = msgp.AppendTime(o, z.LastStarted)
		// string "last_error"
		o = append(o, 0xaa, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72)
		o = msgp.AppendString(o, z.LastError)
		if (zb0001Mask & 0x80) == 0 { // if not omitted
			// string "completed"
			o = append(o, 0xa9, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64)
			o = msgp.AppendArrayHeader(o, uint32(len(z.Completed)))
			for za0001 := range z.Completed {
				o = msgp.AppendTime(o, z.Completed[za0001])
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BucketScanInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	var zb0001Mask uint8 /* 1 bits */
	_ = zb0001Mask
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "pool":
			z.Pool, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Pool")
				return
			}
		case "set":
			z.Set, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Set")
				return
			}
		case "cycle":
			z.Cycle, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Cycle")
				return
			}
		case "ongoing":
			z.Ongoing, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ongoing")
				return
			}
		case "last_update":
			z.LastUpdate, bts, err = msgp.ReadTimeUTCBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "last_started":
			z.LastStarted, bts, err = msgp.ReadTimeUTCBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastStarted")
				return
			}
		case "last_error":
			z.LastError, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastError")
				return
			}
		case "completed":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Completed")
				return
			}
			if cap(z.Completed) >= int(zb0002) {
				z.Completed = (z.Completed)[:zb0002]
			} else {
				z.Completed = make([]time.Time, zb0002)
			}
			for za0001 := range z.Completed {
				z.Completed[za0001], bts, err = msgp.ReadTimeUTCBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Completed", za0001)
					return
				}
			}
			zb0001Mask |= 0x1
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	// Clear omitted fields.
	if (zb0001Mask & 0x1) == 0 {
		z.Completed = nil
	}

	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BucketScanInfo) Msgsize() (s int) {
	s = 1 + 5 + msgp.IntSize + 4 + msgp.IntSize + 6 + msgp.Uint64Size + 8 + msgp.BoolSize + 12 + msgp.TimeSize + 13 + msgp.TimeSize + 11 + msgp.StringPrefixSize + len(z.LastError) + 10 + msgp.ArrayHeaderSize + (len(z.Completed) * (msgp.TimeSize))
	return
}
