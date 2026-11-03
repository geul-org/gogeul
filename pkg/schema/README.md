# schema

`.geul` 바이트열이 형식에 맞는지 검사한다. 위반은 Fact(Rule, Where, Expected, Actual)로 낸다.
규칙의 출처는 명세 `README`(TAG, Length), `OVERVIEW`(Kind별 규칙), `docs/edge`·`docs/value`(Codec 상세), `GEUL-ORG`·`edge/PUBLICATION`(버전 목록)이다.

## 층

| 층 | 근거 | 조회 | 상태 |
|---|---|---|---|
| 1. 0층 형식 | 명세와 core 내장 규칙만 | 없음 | 지금 구현 대상 |
| 2. Codec 규격 | 받은 Codec 규격(직접 코드 상세, ISSUER 형식 레이아웃) | Codec을 인자로 받는다 | 이후 |
| 3. 의미 | 격틀, 선택 제약, 클래스 계층(P279), 공리 | 스키마 문서를 resolve로 받는다 | 이후 |

- 1층은 코덱북 없이 어떤 디코더든 확인할 수 있는 것만 본다(GEUL-ORG 0층). 모르는 Codec의 Payload는 열지 않고 Length로 건너뛴다.
- 2층은 Codec 규격을 받아 그 Codec의 Payload 레이아웃까지 본다. 처음에는 이 라이브러리에 내장한 규격(TOULMIN, `docs/value`의 직접 코드 형식, `docs/edge`의 서술 Edge)을 쓰고, 나중에는 Codec 규격 문서를 받아 쓴다.
- 3층은 기존 설계 그대로다. 격틀은 Defeasible이며 관용 표현, 비유, 허구는 예외(EXCEPT)로 허용한다. 스키마도 GEUL 문서로 공표하고 IMPORT로 가져오며, `0x0000` 핵심 스키마는 엔진에 내장한다.

## 1층 규칙

### F 프레이밍 (README)

| 규칙 | 내용 |
|---|---|
| F1 | 바이트 수가 짝수다(WORD 정렬). |
| F2 | Length가 레코드의 전체 WORD 수와 같고, 레코드가 입력 안에 들어간다. |
| F3 | 1~14 WORD는 직접 길이, 그보다 긴 레코드는 `1110`(전체 16~65535; 15는 확장 WORD를 빼면 14라 직접 길이), 그 위는 `1111`(전체 65537 이상; 65536은 `1110`으로 65535)이다. 더 넓은 표현은 비정규로 거부한다(문법 record.md 정규형). |
| F4 | TAG WORD 접두사가 `000`이다. 그 밖은 Error다(v1 디코더는 `000`이 아닌 TAG WORD를 거부한다, 문법 overview). |
| F5 | Codec이 ISSUER-2/4/16WORD(Value TAG `0x09DL`~`0x09FL`, Edge TAG `0x1BDL`~`0x1BFL`)면 TAG 뒤에 그 폭의 Issuer와 코덱 번호 2 WORD가 있다. Issuer는 값이 들어가는 가장 짧은 폭이다. Reference의 REFERENCE-ISSUER(TAG `0x05DL`~`0x05FL`)는 Payload가 그 Issuer의 것이다. |

### D 문서 구조 (OVERVIEW 구조)

| 규칙 | 내용 |
|---|---|
| D1 | 첫 레코드가 `DOC-1WORD/2WORD/4WORD`이고, 그 길이가 입력 전체와 같으며, 값이 들어가는 가장 좁은 폭이다. |
| D2 | `END`가 마지막 레코드다. 뒤에 아무것도 없다. |
| D3 | HEAD에는 Meta와 HEAD Edge(대상이 `0`(문서 전체)인 `SOURCE`, HEAD `SOURCE`를 대상으로 하는 `QUALIFIER`)만, BODY에는 Meta가 아닌 레코드만 둔다. `BODY`는 한 번만 둔다. HEAD 레코드 뒤에 그 밖의 레코드가 오면 BODY 누락이다. HEAD가 비면 `BODY`를 생략한다(빈 HEAD 뒤 `BODY`는 정규형이 아니다, 문법 document.md·0x1C3L). |
| D4 | `BODY`·`END`는 Payload가 없는 1 WORD다. 문서 안에 `DOC-*`가 다시 나오지 않는다. |

