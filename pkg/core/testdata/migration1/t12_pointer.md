| 이름 | 현행 CCCC | 현행 TAG(L=0) | [제안] C | [제안] 새 TAG | 양화 칸 | 구분 |
|---|---|---|---|---|---:|---|
| ISSUER-POINTER | `0110` | `0030` | `0x03` | `0x003L` | 1 | 확정 |
| DOC-POINTER | `0111` | `0038` | `0x04` | `0x004L` | 1 | 확정 |
| EXTERNAL-POINTER-1WORD | `1000` | `0040` | `0x1E`~`0x22` | `0x01EL`~`0x022L` | 5 | 5Q 확정, 번호 [제안] |
| EXTERNAL-POINTER-2WORD | `1001` | `0048` | `0x23`~`0x27` | `0x023L`~`0x027L` | 5 | 5Q 확정, 번호 [제안] |
| EXTERNAL-POINTER-4WORD | `1010` | `0050` | `0x28`~`0x2C` | `0x028L`~`0x02CL` | 5 | 5Q 확정, 번호 [제안] |
| EXTERNAL-POINTER-8WORD | `1011` | `0058` | 없음 | 없음 | 0 | 삭제 확정 |
| EXTERNAL-ISSUER-1WORD | `1100` | `0060` | 없음 | 없음 | 0 | 삭제 확정 |
| EXTERNAL-ISSUER-2WORD | `1101` | `0068` | `0x2D`~`0x31` | `0x02DL`~`0x031L` | 5 | 5Q 확정, 번호 [제안] |
| EXTERNAL-ISSUER-4WORD | `1110` | `0070` | `0x32`~`0x36` | `0x032L`~`0x036L` | 5 | 5Q 확정, 번호 [제안] |
| EXTERNAL-ISSUER-16WORD | `1111` | `0078` | `0x37`~`0x3B` | `0x037L`~`0x03BL` | 5 | 5Q 확정, 번호 [제안] |
