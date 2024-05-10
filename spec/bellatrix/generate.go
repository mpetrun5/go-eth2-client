// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bellatrix

//nolint:revive
// Need to `go install github.com/ferranbt/fastssz/sszgen@latest` for this to work.
//go:generate rm -f beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go signedbeaconblock_ssz.go
//go:generate sszgen -suffix ssz -include ../phase0,../altair -path . -objs BeaconBlock,BeaconBlockBody,BeaconState,ExecutionPayload,ExecutionPaylodHeader,SignedBeaconBlock
//go:generate goimports -w beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go signedbeaconblock_ssz.go
