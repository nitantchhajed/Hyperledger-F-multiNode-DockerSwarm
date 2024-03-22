package contract

import (
	"errors"
	"fmt"
)

const SHA256_HASH_BASE64_LENGTH = 44
const SIGNATURE_RSA2048_BASE64_LENGTH = 344

func (cd *ContractDefinition) UserRoleDefinition(role string) *ContractUserRoleDefinition {
	for _, ur := range cd.UserRoles {
		if ur.Role == role {
			return &ur
		}
	}
	return nil
}

func (cb *ContractBlock) Validate() error {
	d := cb.Definition

	if cb.ContractFamilyId != d.ContractFamilyId {
		return errors.New("invalid contract family id")
	}

	if cb.ContractTypeId != d.ContractType {
		return errors.New("invalid contract type id")
	}

	if cb.ContractTypeVersion == 0 {
		return errors.New("invalid contract type version")
	}

	if cb.StorageYears < 1 || cb.StorageYears > 30 {
		return errors.New("invalid storage years")
	}

	opt := cb.ContractOptions

	if opt.ExpiryDate != nil {
		if opt.ExpiryDate.Before(cb.SealedOnDate) {
			return errors.New("invalid expiry date")
		}

		if opt.DaysToSign < 1 {
			return errors.New("invalid days to sign")
		}

		if opt.AllowSignatureExtension && opt.MaxDaysToSign < 1 {
			return errors.New("invalid days to sign extension")
		}
	}

	return nil
}

func (sm *SignatureMethod) Validate() error {
	if sm.SignatureType == "" {
		return errors.New("invalid signature type")
	}

	if sm.SignatureType != "advanced" && sm.SignatureType != "qualified" {
		return errors.New("invalid signature type")
	}

	if sm.PackageMethodId != 1 && sm.PackageMethodId != 2 && sm.PackageMethodId != 3 {
		return errors.New("invalid package method id")
	}

	if sm.SignatureProvider != "Subskribo" && sm.SignatureProvider != "Connective" {
		return errors.New("invalid signature provider")
	}

	return nil
}

const (
	InstructionMinLength = 60
)

func (rid *ReleaseInstructionDetail) Validate(evidenceRequired bool) error {
	if len(rid.Instructions) < InstructionMinLength {
		return errors.New("invalid instructions")
	}

	if !rid.IsCustomRelease {
		if rid.StandardReleaseTemplateId < 1 {
			return errors.New("invalid standard release template id")
		}
	} else {
		if rid.NotaryPackage != nil {
			if rid.NotarySignature == "" {
				return errors.New("invalid notary signature")
			}

			if rid.NotaryPackage.ApprovalState != "none" {
				if rid.AcceptancePackage == nil {
					return errors.New("invalid acceptance package")
				}

				if rid.AcceptanceSignature == "" {
					return errors.New("invalid acceptance signature")
				}
			}
		} else {
			// verifiers
			if rid.ConsensusMethod == "" {
				return errors.New("invalid consensus method")
			}

			if evidenceRequired {
				if !rid.IsEvidenceRequiredForRelease {
					return errors.New("invalid evidence required for release flag")
				}
			}
		}
	}
	return nil
}

func (pi *ContractProxyInstructions) Validate() error {

	if pi.VisibleToAll {
		if pi.Instructions == "" {
			return errors.New("invalid instructions")
		}
	}

	if pi.InstructionsHash == "" {
		return errors.New("invalid instructions hash")
	}

	return nil
}

func (cc *ConstructedContentItem) Validate() error {

	if cc.ContentId < 1 {
		return errors.New("invalid content id")
	}

	if cc.PlainHash == "" {
		return errors.New("invalid plain hash")
	}

	return nil
}

func (sp *ContractSignaturePackage) Validate(signatureMethodId int64) error {

	if sp.ContractId < 1 {
		return errors.New("invalid contract id")
	}

	if sp.ContractHash == "" {
		return errors.New("invalid contract hash")
	}

	if sp.UserId == "" {
		return errors.New("invalid user id")
	}

	if sp.UserFullName == "" {
		return errors.New("invalid user full name")
	}

	if sp.DateSigned.IsZero() {
		return errors.New("invalid date signed")
	}

	if sp.SignatureType == "" {
		return errors.New("invalid signature type")
	}

	if signatureMethodId != 3 {
		if sp.SignatureId == "" {
			return errors.New("invalid signature id")
		}

		if sp.IpAddress == "" {
			return errors.New("invalid ip address")
		}

		if sp.SignatureProvider == "" {
			return errors.New("invalid signature provider")
		}

		if err := sp.KeyInfo.Validate(); err != nil {
			return err
		}

		if sp.ContractHash == "" {
			return errors.New("invalid contract hash")
		}

	}

	return nil
}

