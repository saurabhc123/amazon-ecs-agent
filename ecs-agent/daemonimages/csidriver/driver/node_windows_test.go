//go:build windows
// +build windows

// this file has been modified from its original found in:
// https://github.com/kubernetes-sigs/aws-ebs-csi-driver

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"context"
	"testing"

	"github.com/aws/amazon-ecs-agent/ecs-agent/daemonimages/csidriver/driver/internal"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	volumeID = "vol-07faa87d7a3debf5b"
	nvmeName = "/dev/disk/by-id/nvme-Amazon_Elastic_Block_Store_voltest"
)

func TestNodeStageVolume(t *testing.T) {

	var (
		targetPath = "C://csi_proxy_poc//mount"
		devicePath = "/dev/fake"
		// deviceFileInfo = fs.FileInfo(&fakeFileInfo{devicePath, os.ModeDevice})
		stdVolCap = &csi.VolumeCapability{
			AccessType: &csi.VolumeCapability_Mount{
				Mount: &csi.VolumeCapability_MountVolume{
					FsType: FSTypeXfs,
				},
			},
			AccessMode: &csi.VolumeCapability_AccessMode{
				Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
			},
		}
		// With few exceptions, all "success" non-block cases have roughly the same
		// expected calls and only care about testing the FormatAndMountSensitiveWithFormatOptions call. The
		// exceptions should not call this, instead they should define expectMock
		// from scratch.
		//successExpectMock = func(mockMounter MockMounter, mockDeviceIdentifier MockDeviceIdentifier) {
		//	mockMounter.EXPECT().PathExists(gomock.Eq(targetPath)).Return(false, nil)
		//	mockMounter.EXPECT().MakeDir(targetPath).Return(nil)
		//	mockMounter.EXPECT().GetDeviceNameFromMount(targetPath).Return("", 1, nil)
		//	mockMounter.EXPECT().PathExists(gomock.Eq(devicePath)).Return(true, nil)
		//	mockDeviceIdentifier.EXPECT().Lstat(gomock.Eq(devicePath)).Return(deviceFileInfo, nil)
		//	mockMounter.EXPECT().NeedResize(gomock.Eq(devicePath), gomock.Eq(targetPath)).Return(false, nil)
		//}
	)
	testCases := []struct {
		name         string
		request      *csi.NodeStageVolumeRequest
		inFlightFunc func(*internal.InFlight) *internal.InFlight
		expectMock   func(mockMounter MockMounter, mockDeviceIdentifier MockDeviceIdentifier)
		expectedCode codes.Code
	}{
		{
			name: "success normal",
			request: &csi.NodeStageVolumeRequest{
				PublishContext:    map[string]string{DevicePathKey: devicePath},
				StagingTargetPath: targetPath,
				VolumeCapability:  stdVolCap,
				VolumeId:          volumeID,
			},
			//expectMock: func(mockMounter MockMounter, mockDeviceIdentifier MockDeviceIdentifier) {
			//	successExpectMock(mockMounter, mockDeviceIdentifier)
			//	mockMounter.EXPECT().FormatAndMountSensitiveWithFormatOptions(gomock.Eq(devicePath), gomock.Eq(targetPath), gomock.Eq(defaultFsType), gomock.Any(), gomock.Nil(), gomock.Len(0))
			//},
		}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()

			mockMounter := NewMockMounter(mockCtl)
			mockDeviceIdentifier := NewMockDeviceIdentifier(mockCtl)

			inFlight := internal.NewInFlight()
			if tc.inFlightFunc != nil {
				tc.inFlightFunc(inFlight)
			}

			awsDriver := &nodeService{
				mounter:          mockMounter,
				deviceIdentifier: mockDeviceIdentifier,
				inFlight:         inFlight,
			}

			if tc.expectMock != nil {
				tc.expectMock(*mockMounter, *mockDeviceIdentifier)
			}

			_, err := awsDriver.NodeStageVolume(context.TODO(), tc.request)
			if tc.expectedCode != codes.OK {
				expectErr(t, err, tc.expectedCode)
			} else if err != nil {
				t.Fatalf("Expect no error but got: %v", err)
			}
		})
	}
}

func expectErr(t *testing.T, actualErr error, expectedCode codes.Code) {
	if actualErr == nil {
		t.Fatalf("Expect error but got no error")
	}

	status, ok := status.FromError(actualErr)
	if !ok {
		t.Fatalf("Failed to get error status code from error: %v", actualErr)
	}

	if status.Code() != expectedCode {
		t.Fatalf("Expected error code %d, got %d message %s", codes.InvalidArgument, status.Code(), status.Message())
	}
}
