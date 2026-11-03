//ff:func feature=resolve type=helper control=sequence
//ff:what namesPath — where an Issuer's document list (BIND names, the old documents.geul) is kept: `<Issuer>/docs/0x00000001.geul`
package resolve

// namesPath is the stored address of issuer's document list (GEUL-ORG 문서
// 이름; `<Issuer>/docs/` list, listDocPath).
func namesPath(issuer string) string {
	return listDocPath(issuer + "/" + docsDir)
}
