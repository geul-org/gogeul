# resolve

다른 문서를 가리키는 두 경로를 모두 해석한다. 문서를 찾아 받고, 검증하고, 연결한 뒤 그 안의 대상 레코드를 찾는다.

| 경로 | 어디에 | 문서 지정 | 버전 |
|---|---|---|---|
| `IMPORT-DOC-*` + `EXTERNAL-POINTER-*` | HEAD에서 가져오고 BODY에서 IMPORT ID로 가리킨다 | IMPORT ID → Key | 고정 가능 |
| `EXTERNAL-ISSUER-*` (지시선언) | BODY의 Reference 한 개로 완결 | Reference 안의 Issuer + 문서 ID | 최신 |

규칙의 출처는 명세 `OVERVIEW`(IMPORT, Reference), `ISSUER`(가져오기), `GEUL-ORG`(ID 표기, 공표 주소, 위치 해석)다.

```
.geul ──inspect.Decode──▶ HEAD의 IMPORT-DOC-*  ─┐
                       └▶ BODY의 EXTERNAL-ISSUER-* ─┤ Key(Issuer, 문서 ID, 버전 ID)
                                                    ▼
                              문서별 버전 선택(IMPORT 고정 > 잠금 > 최신)
                                                    ▼
                        Path ─▶ Source(파일·pack·zstd pack, 로컬·HTTP) ─▶ 바이트
                                                    │ 내용 해시 검증, 잠금 기록
                                                    ▼
                        Graph: 문서 + 의존 문서 (재귀) ─▶ Target(Reference) ─▶ 대상 레코드
```

## 원칙

- 조회는 명시적 단계다. `core`·`inspect`의 디코드는 조회하지 않는다.
- 위치는 문서에 넣지 않는다. 문서는 ID만 적고, ID를 위치로 바꾸는 것은 이 패키지와 사용자 설정이 한다.
- 받은 곳을 믿지 않는다. 버전 ID의 내용 해시(잠금 기록 또는 서명된 버전 목록)로 검증하므로 미러·캐시·CDN 어디서 받아도 된다.
- 네트워크는 HTTP Source를 주입했을 때만 쓴다. 기본은 오프라인(아카이브)이다.
- 실제로 쓰는 참조만 조회한다. 깊이·문서 수·크기에 상한을 두고 설정으로 바꿀 수 있다.

## 주소

`Key{Issuer, Doc, Version}`는 두 곳에서 만든다. Issuer 폭은 2/4/16 WORD, 문서 ID는 4 WORD(코덱은 2 WORD 코덱 ID), 버전 ID는 2 WORD 순번이며 0이면 최신이다.

| 출처 | 레이아웃 | Key |
|---|---|---|
| `IMPORT-DOC` (Meta TAG `0x1E4L`) | IMPORT ID + 대상 포인터(DOC-POINTER 또는 CODEC-POINTER) | 그대로 |
| `EXTERNAL-ISSUER-*` (Reference TAG `0x02FL`·`0x034L`·`0x039L`, 5Q 체계 9~11) | TAG + Issuer + 문서 ID(4) + LOCAL ID(나머지) | 버전 ID 0(최신) |

- `EXTERNAL-ISSUER-*`는 BODY의 독립 Reference일 수도, Edge Payload에 내장된 Reference일 수도 있다. 독립 레코드는 코덱북 없이 모은다. 내장된 것은 `inspect.Layout`이 Record로 표시한 부분 중 Kind가 Reference인 것을 모은다. 레이아웃을 모르는 Edge 안의 참조는 찾지 못하며 미해결 목록에 그 Edge를 남긴다.
- Quantifier가 Variable이면 대상은 그 문서의 LOCAL 변수다. 해석 방법은 같다.

### 공표 주소 (migration.md §3.13)

ID는 소문자 `0x` + A-F 대문자 hex로 쓰고(§1.1 120), 자릿수 = 폭 × 4다. 주소 입력(`ParseAddr`, `Dir.Find`, `Dir.Lookup`)은 소문자·섞인 표기도 받아 대문자로 정규화해 찾는다(§1.1 122). 문서 ID는 16자리, 코덱 ID는 8자리, 버전 ID는 8자리다. 쓸 때 Issuer는 발급 폭 자릿수(2 WORD 8자리, 4 WORD 16자리, 16 WORD 64자리)로 쓰고, 읽을 때는 8·16·64자리 0 채움 표기를 같은 Issuer로 받는다(`ParseAddr`가 정규화).

