// Used by chaincode on the blockchain and by the server
//
// Hashes are SHA256 Base64 encoded from JSON representation of data
// IMPORTANT: hashes are not derived from the dto/structs
//
// The immutable portion of the contract is the essence of a smart contract.
// Any time a change of state occurs in a contract, a stringified JSON of
// the immutable contract is supplied as an argument to a blockchain method,
// hashed and validated against what was published on blockchain and then parsed.
//
// The rules and definition combined with current state is how chaincode validates a proposed change.
// Validation occurs server side prior to submission to blockchain and in the chaincode.
//
// Stringified Json including formatting must be repeatable, so that hashes of blocks use in its construction are identical.
// Once instantiated, then the stringified JSON is stored in the database as a read only string,
// and thus hashing it always produces the same value.
// Parsing and then restringifying it for hash comparison is not dependable.
// Alternatives are Bencode (used by bit-torent) and cannonical JSON implementations.
package contract

import (
	"time"
)

const (
	Unknown = ""
	None    = ""

	Agreement       = "agreement"
	Approver        = "approver"
	Beneficiary     = "beneficiary"
	Contractual     = "contractual"
	Creator         = "creator"
	Notary          = "notary"
	Notifier        = "notifier"
	Proxy           = "proxy"
	ServiceProvider = "service-provider"
	Verifier        = "verifier"
	Signatory       = "signatory"
)

type SignPackageMethodId int64

const (
	SignPackageMethodId_Unknown SignPackageMethodId = 0

	// no constructed content items,
	// signs against viewing of contract block which contains the contract content items
	SignPackageMethodId_OriginalContent SignPackageMethodId = 1

	// signee views constructed content without signature placeholders
	// and signs against contract block which contains hash of constructed content
	// constructed content is a content item appended with the contract details
	SignPackageMethodId_Constructed SignPackageMethodId = 2

	// Placeholders, hashes, and contents are embedded in the PDF document
	// Once all signatures are collected, the PDF is saved as a Finalized PDF to the contract vault
	// The pdf contains all the information needed to validate the signatures.
	SignPackageMethodId_Embedded SignPackageMethodId = 3
)

// A container for the immutable portion of a contract.
// When ready for instantiation, it is sealed and hence forth read only.
// The Json is recorded into a read only record in the database using Contract ID as primary key,
//
//	and a hash of the Json is anchored onto the blockchain.
//
// The Json representation 'is the authorative' record for the contract.
type ImmutableContract struct { // contract_container
	Contract               ContractBlock           `json:"contract"`
	ContractHash           string                  `json:"contract_hash"`
	ContractSignatures     ContractSignatures      `json:"contract_signatures"`
	ContractSignaturesHash string                  `json:"contract_signatures_hash"`
	FinalizedContent       *ConstructedContentItem `json:"finalized_content"`
	SealedOnDate           time.Time               `json:"sealed_on_date"`
}

// Everything about a contract prior to consent phase.
// When the contract moves to consent (signatures), it is sealed and is read only.
// A hash of this block is chained into the contract signature block, and bound with each signature;
//
//	so if changed after any signature, then the signature would be invalidated.
type ContractBlock struct {
	ContractID          int64  `json:"contract_id"`
	SchemaVersion       int64  `json:"schema_version"` // version of this contract block schema
	Language            string `json:"language"`
	ContractFamilyId    int64  `json:"contract_family_id"`    // Must match what is in the definition.
	ContractTypeId      int64  `json:"contract_type_id"`      // The contract type for each contract is unique and never changes.  Must match what is in the definition.
	ContractTypeVersion int64  `json:"contract_type_version"` // version of the contract type used
	CreatedWithTierId   int64  `json:"created_with_tier_id"`  // the tier id of the user who created this contract
	ContractName        string `json:"contract_name"`         // localized contract name for the language used

	DisplayName string `json:"display_name"` // Name for this contract as entered by author
	Description string `json:"description"`  // Optinal description for this contract as entered by author

	Organizations    []ContractOrganization               `json:"organizations"`
	VirtualPositions []ContractParticipantVirtualPosition `json:"virtual_positions"`
	Participants     []ContractParticipant                `json:"participants"`

	ContractOptions ContractOptions `json:"contract_options"`
	ContentItems    []ContentItem   `json:"content_items"`
	SignatureMethod SignatureMethod `json:"signature_method"`
	StorageYears    int64           `json:"storage_years"`

	ProxyInstructions   *ContractProxyInstructions `json:"proxy_instructions"`   // only for a conditional release contract wih a proxy beneficiary
	ReleaseInstructions *ReleaseInstructionDetail  `json:"release_instructions"` // null if not a conditional release contract

	// If notary approval required for release instructions,
	// then the release instructions are hashed using SHA265 and the hash signed by the subskribo managed private key of the notary.
	// The encrypted hash is saved here using base64 encoding.
	// To verify, the public key of the notary is fetched from the discovery service using the SigningKeyIdentity field.
	// ReleaseInstructionsSignature *string `json:"release_instructions_signature"`

	// Charges and payment transactions are not shown in the contract for privacy reasons.
	// They are stored in the database and referenced by the contract id.
	// This is only for charges prior to sealing the contract block'
	// Hash is generated from ContractBlockPayment json
	ContractPaymentHash string `json:"contract_payment_hash"`

	Definition        ContractDefinition `json:"definition"`         // the definition used to validate required fields and values
	DefinitionVersion int64              `json:"definition_version"` // version of the definition schema used

	SealedOnDate time.Time `json:"sealed_on_date"` // Date contract is sealed and ready for consent
}

