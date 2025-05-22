package constants

// infrastructure
const (
	ErrConnectionToDB     = "failed to connect to db: %w"
	ErrLoadENV            = "failed to load .env file: %w"
	ErrConvertStringToInt = "error when convert string to int: %w"
)

// grpc client
const (
	ErrConnectionToGrpcClient = "error when connect to grpc client service '%s', err: %v"
)

const (
	SUCCESS = "Success"
	FAILED  = "Failed"
	ERROR   = "Error"
)

const (
	RESPONSE = "response"
	REQUEST  = "request"
	SERVICE  = "service"
	QUERY    = "query"
)

const (
	InfoAuthLoggedIn            = "IF1015"
	InfoAuthLoggedOut           = "IF1016"
	InfoCreated                 = "IF1001"
	InfoUpdated                 = "IF1002"
	InfoDeleted                 = "IF1003"
	InfoUploaded                = "IF1004"
	InfoDownloaded              = "IF1005"
	InfoRetrieved               = "IF1006"
	InfoAccountRegistered       = "IF1011"
	InfoAccountVerified         = "IF1012"
	InfoAccountPasswordChanged  = "IF1013"
	InfoAccountPasswordResetted = "IF1014"
)

const (
	InfoAuthLoggedInDescription            = "Login Succeed"
	InfoAuthLoggedOutDescription           = "Logout Succeed"
	InfoCreatedDescription                 = "Resource Created"
	InfoUpdatedDescription                 = "Resource Updated"
	InfoDeletedDescription                 = "Resource Deleted"
	InfoUploadedDescription                = "Resource Updated"
	InfoDownloadedDescription              = "Resource Downloaded"
	InfoRetrievedDescription               = "Resource Retrieved"
	InfoAccountRegisteredDescription       = "Account Registered"
	InfoAccountVerifiedDescription         = "Account Verified"
	InfoAccountPasswordChangedDescription  = "Account Password Changed"
	InfoAccountPasswordResettedDescription = "Account Password Resetted"
)

const (
	InfoCreatedMessage                 = "Proses Penyimpanan Data Berhasil"
	InfoUpdatedMessage                 = "Proses Perubahan Data Berhasil"
	InfoDeletedMessage                 = "Proses Penghapusan Data Berhasil"
	InfoUploadedMessage                = "Proses Unggah Data Berhasil"
	InfoRetrievedMessage               = "Proses Baca Data Berhasil"
	InfoDownloadedMessage              = "Proses Unduh Data Berhasil"
	InfoAccountRegisteredMessage       = "Proses Pendaftaran Akun Berhasil"
	InfoAccountVerifiedMessage         = "Proses Verifikasi Akun Berhasil"
	InfoAccountPasswordChangedMessage  = "Proses Perubahan Password Berhasil"
	InfoAccountPasswordResettedMessage = "Proses Reset Password Berhasil"
)

const (
	InfoSucessGetJabatan = "Get data jabatan berhasil"
)

const (
	InfoSucessGetPrivilege = "data privilege ditemukan"
)

const (
	InfoSucessGetAplikasi = "data aplikasi ditemukan"
)

const (
	DefaultFormatDate      = "2006-01-02"
	DefaultFormatTimeStamp = "2006-01-02 15:04:05"
)

const (
	TokenType = "Bearer"
)

const (
	Pong = "pong!"
)

const (
	Healthy    = "healthy"
	NotHealthy = "not healthy"
)

const (
	Connected    = "connected"
	Disconnected = "disconnected"
)

type contextKey int

const (
	ContextKeyFiberContext contextKey = iota
	ContextKeyTRID
	ContextKeyTCID
	ContextKeyRequest
	ContextKeyTimeExecutionStart
	ContextKeyTimeExecutionEnd
	ContextKeyConfigToken
)

const (
	PrefixForTRID = "TR"
	PrefixForTCID = "TC"
)

// Header type
const (
	ContentType = "Content-Type"
)

// Token
const (
	AuthorizationHeaderKey = "Authorization"
	AuthorizationType      = "Bearer"
	MinSecretKeySize       = 32
)

// Header
const (
	ProductIdHeaderKey  = "product_id"
	DeviceIdHeaderKey   = "device_id"
	RegisterIdHeaderKey = "register_id"
	ClientIpHeaderKey   = "client_ip"
	RememberMeHeaderKey = "remember_me"
)

const (
	DefaultCacheKeyATPrefix = "auth_access_token_%s"
	DefaultCacheKeyRTPrefix = "auth_refresh_token_%s"
)

const (
	SsoKey               = "sso_%s_%s"
	LoginRateLimitKey    = "login_threshold_%s"
	ForgotPasswordOtpKey = "forgot_password_otp_%s"
)
