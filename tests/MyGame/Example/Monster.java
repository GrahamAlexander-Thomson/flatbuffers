// automatically generated, do not modify

package MyGame.Example;

import java.nio.*;
import java.lang.*;
import java.util.*;
import com.google.flatbuffers.*;

@SuppressWarnings("unused")
/**
 * an example documentation comment: monster object
 */
public final class Monster extends Table {
  public static Monster getRootAsMonster(ByteBuffer _bb) { return getRootAsMonster(_bb, new Monster()); }
  public static Monster getRootAsMonster(ByteBuffer _bb, Monster obj) { _bb.order(ByteOrder.LITTLE_ENDIAN); return (obj.__init(_bb.getInt(_bb.position()) + _bb.position(), _bb)); }
  public static boolean MonsterBufferHasIdentifier(ByteBuffer _bb) { return __has_identifier(_bb, "MONS"); }
  public Monster __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  public Vec3 pos() { return pos(new Vec3()); }
  public Vec3 pos(Vec3 obj) { int o = __offset(4); return o != 0 ? obj.__init(o + bb_pos, bb) : null; }
  public short mana() { int o = __offset(6); return o != 0 ? bb.getShort(o + bb_pos) : 150; }
  public boolean mutateMana(short mana) { int o = __offset(6); if (o != 0) { bb.putShort(o + bb_pos, mana); return true; } else { return false; } }
  public short hp() { int o = __offset(8); return o != 0 ? bb.getShort(o + bb_pos) : 100; }
  public boolean mutateHp(short hp) { int o = __offset(8); if (o != 0) { bb.putShort(o + bb_pos, hp); return true; } else { return false; } }
  public String name() { int o = __offset(10); return o != 0 ? __string(o + bb_pos) : null; }
  public ByteBuffer nameAsByteBuffer() { return __vector_as_bytebuffer(10, 1); }
  public int inventory(int j) { int o = __offset(14); return o != 0 ? bb.get(__vector(o) + j * 1) & 0xFF : 0; }
  public int inventoryLength() { int o = __offset(14); return o != 0 ? __vector_len(o) : 0; }
  public ByteBuffer inventoryAsByteBuffer() { return __vector_as_bytebuffer(14, 1); }
  public boolean mutateInventory(int j, int inventory) { int o = __offset(14); if (o != 0) { bb.put(__vector(o) + j * 1, (byte)inventory); return true; } else { return false; } }
  public byte color() { int o = __offset(16); return o != 0 ? bb.get(o + bb_pos) : 8; }
  public boolean mutateColor(byte color) { int o = __offset(16); if (o != 0) { bb.put(o + bb_pos, color); return true; } else { return false; } }
  public byte testType() { int o = __offset(18); return o != 0 ? bb.get(o + bb_pos) : 0; }
  public boolean mutateTestType(byte test_type) { int o = __offset(18); if (o != 0) { bb.put(o + bb_pos, test_type); return true; } else { return false; } }
  public Table test(Table obj) { int o = __offset(20); return o != 0 ? __union(obj, o) : null; }
  public Test test4(int j) { return test4(new Test(), j); }
  public Test test4(Test obj, int j) { int o = __offset(22); return o != 0 ? obj.__init(__vector(o) + j * 4, bb) : null; }
  public int test4Length() { int o = __offset(22); return o != 0 ? __vector_len(o) : 0; }
  public String testarrayofstring(int j) { int o = __offset(24); return o != 0 ? __string(__vector(o) + j * 4) : null; }
  public int testarrayofstringLength() { int o = __offset(24); return o != 0 ? __vector_len(o) : 0; }
  /**
   * an example documentation comment: this will end up in the generated code
   * multiline too
   */
  public Monster testarrayoftables(int j) { return testarrayoftables(new Monster(), j); }
  public Monster testarrayoftables(Monster obj, int j) { int o = __offset(26); return o != 0 ? obj.__init(__indirect(__vector(o) + j * 4), bb) : null; }
  public int testarrayoftablesLength() { int o = __offset(26); return o != 0 ? __vector_len(o) : 0; }
  public Monster enemy() { return enemy(new Monster()); }
  public Monster enemy(Monster obj) { int o = __offset(28); return o != 0 ? obj.__init(__indirect(o + bb_pos), bb) : null; }
  public int testnestedflatbuffer(int j) { int o = __offset(30); return o != 0 ? bb.get(__vector(o) + j * 1) & 0xFF : 0; }
  public int testnestedflatbufferLength() { int o = __offset(30); return o != 0 ? __vector_len(o) : 0; }
  public ByteBuffer testnestedflatbufferAsByteBuffer() { return __vector_as_bytebuffer(30, 1); }
  public Monster testnestedflatbufferAsMonster() { return testnestedflatbufferAsMonster(new Monster()); }
  public Monster testnestedflatbufferAsMonster(Monster obj) { int o = __offset(30); return o != 0 ? obj.__init(__indirect(__vector(o)), bb) : null; }
  public boolean mutateTestnestedflatbuffer(int j, int testnestedflatbuffer) { int o = __offset(30); if (o != 0) { bb.put(__vector(o) + j * 1, (byte)testnestedflatbuffer); return true; } else { return false; } }
  public Stat testempty() { return testempty(new Stat()); }
  public Stat testempty(Stat obj) { int o = __offset(32); return o != 0 ? obj.__init(__indirect(o + bb_pos), bb) : null; }
  public boolean testbool() { int o = __offset(34); return o != 0 ? 0!=bb.get(o + bb_pos) : false; }
  public boolean mutateTestbool(boolean testbool) { int o = __offset(34); if (o != 0) { bb.put(o + bb_pos, (byte)(testbool ? 1 : 0)); return true; } else { return false; } }
  public int testhashs32Fnv1() { int o = __offset(36); return o != 0 ? bb.getInt(o + bb_pos) : 0; }
  public boolean mutateTesthashs32Fnv1(int testhashs32_fnv1) { int o = __offset(36); if (o != 0) { bb.putInt(o + bb_pos, testhashs32_fnv1); return true; } else { return false; } }
  public long testhashu32Fnv1() { int o = __offset(38); return o != 0 ? (long)bb.getInt(o + bb_pos) & 0xFFFFFFFFL : 0; }
  public boolean mutateTesthashu32Fnv1(long testhashu32_fnv1) { int o = __offset(38); if (o != 0) { bb.putInt(o + bb_pos, (int)testhashu32_fnv1); return true; } else { return false; } }
  public long testhashs64Fnv1() { int o = __offset(40); return o != 0 ? bb.getLong(o + bb_pos) : 0; }
  public boolean mutateTesthashs64Fnv1(long testhashs64_fnv1) { int o = __offset(40); if (o != 0) { bb.putLong(o + bb_pos, testhashs64_fnv1); return true; } else { return false; } }
  public long testhashu64Fnv1() { int o = __offset(42); return o != 0 ? bb.getLong(o + bb_pos) : 0; }
  public boolean mutateTesthashu64Fnv1(long testhashu64_fnv1) { int o = __offset(42); if (o != 0) { bb.putLong(o + bb_pos, testhashu64_fnv1); return true; } else { return false; } }
  public int testhashs32Fnv1a() { int o = __offset(44); return o != 0 ? bb.getInt(o + bb_pos) : 0; }
  public boolean mutateTesthashs32Fnv1a(int testhashs32_fnv1a) { int o = __offset(44); if (o != 0) { bb.putInt(o + bb_pos, testhashs32_fnv1a); return true; } else { return false; } }
  public long testhashu32Fnv1a() { int o = __offset(46); return o != 0 ? (long)bb.getInt(o + bb_pos) & 0xFFFFFFFFL : 0; }
  public boolean mutateTesthashu32Fnv1a(long testhashu32_fnv1a) { int o = __offset(46); if (o != 0) { bb.putInt(o + bb_pos, (int)testhashu32_fnv1a); return true; } else { return false; } }
  public long testhashs64Fnv1a() { int o = __offset(48); return o != 0 ? bb.getLong(o + bb_pos) : 0; }
  public boolean mutateTesthashs64Fnv1a(long testhashs64_fnv1a) { int o = __offset(48); if (o != 0) { bb.putLong(o + bb_pos, testhashs64_fnv1a); return true; } else { return false; } }
  public long testhashu64Fnv1a() { int o = __offset(50); return o != 0 ? bb.getLong(o + bb_pos) : 0; }
  public boolean mutateTesthashu64Fnv1a(long testhashu64_fnv1a) { int o = __offset(50); if (o != 0) { bb.putLong(o + bb_pos, testhashu64_fnv1a); return true; } else { return false; } }
  public boolean testarrayofbools(int j) { int o = __offset(52); return o != 0 ? 0!=bb.get(__vector(o) + j * 1) : false; }
  public int testarrayofboolsLength() { int o = __offset(52); return o != 0 ? __vector_len(o) : 0; }
  public ByteBuffer testarrayofboolsAsByteBuffer() { return __vector_as_bytebuffer(52, 1); }
  public boolean mutateTestarrayofbools(int j, boolean testarrayofbools) { int o = __offset(52); if (o != 0) { bb.put(__vector(o) + j * 1, (byte)(testarrayofbools ? 1 : 0)); return true; } else { return false; } }
  public float testf() { int o = __offset(54); return o != 0 ? bb.getFloat(o + bb_pos) : 3.14159f; }
  public boolean mutateTestf(float testf) { int o = __offset(54); if (o != 0) { bb.putFloat(o + bb_pos, testf); return true; } else { return false; } }