주소는 다음 10가지뿐이다(`<기준>/.well-known/geul/` 아래). 그 밖의 꼴(`<I>/docs/`·`<I>/codecs/` 목록 포함, §1.1 101), 이름 주소(`geul.org/triple`), 도메인 정규화는 공표 주소가 아니다(`ParseAddr` 오류, 조회 404). Issuer 이름은 CLI·해석기가 Issuer 목록의 `BIND`로 ID로 바꾼다(`Dir.Find`, `Dir.Lookup`). 문서 이름 색인은 없다(문서·코덱은 ID로만 찾는다).

| 대상 | 최신 | 버전 | 버전 목록 | 함수 |
|---|---|---|---|---|
| Issuer 목록 | `issuers/latest.geul` | `issuers/<V>.geul` | `issuers/versions.geul` | `Issuers`, `IssuersVersions` |
| 코덱 | `<I>/codecs/<C>/latest.geul` | `<I>/codecs/<C>/<V>.geul` | `<I>/codecs/<C>/versions.geul` | `Codec`, `CodecVersions` |
| 문서 | `<I>/docs/<D>/latest.geul` | `<I>/docs/<D>/<V>.geul` | `<I>/docs/<D>/versions.geul` | `Doc`, `DocVersions` |
| 내용 해시 | | `<I>/hashes/<H>.geul` | | `Hash` |

- `ParseAddr`는 `file://…`, `https://…`, `.well-known/geul/…`, 맨 경로를 받아 `Addr`로, `Addr.Path()`는 저장 경로로 바꾼다. `Fetch(ctx, src, 주소 문자열)`·`FetchAddr`가 바이트를 낸다.
- `latest.geul`은 저장하지 않는 별칭이다. 문서·코덱·Issuer 목록 모두 `versions.geul`에서 폐기 권고·삭제가 아닌 가장 높은 버전을 찾아 그 바이트를 낸다(리다이렉트 없음).
- 목록은 `issuers/` 하나다. 버전 1부터이고 예약 문서 ID `0xFFFFFFFFFFFFFFFF`(Issuer `0x00000000`, `core.DocIDIssuers`)로 버전 목록을 갖는다. `versions.geul`의 `PUBLICATION`은 `DOC-POINTER(0x00000000, 예약 ID, 버전)`다. 이 예약 ID는 모든 Issuer에서 `<I>/docs/<D>/` 주소로 받지 않고(404) 일반 문서로 공표하지 않는다. 문서 ID `0x0000000000000000`도 모든 Issuer의 예약 ID로, 용어 정의 색인 자리(`core.DocIDTermIndex`)이며 보통 문서처럼 공표·조회한다(§1.1 121). 공표는 `Dir.PublishList`다. `<I>/docs/`·`<I>/codecs/` 목록은 없다(§1.1 101). `Dir.Documents`는 저장 폴더(pack 인벤토리)에서 문서·코덱 ID를 나열한다.
- 코덱 문서의 `DOC-ID`는 `CODEC-POINTER(Issuer, 코덱 ID, 버전 0)`이고 `<I>/codecs/<C>/`에 놓인다(`Dir.PublishCodec`).
- Issuer는 코덱 ID를 Kind와 상관없이 자기 `codecs/` 안에서 겹치지 않게 배정한다(§1.1 103). 그래서 두 가지를 기계로 검사한다. 해석: 레코드로 찾은 코덱 문서의 `DECLARE` 대상 Kind가 레코드 Kind와 다르면 오류(`checkDeclareKind`, `Dir.CodecSpecs`·`Resolver.CodecSpecs`). 공표: 같은 Issuer·같은 코덱 ID에 이미 있는(삭제되지 않은) 버전과 `DECLARE` 대상 TAG가 다르면 거부(`Dir.sameCodec`, `Publish`·`PublishCodec`).
- 옛 경로(`issuers.geul`, `documents.geul`, `h/`, `<I>/<문서 ID>/`, `<I>/<문서 ID>.geul`)는 읽지도 쓰지도 않는다.

