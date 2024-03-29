package tables

import (
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainmetrics "github.com/bartosian/suimon/internal/core/domain/metrics"
)

var (
	ColumnsConfigProtocol = ColumnsConfig{
		enums.ColumnNameMinSupportedProtocolVersion:                         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameMaxSupportedProtocolVersion:                         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameProtocolVersion:                                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagAcceptZkloginInMultisig:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagAdvanceEpochStartTimeInSafeMode:          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagAdvanceToHighestSupportedProtocolVersion: NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagAllowReceivingObjectID:                   NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagBanEntryInit:                             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagCommitRootStateDigest:                    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagConsensusOrderEndOfEpochLast:             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagDisableInvariantViolationCheckInSwapLoc:  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagDisallowAddingAbilitiesOnUpgrade:         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagDisallowChangeStructTypeParamsOnUpgrade:  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagEnableEffectsV2:                          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagEnableJwkConsensusUpdates:                NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagEndOfEpochTransactionSupported:           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagHardenedOtwCheck:                         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagIncludeConsensusDigestInPrologue:         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagLoadedChildObjectFormat:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagLoadedChildObjectFormatType:              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagLoadedChildObjectsFixed:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagMissingTypeIsCompatibilityError:          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagNarwhalCertificateV2:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagNarwhalHeaderV2:                          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagNarwhalNewLeaderElectionSchedule:         NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagNarwhalVersionedMetadata:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagNoExtraneousModuleBytes:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagPackageDigestHashModule:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagPackageUpgrades:                          NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagRandomBeacon:                             NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagReceiveObjects:                           NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagRecomputeHasPublicTransferInExecution:    NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagScoringDecisionWithValidityCutoff:        NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagSharedObjectDeletion:                     NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagSimpleConservationChecks:                 NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagSimplifiedUnwrapThenDelete:               NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagThroughputAwareConsensusSubmission:       NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagTxnBaseCostAsMultiplier:                  NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagUpgradedMultisigSupported:                NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagVerifyLegacyZkloginAddress:               NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
		enums.ColumnNameFeatureFlagZkLoginAuth:                              NewDefaultColumnConfig(text.AlignCenter, text.AlignCenter, false),
	}

	RowsConfigProtocol = RowsConfig{
		0: {
			enums.ColumnNameMinSupportedProtocolVersion,
			enums.ColumnNameMaxSupportedProtocolVersion,
			enums.ColumnNameProtocolVersion,
			enums.ColumnNameFeatureFlagAcceptZkloginInMultisig,
			enums.ColumnNameFeatureFlagAdvanceEpochStartTimeInSafeMode,
			enums.ColumnNameFeatureFlagAdvanceToHighestSupportedProtocolVersion,
			enums.ColumnNameFeatureFlagAllowReceivingObjectID,
		},
		1: {
			enums.ColumnNameFeatureFlagBanEntryInit,
			enums.ColumnNameFeatureFlagCommitRootStateDigest,
			enums.ColumnNameFeatureFlagConsensusOrderEndOfEpochLast,
			enums.ColumnNameFeatureFlagDisableInvariantViolationCheckInSwapLoc,
			enums.ColumnNameFeatureFlagDisallowAddingAbilitiesOnUpgrade,
			enums.ColumnNameFeatureFlagDisallowChangeStructTypeParamsOnUpgrade,
			enums.ColumnNameFeatureFlagEnableEffectsV2,
		},
		2: {
			enums.ColumnNameFeatureFlagEnableJwkConsensusUpdates,
			enums.ColumnNameFeatureFlagEndOfEpochTransactionSupported,
			enums.ColumnNameFeatureFlagHardenedOtwCheck,
			enums.ColumnNameFeatureFlagIncludeConsensusDigestInPrologue,
			enums.ColumnNameFeatureFlagLoadedChildObjectFormat,
			enums.ColumnNameFeatureFlagLoadedChildObjectFormatType,
			enums.ColumnNameFeatureFlagLoadedChildObjectsFixed,
		},
		3: {
			enums.ColumnNameFeatureFlagMissingTypeIsCompatibilityError,
			enums.ColumnNameFeatureFlagNarwhalCertificateV2,
			enums.ColumnNameFeatureFlagNarwhalHeaderV2,
			enums.ColumnNameFeatureFlagNarwhalNewLeaderElectionSchedule,
			enums.ColumnNameFeatureFlagNarwhalVersionedMetadata,
			enums.ColumnNameFeatureFlagNoExtraneousModuleBytes,
			enums.ColumnNameFeatureFlagPackageDigestHashModule,
		},
		4: {
			enums.ColumnNameFeatureFlagPackageUpgrades,
			enums.ColumnNameFeatureFlagRandomBeacon,
			enums.ColumnNameFeatureFlagReceiveObjects,
			enums.ColumnNameFeatureFlagRecomputeHasPublicTransferInExecution,
			enums.ColumnNameFeatureFlagScoringDecisionWithValidityCutoff,
			enums.ColumnNameFeatureFlagSharedObjectDeletion,
			enums.ColumnNameFeatureFlagSimpleConservationChecks,
		},
		5: {
			enums.ColumnNameFeatureFlagSimplifiedUnwrapThenDelete,
			enums.ColumnNameFeatureFlagThroughputAwareConsensusSubmission,
			enums.ColumnNameFeatureFlagTxnBaseCostAsMultiplier,
			enums.ColumnNameFeatureFlagUpgradedMultisigSupported,
			enums.ColumnNameFeatureFlagVerifyLegacyZkloginAddress,
			enums.ColumnNameFeatureFlagZkLoginAuth,
		},
	}
)

