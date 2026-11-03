# toulmin

TOULMIN 코덱(명세: geul.org 문법 [TOULMIN](https://geul.org/ko/grammar/edge/toulmin/))을 구현한다. toulmin 엔진의 판정·실행 그래프를 GEUL 레코드로 쓰고 읽는다.

| 묶음 | 내용 |
|---|---|
| 코드표 | NodeType, RoleCode, LevelCode, PropCode, ControlCode, SectionCode, 플래그, 비교·논리·양화 연산자, 필드 타입 |
| 코덱 대응 | 직접 TOULMIN 코덱(Edge C `0x60`~`0x76`)과 EDGE-ISSUER-2WORD(TAG `0x1BDL`, Issuer `0x00000000`) 코덱 번호 2 WORD(`0x0000C000`~) ↔ 코드표 |
| 값 | VFix(판정값·Qualifier·공격 세기 ↔ SIGNED-DECIMAL, 공격 세기는 ATTACKS의 F_FORCE), Gate(연산자 + UNSIGNED-DECIMAL), Span(SOURCE가 선언한 폭), ID·REF(LOCAL Reference 레코드) |
| 레코드 | NODE, LINK, 제어(SECTION, CASE-BEGIN/END, PROSE), TEXT, SOURCE |
| 필드 도우미 | `PutID`/`TakeID`, `PutValue`/`TakeValue`, `PutSpan`/`TakeSpan`, `TakeWord`, `ExpectEnd`, `EdgeRecord`, `RecordRef` |
| 버전 | `SpecVersion`(= `core.SpecVersion`, 1), `ProfileString`/`CheckProfile`("tangeul-1.0") |

- core 위에서 동작하며 tangl AST에 의존하지 않는다.
- EXPR(PREDICATE-DECLARE 조건식)과 AST ↔ 코드 변환은 tangl AST에 묶여 있어 toulmin(`pkg/tangeul`)에 남긴다.