func (k *KeyInfo) Validate() error {

	if k.X509Certificate == "" {
		if k.KeyId == "" {
			return errors.New("invalid key id")
		}

		if k.KeyType == "" {
			return errors.New("invalid key type")
		}

		if k.KeySource == "" {
			return errors.New("invalid key source")
		}
	}

	return nil
}

func (c *ImmutableContract) ValidateSignaturesComplete() error {
	if c == nil {
		return fmt.Errorf("contract container is nil for method: ValidateSignaturesComplete")
	}

	if len(c.ContractSignatures.ContractHash) != SHA256_HASH_BASE64_LENGTH {
		return fmt.Errorf("contract signatures block does not have a contract hash set, or is not of correct length")
	}

	signedPackCount := len(c.ContractSignatures.Signatures)
	if signedPackCount == 0 {
		return fmt.Errorf("contract signatures block does not have any signature packages set")
	}

	signatoryCount := c.Contract.GetSignatoryCountFromParticipants()

	if signedPackCount != signatoryCount {
		return fmt.Errorf("contract signatures block does not have the same number of signature packages as there are signatories")
	}

	isEmbedded := c.Contract.SignatureMethod.PackageMethodId == int64(SignPackageMethodId_Embedded)
	requireConstructedContent := c.Contract.SignatureMethod.PackageMethodId == int64(SignPackageMethodId_Constructed)

	if isEmbedded {
		// if c.ContractSignatures.FinalizedContent == nil || c.ContractSignatures.FinalizedContent.ContentId == 0 {
		// 	return fmt.Errorf("contract signatures block does not have a finalized content item set")
		// }
	}

	// constructedContentItemId := int64(0)

	if requireConstructedContent {
		// if len(c.ContractSignatures.ConstructedContentItems) == 0 {
		// 	return fmt.Errorf("contract signatures block does not have any constructed content items set")
		// }
		// myItem := &c.ContractSignatures.ConstructedContentItems[0]
		// constructedContentItemId = myItem.ContentId
		// if constructedContentItemId == 0 || len(myItem.PlainHash) != SHA256_HASH_BASE64_LENGTH {
		// 	return fmt.Errorf("contract signatures block does not have content id or else required parameters set for constructed content item")
		// }
	}

	found := false
	for _, p := range c.Contract.Participants {
		if p.IsRole(Signatory) {
			for _, sp := range c.ContractSignatures.Signatures {
				if sp.ContractSignaturePackage.UserId == p.UserId {
					found = true
					if sp.ContractSignaturePackage.UserFullName == "" {
						return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "full name not set")
					}
					if sp.ContractSignaturePackage.DateSigned.IsZero() {
						return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "signed on data not set")
					}
					if len(sp.ContractSignaturePackage.ContractHash) != SHA256_HASH_BASE64_LENGTH {
						return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "package hash no set or incorrect length")
					} else if c.ContractSignatures.ContractHash != sp.ContractSignaturePackage.ContractHash {
						return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "contract hash does not match contract hash in signature package")
					}
					if sp.ContractSignaturePackage.ContractId != c.Contract.ContractID {
						return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "contract id does not match contract id in signature package")
					}

					if requireConstructedContent {
						// if sp.ContractSignaturePackage.ConstructedContentId != constructedContentItemId {
						// 	return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "constructed content id is not set correctly")
						// }
					}

					if !isEmbedded {
						// to do: validate key information set

						if len(sp.ContractSignaturePackageHash) != SHA256_HASH_BASE64_LENGTH {
							return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "package hash not of correct length")
						}

						if sp.Signature == "" {
							return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "missing signature")
						} else if len(sp.Signature) != SIGNATURE_RSA2048_BASE64_LENGTH {
							return formatSignedPackageErr(sp.ContractSignaturePackage.UserId, sp.ContractSignaturePackage.UserFullName, "signature is not of correct length ")
						}
					}
				}

				if !found {
					return fmt.Errorf("contract signatures block does not have a signature package for signatory user id '%v'", p.UserId)
				}
			}
		}
	}

	return nil
}

func formatSignedPackageErr(userId string, userName string, msgPostfix string) error {
	return fmt.Errorf("contract signature package for user id '%v' and name '%v', has error: %v", userId, userName, msgPostfix)
}

func (c *ContractBlock) GetSignatoryCountFromParticipants() int {
	if c == nil {
		return 0
	}

	count := 0
	for _, p := range c.Participants {
		if p.IsRole(Signatory) {
			count++
		}
	}

	return count
}