func GetProtocolColumnValues(metrics *domainmetrics.Metrics) (ColumnValues, error) {
	return ColumnValues{
		enums.ColumnNameIndex:                                               1,
		enums.ColumnNameMinSupportedProtocolVersion:                         metrics.Protocol.MinSupportedProtocolVersion,
		enums.ColumnNameMaxSupportedProtocolVersion:                         metrics.Protocol.MaxSupportedProtocolVersion,
		enums.ColumnNameProtocolVersion:                                     metrics.Protocol.ProtocolVersion,
		enums.ColumnNameFeatureFlagAcceptZkloginInMultisig:                  metrics.Protocol.FeatureFlags.AcceptZkloginInMultisig,
		enums.ColumnNameFeatureFlagAdvanceEpochStartTimeInSafeMode:          metrics.Protocol.FeatureFlags.AdvanceEpochStartTimeInSafeMode,
		enums.ColumnNameFeatureFlagAdvanceToHighestSupportedProtocolVersion: metrics.Protocol.FeatureFlags.AdvanceToHighestSupportedProtocolVersion,
		enums.ColumnNameFeatureFlagAllowReceivingObjectID:                   metrics.Protocol.FeatureFlags.AllowReceivingObjectID,
		enums.ColumnNameFeatureFlagBanEntryInit:                             metrics.Protocol.FeatureFlags.BanEntryInit,
		enums.ColumnNameFeatureFlagCommitRootStateDigest:                    metrics.Protocol.FeatureFlags.CommitRootStateDigest,
		enums.ColumnNameFeatureFlagConsensusOrderEndOfEpochLast:             metrics.Protocol.FeatureFlags.ConsensusOrderEndOfEpochLast,
		enums.ColumnNameFeatureFlagDisableInvariantViolationCheckInSwapLoc:  metrics.Protocol.FeatureFlags.DisableInvariantViolationCheckInSwapLoc,
		enums.ColumnNameFeatureFlagDisallowAddingAbilitiesOnUpgrade:         metrics.Protocol.FeatureFlags.DisallowAddingAbilitiesOnUpgrade,
		enums.ColumnNameFeatureFlagDisallowChangeStructTypeParamsOnUpgrade:  metrics.Protocol.FeatureFlags.DisallowChangeStructTypeParamsOnUpgrade,
		enums.ColumnNameFeatureFlagEnableEffectsV2:                          metrics.Protocol.FeatureFlags.EnableEffectsV2,
		enums.ColumnNameFeatureFlagEnableJwkConsensusUpdates:                metrics.Protocol.FeatureFlags.EnableJwkConsensusUpdates,
		enums.ColumnNameFeatureFlagEndOfEpochTransactionSupported:           metrics.Protocol.FeatureFlags.EndOfEpochTransactionSupported,
		enums.ColumnNameFeatureFlagHardenedOtwCheck:                         metrics.Protocol.FeatureFlags.HardenedOtwCheck,
		enums.ColumnNameFeatureFlagIncludeConsensusDigestInPrologue:         metrics.Protocol.FeatureFlags.IncludeConsensusDigestInPrologue,
		enums.ColumnNameFeatureFlagLoadedChildObjectFormat:                  metrics.Protocol.FeatureFlags.LoadedChildObjectFormat,
		enums.ColumnNameFeatureFlagLoadedChildObjectFormatType:              metrics.Protocol.FeatureFlags.LoadedChildObjectFormatType,
		enums.ColumnNameFeatureFlagLoadedChildObjectsFixed:                  metrics.Protocol.FeatureFlags.LoadedChildObjectsFixed,
		enums.ColumnNameFeatureFlagMissingTypeIsCompatibilityError:          metrics.Protocol.FeatureFlags.MissingTypeIsCompatibilityError,
		enums.ColumnNameFeatureFlagNarwhalCertificateV2:                     metrics.Protocol.FeatureFlags.NarwhalCertificateV2,
		enums.ColumnNameFeatureFlagNarwhalHeaderV2:                          metrics.Protocol.FeatureFlags.NarwhalHeaderV2,
		enums.ColumnNameFeatureFlagNarwhalNewLeaderElectionSchedule:         metrics.Protocol.FeatureFlags.NarwhalNewLeaderElectionSchedule,
		enums.ColumnNameFeatureFlagNarwhalVersionedMetadata:                 metrics.Protocol.FeatureFlags.NarwhalVersionedMetadata,
		enums.ColumnNameFeatureFlagNoExtraneousModuleBytes:                  metrics.Protocol.FeatureFlags.NoExtraneousModuleBytes,
		enums.ColumnNameFeatureFlagPackageDigestHashModule:                  metrics.Protocol.FeatureFlags.PackageDigestHashModule,
		enums.ColumnNameFeatureFlagPackageUpgrades:                          metrics.Protocol.FeatureFlags.PackageUpgrades,
		enums.ColumnNameFeatureFlagRandomBeacon:                             metrics.Protocol.FeatureFlags.RandomBeacon,
		enums.ColumnNameFeatureFlagReceiveObjects:                           metrics.Protocol.FeatureFlags.ReceiveObjects,
		enums.ColumnNameFeatureFlagRecomputeHasPublicTransferInExecution:    metrics.Protocol.FeatureFlags.RecomputeHasPublicTransferInExecution,
		enums.ColumnNameFeatureFlagScoringDecisionWithValidityCutoff:        metrics.Protocol.FeatureFlags.ScoringDecisionWithValidityCutoff,
		enums.ColumnNameFeatureFlagSharedObjectDeletion:                     metrics.Protocol.FeatureFlags.SharedObjectDeletion,
		enums.ColumnNameFeatureFlagSimpleConservationChecks:                 metrics.Protocol.FeatureFlags.SimpleConservationChecks,
		enums.ColumnNameFeatureFlagSimplifiedUnwrapThenDelete:               metrics.Protocol.FeatureFlags.SimplifiedUnwrapThenDelete,
		enums.ColumnNameFeatureFlagThroughputAwareConsensusSubmission:       metrics.Protocol.FeatureFlags.ThroughputAwareConsensusSubmission,
		enums.ColumnNameFeatureFlagTxnBaseCostAsMultiplier:                  metrics.Protocol.FeatureFlags.TxnBaseCostAsMultiplier,
		enums.ColumnNameFeatureFlagUpgradedMultisigSupported:                metrics.Protocol.FeatureFlags.UpgradedMultisigSupported,
		enums.ColumnNameFeatureFlagVerifyLegacyZkloginAddress:               metrics.Protocol.FeatureFlags.VerifyLegacyZkloginAddress,
		enums.ColumnNameFeatureFlagZkLoginAuth:                              metrics.Protocol.FeatureFlags.ZkloginAuth,
	}, nil
}