## 설정

위치는 해석하는 쪽의 설정이 정한다. 규격은 명세 `GEUL-ORG` 위치 해석 절이다.

```go
type Config struct { // 설계 전체. 지금은 Proxy·Private만 있다(구현 상태)
	Proxy    []Entry   // GEULPROXY. 기본 https://geul.org,registry
	Private  []Pattern // GEULPRIVATE
	NoLog    []Pattern // GEULNOLOG. 비면 Private
	Insecure []Pattern // GEULINSECURE
	Log      *LogConfig // GEULLOG. nil이면 off
	Work     *Workspace // geul.work (GEULWORK)
}
type Entry struct {
	Name        string // 적힌 그대로. Node.Source에 남는다
	Source      Source // https://·file:// → Dir(저장 형태는 읽을 때 판별), registry → Registry
	Public      bool   // https://geul.org, registry. Private 문서는 건너뛴다
	Fallthrough bool   // 뒤 구분자가 | 이면 어떤 오류에도 다음 출처로
	Off         bool   // off: 여기서 멈춘다
}
func ConfigFromEnv(getenv func(string) string) (Config, error)
```

- `Pattern`은 `<Issuer>[/<문서 ID>]` 또는 `*`다. 마디 전체 비교만 한다.
- 환경변수와 사용자 설정 파일을 읽어 `Config`를 만든다. 환경변수가 우선한다.
- 작업공간(`geul.work`)은 `use <파일|디렉터리>`, `replace <패턴>[@<버전 ID>] => <로컬 파일 | Issuer/문서 ID[@버전 ID]>`를 읽는다. `GEULWORK=off`면 쓰지 않고, 비면 현재 디렉터리부터 위로 찾는다.

## Source

```go
type Source interface {
	Fetch(ctx context.Context, path string) ([]byte, error) // path는 .well-known/geul/ 아래 상대 경로
}
```

- 해석기는 최신 주소(`<문서 ID>.geul`, 302)를 쓰지 않는다. 버전 목록을 받아 버전을 고르고 그 정본 주소를 받는다. 그래서 `Dir`과 `HTTP`가 같은 순서로 동작하고, 받은 내용은 언제나 목록의 내용 해시로 검증된다.

| 구현 | 출처 항목 | 내용 |
|---|---|---|
| `Dir` | `file://<디렉터리>`, `https://<기준>` | 아카이브 하나. `Root`만 주면 디렉터리 트리 `<Root>/.well-known/geul/`이고, `Store`를 주면 그 Store를 읽는다. `Root`와 `Store`를 함께 주면 Store가 없는 경로를 `Root`의 낱개 파일에서 읽는다. `Publish`로 문서를 공표한다: 버전 배정, 문서와 `h/` 사본 저장, 새 목록 문서(`PUBLICATION` 덧붙임, `PARENT` = 직전 목록) 저장, 직전 목록은 `h/`에 남는다. 출처 목록의 항목은 읽기 전용이다. |
| `HTTP` | (`https://`의 낱개 GET) | 기준 주소의 `.well-known/geul/<경로>`를 한 번 GET 한다. 평문 `http://`는 `GEULINSECURE`가 생기기 전까지 받지 않는다. |
| `Registry` | `registry` | geul.org 등록부에서 Issuer의 기준 주소를 받아 `HTTP`로 받는다. 등록부 API가 정해지기 전까지 모든 조회가 '없음'이다. |
| `Memory` | (테스트) | 경로 → 바이트 맵. |

- 오류 분류: `404`·`410`은 "없음"(다음 출처로), 그 밖은 "고장"(구분자가 `|`일 때만 다음 출처로). `Dir`의 파일 없음은 "없음"이다.

## 저장 형태 판별

출처 목록의 `file://`와 `https://`는 어디서든 낱개 `.geul` 파일, 무압축 pack 세트, zstd pack 세트를 읽는다. 형태는 적지 않는다. 항목을 파싱할 때는 아무것도 읽지 않고, 읽을 때 찾는다(`site`). 데이터 폴더는 `pack.OpenDir`로 연 `Archive`가 맡고 규칙은 같다.