type ContractOptions struct {
	ExpiryDate    *time.Time `json:"expiry_date"`    // null if no expiry date
	EffectiveDate *time.Time `json:"effective_date"` // null if effective immeadiately upon consent

	// If true, then author can void contract prior to instantiation.
	VoidableByAuthorPriorToInstantiation bool `json:"voidable_by_author_prior_to_instantiation"`

	// If true, then author can void contract after instantiation.
	// Only applicable for conditional releases contracts if not yet released.
	VoidableByAuthor bool `json:"voidable_by_author"`

	// If true, then participants who have been granted this right can void contract.
	// Only applicable for conditional releases contracts if not yet released.
	VoidableByParticipants bool `json:"voidable_by_participants"`

	// If true, then notary, if one of the participants, can void contract.
	VoidableByNotary bool `json:"voidable_by_notary"`

	DaysToSign              int64 `json:"days_to_sign"`              // Number of days to sign contract once enters consent phase.
	MaxDaysToSign           int64 `json:"max_days_to_sign"`          // Max number of days which can be extended to sign contract once enters consent phase.
	AllowSignatureExtension bool  `json:"allow_signature_extension"` // If true, then author can extend signature expiry date.

	// if true then min kyc level for each role is overridden by MinKycLevelForAllRoles
	IsMinKycLevelForAllRoles bool  `json:"is_min_kyc_level_for_all_roles"`
	MinKycLevelForAllRoles   int64 `json:"min_kyc_level_for_all_roles"`

	// Key/Value pair options for this contract as entered by author.
	// Supports additional options in future without schema change.
	Options []ContractOption `json:"options"`
}

type ContractUserRoleDefinition struct {
	Role               string `json:"role"`
	Min                int64  `json:"min"`
	Max                int64  `json:"max"`
	IncludeRoleInCount string `json:"include_role_in_count"`
	MinKycLevel        int64  `json:"min_kyc_level"`
}

type ContractOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ContractParticipant struct {
	UserId          string                        `json:"user_id"` // user type + id, example c102, n48, ...
	Roles           []string                      `json:"roles"`   // creator, beneficiary, notary, approver, ...
	FullName        string                        `json:"full_name"`
	CanVoidContract bool                          `json:"can_void_contract"` // true if this participant can void the contract.
	KycLevel        int64                         `json:"kyc_level"`         // KYC level of the participant
	IdentityClaims  []ContractIdentityClaim       `json:"identity_claims"`
	Positions       []ContractParticipantPosition `json:"positions"`
}

type ContractIdentityClaim struct {
	IdentityClaimId int64  `json:"identity_claim_id"` // unique id for this claim
	Claim           string `json:"claim"`             // required. name, mobile-phone, etc
	Value           string `json:"value"`             // '...' if value is private. A separate record is kept of all current identity values for notary use
	Verifier        string `json:"verifier"`          // the identity of the verifier, Subskribo, ItsMe, etc.
	KycLevel        int64  `json:"kyc_level"`         // KYC level of the claim
}

