/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

/* Fixed Data in ROM - Field and Curve parameters */

package  MF256E

var Modulus= [...]Chunk {0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFA7FFFF}
const MConst Chunk=0xFFA80000

const CURVE_A int= -1
var CURVE_B = [...]Chunk {0x350A,0x0,0x0,0x0,0x0}
var CURVE_Order=[...]Chunk {0xD92EDED8EC7BAB,0xBBAFB86733C966,0xFFFFFFFFFFB154,0xFFFFFFFFFFFFFF,0x3FE9FFFF}
var CURVE_Gx =[...]Chunk {0x1,0x0,0x0,0x0,0x0}
var CURVE_Gy =[...]Chunk {0xEAA722F2F3C908,0x5E648DFEA68D7D,0xF3DB2C1AACA0C0,0xF8CC4D5AEAEBEE,0xDAD8D4F8}
