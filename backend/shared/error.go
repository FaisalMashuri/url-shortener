package shared

// variabel const untuk error response
const (
	//variabel const untuk error pada autentikasi
	RespSuccess = "0200"
)

// variabel const untuk error response
const (
	//variabel const untuk error pada autentikasi
	ErrCodeCredentialsRequired = "0101"
	ErrCodeErrorSaveCredential = "0102"

	//variabel const untuk error pada user
	ErrCodeFailCreateUser          = "0201"
	ErrCodeFailUpdateUser          = "0202"
	ErrCodeFailDeleteUser          = "0203"
	ErrCodeFailGetUserByPN         = "0204"
	ErrCodeFailGetAllUser          = "0205"
	ErrCodeUserIdNotFound          = "0206"
	ErrCodeUserIDAlreadyRegistered = "0207"
	ErrCodeUserPnAlreadyRegistered = "0208"

	//variabel const untuk error pada submission
	ErrCodeFailCreateSubmission          = "0301"
	ErrCodeFailUpdateSubmission          = "0302"
	ErrCodeFailDeleteSubmission          = "0303"
	ErrCodeFailGetSubmissionById         = "0304"
	ErrCodeFailGetAllSubmission          = "0305"
	ErrCodeSubmissionIdNotFound          = "0306"
	ErrCodeSubmissionIDAlreadyRegistered = "0307"
	ErrMimeTypeFileNotSupport            = "0308"
	ErrFileSizeTooLarge                  = "0309"
	//variabel const untuk error pada partner
	ErrCodeFailUpdatePartner             = "0402"
	ErrCodeFailDeletePartner             = "0403"
	ErrCodeFailGetPartnerById            = "0404"
	ErrCodeFailGetAllPartner             = "0405"
	ErrCodePartnerIdNotFound             = "0406"
	ErrCodePartnerIDAlreadyRegistered    = "0407"
	ErrCodeProviderIDAlreadyRegistered   = "0408"
	ErrCodePartnerFieldValidationError   = "0409"
	ErrCodeFailGetTokenPartner           = "0410"
	ErrCodeFailCreatePartner             = "0411"
	ErrDocumentSubmissionPartnerNotFound = "0412"

	//variabel const untuk error pada product
	ErrCodeFailCreateProduct          = "0601"
	ErrCodeFailUpdateProduct          = "0602"
	ErrCodeFailDeleteProduct          = "0603"
	ErrCodeFailGetProductById         = "0604"
	ErrCodeFailGetAllProduct          = "0605"
	ErrCodeProductNotFound            = "0606"
	ErrCodeProductIDAlreadyRegistered = "0607"

	//variabel const untuk error pada briva
	ErrCodeFailCallBackUpdateStatusCorpcode = "0702"
	ErrCodeFailCallBackCorpcodeNotFound     = "0704"

	//other errors
	Unauthorized             = "0401"
	ErrCodeServerError       = "0500"
	ErrCodeTimeout           = "0501"
	ErrCodeErrorValidation   = "0502"
	ErrFileLarge             = "0503"
	ErrInvalidFile           = "0504"
	ErrFileEmpty             = "0505"
	ErrUnavailableService    = "0506"
	ErrInvalidParam          = "0507"
	ErrDataNotExist          = "0509"
	ErrUnexpectedError       = "0510"
	ErrInvalidAccount        = "0511"
	ErrInactiveAccount       = "0512"
	ErrTimeOut               = "0513"
	ErrInvalidRequestFamily  = "0514"
	ErrInvalidFieldFormat    = "0515"
	ErrInvalidFieldMandatory = "0516"
	ErrCodeEmptyParameter    = "0517"
	ErrDocStatus             = "0518"
)

var ErrMessages = map[string]string{
	ErrDocStatus: "Document status %s does not have request status 'Berhasil'",
}