// only included if position is current
type ContractParticipantPosition struct {

	// if negative, then this is a virtual organization added by the author for this contract
	OrgId         int64  `json:"org_id"`
	OrgLegalName  string `json:"org_legal_name"`
	OrgCommonName string `json:"org_common_name"`
	Position      string `json:"position"`
	VerLevel      int64  `json:"ver_level"`

	// true if the position was added by the author for this contract
	IsVirtual bool `json:"is_virtual"`
}

type ContentItem struct {
	ContentId       int64  `json:"content_id"`     // the primary key of content item in database
	ItemRole        string `json:"item_role"`      // example: conditional_release, agreement, ...
	PlainHash       string `json:"plain_hash"`     // SHA256 hash of content item prior to encryption
	PlainSaltBase64 string `json:"plain_salt"`     // salt used to encrypt content item
	EncryptedHash   string `json:"encrypted_hash"` // SHA256 hash of content item after encryption
}

type ConstructedContentItem struct {
	ContentId      int64   `json:"content_id"`       // the primary key of constructed content item in database
	OrigContentIds []int64 `json:"orig_content_ids"` // primary keys of original contract content items in database
	ItemRole       string  `json:"item_role"`        // constructed-agreement or finalized-agreement
	PlainHash      string  `json:"plain_hash"`       // SHA256 hash of content item prior to encryption
	EncryptedHash  string  `json:"encrypted_hash"`   // SHA256 hash of content item after encryption
	ConstructTypes string  `json:"construct_types"`  // example: "signature-placeholders,terms"
}

type ReleaseInstructionDetail struct {
	Instructions              string `json:"instructions"` // instructions for conditional release.
	IsCustomRelease           bool   `json:"is_custom_release"`
	StandardReleaseTemplateId int64  `json:"standard_release_template_id"`

	// only applicable if notary used with custom release instructions
	NotaryPackage *NotaryInstructPackage `json:"notary_package"`

	// if acceptance required from notary response, this is recorded here
	AcceptancePackage *CreatorAcceptancePackage `json:"acceptance_package"`

	NotarySignature     string `json:"notary_signature"`     // signature of notary for notary package if applicable
	AcceptanceSignature string `json:"acceptance_signature"` // creator signature of acceptance of notary response if applicable

	// value is set from contract definition and repeated here for convenience
	IsEvidenceRequiredForRelease bool `json:"is_evidence_required_for_release"`

	// only applicable when verifiers are used for release
	ConsensusMethod string `json:"consensus_method"`

	// minimum number of verifiers required for consensus for release
	// this is derived from consensus method and number of verifiers and finalized when draft contract ready for sealing.
	MinVerifiersForConsensus int64 `json:"min_verifiers_for_consensus"`
}

// required to be signed by notary if custom release instructions are used with a notary
type NotaryInstructPackage struct {
	ContractId           int64   `json:"contract_id"`
	RequestId            string  `json:"request_id"` // id of the request to notary
	NotaryId             string  `json:"notary_id"`
	SuppliedInstructions string  `json:"supplied_instructions"`
	ApprovedInstructions string  `json:"notary_instructions"`
	MessageToNotary      string  `json:"message_to_notary"`
	MessageFromNotary    string  `json:"message_from_notary"`
	ApprovalPayTransId   int64   `json:"approval_pay_trans_id"` // payment transaction from fee paid to notary for approval
	ApprovalState        string  `json:"approval_state"`        // none, approved, rejected, acceptance-required
	KeyInfo              KeyInfo `json:"key_info"`
	AdditionalFeeCents   int64   `json:"additional_fee_cents"`
	// if true, then AdditionalFeeCents will be 0 and the fee will be stored outside the package.for privacy reasons.
	AddFeeStoredOutside bool      `json:"add_fee_stored_outside"`
	SubmittedDate       time.Time `json:"submitted_date"`
	SealedOnDate        time.Time `json:"sealed_on_date"`
}

// required to be signed by creator if custom release instructions are used with a notary
type CreatorAcceptancePackage struct {
	ContractId          int64     `json:"contract_id"`
	NotarySignatureHash string    `json:"notary_signature_hash"`
	KeyInfo             KeyInfo   `json:"key_info"`
	AcceptedOnDate      time.Time `json:"accepted_on_date"`
}

