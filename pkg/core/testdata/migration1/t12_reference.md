| 현행 TAG(L=0) | CCCC | QQQ | 현행 뜻 | 새 C | 새 TAG | 구분 |
|---|---|---|---|---|---|---|
| `0000` | `0000` | `000`(표지 번호) | 자리 표지 ABSENT | `0x00` | `0x000L` | 확정 |
| `0001` | `0000` | `001`(표지 번호) | 자리 표지 UNKNOWN-VALUE | `0x01` | `0x001L` | 확정 |
| `0002` | `0000` | `010`(표지 번호) | 자리 표지 NO-VALUE | `0x02` | `0x002L` | 확정 |
| `0003`~`0007` | `0000` | `011`~`111`(표지 번호) | 자리 표지 예약 5칸 | 없음 | 없음 | 확정(삭제) |
| `0008` | `0001` | `000` | LOCAL × Identified | `0x05` | `0x005L` | 확정 |
| `0009` | `0001` | `001` | LOCAL × Generic | `0x06` | `0x006L` | 확정 |
| `000A` | `0001` | `010` | LOCAL × Free-choice | `0x07` | `0x007L` | 확정 |
| `000B` | `0001` | `011` | LOCAL × Universal | `0x08` | `0x008L` | 확정 |
| `000C` | `0001` | `100` | LOCAL × Variable | `0x09` | `0x009L` | 확정 |
| `000D`~`000F` | `0001` | `101`~`111` | LOCAL × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0010` | `0010` | `000` | GEUL × Identified | `0x0A` | `0x00AL` | 계산 |
| `0011` | `0010` | `001` | GEUL × Generic | `0x0B` | `0x00BL` | 계산 |
| `0012` | `0010` | `010` | GEUL × Free-choice | `0x0C` | `0x00CL` | 계산 |
| `0013` | `0010` | `011` | GEUL × Universal | `0x0D` | `0x00DL` | 계산 |
| `0014` | `0010` | `100` | GEUL × Variable | `0x0E` | `0x00EL` | 계산 |
| `0015`~`0017` | `0010` | `101`~`111` | GEUL × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0018` | `0011` | `000` | GEUL-SIDX × Identified | `0x0F` | `0x00FL` | 계산 |
| `0019` | `0011` | `001` | GEUL-SIDX × Generic | `0x10` | `0x010L` | 계산 |
| `001A` | `0011` | `010` | GEUL-SIDX × Free-choice | `0x11` | `0x011L` | 계산 |
| `001B` | `0011` | `011` | GEUL-SIDX × Universal | `0x12` | `0x012L` | 계산 |
| `001C` | `0011` | `100` | GEUL-SIDX × Variable | `0x13` | `0x013L` | 계산 |
| `001D`~`001F` | `0011` | `101`~`111` | GEUL-SIDX × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0020` | `0100` | `000` | WIKIDATA-QID × Identified | `0x14` | `0x014L` | 계산 |
| `0021` | `0100` | `001` | WIKIDATA-QID × Generic | `0x15` | `0x015L` | 계산 |
| `0022` | `0100` | `010` | WIKIDATA-QID × Free-choice | `0x16` | `0x016L` | 계산 |
| `0023` | `0100` | `011` | WIKIDATA-QID × Universal | `0x17` | `0x017L` | 계산 |
| `0024` | `0100` | `100` | WIKIDATA-QID × Variable | `0x18` | `0x018L` | 계산 |
| `0025`~`0027` | `0100` | `101`~`111` | WIKIDATA-QID × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0028` | `0101` | `000` | WIKIDATA-PID × Identified | `0x19` | `0x019L` | 계산 |
| `0029` | `0101` | `001` | WIKIDATA-PID × Generic | `0x1A` | `0x01AL` | 계산 |
| `002A` | `0101` | `010` | WIKIDATA-PID × Free-choice | `0x1B` | `0x01BL` | 계산 |
| `002B` | `0101` | `011` | WIKIDATA-PID × Universal | `0x1C` | `0x01CL` | 계산 |
| `002C` | `0101` | `100` | WIKIDATA-PID × Variable | `0x1D` | `0x01DL` | 계산 |
| `002D`~`002F` | `0101` | `101`~`111` | WIKIDATA-PID × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0030` | `0110` | `000` | ISSUER-POINTER (양화 없음, 1칸) | `0x03` | `0x003L` | 확정 |
| `0031`~`0034` | `0110` | `001`~`100` | ISSUER-POINTER × 비 Identified | 없음 | 없음 | 확정(현행도 invalid) |
| `0035`~`0037` | `0110` | `101`~`111` | ISSUER-POINTER × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0038` | `0111` | `000` | DOC-POINTER (양화 없음, 1칸) | `0x04` | `0x004L` | 확정 |
| `0039`~`003C` | `0111` | `001`~`100` | DOC-POINTER × 비 Identified | 없음 | 없음 | 확정(현행도 invalid) |
| `003D`~`003F` | `0111` | `101`~`111` | DOC-POINTER × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0040` | `1000` | `000` | EXTERNAL-POINTER-1WORD × Identified | `0x1E` | `0x01EL` | 5Q는 확정, 번호 [제안] |
| `0041` | `1000` | `001` | EXTERNAL-POINTER-1WORD × Generic | `0x1F` | `0x01FL` | 5Q는 확정, 번호 [제안] |
| `0042` | `1000` | `010` | EXTERNAL-POINTER-1WORD × Free-choice | `0x20` | `0x020L` | 5Q는 확정, 번호 [제안] |
| `0043` | `1000` | `011` | EXTERNAL-POINTER-1WORD × Universal | `0x21` | `0x021L` | 5Q는 확정, 번호 [제안] |
| `0044` | `1000` | `100` | EXTERNAL-POINTER-1WORD × Variable | `0x22` | `0x022L` | 5Q는 확정, 번호 [제안] |
| `0045`~`0047` | `1000` | `101`~`111` | EXTERNAL-POINTER-1WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0048` | `1001` | `000` | EXTERNAL-POINTER-2WORD × Identified | `0x23` | `0x023L` | 5Q는 확정, 번호 [제안] |
| `0049` | `1001` | `001` | EXTERNAL-POINTER-2WORD × Generic | `0x24` | `0x024L` | 5Q는 확정, 번호 [제안] |
| `004A` | `1001` | `010` | EXTERNAL-POINTER-2WORD × Free-choice | `0x25` | `0x025L` | 5Q는 확정, 번호 [제안] |
| `004B` | `1001` | `011` | EXTERNAL-POINTER-2WORD × Universal | `0x26` | `0x026L` | 5Q는 확정, 번호 [제안] |
| `004C` | `1001` | `100` | EXTERNAL-POINTER-2WORD × Variable | `0x27` | `0x027L` | 5Q는 확정, 번호 [제안] |
| `004D`~`004F` | `1001` | `101`~`111` | EXTERNAL-POINTER-2WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0050` | `1010` | `000` | EXTERNAL-POINTER-4WORD × Identified | `0x28` | `0x028L` | 5Q는 확정, 번호 [제안] |
| `0051` | `1010` | `001` | EXTERNAL-POINTER-4WORD × Generic | `0x29` | `0x029L` | 5Q는 확정, 번호 [제안] |
| `0052` | `1010` | `010` | EXTERNAL-POINTER-4WORD × Free-choice | `0x2A` | `0x02AL` | 5Q는 확정, 번호 [제안] |
| `0053` | `1010` | `011` | EXTERNAL-POINTER-4WORD × Universal | `0x2B` | `0x02BL` | 5Q는 확정, 번호 [제안] |
| `0054` | `1010` | `100` | EXTERNAL-POINTER-4WORD × Variable | `0x2C` | `0x02CL` | 5Q는 확정, 번호 [제안] |
| `0055`~`0057` | `1010` | `101`~`111` | EXTERNAL-POINTER-4WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0058`~`005C` | `1011` | `000`~`100` | EXTERNAL-POINTER-8WORD × 5Q | 없음 | 없음 | 확정(폭 삭제). IMPORT ID는 4 WORD까지 |
| `005D`~`005F` | `1011` | `101`~`111` | EXTERNAL-POINTER-8WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0060`~`0064` | `1100` | `000`~`100` | EXTERNAL-ISSUER-1WORD × 5Q | 없음 | 없음 | 확정(Issuer 폭 삭제). Issuer를 2 WORD로 넓혀 `EXTERNAL-ISSUER-2WORD`로 옮긴다(Length +1) |
| `0065`~`0067` | `1100` | `101`~`111` | EXTERNAL-ISSUER-1WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0068` | `1101` | `000` | EXTERNAL-ISSUER-2WORD × Identified | `0x2D` | `0x02DL` | 5Q는 확정, 번호 [제안] |
| `0069` | `1101` | `001` | EXTERNAL-ISSUER-2WORD × Generic | `0x2E` | `0x02EL` | 5Q는 확정, 번호 [제안] |
| `006A` | `1101` | `010` | EXTERNAL-ISSUER-2WORD × Free-choice | `0x2F` | `0x02FL` | 5Q는 확정, 번호 [제안] |
| `006B` | `1101` | `011` | EXTERNAL-ISSUER-2WORD × Universal | `0x30` | `0x030L` | 5Q는 확정, 번호 [제안] |
| `006C` | `1101` | `100` | EXTERNAL-ISSUER-2WORD × Variable | `0x31` | `0x031L` | 5Q는 확정, 번호 [제안] |
| `006D`~`006F` | `1101` | `101`~`111` | EXTERNAL-ISSUER-2WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0070` | `1110` | `000` | EXTERNAL-ISSUER-4WORD × Identified | `0x32` | `0x032L` | 5Q는 확정, 번호 [제안] |
| `0071` | `1110` | `001` | EXTERNAL-ISSUER-4WORD × Generic | `0x33` | `0x033L` | 5Q는 확정, 번호 [제안] |
| `0072` | `1110` | `010` | EXTERNAL-ISSUER-4WORD × Free-choice | `0x34` | `0x034L` | 5Q는 확정, 번호 [제안] |
| `0073` | `1110` | `011` | EXTERNAL-ISSUER-4WORD × Universal | `0x35` | `0x035L` | 5Q는 확정, 번호 [제안] |
| `0074` | `1110` | `100` | EXTERNAL-ISSUER-4WORD × Variable | `0x36` | `0x036L` | 5Q는 확정, 번호 [제안] |
| `0075`~`0077` | `1110` | `101`~`111` | EXTERNAL-ISSUER-4WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
| `0078` | `1111` | `000` | EXTERNAL-ISSUER-16WORD × Identified | `0x37` | `0x037L` | 5Q는 확정, 번호 [제안] |
| `0079` | `1111` | `001` | EXTERNAL-ISSUER-16WORD × Generic | `0x38` | `0x038L` | 5Q는 확정, 번호 [제안] |
| `007A` | `1111` | `010` | EXTERNAL-ISSUER-16WORD × Free-choice | `0x39` | `0x039L` | 5Q는 확정, 번호 [제안] |
| `007B` | `1111` | `011` | EXTERNAL-ISSUER-16WORD × Universal | `0x3A` | `0x03AL` | 5Q는 확정, 번호 [제안] |
| `007C` | `1111` | `100` | EXTERNAL-ISSUER-16WORD × Variable | `0x3B` | `0x03BL` | 5Q는 확정, 번호 [제안] |
| `007D`~`007F` | `1111` | `101`~`111` | EXTERNAL-ISSUER-16WORD × 예약 | 없음 | 없음 | 확정(QQQ 삭제) |
