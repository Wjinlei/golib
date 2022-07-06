package os

import "os"

// 判断是否具有任意权限
func IsAny(mode os.FileMode) bool {
	return mode&0777 != 0
}

// 判断UGO位上是否任意一位具有r权限
func IsReadAny(mode os.FileMode) bool {
	return mode&0444 != 0
}

// 判断UGO位上是否任意一位具有w权限
func IsWriteAny(mode os.FileMode) bool {
	return mode&0222 != 0
}

// 判断UGO位上是否任意一位具有x权限
func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

// 判断User位上是否具有r权限
func IsReadOwner(mode os.FileMode) bool {
	return mode&0400 != 0
}

// 判断Group位上是否具有r权限
func IsReadGroup(mode os.FileMode) bool {
	return mode&0040 != 0
}

// 判断Other位上是否具有r权限
func IsReadOther(mode os.FileMode) bool {
	return mode&0004 != 0
}

// 判断User位上是否具有w权限
func IsWriteOwner(mode os.FileMode) bool {
	return mode&0200 != 0
}

// 判断Group位上是否具有w权限
func IsWriteGroup(mode os.FileMode) bool {
	return mode&0020 != 0
}

// 判断Other位上是否具有w权限
func IsWriteOther(mode os.FileMode) bool {
	return mode&0002 != 0
}

// 判断User位上是否具有x权限
func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}

// 判断Group位上是否具有x权限
func IsExecGroup(mode os.FileMode) bool {
	return mode&0010 != 0
}

// 判断Other位上是否具有x权限
func IsExecOther(mode os.FileMode) bool {
	return mode&0001 != 0
}

// 判断UGO位上是否都具有r权限
func IsReadAll(mode os.FileMode) bool {
	return mode&0444 == 0444
}

// 判断UGO位上是否都具有w权限
func IsWriteAll(mode os.FileMode) bool {
	return mode&0222 == 0222
}

// 判断UGO位上是否都具有x权限
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}