// Each contract has a definition that defines the required fields and values.
// The current version for the contract being created is captured at time the contract block is sealed.
type ContractDefinition struct {
	ContractFamilyId int64 `json:"contract_family_id"`

	// The contract type for each contract is unique and never changes.
	// This is not the same as a contract family, which is a group of simular contract types.
	ContractType int64 `json:"contract_type"`

	// Anytime a change is made to a particualar contract type, the version is incremented.
	ContractTypeVersion int64 `json:"contract_type_version"`
	// The version of the contract definition schema used to validate the contract.
	// this is used to route the contract to the correct validation service.
	SchemaVersion int64 `json:"schema_version"`

	ContractNameEnglish string `json:"contract_name_english"`

	Options ContractDefinitionOptions `json:"options"`

	// List of user roles and their min/max counts
	// example: if this contract does not use a notary, then the notary role has a max value of 0.
	// example: if at least 1 beneficiary is required, but no more than 10, then the beneficiary role has a min value of 1 and a max value of 10.
	UserRoles []ContractUserRoleDefinition `json:"user_roles"`

	// todo: add any additional validation rules

	// there are two methods for adding additional validation rules:
	// method 1: Attach a schema version attribute to the rule, so only applied to contracts using a schema version >= the attribute.
	// method 2: Route to validation service based on the contract schema version
	// which method we settle on, depends on how the version routing is developed on blockchain.
	// method 2 is easiest to implement and document in version history for smart contract, but method 1 reduces code duplication.
	// In all cases a new build is required when adding new validation rules.
}

type ContractDefinitionOptions struct {
	RequiredOptions   []string `json:"required_options"`   // a list of keys that must be present in the options list with non-empty values.
	DisallowedOptions []string `json:"disallowed_options"` // a list of keys that are ignored in the options list.

	// if true, if an option key is not in required or disallowed lists, then validation fails.
	// this is a failsafe that can be used to protect from mistakes when making a breaking change
	FailIfUnspecifiedOptions bool `json:"fail_if_unspecified_options"`

	EvidenceRequiredForConditionalRelease     bool `json:"evidence_required_for_conditional_release"`       // if true, then evidence must be entered by verifier for conditional release.
	AllowVoidableByAuthor                     bool `json:"allow_voidable_by_author"`                        // does not require consensus
	AllowVoidableByAuthorPriorToInstantiation bool `json:"allow_voidable_by_author_prior_to_instantiation"` // does not require consensus
	AllowVoidableByNotary                     bool `json:"allow_voidable_by_notary"`
	AllowServiceProvider                      bool `json:"allow_service_provider"`      // if true, then contract can be created by a service provider
	AllowNotaryAsBeneficiary                  bool `json:"allow_notary_as_beneficiary"` // if true, then notary can be a beneficiary

	// Requires consensus of all participants who are flagged with this ability.
	// If a consensual agreement, includes all participants with a contractual role even if not flagged as such.
	AllowVoidableByParticipants bool `json:"allow_voidable_by_participants"`

	// if true, then contract can be created with a single min kyc level for all participants by the author
	// this overrides values set for each role in definition.
	AllowMinKycLevelForAllRoles bool `json:"allow_min_kyc_level_for_all_roles"`
}

type SignedContractSignature struct {
	ContractSignaturePackage     ContractSignaturePackage `json:"contract_signature_package"`
	ContractSignaturePackageHash string                   `json:"contract_signature_package_hash"` // hash of the contract signature package
	Signature                    string                   `json:"signature"`                       // signature of the contract signature package hash
}

type ContractSignatures struct {

	// this chains the content block into this block, so that the signature block is immutable
	// it is also bound into each signature as part of the signing process.
	ContractHash string `json:"contract_hash"` // Hash encoded to Base64 of contract block from its JSON representation

	// if true, then this contract requires approval by approvers.
	// This is derived from the contract participants block by examing the roles of each participant.
	HasApprovers         bool       `json:"has_approvers"`
	ApproverSealedOnDate *time.Time `json:"approver_sealed_on_date"` // The date that all approvers (if any) have approved the contract.

	// The date that all signatures are complete and the contract is ready to be instantiated.
	SealedOnDate time.Time `json:"sealed_on_date"`

	Signatures []SignedContractSignature `json:"signatures"`
}