### R Reference (OVERVIEW Reference)

| 규칙 | 내용 |
|---|---|
| R1 | Reference TAG `0x000L`~`0x002L`은 자리 표지 ABSENT·UNKNOWN-VALUE·NO-VALUE(WORD `0x0000`·`0x0010`·`0x0020`)다. 레코드 자리(HEAD·BODY)의 자리 표지는 Error다(표지는 그것을 받는 레코드 필드 안에서만 쓴다). Payload가 있는 표지 C와 빈칸(`0x003L`·`0x004L`, `0x01EL`·`0x01FL`, `0x03EL`·`0x03FL`, `0x043L`~`0x05CL`)은 보지 않는다(migration.md §1.1 112). 5Q 체계는 섹션 0 `0x005L`~`0x01DL`(체계 C/5)과 섹션 1 `0x020L`~`0x03DL`(체계 6 + C/5)이고 Quantifier는 C%5다(C는 섹션 안 5비트 오프셋). |
| R2 | 대상 ID는 가장 짧은 폭이다. 값 0은 `0x0000` 1 WORD이고, 그 밖에는 선행 `0x0000` WORD가 없다. |
| R3 | `EXTERNAL-POINTER-*`: IMPORT ID가 들어가는 가장 좁은 Codec이고, 그 IMPORT ID를 배정한 `IMPORT-DOC`이 HEAD에 있다. 대상 ID는 `0`이 아니다(문서 자신은 `DOC-POINTER`). |
| R4 | `EXTERNAL-ISSUER-*`(체계 9~11, 2/4/16 WORD): Issuer 폭이 체계와 맞고 값이 들어가는 가장 짧은 폭이며, 문서 ID 4 WORD 뒤에 대상 ID가 1 WORD 이상 있다. 대상 ID는 `0`이 아니다(문서 자신은 `DOC-POINTER`). 문서 ID `0`은 Issuer의 용어 정의 색인 문서다. |
| R5 | `ISSUER-POINTER`(`0x040L`)·`DOC-POINTER`(`0x041L`)·`CODEC-POINTER`(`0x042L`): Issuer 폭(Length에서 고정 필드 — DOC-POINTER 문서 ID 4 + 버전 ID 2, CODEC-POINTER 코덱 ID 2 + 버전 ID 2 — 를 뺀 나머지)은 2·4·16 WORD 중 값이 들어가는 가장 짧은 폭이다. 문서 ID `0`은 Issuer의 용어 정의 색인 문서, `0xFFFFFFFFFFFFFFFF`는 Issuer 목록(`issuers/`) 예약 ID다(docs·codecs 목록은 없다, migration.md §1.1 101). |

### V Value (OVERVIEW Value)

