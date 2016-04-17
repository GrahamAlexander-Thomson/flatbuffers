// automatically generated, do not modify

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
/// an example documentation comment: monster object
type Monster struct {
	_tab flatbuffers.Table
}

func GetRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Monster{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Monster) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Monster) Pos(obj *Vec3) *Vec3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vec3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Mana() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 150
}

func (rcv *Monster) Hp() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 100
}

func (rcv *Monster) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Inventory(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j * 1))
	}
	return 0
}

func (rcv *Monster) InventoryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) InventoryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Color() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 8
}

func (rcv *Monster) TestType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Test(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *Monster) Test4(obj *Test, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
	if obj == nil {
		obj = new(Test)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) Test4Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testarrayofstring(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j * 4))
	}
	return nil
}

func (rcv *Monster) TestarrayofstringLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// an example documentation comment: this will end up in the generated code
/// multiline too
func (rcv *Monster) Testarrayoftables(obj *Monster, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(Monster)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) TestarrayoftablesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Enemy(obj *Monster) *Monster {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Monster)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Testnestedflatbuffer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j * 1))
	}
	return 0
}

func (rcv *Monster) TestnestedflatbufferLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) TestnestedflatbufferBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Testempty(obj *Stat) *Stat {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Stat)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Testbool() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashs32Fnv1() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashu32Fnv1() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashs64Fnv1() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashu64Fnv1() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashs32Fnv1a() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashu32Fnv1a() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashs64Fnv1a() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testhashu64Fnv1a() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Testarrayofbools(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j * 1))
	}
	return 0
}

func (rcv *Monster) TestarrayofboolsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testf() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 3.14159
}

func (rcv *Monster) Testarrayofbytes(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j * 1))
	}
	return 0
}

func (rcv *Monster) TestarrayofbytesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) TestarrayofbytesBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Testarrayofbools1(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j * 1))
	}
	return 0
}

func (rcv *Monster) Testarrayofbools1Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testarrayofshorts(j int) int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt16(a + flatbuffers.UOffsetT(j * 2))
	}
	return 0
}

func (rcv *Monster) TestarrayofshortsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testarrayofints(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j * 4))
	}
	return 0
}

func (rcv *Monster) TestarrayofintsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testarrayoflongs(j int) int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j * 8))
	}
	return 0
}

func (rcv *Monster) TestarrayoflongsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testarrayoffloats(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j * 4))
	}
	return 0
}

func (rcv *Monster) TestarrayoffloatsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testarrayofdoubles(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j * 8))
	}
	return 0
}