// The signature by a user for the contract and optional signatures for each content item
type ContractSignaturePackage struct {
	SignatureId       string             `json:"signature_id"`
	ContractId        int64              `json:"contract_id"`
	ContractHash      string             `json:"contract_hash"`
	UserId            string             `json:"user_id"` // composite key of user id, example c102, n48, ...
	UserFullName      string             `json:"user_full_name"`
	DateSigned        time.Time          `json:"date_signed"`
	IpAddress         string             `json:"ip_address"`
	SignatureProvider string             `json:"signature_provider"` // where the signature was done, such as Subskribo, ItsMe, Connective, etc"
	SignatureType     string             `json:"signature_type"`     // the type of signature, such as qualified, advanced, etc"
	KeyInfo           KeyInfo            `json:"key_info"`           // the key info used for the signature
	IsApprover        bool               `json:"is_approver"`        // if true, then this signee is an approver, they must sign before non-approvers can sign
	ContentSignatures []ContentSignature `json:"content_signatures"` // signatures for each content item
}

// Used for signatures bound to content items (such as a document)
// Signatures may be from third party providers
type ContentSignature struct {
	ContentId int64 `json:"content__id"`
	// additional fields to be added
}

// Charges and payment information sealed into a readonly record in database upon moving to consent phase.
// This is NOT included in the immutable contract, rather only a hash of this is included to show immutability.
// A reference to any changes to charges and payments after going live are not included in the immutable contract.
// the record is saved in database with a primary key of the contract id
type ContractGoLiveFeesRecord struct {
	ContractId   int64 `json:"contract_id"`
	ContractHash int64 `json:"contract_hash"`

	SealedOnDate time.Time `json:"sealed_on_date"`

	Payments []PaymentRecord `json:"payments"`
	Charges  []ChargeRecord  `json:"charges"`
}

type PaymentRecord struct {
	Id              int64  `json:"id"`
	SummaryLine     string `json:"summary_line"`
	AmountInCredits int64  `json:"amount_in_credits"` // 1 euro = 100 credits
	IncludesVat     bool   `json:"includes_vat"`
}

type ChargeRecord struct {
	Id              int64  `json:"id"`
	SummaryLine     string `json:"summary_line"`
	AmountInCredits int64  `json:"amount_in_credits"` // 1 euro = 100 credits
	FeeType         string `json:"fee_type"`
	IncludesVat     bool   `json:"includes_vat"`
}

type KeyInfo struct {
	KeyId           string `json:"key_id"`
	KeyType         string `json:"key_type"`
	KeySource       string `json:"key_source"`       // azure || local
	KeyFingerprint  string `json:"key_fingerprint"`  // optional, if available
	X509Certificate string `json:"x509_certificate"` // the x509 certificate used if supplied
}

// a separate list is kept for merging into the contract positions when sealing.
// just before sealing, this list is trimmed of any positions that do not have a corresponding organization in the contract.
type ContractParticipantVirtualPosition struct {
	// if negative, then this is a virtual organization added by the author for this contract
	OrgId    int64  `json:"org_id"`
	UserId   string `json:"user_id"`
	Position string `json:"position"`
}

type ContractOrganization struct {
	OrgId int64 `json:"org_id"`

	// true if this organization was added by the author for this contract
	// a virtual organization has a negative org_id
	IsVirtual      bool   `json:"is_virtual"`
	LegalName      string `json:"legal_name"`
	CommonName     string `json:"common_name"`
	OrgType        string `json:"org_type"`
	Signatories    string `json:"signatories"`
	NonSignatories string `json:"non_signatories"`
}

// populate and test with enums used for each property
type SignatureMethod struct {
	PackageMethodId   int64  `json:"package_method_id"`
	SignatureType     string `json:"signature_type"`
	SignatureProvider string `json:"signature_provider"`
}

type ContractProxyInstructions struct {
	// instructions for the proxy beneficiary, only populated if visible to all
	// otherwise it is saved to database and only visible to the creator, notary, proxy beneficiary
	Instructions     string `json:"instructions"`
	VisibleToAll     bool   `json:"visible_to_all"`
	InstructionsHash string `json:"instructions_hash"` // hash of the instructions
}

func (c *ContractParticipant) IsRole(role string) bool {
	if c == nil {
		return false
	}

	if role == "" {
		return false
	}

	for _, r := range c.Roles {
		return r == role
	}

	return false
}
