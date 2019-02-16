// +build !s390x !arm
//
// SPDX-License-Identifier: Apache-2.0
//

package main

func archConvertStatFs(cgroupFsType int) int64 {
	return int64(cgroupFsType)
}
