package account

type DocumentType uint

const (
	DocumentCPF = iota + 1
	DocumentCNPJ
)

func (dt DocumentType) String() string {
	switch dt {
	case DocumentCPF:
		return "document_cpf"
	case DocumentCNPJ:
		return "document_cnpj"
	default:
		return "unknown_document_type"
	}
}

func (dt DocumentType) Valid() bool {
	switch dt {
	case DocumentCPF, DocumentCNPJ:
		return true
	}

	return false
}

var documentTypeMap = map[string]DocumentType{
	"document_cpf":  DocumentCPF,
	"document_cnpj": DocumentCNPJ,
}

func ParseDocumentType(document string) (DocumentType, error) {
	value, ok := documentTypeMap[document]
	if ok {
		return value, nil
	}

	return 0, ErrUnknownDocumentType
}