| 규칙 | 내용 |
|---|---|
| V2 | 고정 폭 직접 코드는 Payload 폭이 맞다. BOOL 1(`0x0000`·`0x0001`만), BINARY-FLOAT 1/2/4/8, GREGORIAN-DATE 2, CIVIL-TIME-NANOSECONDS 4(`[0, 86400×10^9)`), POSIX-TIMESTAMP-SECONDS 4, DURATION-SECONDS 4, LANGUAGE-SHORT-CODE 1, UUID 8, SHA-256-DIGEST 16, SHA1-DIGEST 10, DURATION-NANOSECONDS 4. BIT-VECTOR는 bit 수 N 2 WORD 뒤 정확히 ceil(N/16) WORD이고 마지막 WORD의 쓰지 않는 하위 bit가 0이다(§3.14로 직접 코드가 된 Value). |
| V3 | SIGNED-INTEGER·UNSIGNED-INTEGER와 DECIMAL의 coefficient는 최소 WORD다. DECIMAL은 exponent 2 WORD + coefficient 1 WORD 이상이다. |
| V4 | 홀짝 코드 쌍(텍스트, URI, 언어 태그, EDTF, GEDCOM-DATE, UTF8-LANGUAGE-TEXT의 텍스트 부분, OPAQUE-BYTES)에서 홀수 코드는 Payload가 1 WORD 이상이고 마지막 하위 byte가 `0x00`이다. |
| V5 | UTF8-TEXT는 well-formed UTF-8이다(선두 U+FEFF도 문자로 허용한다). URI·URI-RELATIVE-REF·LANGUAGE-BCP47-CODE·EDTF는 ASCII이고 빈 문자열이 아니다. GEDCOM-DATE는 ASCII 0x20~0x7E이고 비어 있지 않으며 앞뒤 공백이 없다. |
| V6 | EDTF는 GREGORIAN-DATE로 정확히 쓸 수 있는 날짜(한정자·미상 자릿수·구간 없는 `[-]YYYY-MM-DD`)가 아니다. |
| V7 | UTF8-LANGUAGE-TEXT는 완전한 LANGUAGE-SHORT-CODE 또는 LANGUAGE-BCP47-CODE 레코드로 시작하고 그 레코드도 Value 규칙을 지킨다. 나머지는 UTF8-TEXT처럼 V4·V5를 따른다. |
| V8 | COORDINATE-EARTH는 Payload가 10 WORD(위도·경도 binary64, 불확도 2 WORD binary32) 뒤 Unit 레코드, COORDINATE-GEO는 천체 Reference 레코드 하나(globe 닫힌 목록: 5Q 체계의 Identified 코드 `0x005L`·`0x00AL`·`0x00FL`·`0x014L`·`0x020L`·`0x025L`·`0x02AL`·`0x02FL`·`0x034L`·`0x039L`, WIKIDATA-PID·포인터·REFERENCE-ISSUER 제외, 문법 0x065L) 뒤 같은 배치다. 위도·경도는 유한하고 위도 [-90, 90], 경도 [-180, 180](Earth)·[-360, 360](Geo)이다. 불확도는 유한한 +0 이상(-0·NaN·±Inf 아님)이다. Unit은 Quantifier Identified인 Reference 또는 ABSENT(불확도 미상)이고, ABSENT면 불확도는 +0(`0000 0000`)이다. |
| V9 | QUANTITY는 수치 amount, Reference(자리 표지 제외) 또는 ABSENT(`0x0000`) unit, 수치 또는 ABSENT lower·upper의 네 레코드뿐이다. 수치는 SIGNED-INTEGER·UNSIGNED-INTEGER·BINARY-FLOAT·FRACTION·SIGNED-DECIMAL·UNSIGNED-DECIMAL의 닫힌 목록이다. 있는 수치는 정규형이다: 정수·coefficient 최소 WORD, BINARY-FLOAT 1/2/4/8 WORD의 유한값(NaN·±Infinity 아님), FRACTION은 V10(value/numeric/QUANTITY). |
| V10 | FRACTION은 SIGNED-INTEGER 또는 SIGNED-DECIMAL numerator와 UNSIGNED-INTEGER 또는 UNSIGNED-DECIMAL denominator의 두 레코드뿐이다. 성분은 최소 WORD이고, 지수 0인 DECIMAL 성분(정수 원문은 INTEGER)과 값이 0인 분모는 Error다. 약분 여부는 보지 않는다(표기 보존, value/numeric/FRACTION). |

- ISSUER 형식 Value는 1층에서 Issuer·형식 번호의 존재(F5)만 본다. 레이아웃은 2층이다.

### E Edge (OVERVIEW Edge)

