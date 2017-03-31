#include "arch.h"
#include "fp_C41417.h"

/* Curve C41417 */

#if CHUNK==32
const BIG_416 Modulus_C41417= {0x1FFFFFEF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0xFF};
const chunk MConst_C41417=0x11;
#endif

#if CHUNK==64
const BIG_416 Modulus_C41417= {0xFFFFFFFFFFFFFEF,0xFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFFF,0x3FFFFFFFFFFFFF};
const chunk MConst_C41417=0x11;
#endif