경로 하나를 찾는 순서는 다음과 같다. `pack.Archive`의 읽기 순서를 따르고, 맨 뒤에 낱개 파일이 붙는다.

| 경로 | 찾는 곳 |
|---|---|
| `<Issuer>/…`(`hashes/` 포함) | `<Issuer>/latest.txt`의 세트 → 맨 위 `latest.txt`의 세트 → 낱개 파일 |
| `issuers/…` | 맨 위 세트 → `0x00000000` 세트 → 그 밖의 세트 → 낱개 파일 |

- **세트의 판별은 `latest.txt`뿐이다.** `latest.txt`가 없으면(파일 없음, `404`·`410`) 세트가 없는 것이다. 있는데 읽히지 않거나 가리키는 매니페스트가 없으면 고장이다. 고장난 세트를 낱개 파일로 덮지 않으며, `,` 뒤에서는 거기서 멈추고 `|` 뒤에서는 다음 출처로 간다. 다른 Issuer의 세트가 고장나도 이 Issuer의 조회는 영향을 받지 않는다.
- **혼재.** 경로마다 찾으므로 Issuer별로 세트와 파일을 섞을 수 있고, 같은 Issuer의 세트 옆에 낱개 문서를 둘 수도 있다. 둘 다 있는 경로는 세트가 이긴다.
- **그 밖의 세트**는 폴더면 Issuer 폴더를 나열해 찾는다. 서버는 나열이 없으므로 Issuer 목록(`issuers/latest.geul`)이 묶은 Issuer와 이미 연 세트만 본다. `PARENT` 사슬은 목록을 읽을 때 세트가 열리므로 영향이 없다.
- 세트는 한 번 열면 그 출처 값이 사는 동안 그 세대로 읽는다. '없음'도 기억하고, 고장은 기억하지 않는다.
- 나열(`Issuers`, `Documents`, `Tree`)은 로컬에서만 된다. 원격은 오류다. 출처 목록 항목에 쓰기는 거부한다.
- `Dir.Fetch`는 Store가 `Source`이면 호출자의 context를 넘긴다. 그래서 `Find`·`At`·`Resolver`가 세 형태에서 같은 코드로 돈다.

원격 비용(`httptest` 위 공표 표본으로 측정, 위키데이터 세트는 문서 약 9,600개):

| 조회 | 요청 | 받은 바이트 |
|---|---|---|
| 첫 문서(세트 열기 포함, Q42) | 6: `latest.txt` 1, 매니페스트 1, 색인 Range 1, 사전 Range 1, 목록·문서 Range 2 | 158,743 |
| 같은 세트의 다음 문서 | Range 1~2 | 문서 크기 |
| 세트에 없는 문서 | 낱개 GET 1(`404`), 처음이면 맨 위 `latest.txt` 1 | 수백 |
| 낱개 파일 트리 | Issuer마다 처음 한 번 탐침 2(`<Issuer>/latest.txt`, `latest.txt`), 뒤로 문서당 GET 1 | 문서 크기 |

전체 세트를 받지 않는다. `o/`의 객체는 언제나 Range로만 읽고, 서버가 Range를 무시하고 `200`을 주면 `pack.ErrRange`로 거부한다(고장). 힙은 매니페스트, 사전, 디렉터리 블록 캐시(세트당 최대 `pack.DefaultCacheBlocks`)다.

`pkg/resolve`가 `pkg/pack`을 import한다. `file://`와 `https://`가 등록이나 주입 없이 이 패키지만으로 세 형태를 읽게 하려는 것이다. `pkg/pack`은 `pkg/resolve`에 의존하지 않고, `ErrNotFound`·`ErrVerify`는 `pkg/pack`이 정의해 `resolve`가 같은 값으로 이름을 붙인다.

## 위치 해석

Key 하나를 받는 순서(GEUL-ORG 위치 해석 순서).

