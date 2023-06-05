package lvm

import (
	"fmt"
	"testing"

	"github.com/caoyingjunz/csi-driver-localstorage/pkg/cache"
	"github.com/container-storage-interface/spec/lib/go/csi"
)

func TestCreateLV(t *testing.T) {
	pathPrefix := "/dev"
	volumeID := "123456789"
	vgname := "vg1"
	req := &csi.CreateVolumeRequest{
		Name: "testLV",
		CapacityRange: &csi.CapacityRange{
			RequiredBytes: 1000,
		},
		Parameters: make(map[string]string),
	}
	req.Parameters["vgname"] = vgname

	volume, vgname, _ := NewVolumeForCreate(pathPrefix, volumeID, req)
	fmt.Println(volume, vgname)

	CreateLogicalVolume(volume, vgname)

	exist, _ := CheckVolumeExists(volume)
	if !exist {
		t.Fatal()
	}
}

func TestRemoveLV(t *testing.T) {
	volume := &cache.Volume{
		VolPath: "/dev/vg1/testLV",
	}

	RemoveLogicalVolume(volume)

	exist, _ := CheckVolumeExists(volume)
	if exist {
		t.Fatal()
	}
}
