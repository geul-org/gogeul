# migration1 기준 표 (plan2)

`draft/transform/gen/plan2/`(gen_plan2.py, layout2.py)가 만든 구역 재배치 plan2(2026-11-03) 표의 사본이다. 시험 기대값은 이 표에서 만들고 코드 상수에서 만들지 않는다.

- `vector2.tsv`: 현행 TAG WORD 전수(유효한 3,315값)의 현행 → plan2 값. `-`는 TAG만으로 옮기지 않는 값(EXTERNAL-POINTER-8WORD 삭제, 1 WORD Issuer 형식은 2 WORD로 합침). CODEC-POINTER, REFERENCE-ISSUER와 §3.14 직접 코드는 신설이라 옛 값이 없다.
- `s314_p2.tsv`: §3.14 GEUL Issuer 형식 25개(Value 4, Edge 21)의 plan2 직접 코드 TAG와 코덱 ID.
- `p1_to_p2.tsv`: plan1 배정 칸 231개 전수(직접 코드, 포인터, ISSUER, Edge Reserved)의 plan1 → plan2 코드와 Kind.
- `layout2.md`: §3.15 섹션 표(섹션 → Kind)와 배정 칸 이름 표.
- `example2.json`: 문법 `example.md` 문서의 plan2 WORD와 내용 해시. plan1 예시(gen12.json `ex_new`)의 레코드 머리 TAG와 내장 TAG(DOC-POINTER, LANGUAGE-BCP47-CODE, EDTF)를 `p1_to_p2.tsv`로 옮기고 해시는 END 앞 WORD의 big-endian SHA-256으로 따로 계산했다.
