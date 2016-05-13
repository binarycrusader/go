// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "asm_sparc64.h"

TEXT _rt0_sparc64_solaris(SB),NOSPLIT|NOFRAME,$0
	MOVD	$main(SB), R27
	JMPL	R27, ZR

TEXT main(SB),NOSPLIT|NOFRAME,$0
	MOVD	$(8+176)(BSP), R8 // argv
	MOVD	176(BSP), R9 // argc
	MOVD	$runtime·rt0_go(SB), R27
	JMPL	R27, ZR
