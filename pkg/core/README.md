# core

엔진 내장 규칙을 구현한다. 다른 패키지는 모두 core 위에 올라간다.

## 범위

- 프레이밍: TAG WORD와 Length 확장만으로 레코드 경계를 자른다. 스키마가 필요 없다.
- TAG WORD `000K KKCC CCCC LLLL`(migration1): Kind는 KKK 표(`KindOf`: 000 Reference, 001 Value, 010~110 Edge, 111 Meta), Codec은 영역 안 번호(`Record.Codec uint16`; Reference·Value·Meta `0x00`~`0x3F`, Edge `0x00`~`0x13F`), `TargetWord`/`SplitTarget`이 (Kind, Codec) ↔ CODEC 대상 WORD를 바꾼다.
- ISSUER 코드는 영역마다 마지막 셋(2/4/16 WORD, `IssuerWidth`). Value·Edge는 TAG 뒤 Issuer + 코덱 번호 2 WORD(`Record.Format uint32`, `FormatWords`), Reference의 REFERENCE-ISSUER는 Payload가 Issuer의 것이다. Issuer는 2/4/16 WORD 정규형(`IssuerCanonical`), 더 넓은 0 채움은 같은 Issuer(`IssuerNormal`).
- Reference C: 자리 표지 0~2, ISSUER-POINTER 3, DOC-POINTER 4, 5Q 체계 `0x05`~`0x3B`(`RefScheme`·`RefQuantifier`), CODEC-POINTER `0x3C`, REFERENCE-ISSUER `0x3D`~`0x3F`.
- 목록 예약 문서 ID: `DocIDIssuers`(Issuer 목록 `issuers/` 하나, §1.1 101).
- 문서 구조: `DOC-*`, HEAD(Meta, 문서 전체 SOURCE와 그 QUALIFIER — `HeadEdge`), BODY, `END`. 필드 폭은 Codec과 Length가 정한다(문서 단위 폭 선언 없음).
- 모델: Reference, Value, Edge, Meta. 모르는 레코드는 원본 WORD를 그대로 보존한다.
- 읽기: `Decode(bytes) → Doc, warns, err`
- 쓰기: 빌더가 ID와 Length를 정하고 정규형(가장 짧은 폭) 바이트로 쓴다.
- 내용 해시: 검증 클래스 레코드를 뺀 SHA-256(OVERVIEW 계보). 버전 ID(2 WORD 순번)는 공표처가 배정하며 core가 계산하지 않는다.

## 구현된 것

| 파일 | 내용 |
|---|---|
| `record.go`, `encode_record.go`, `decode_record.go` | 레코드 하나: TAG WORD, Length(확장 포함, 정규형 검사), ISSUER 확장, Payload |
| `wrap_document.go`, `split_document.go` | DOC 컨테이너: 가장 좁은 DOC-* 로 감싸기, HEAD/BODY/END 나누기와 배치 검사 |
| `*_value.go`, `*_language_text.go`, `*_coordinate.go`, `encode_posix_seconds.go` | Value: 바이트 Value(UTF8-TEXT·URI·BCP47 등 홀짝 쌍), UNSIGNED-INTEGER, SIGNED-/UNSIGNED-DECIMAL, GREGORIAN-DATE·POSIX-TIMESTAMP-SECONDS, UTF8-LANGUAGE-TEXT, COORDINATE-EARTH·GEO |
| `*_doc_pointer.go`, `*_issuer_pointer.go`, `*_codec_pointer.go`, `*_target_pointer.go`, `encode_doc_id*.go`, `decode_doc_id_codec.go`, `*_spec_version.go` | DOC-·ISSUER-·CODEC-POINTER Reference, 문서·코덱 대상 포인터, Meta DOC-ID(코덱 문서는 CODEC-POINTER)·SPEC-VERSION |
| `*_publication.go`, `publication_tail.go`, `*_bind.go` | 버전 목록 Edge(PUBLICATION·DEPRECATION·TOMBSTONE), BIND(Issuer 목록·코덱 규격의 이름. 버전 목록에는 없다, §1.1 102) |
| `*_declare.go`, `*_field_decl.go`, `*_enum_decl.go`, `*_codec_decl.go`, `codec_matches.go` | Codec 규격 어휘 DECLARE·FIELD·ENUM·CODEC, CODEC 대상 대조 |
| `encode_id.go`, `decode_id.go`, `width*.go`, `min_width.go` | 고정 폭(1/2/4/8 WORD) 필드와 폭 코드 (DOC-* 길이, TOULMIN SPAN) |
| `ref_codec.go`, `encode_local_ref.go`, `decode_local_ref.go` | Reference Codec, LOCAL Reference (Edge Id·문서 안 언급) |
| `kind*.go`, `tag_*.go`, `split_target.go`, `target_word.go`, `codec_max.go` | KKK 표, 영역 시작 TAG, (Kind, Codec) ↔ TAG·대상 WORD |
| `meta_codec.go`, `value_codec.go`, `edge_codec.go`, `codec_issuer.go`, `issuer_*.go`, `list_doc_id.go` | 사용하는 Meta·Value·Edge·ISSUER 코드(§3.14로 직접 코드가 된 OPAQUE-BYTES `0x05EL`/`0x05FL`, BIT-VECTOR `0x060L`, DURATION-NANOSECONDS `0x061L` 포함), Issuer 폭·정규형, 목록 예약 ID. 다른 Issuer의 Value ISSUER 형식 이름은 `issuer_format_names.go`(GEUL Issuer 행은 없다) |
| `testdata/migration1/` | 기준 벡터(`vector.tsv`, `gen12.json`, `t12_*.md`, §3.14 표 `s314.tsv`). 시험 기대값은 코드 상수가 아니라 이 표에서 온다 |
| `content_hash.go`, `strip_validation.go` | 내용 해시: 검증 클래스 Meta를 빼고 DOC-* 길이를 다시 계산한 뒤 END 직전까지의 SHA-256 |

- toulmin(`draft/toulmin`, 브랜치 `geul-t002`)이 이 패키지를 import해 `.geul`을 읽고 쓴다.
- 아직 없는 것: EXTERNAL-ISSUER·EXTERNAL-POINTER 인코딩, EDTF·GEDCOM-DATE 해석(바이트 Value로만 다룬다), 정규형 전체 규칙, 서명.

## 원칙

- 외부 의존과 네트워크 I/O가 없다.
- 디코딩은 IMPORT를 조회하지 않는다. 풀리지 않은 외부 참조는 IMPORT ID와 대상 ID만 가진 채 둔다.
- 쓰기와 읽기를 함께 만들고 `Decode(Encode(x)) == x` 왕복으로 검증한다.

## 보류

- 서명: 검증 클래스 레코드(`SIGNATURE`)는 해석하지 않고 보존만 한다. 내용 해시 계산에서 빼는 규칙만 지킨다.
- 정규형 규칙(폭, Length 확장 조건, 레코드 순서)이 정해지기 전에는 원본 바이트 보존 왕복으로 검증한다.
