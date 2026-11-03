| 현행 Codec | 현행 TAG(L=0) | 이름 | 새 코드 | 새 TAG |
|---|---|---|---|---|
| `0x00` | `0800` | BOOL | `0x40` | `0x040L` |
| `0x01` | `0801` | SIGNED-INTEGER | `0x41` | `0x041L` |
| `0x02` | `0802` | UNSIGNED-INTEGER | `0x42` | `0x042L` |
| `0x03` | `0803` | BINARY-FLOAT | `0x43` | `0x043L` |
| `0x04` | `0804` | COORDINATE-EARTH | `0x44` | `0x044L` |
| `0x05` | `0805` | COORDINATE-GEO | `0x45` | `0x045L` |
| `0x06` | `0806` | QUANTITY | `0x46` | `0x046L` |
| `0x07` | `0807` | FRACTION | `0x47` | `0x047L` |
| `0x08`~`0x0F` | `0808`~`080F` | | 없음 | 없음 |
| `0x10` | `0810` | SIGNED-DECIMAL | `0x48` | `0x048L` |
| `0x11` | `0811` | UNSIGNED-DECIMAL | `0x49` | `0x049L` |
| `0x12`~`0x1F` | `0812`~`081F` | | 없음 | 없음 |
| `0x20` | `0820` | GREGORIAN-DATE | `0x4A` | `0x04AL` |
| `0x21` | `0821` | CIVIL-TIME-NANOSECONDS | `0x4B` | `0x04BL` |
| `0x22` | `0822` | POSIX-TIMESTAMP-SECONDS | `0x4C` | `0x04CL` |
| `0x23` | `0823` | DURATION-SECONDS | `0x4D` | `0x04DL` |
| `0x24` | `0824` | EDTF | `0x4E` | `0x04EL` |
| `0x25` | `0825` | EDTF | `0x4F` | `0x04FL` |
| `0x26` | `0826` | GEDCOM-DATE | `0x50` | `0x050L` |
| `0x27` | `0827` | GEDCOM-DATE | `0x51` | `0x051L` |
| `0x28`~`0x3F` | `0828`~`083F` | | 없음 | 없음 |
| `0x40` | `0840` | UTF8-TEXT-EVEN | `0x52` | `0x052L` |
| `0x41` | `0841` | UTF8-TEXT-ODD | `0x53` | `0x053L` |
| `0x42`~`0x43` | `0842`~`0843` | | 없음 | 없음 |
| `0x44` | `0844` | URI-EVEN | `0x54` | `0x054L` |
| `0x45` | `0845` | URI-ODD | `0x55` | `0x055L` |
| `0x46` | `0846` | LANGUAGE-BCP47-CODE-EVEN | `0x56` | `0x056L` |
| `0x47` | `0847` | LANGUAGE-BCP47-CODE-ODD | `0x57` | `0x057L` |
| `0x48` | `0848` | LANGUAGE-SHORT-CODE | `0x5A` | `0x05AL` |
| `0x49` | `0849` | | 없음 | 없음 |
| `0x4A` | `084A` | UTF8-LANGUAGE-TEXT | `0x58` | `0x058L` |
| `0x4B` | `084B` | UTF8-LANGUAGE-TEXT | `0x59` | `0x059L` |
| `0x4C`~`0x5F` | `084C`~`085F` | | 없음 | 없음 |
| `0x60` | `0860` | UUID | `0x5B` | `0x05BL` |
| `0x61` | `0861` | SHA-256-DIGEST | `0x5C` | `0x05CL` |
| `0x62` | `0862` | SHA1-DIGEST | `0x5D` | `0x05DL` |
| `0x63`~`0x7B` | `0863`~`087B` | | 없음 | 없음 |
| `0x7C` | `087C` | ISSUER-1WORD | 삭제, 번호는 예약 | 없음 |
| `0x7D` | `087D` | VALUE-ISSUER-2WORD | `0x7D` | `0x07DL` |
| `0x7E` | `087E` | VALUE-ISSUER-4WORD | `0x7E` | `0x07EL` |
| `0x7F` | `087F` | VALUE-ISSUER-16WORD | `0x7F` | `0x07FL` |