| 규칙 | 내용 |
|---|---|
| E1 | Payload 첫 요소가 LOCAL Reference(Quantifier `000`)다. 그 LOCAL ID가 Edge Id이며 `0`이 아니다(`0`은 문서 자신). |
| E2 | Edge Id는 문서 안에서 중복되지 않는다. HEAD Edge와 BODY는 한 번호 공간이다. |
| E4 | core 내장 Codec은 레이아웃이 맞다. `PUBLICATION`·`DEPRECATION`·`TOMBSTONE`(Edge Id 뒤 22·6·7 WORD), `DECLARE`·`FIELD`·`ENUM`·`CODEC`(edge/DECLARE·FIELD·ENUM·CODEC). |
| E5 | CODEC의 필드는 같은 문서의 FIELD이고, 한 필드에 같은 대상을 두 번 두지 않으며, 날값 필드(폭 ≥ 1)에는 CODEC이 하나뿐이고 와일드카드(LLLL `1111`)나 자리 표지(`0000`·`0010`·`0020`)가 아니다. 자리 표지 대상만 있는 필드는 없다(edge/CODEC 선택 필드와 자리 표지). 대상 WORD는 LLLL `0000`이거나 `1111` 와일드카드다(E4). |
| E6 | (Warn) Object가 NO-VALUE(`0x0002`)인 TRIPLE과 같은 Subject·Predicate·rank, 같은 QUALIFIER 묶음(Property=Value)의 TRIPLE이 다른 Object(값 또는 UNKNOWN-VALUE)를 가지면 경고한다. 거부하지 않는다(위키데이터에도 있는 상태, 읽는 쪽이 rank로 고른다). TRIPLE·QUALIFIER는 문법이 고정한 레이아웃으로 자른다. |
| E7 | TRIPLE·QUALIFIER의 Predicate·Subject(Property·Target)에는 Reference Codec `0000`(자리 표지)이 오지 않고, Object(Value)에는 `0x0001` UNKNOWN-VALUE·`0x0002` NO-VALUE만 온다(ABSENT·예약 표지·Payload 있는 Codec `0000`은 Error). 두 Codec의 필드는 문법이 고정하므로 1층에서 본다. 다른 Codec 필드의 표지 수용은 규격 문서(2층) 몫이다. |
| E8 | 근거 Edge `SOURCE`(`0x20`, edge/SOURCE): Edge Id 뒤에 대상·원천 레코드 둘뿐이다. 대상은 LOCAL Reference이고 `0`(문서 전체)이거나 BODY의 Edge Id다. 원천은 Reference이고, 자리 표지는 `0x0000` ABSENT와 `0x0001` UNKNOWN-VALUE만 받는다(Value·NO-VALUE·예약 표지는 Error). 대상이 `0`인 `SOURCE`와 HEAD `SOURCE`를 대상으로 하는 `QUALIFIER`가 BODY에 있으면 Error다(HEAD에 둔다). 레이아웃을 문법이 고정하므로 1층에서 본다. |
| E9 | ENUM(edge/ENUM)의 필드는 같은 문서의 날값 FIELD(폭 ≥ 1)이고, 코드값은 정확히 그 폭이며, 한 필드에서 같은 코드값을 두 번 선언하지 않는다. 디코딩되지 않는 ENUM은 E4다. |

### M Meta (OVERVIEW Meta)

| 규칙 | 내용 |
|---|---|
| M2 | `DOC-ID`는 하나이고, Payload는 `DOC-POINTER`(Issuer, 문서 ID ≠ 0, 버전 0) Reference 하나다. |
| M3 | `IMPORT-DOC`: Payload는 `DOC-POINTER` Reference 하나와 그 뒤의 IMPORT ID다. IMPORT ID가 최소 폭이고 문서 안에서 중복되지 않는다. |
| M4 | `SPEC-VERSION`·`PROFILE`·`DEFAULT-LANGUAGE`은 각각 하나다. `SPEC-VERSION`의 Payload는 내장 Value 없이 가장 짧은 폭의 unsigned 정수다(v1 = `0x0001`). 나머지는 내장 Value가 해당 형식이다. |
| M5 | `EFFECTIVE-FROM`·`EFFECTIVE-UNTIL`은 각각 하나이고 GREGORIAN-DATE 또는 POSIX-TIMESTAMP-SECONDS다. 둘 다 있으면 FROM이 UNTIL보다 앞선다. |
| M6 | `PARENT`는 SHA-256-DIGEST Value 하나다. M4~M6이 검사하는 Meta의 내장 Value는 V 규칙도 통과한다. 그 밖의 Meta Payload는 2층이다. |
| M7 | `DELTA-BASE`는 2 WORD 버전 ID ≥ 1, `DELTA-REMOVE`는 서로 다른 LOCAL Reference 1개 이상, `NEXT-LOCAL-ID`는 최소 폭 unsigned ≥ 1이다(meta/lineage). 놓일 자리(diff 문서, 버전 목록)는 아직 검사하지 않는다. |

### P 버전 목록 문서 (edge/PUBLICATION)

`DOC-ID`가 없고 BODY가 `PUBLICATION`으로 시작하는 문서(`<I>/docs/<D>/versions.geul`, `<I>/codecs/<C>/versions.geul`, 목록의 `versions.geul`)에 적용한다. 목록이 다루는 문서는 첫 `PUBLICATION`의 문서 ID다.