1. **작업공간**: `replace`가 맞으면 그 대상, `use` 문서의 `DOC-ID`가 맞으면 그 파일. 결과 Node에 `Replaced`를 표시한다.
2. **캐시(`Dir`)**: 버전을 고정한 Key와 `<I>/hashes/` 경로만. 불변 주소라서 먼저 본다.
3. **출처 목록**: `Proxy`를 차례로. `Private`에 맞으면 `https://geul.org` 항목과 `registry`를 건너뛴다. `off`에서 멈춘다.
4. **버전 로그**: 최신을 받았으면 `Log`로 확인한다. `NoLog`에 맞거나 `Log`가 nil이면 건너뛴다. 로그에 닿지 못하면 실패다.

- 16 WORD Issuer는 등록부가 없어 `registry`가 404를 준다. 출처 목록이나 작업공간으로 찾는다.
- 받은 정본은 캐시에 `<I>/docs/<문서 ID>/<버전 ID>.geul`과 `<I>/hashes/<내용 해시>.geul`로 넣는다. 작업공간 문서는 캐시에 넣지 않는다.

## 버전 선택

각 `IMPORT-DOC-*`은 자기가 고정한 판본에 바인딩된다. 새 버전은 호환을 약속하지 않으므로(OVERVIEW 계보 개정) 그래프 전체에서 버전을 하나로 합치지 않는다. 같은 문서의 여러 판본이 한 그래프에 함께 있을 수 있고, Node는 (문서, 버전)마다 하나다.

1. 작업공간의 `use`·`replace`가 맞으면 그것.
2. `IMPORT-DOC-*`이 버전을 고정했으면 그 버전.
3. 고정이 없으면 잠금 기록의 버전.
4. 잠금도 없으면 최신을 받고(버전 로그로 확인), 받은 버전 ID와 내용 해시를 잠금에 남긴다.
5. 시점 T를 정한 해석(`asOf`)이면 최신 대신 T에 시행 중인 판본을 고른다(`EFFECTIVE-FROM`·`EFFECTIVE-UNTIL`, 조문별로는 대상 Edge에 붙은 QUALIFIER P580·P582).

- `EXTERNAL-ISSUER-*`는 명세상 "최신을 따른다". 이 패키지에서는 4번 규칙 안의 "최신"으로 다룬다. 같은 문서를 IMPORT로도 가져왔다면 IMPORT의 버전을 쓴다. 잠금 덕분에 지시선언도 재현된다.
- 지시선언이 가리키는 문서의 IMPORT는 그 문서의 대상을 실제로 따라갈 때만 내려간다(지연 해석).

## 검증

- **내용 해시**: 여는 레코드(`DOC-*`)부터 `END` 직전까지의 SHA-256이다. 검증 클래스(`100`) Meta는 계산에서 빼고, 여는 레코드의 Length와 폭은 뺀 길이에 맞춰 다시 계산한다(ISSUER 서명 절).
- **버전 ID → 내용 해시 대응**: 잠금 기록에 있으면 그것을 쓴다. 없으면 받은 출처의 버전 목록에서 얻고 `Log`로 확인한 뒤 잠금에 남긴다(Go `go.sum`). `NoLog`에 맞거나 `Log`가 nil이면 출처의 버전 목록을 믿고 잠금에 남기며 경고한다.
- 받은 바이트의 내용 해시가 대응과 같아야 한다. 다르면 거부한다. 잠금과 버전 목록이 다르면 거부한다(같은 버전 ID, 다른 내용).
- 받은 문서의 `DOC-ID`는 Key의 Issuer·문서 ID와 같아야 한다.
- 작업공간에서 로컬 파일로 대체한 문서는 고정 버전과 달라도 쓰고 경고한다. `Replaced`로 표시하며 잠금 기록은 바꾸지 않는다.
- 최신을 받은 Key는 받은 버전 ID와 계산한 내용 해시를 잠금 기록에 남기고 경고한다. 다음 해석부터는 잠금의 버전으로 고정된다.
- 서명 검증(검증 클래스 Meta, Issuer 키)은 이 패키지 밖이다. 검증기를 주입할 수 있게만 둔다.

## 해석 결과

