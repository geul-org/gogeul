# migration1 기준 벡터

`draft/transform/gen/`(gen12.py, vec.py)이 2026-11-03 migration1 수정 전 상태에서 만든 표의 사본이다. 근거는 `files/geul/migration.md` §3이다.

- `vector.tsv`: 현행 TAG WORD 전수(0x0000~0x1FFF 중 유효한 값)의 현행 → 새 값. `-`는 TAG만으로 옮기지 않는 값(EXTERNAL-POINTER-8WORD 삭제, 1 WORD Issuer 형식은 2 WORD로 합침). CODEC-POINTER `0x03CL`, REFERENCE-ISSUER `0x03DL~0x03FL`은 신설이라 옛 값이 없다.
- `gen12.json`: 번호 검사 결과, 예시 문서(`example.md`) 새 WORD와 내용 해시.
- `t12_*.md`: 계획서 표.

시험 기대값은 이 표에서 만들고 코드 상수에서 만들지 않는다.
- `s314.tsv`: 계획서 §3.14 표(2026-11-03 사용자 결정, §1.1 97~99)를 옮겨 적은 것. GEUL Issuer가 ISSUER 형식으로 쓰던 Value 4개·Edge 21개의 새 직접 코드 TAG와 코덱 ID. `vector.tsv`는 이 결정 전에 만들어져 이 25개를 담지 않는다.
