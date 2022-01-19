package main

import (
	"fmt"

	"github.com/rahulrishireddy/calculator/Operations"
)

func main() {

	Operations.AddInt8(100, 20)
	Operations.AddInt16(-130, 150)
	Operations.AddInt32(-180, 200)
	Operations.AddInt64(293, 298)
	Operations.AddUint8(100, 20)
	Operations.AddUint16(100, 150)
	Operations.AddUint32(200, 220)
	Operations.AddUint64(170, 220)
	Operations.AddFloat32(220.24, 210.11)
	Operations.AddFloat64(132.24, 156.222)

	Operations.SubInt8(100, 120)
	Operations.SubInt16(-130, 150)
	Operations.SubInt32(-180, 200)
	Operations.SubInt64(293, 298)
	Operations.SubUint8(100, 20)
	Operations.SubUint16(100, 150)
	Operations.SubUint32(200, 220)
	Operations.SubUint64(170, 220)
	Operations.SubFloat32(220.24, 210.11)
	Operations.SubFloat64(132.24, 156.222)

	Operations.MulInt8(100, 120)
	Operations.MulInt16(-130, 150)
	Operations.MulInt32(-180, 200)
	Operations.MulInt64(293, 298)
	Operations.MulUint8(100, 20)
	Operations.MulUint16(100, 150)
	Operations.MulUint32(200, 220)
	Operations.MulUint64(170, 220)
	Operations.MulFloat32(220.24, 210.11)
	Operations.MulFloat64(132.24, 156.222)

	Operations.DivisionInt8(100, 120)
	Operations.DivisionInt16(-130, 150)
	Operations.DivisionInt32(-180, 200)
	Operations.DivisionInt64(293, 298)
	Operations.DivisionUint8(100, 20)
	Operations.DivisionUint16(100, 150)
	Operations.DivisionUint32(200, 220)
	Operations.DivisionUint64(170, 220)
	Operations.DivisionFloat32(220.24, 210.11)
	Operations.DivisionFloat64(132.24, 156.222)

	fmt.Println("program is finished")
}
