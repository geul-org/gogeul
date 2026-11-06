# toulmin

TOULMIN 코덱(명세: geul.org 문법 [TOULMIN](https://geul.org/ko/grammar/toulmin/))을 구현한다. toulmin 엔진의 판정·실행 그래프를 GEUL 레코드로 쓰고 읽는다.

| 묶음 | 내용 |
|---|---|
| 코드표 | NodeType, RoleCode, LevelCode, PropCode, ControlCode, SectionCode, 플래그, 비교·논리·양화 연산자, 필드 타입 |
| 코덱 대응 | 직접 TOULMIN 코덱: Edge Codec `0x60`~`0x79`(TAG `0x100L`~`0x119L`)와, 옛 GEUL ISSUER 형식 `0xC000`~`0xC029`에서 옮긴 21종 Edge Codec `0x80`~`0x94`(TAG `0x120L`~`0x134L`, 섹션 9 "TOULMIN 형식") ↔ 코드표. ISSUER 형식 레코드는 쓰지 않는다 |
| 값 | VFix(판정값·Qualifier·공격 세기 ↔ SIGNED-DECIMAL, 공격 세기는 ATTACKS의 F_FORCE), Gate(연산자 + UNSIGNED-DECIMAL), Span(SOURCE가 선언한 폭), ID·REF(LOCAL Reference 레코드) |
| 레코드 | NODE, LINK, 제어(SECTION, CASE-BEGIN/END, PROSE), TEXT, SOURCE |
| 필드 도우미 | `PutID`/`TakeID`, `PutValue`/`TakeValue`, `PutSpan`/`TakeSpan`, `TakeWord`, `ExpectEnd`, `EdgeRecord`, `RecordRef` |
| 버전 | `SpecVersion`(= `core.SpecVersion`, 1), `ProfileString`/`CheckProfile`("tangeul-1.0") |

- core 위에서 동작하며 tangl AST에 의존하지 않는다.
- EXPR(PREDICATE-DECLARE 조건식)과 AST ↔ 코드 변환은 tangl AST에 묶여 있어 toulmin(`pkg/tangeul`)에 남긴다.
