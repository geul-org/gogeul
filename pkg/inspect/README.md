# inspect

GEUL 문서(`.geul`)를 해석된 문서 모델로 읽는다. 문서를 보여 주는 도구(뷰어 ve 등)가 같은 해석을 쓰도록 한 곳에 둔다.

```go
d, err := inspect.Decode(b) // Document{Head, Body, Profile, Toulmin}
```

| Entry 필드 | 내용 |
|---|---|
| `Record`, `Words` | 원본 레코드와 그 워드(TAG WORD 포함) |
| `Name` | 스펙 이름(`core.CodecName`, `toulmin.RecordName`). 모르면 `""` |
| `Data` | 해석 내용. 아래 표 |
| `ID`, `Refs` | Edge Id와 이 레코드가 참조하는 Edge Id(0 제외, 필드 순서) |

| Data | 레코드 |
|---|---|
| `Value` | HEAD의 SPEC-VERSION(UNSIGNED-INTEGER로 표시)·PROFILE·DEFAULT-LANGUAGE, BODY의 독립 Value |
| `DocID`, `ImportDoc` | DOC-ID, IMPORT-DOC (버전 ID 2 WORD, 0이면 최신) |
| `Source`, `toulmin.TextPacket`, `ControlPacket`, `NodePacket`, `EdgePacket` | PROFILE이 `tangeul-1.x`일 때의 TOULMIN 레코드 |

- **무손실 원칙.** TOULMIN 레코드는 다시 인코딩한 결과가 원래 워드와 같을 때만 `Data`를 채운다. 그 밖에는 `Data`가 nil이고 원본 레코드만 남는다.
- 필드 폭은 Codec과 Length가 정한다. Edge Id와 REF는 LOCAL Reference 레코드다. SOURCE가 선언한 SPAN 폭은 그 뒤 레코드에 적용한다.
- PREDICATE-DECLARE(조건식, tangl AST)는 해석하지 않는다.

## 워드 배치

`Layout(e) []Part`는 레코드의 워드를 이름 붙은 조각으로 나눈다. 조각을 이어 붙이면 `e.Words`가 된다.

- 머리(`Head`): TAG, LENGTH(Length 확장), ISSUER, FORMAT.
- Payload: TOULMIN 레코드는 필드 단위(ID, header, src, dst, flags, span.off …). 내장 레코드(LOCAL Reference, Value)는 `Record`로 표시하며 첫 워드가 그 레코드의 TAG다.
- SPAN 폭은 Entry에 없으므로 1/2/4/8 중 Payload를 정확히 다 쓰는 폭을 고른다. 맞는 배치가 없으면 Edge는 앞의 LOCAL Edge Id, Meta는 내장 Value만 떼고 나머지는 `payload` 한 조각이다.