func (rcv *Monster) TestarrayofdoublesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MonsterStart(builder *flatbuffers.Builder) { builder.StartObject(33) }
func MonsterAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) { builder.PrependStructSlot(0, flatbuffers.UOffsetT(pos), 0) }
func MonsterAddMana(builder *flatbuffers.Builder, mana int16) { builder.PrependInt16Slot(1, mana, 150) }
func MonsterAddHp(builder *flatbuffers.Builder, hp int16) { builder.PrependInt16Slot(2, hp, 100) }
func MonsterAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(name), 0) }
func MonsterAddInventory(builder *flatbuffers.Builder, inventory flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(inventory), 0) }
func MonsterStartInventoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(1, numElems, 1)
}
func MonsterAddColor(builder *flatbuffers.Builder, color int8) { builder.PrependInt8Slot(6, color, 8) }
func MonsterAddTestType(builder *flatbuffers.Builder, testType byte) { builder.PrependByteSlot(7, testType, 0) }
func MonsterAddTest(builder *flatbuffers.Builder, test flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(test), 0) }
func MonsterAddTest4(builder *flatbuffers.Builder, test4 flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(test4), 0) }
func MonsterStartTest4Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 2)
}
func MonsterAddTestarrayofstring(builder *flatbuffers.Builder, testarrayofstring flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(testarrayofstring), 0) }
func MonsterStartTestarrayofstringVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func MonsterAddTestarrayoftables(builder *flatbuffers.Builder, testarrayoftables flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(testarrayoftables), 0) }
func MonsterStartTestarrayoftablesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func MonsterAddEnemy(builder *flatbuffers.Builder, enemy flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(enemy), 0) }
func MonsterAddTestnestedflatbuffer(builder *flatbuffers.Builder, testnestedflatbuffer flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(testnestedflatbuffer), 0) }
func MonsterStartTestnestedflatbufferVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(1, numElems, 1)
}
func MonsterAddTestempty(builder *flatbuffers.Builder, testempty flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(testempty), 0) }
func MonsterAddTestbool(builder *flatbuffers.Builder, testbool byte) { builder.PrependByteSlot(15, testbool, 0) }
func MonsterAddTesthashs32Fnv1(builder *flatbuffers.Builder, testhashs32Fnv1 int32) { builder.PrependInt32Slot(16, testhashs32Fnv1, 0) }
func MonsterAddTesthashu32Fnv1(builder *flatbuffers.Builder, testhashu32Fnv1 uint32) { builder.PrependUint32Slot(17, testhashu32Fnv1, 0) }
func MonsterAddTesthashs64Fnv1(builder *flatbuffers.Builder, testhashs64Fnv1 int64) { builder.PrependInt64Slot(18, testhashs64Fnv1, 0) }
func MonsterAddTesthashu64Fnv1(builder *flatbuffers.Builder, testhashu64Fnv1 uint64) { builder.PrependUint64Slot(19, testhashu64Fnv1, 0) }
func MonsterAddTesthashs32Fnv1a(builder *flatbuffers.Builder, testhashs32Fnv1a int32) { builder.PrependInt32Slot(20, testhashs32Fnv1a, 0) }
func MonsterAddTesthashu32Fnv1a(builder *flatbuffers.Builder, testhashu32Fnv1a uint32) { builder.PrependUint32Slot(21, testhashu32Fnv1a, 0) }
func MonsterAddTesthashs64Fnv1a(builder *flatbuffers.Builder, testhashs64Fnv1a int64) { builder.PrependInt64Slot(22, testhashs64Fnv1a, 0) }
func MonsterAddTesthashu64Fnv1a(builder *flatbuffers.Builder, testhashu64Fnv1a uint64) { builder.PrependUint64Slot(23, testhashu64Fnv1a, 0) }
func MonsterAddTestarrayofbools(builder *flatbuffers.Builder, testarrayofbools flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(24, flatbuffers.UOffsetT(testarrayofbools), 0) }
func MonsterStartTestarrayofboolsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(1, numElems, 1)
}
func MonsterAddTestf(builder *flatbuffers.Builder, testf float32) { builder.PrependFloat32Slot(25, testf, 3.14159) }
func MonsterAddTestarrayofbytes(builder *flatbuffers.Builder, testarrayofbytes flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(26, flatbuffers.UOffsetT(testarrayofbytes), 0) }
func MonsterStartTestarrayofbytesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(1, numElems, 1)
}
func MonsterAddTestarrayofbools1(builder *flatbuffers.Builder, testarrayofbools1 flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(27, flatbuffers.UOffsetT(testarrayofbools1), 0) }
func MonsterStartTestarrayofbools1Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(1, numElems, 1)
}
func MonsterAddTestarrayofshorts(builder *flatbuffers.Builder, testarrayofshorts flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(28, flatbuffers.UOffsetT(testarrayofshorts), 0) }
func MonsterStartTestarrayofshortsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(2, numElems, 2)
}
func MonsterAddTestarrayofints(builder *flatbuffers.Builder, testarrayofints flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(29, flatbuffers.UOffsetT(testarrayofints), 0) }
func MonsterStartTestarrayofintsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func MonsterAddTestarrayoflongs(builder *flatbuffers.Builder, testarrayoflongs flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(30, flatbuffers.UOffsetT(testarrayoflongs), 0) }
func MonsterStartTestarrayoflongsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(8, numElems, 8)
}
func MonsterAddTestarrayoffloats(builder *flatbuffers.Builder, testarrayoffloats flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(31, flatbuffers.UOffsetT(testarrayoffloats), 0) }
func MonsterStartTestarrayoffloatsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func MonsterAddTestarrayofdoubles(builder *flatbuffers.Builder, testarrayofdoubles flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(32, flatbuffers.UOffsetT(testarrayofdoubles), 0) }
func MonsterStartTestarrayofdoublesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(8, numElems, 8)
}
func MonsterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