```go
type DocKey struct{ Issuer, Doc string } // hex 정규형. 버전과 무관한 문서 정체

type Graph struct {
	Root  *Node
	Docs  map[DocKey]*Node // 문서마다 하나. IMPORT와 지시선언이 같은 Node를 공유한다
}
type Node struct {
	Key        Key
	Version    uint32              // 실제로 받은 버전 ID
	Hash       [32]byte            // 계산한 내용 해시
	Doc        inspect.Document
	Imports    map[uint64]*Node    // IMPORT ID → 의존 문서
	Replaced   string              // 작업공간에서 대체됐으면 대체 대상(파일 경로 또는 Key), 아니면 ""
	Source     string              // 받은 출처(디버깅·감사용)
	Unresolved []Miss              // 찾지 못한 참조와 이유
}

func (g *Graph) Target(from *Node, ref core.Record) (*Node, inspect.Entry, error)
```

- 재귀로 내려간다. 같은 문서(DocKey)는 한 번만 받고 공유한다. 순환은 DocKey 경로로 찾아 거부한다.
- `Target`은 Reference 하나를 대상 레코드로 바꾼다. 참조가 놓인 문서(`from`)가 기준이다.
  - `LOCAL` → `from` 문서의 Edge Id(LOCAL) 색인.
  - `EXTERNAL-POINTER-*` = IMPORT ID + LOCAL ID → `from.Imports[IMPORT ID]`의 색인.
  - `EXTERNAL-ISSUER-*` = Issuer + 문서 ID + LOCAL ID → `Docs[DocKey]`의 색인. 아직 받지 않았으면 버전 선택 규칙으로 받는다.
- 지시선언이 가리키는 문서 안의 `EXTERNAL-POINTER-*`는 그 문서의 IMPORT ID 공간을 쓴다. 기준 문서가 바뀌므로 `Target`의 결과 Node를 다음 `from`으로 넘긴다.
- 찾지 못한 대상은 오류가 아니라 미해결 목록으로 돌려준다. 뷰어는 그대로 보여 주고, 실행기는 거부한다.

## 의존성 해석

- 문서는 HEAD의 `IMPORT-DOC-*`으로 의존 문서를 적는다(Go `require`).
- 새 버전은 호환을 약속하지 않는다. 같은 문서를 여러 버전으로 요구하면 각 판본을 따로 받아 각 IMPORT에 바인딩한다. 최소 버전 선택은 쓰지 않는다.
- 버전 ID로 고정한 가져오기는 재현된다. 버전 ID가 0이면 잠금 기록이 재현을 맡는다.

## 구현 상태

| 됨 | 파일 |
|---|---|
| Key, DocKey, 주소(`versionPath`, `listPath`, `hashPath`, `FormatVersion`) | `key*.go`, `doc_key.go`, `*_path.go`, `format_version.go` |
| 버전 목록 문서(edge/PUBLICATION): 읽기·쓰기, DEPRECATION·TOMBSTONE 반영(BODY는 이 셋만. `BIND`는 오류, §1.1 102) | `version_entry.go`, `version_log.go`, `decode_version_list.go`, `encode_version_list.go`, `apply_list_record.go`, `add_publication.go`, `mark_version.go`, `pick_version.go` |
| Source: `Memory`, `Dir`(+`Publish`), `HTTP`, `Registry`(빈 구현) | `memory*.go`, `dir*.go`, `http*.go`, `registry*.go` |
| 저장 형태 판별: 세트를 열며 찾는 `site`, Store 옆의 낱개 파일 `overlay` | `site*.go`, `local_site.go`, `remote_site.go`, `sub_store*.go`, `overlay*.go` |
| 설정: `ConfigFromEnv`(GEULPROXY 기본 `https://geul.org,registry`, GEULPRIVATE), `Pattern` | `config*.go`, `parse_entry.go`, `parse_proxy.go`, `*pattern*.go` |
| `Resolver.Fetch`: 출처 순서, `,`/`\|`, Private 건너뛰기, `off`, 검증 실패는 즉시 중단 | `resolver_fetch.go`, `fetch_from.go`, `verify_doc.go` |
| `Resolver.Resolve`: IMPORT-DOC 재귀, 각 IMPORT를 자기 판본에 바인딩, (문서, 버전)마다 Node 하나 | `resolver_resolve.go`, `resolver_collect.go`, `selection*.go`, `node_key.go`, `link_imports.go` |
| `Resolver.AsOf`: 최신(버전 0) IMPORT를 그 시점에 시행 중인 판본에 바인딩(`EFFECTIVE-FROM`·`EFFECTIVE-UNTIL`) | `resolver_resolve_key.go`, `fetch_as_of.go`, `effective_*.go`, `meta_time.go`, `core/decode_time_value.go` |
| 내용 해시 | `core/content_hash.go` |