  public static void startMonster(FlatBufferBuilder builder) { builder.startObject(26); }
  public static void addPos(FlatBufferBuilder builder, int posOffset) { builder.addStruct(0, posOffset, 0); }
  public static void addMana(FlatBufferBuilder builder, short mana) { builder.addShort(1, mana, 150); }
  public static void addHp(FlatBufferBuilder builder, short hp) { builder.addShort(2, hp, 100); }
  public static void addName(FlatBufferBuilder builder, int nameOffset) { builder.addOffset(3, nameOffset, 0); }
  public static void addInventory(FlatBufferBuilder builder, int inventoryOffset) { builder.addOffset(5, inventoryOffset, 0); }
  public static int createInventoryVector(FlatBufferBuilder builder, byte[] data) { builder.startVector(1, data.length, 1); for (int i = data.length - 1; i >= 0; i--) builder.addByte(data[i]); return builder.endVector(); }
  public static void startInventoryVector(FlatBufferBuilder builder, int numElems) { builder.startVector(1, numElems, 1); }
  public static void addColor(FlatBufferBuilder builder, byte color) { builder.addByte(6, color, 8); }
  public static void addTestType(FlatBufferBuilder builder, byte testType) { builder.addByte(7, testType, 0); }
  public static void addTest(FlatBufferBuilder builder, int testOffset) { builder.addOffset(8, testOffset, 0); }
  public static void addTest4(FlatBufferBuilder builder, int test4Offset) { builder.addOffset(9, test4Offset, 0); }
  public static void startTest4Vector(FlatBufferBuilder builder, int numElems) { builder.startVector(4, numElems, 2); }
  public static void addTestarrayofstring(FlatBufferBuilder builder, int testarrayofstringOffset) { builder.addOffset(10, testarrayofstringOffset, 0); }
  public static int createTestarrayofstringVector(FlatBufferBuilder builder, int[] data) { builder.startVector(4, data.length, 4); for (int i = data.length - 1; i >= 0; i--) builder.addOffset(data[i]); return builder.endVector(); }
  public static void startTestarrayofstringVector(FlatBufferBuilder builder, int numElems) { builder.startVector(4, numElems, 4); }
  public static void addTestarrayoftables(FlatBufferBuilder builder, int testarrayoftablesOffset) { builder.addOffset(11, testarrayoftablesOffset, 0); }
  public static int createTestarrayoftablesVector(FlatBufferBuilder builder, int[] data) { builder.startVector(4, data.length, 4); for (int i = data.length - 1; i >= 0; i--) builder.addOffset(data[i]); return builder.endVector(); }
  public static void startTestarrayoftablesVector(FlatBufferBuilder builder, int numElems) { builder.startVector(4, numElems, 4); }
  public static void addEnemy(FlatBufferBuilder builder, int enemyOffset) { builder.addOffset(12, enemyOffset, 0); }
  public static void addTestnestedflatbuffer(FlatBufferBuilder builder, int testnestedflatbufferOffset) { builder.addOffset(13, testnestedflatbufferOffset, 0); }
  public static int createTestnestedflatbufferVector(FlatBufferBuilder builder, byte[] data) { builder.startVector(1, data.length, 1); for (int i = data.length - 1; i >= 0; i--) builder.addByte(data[i]); return builder.endVector(); }
  public static void startTestnestedflatbufferVector(FlatBufferBuilder builder, int numElems) { builder.startVector(1, numElems, 1); }
  public static void addTestempty(FlatBufferBuilder builder, int testemptyOffset) { builder.addOffset(14, testemptyOffset, 0); }
  public static void addTestbool(FlatBufferBuilder builder, boolean testbool) { builder.addBoolean(15, testbool, false); }
  public static void addTesthashs32Fnv1(FlatBufferBuilder builder, int testhashs32Fnv1) { builder.addInt(16, testhashs32Fnv1, 0); }
  public static void addTesthashu32Fnv1(FlatBufferBuilder builder, long testhashu32Fnv1) { builder.addInt(17, (int)testhashu32Fnv1, 0); }
  public static void addTesthashs64Fnv1(FlatBufferBuilder builder, long testhashs64Fnv1) { builder.addLong(18, testhashs64Fnv1, 0); }
  public static void addTesthashu64Fnv1(FlatBufferBuilder builder, long testhashu64Fnv1) { builder.addLong(19, testhashu64Fnv1, 0); }
  public static void addTesthashs32Fnv1a(FlatBufferBuilder builder, int testhashs32Fnv1a) { builder.addInt(20, testhashs32Fnv1a, 0); }
  public static void addTesthashu32Fnv1a(FlatBufferBuilder builder, long testhashu32Fnv1a) { builder.addInt(21, (int)testhashu32Fnv1a, 0); }
  public static void addTesthashs64Fnv1a(FlatBufferBuilder builder, long testhashs64Fnv1a) { builder.addLong(22, testhashs64Fnv1a, 0); }
  public static void addTesthashu64Fnv1a(FlatBufferBuilder builder, long testhashu64Fnv1a) { builder.addLong(23, testhashu64Fnv1a, 0); }
  public static void addTestarrayofbools(FlatBufferBuilder builder, int testarrayofboolsOffset) { builder.addOffset(24, testarrayofboolsOffset, 0); }
  public static int createTestarrayofboolsVector(FlatBufferBuilder builder, boolean[] data) { builder.startVector(1, data.length, 1); for (int i = data.length - 1; i >= 0; i--) builder.addBoolean(data[i]); return builder.endVector(); }
  public static void startTestarrayofboolsVector(FlatBufferBuilder builder, int numElems) { builder.startVector(1, numElems, 1); }
  public static void addTestf(FlatBufferBuilder builder, float testf) { builder.addFloat(25, testf, 3.14159f); }
  public static int endMonster(FlatBufferBuilder builder) {
    int o = builder.endObject();
    builder.required(o, 10);  // name
    return o;
  }
  public static void finishMonsterBuffer(FlatBufferBuilder builder, int offset) { builder.finish(offset, "MONS"); }
};