| 규칙 | 내용 |
|---|---|
| P1 | `IMPORT-DOC`이 없고, BODY에는 그 문서의 `PUBLICATION`·`DEPRECATION`·`TOMBSTONE`만 있다. `BIND`(문서 이름)도 P1 위반이다(문서 이름 없음, migration.md §1.1 102). |
| P2 | Edge Id는 1부터 이어지는 순번이다. |
| P3 | 문서 ID마다 PUBLICATION의 버전 ID가 1부터 빠짐없이 오름차순이고, 같은 내용 해시를 다시 공표하지 않는다. |
| P4 | DEPRECATION·TOMBSTONE은 앞에 나온 PUBLICATION을 가리킨다. |

- 1층은 서명을 확인하지 않는다. 서명 검증과 `PARENT` 사슬 확인은 resolve와 검증기(ISSUER 서명 절)의 일이다.

## 결과

```go
type Severity int // Error, Warn

type Where struct {
	Word   int    // 문서 처음부터의 WORD 오프셋
	Record int    // 레코드 순번 (0 = DOC-*)
	Part   string // 레코드 안의 위치 ("Length", "Edge Id", "Payload[3]" …)
}

type Fact struct {
	Rule     string // "F3", "R2" …
	Severity Severity
	Where    Where
	Expected string
	Actual   string
}

type Report struct {
	Facts   []Fact
	Verdict Verdict // Pass, Fail, Review
}

func Check(b []byte, opt Options) Report

type Options struct {
	Strict bool    // Warn을 Error로
	Codecs []Codec // 2층. 비면 1층만
}
```

- Error가 하나라도 있으면 Fail, Warn만 있으면 Review, 없으면 Pass다.
- Warn은 경고 규칙(F4 Standard 접두사, M4 모르는 SPEC-VERSION, E6 NO-VALUE 모순)이 낸다. `Strict`이면 Error다. 배정되지 않은 코드를 예약(Warn)으로 보고하던 규칙(V1, E3, M1, R1의 빈칸과 Payload 있는 표지)은 지웠다(migration.md §1.1 112). 빈 코드는 코덱 조회(404)에서 드러나고, 변환 검증은 migvalid `new-unassigned`가 맡는다.

## 원칙

- 조회하지 않는다. 네트워크·파일 I/O가 없다.
- 첫 위반에서 멈추지 않고 모은다. 레코드의 Length를 읽을 수 있으면 다음 레코드로 건너뛰어 계속 검사한다. Length를 믿을 수 없으면(F2) 거기서 멈춘다.
- core의 디코드 함수를 재사용하되, core는 첫 오류에서 멈추므로 문서 순회는 schema가 직접 한다.
- 규칙 ID는 이 README의 표와 1:1이다. 규칙을 바꾸면 표를 먼저 고친다.
- 2층 Codec은 인터페이스로 받는다. 대략 `Match(r core.Record) bool`과 `Check(r core.Record, at Where) []Fact`다.

## 구현 상태

- 1층 규칙 표 전체를 구현했다(`Check`). 예외: EDTF·GEDCOM 날짜의 전체 문법과 BCP 47 tag 문법 검사(미결정. V6의 평범한 날짜 금지는 구현했다. 라이브러리는 외부 의존성이 없다). UTF8-NFC-TEXT(`0x42`/`0x43`)는 2026-10-26 폐기되어 배정이 없다(보지 않는다).
- 2층은 `Options.Codecs`로 받은 Codec을 모든 레코드에 돌리는 틀만 있다. 내장 규격(TOULMIN, `docs/value`, `docs/edge`)은 아직 없다.
- 3층(격틀, 선택 제약)은 없다.
- core에 버전 목록 레코드 디코드(`core.DecodePublication`)를 넣어 E4·P 규칙이 쓴다.
- 지금 있는 `.geul`(toulmin testdata, 견본 공표 문서, core.geul)은 1층 위반 0건이다.

## 미결정

- URI·LANGUAGE-BCP47-CODE 문법 검사의 엄격도(RFC 3986 전체 문법, IANA 레지스트리 대조 여부).
- Codec 규격 문서의 형식. 2층이 받을 Codec 규격을 GEUL 문서로 어떻게 표현할지.
- 서술 Edge(REFERENCE-NODE 등)의 레이아웃을 1층에 둘지 2층에 둘지. 지금은 core 내장 Codec만 1층이다.
