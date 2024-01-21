package metrics

// Protocol represents the protocol information of the Sui blockchain network.
// It includes the minimum and maximum supported protocol versions, the current protocol version, and the feature flags.
type Protocol struct {
	MinSupportedProtocolVersion string       `json:"minSupportedProtocolVersion"`
	MaxSupportedProtocolVersion string       `json:"maxSupportedProtocolVersion"`
	ProtocolVersion             string       `json:"protocolVersion"`
	FeatureFlags                FeatureFlags `json:"featureFlags"`
}

// FeatureFlags represents the various feature flags of the Sui blockchain network.
// Each flag represents a specific feature of the network that can be toggled on or off.
type FeatureFlags struct {
	AcceptZkloginInMultisig                  bool `json:"accept_zklogin_in_multisig"`
	AdvanceEpochStartTimeInSafeMode          bool `json:"advance_epoch_start_time_in_safe_mode"`
	AdvanceToHighestSupportedProtocolVersion bool `json:"advance_to_highest_supported_protocol_version"`
	AllowReceivingObjectId                   bool `json:"allow_receiving_object_id"`
	BanEntryInit                             bool `json:"ban_entry_init"`
	CommitRootStateDigest                    bool `json:"commit_root_state_digest"`
	ConsensusOrderEndOfEpochLast             bool `json:"consensus_order_end_of_epoch_last"`
	DisableInvariantViolationCheckInSwapLoc  bool `json:"disable_invariant_violation_check_in_swap_loc"`
	DisallowAddingAbilitiesOnUpgrade         bool `json:"disallow_adding_abilities_on_upgrade"`
	DisallowChangeStructTypeParamsOnUpgrade  bool `json:"disallow_change_struct_type_params_on_upgrade"`
	EnableEffectsV2                          bool `json:"enable_effects_v2"`
	EnableJwkConsensusUpdates                bool `json:"enable_jwk_consensus_updates"`
	EndOfEpochTransactionSupported           bool `json:"end_of_epoch_transaction_supported"`
	HardenedOtwCheck                         bool `json:"hardened_otw_check"`
	IncludeConsensusDigestInPrologue         bool `json:"include_consensus_digest_in_prologue"`
	LoadedChildObjectFormat                  bool `json:"loaded_child_object_format"`
	LoadedChildObjectFormatType              bool `json:"loaded_child_object_format_type"`
	LoadedChildObjectsFixed                  bool `json:"loaded_child_objects_fixed"`
	MissingTypeIsCompatibilityError          bool `json:"missing_type_is_compatibility_error"`
	NarwhalCertificateV2                     bool `json:"narwhal_certificate_v2"`
	NarwhalHeaderV2                          bool `json:"narwhal_header_v2"`
	NarwhalNewLeaderElectionSchedule         bool `json:"narwhal_new_leader_election_schedule"`
	NarwhalVersionedMetadata                 bool `json:"narwhal_versioned_metadata"`
	NoExtraneousModuleBytes                  bool `json:"no_extraneous_module_bytes"`
	PackageDigestHashModule                  bool `json:"package_digest_hash_module"`
	PackageUpgrades                          bool `json:"package_upgrades"`
	RandomBeacon                             bool `json:"random_beacon"`
	ReceiveObjects                           bool `json:"receive_objects"`
	RecomputeHasPublicTransferInExecution    bool `json:"recompute_has_public_transfer_in_execution"`
	ScoringDecisionWithValidityCutoff        bool `json:"scoring_decision_with_validity_cutoff"`
	SharedObjectDeletion                     bool `json:"shared_object_deletion"`
	SimpleConservationChecks                 bool `json:"simple_conservation_checks"`
	SimplifiedUnwrapThenDelete               bool `json:"simplified_unwrap_then_delete"`
	ThroughputAwareConsensusSubmission       bool `json:"throughput_aware_consensus_submission"`
	TxnBaseCostAsMultiplier                  bool `json:"txn_base_cost_as_multiplier"`
	UpgradedMultisigSupported                bool `json:"upgraded_multisig_supported"`
	VerifyLegacyZkloginAddress               bool `json:"verify_legacy_zklogin_address"`
	ZkloginAuth                              bool `json:"zklogin_auth"`
}