아직 없는 것:

1. 잠금 기록, 작업공간(`geul.work`), 캐시(받은 문서를 `Dir`에 저장), `GEULLOG`·`GEULNOLOG`·`GEULINSECURE`.
2. 조문별 시행 기간(대상 Edge의 QUALIFIER P580·P582). 시행일 없는 판본은 처음부터 시행 중으로 본다(공표 시각 `RECORDED-AT`은 아직 쓰지 않는다).
3. 상한(깊이·문서 수·크기). 순환은 같은 Key를 한 번만 받으므로 끝나며, 거부하지 않는다.
4. `core`에 지시·지시선언 Reference 디코드(Issuer·문서 ID·IMPORT ID·LOCAL ID 분리). 지금은 LOCAL만 있다.
5. 참조 수집(독립 Reference + `inspect.Layout`의 내장 Reference), `EXTERNAL-ISSUER-*` 해석, `Target`.
6. 버전 목록 서명·`PARENT` 사슬 확인, `Registry` 조회(그 API·형식이 정해진 뒤).

## 임시 공표 위치

geul.org가 서비스되기 전까지 geul.org 웹사이트 저장소의 `artifacts/geul/.well-known/geul/`을 공표 트리로 쓴다. `GEULPROXY=file://<geul-dev>/draft/geul-org/artifacts/geul,off`로 가리킨다. 기본값은 그대로 `https://geul.org`다.

- migration1 뒤에는 Issuer `0x00000000`(geul.org)의 규격 세트와 Issuer `0x00000001`(wikidata.geul.org)의 위키데이터 세트가 Issuer별 pack 세트로 들어간다(5·6단계 변환 산출물).
- `TestResolvePublished`, `TestDirFindPublished`가 이 트리를 `GEULPROXY`로 읽는다(뒤의 것은 같은 트리를 HTTP Range로도 읽어 대조한다). 이 모듈만 따로 받았거나 `0x00000000/latest.txt`(뒤의 것은 `0x00000001/latest.txt`도)가 없으면 건너뛴다. 전체 세트를 훑지 않고 점 조회만 한다.
- 공표는 이 트리를 데이터 폴더로 한 `geul publish`, 또는 `resolve.Dir{Root: "<geul-dev>/draft/geul-org/artifacts/geul", Store: <pack.OpenDir로 연 세트>}.Publish(bytes)`로 한다.

## 범위 밖

- TANGEUL 원문의 패키지 경로(`geul/schema/frame`)를 Key로 바꾸는 매핑. 인코더(toulmin)의 설정이 맡는다. 문서에는 Key만 들어간다.
- 서명 검증, 버전 목록 서명 확인.

## 미결정

- 지시선언의 "최신"을 잠금으로 고정하는 이 패키지의 해석이 명세 의도와 맞는지. 맞다면 ISSUER의 가져오기 절에 잠금을 적는다.
- 잠금 기록의 형식과 위치. 작업공간 파일 옆에 둘지, 사용자 설정 파일의 위치와 함께 정한다. 텍스트 한 줄 `<Issuer>/<문서 ID> <버전 ID> <내용 해시>`(Go `go.sum`)로 둘지, GEUL 문서로 둘지.
- geul.org 등록부 조회 API(Issuer ID → 기준 주소).
- 위치 제공 Issuer를 HEAD에 적는 Meta(서명키 조회처와 같은 방식).
- Issuer 이름 레이블과 alias, 도메인 이름 주소(GEUL-ORG 도메인 이름)를 Key로 바꾸는 입력을 받을지.
- 텍스트 표현과 콘텐츠 협상(`Accept`에 따라 바이너리, 텍스트, HTML 보기).
- 덤프 주기와 형식.
