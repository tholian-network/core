package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupPrograms() *Programs {

	programs := NewPrograms()
	programs.Add(structs.Program{
		PID:     10,
		Name:    "kitty",
		Command: "/usr/bin/kitty",
		Arguments: []string{
			"/usr/bin/kitty",
			"--working-directory=/home/sandbox",
		},
		Folder: "/home/sandbox",
		Environment: map[string]string{
			"USER": "cookiengineer",
		},
		Filesystem: []string{
			"/usr/bin/kitty",
			"/usr/lib/libdrm_amdgpu.so.1.124.0",
			"/usr/lib/libdrm_intel.so.1.124.0",
			"/usr/lib/libdrm.so.2.124.0",
			"/usr/lib/libedit.so.0.0.75",
			"/usr/lib/libelf-0.192.so",
			"/usr/lib/libgallium-25.0.4-arch1.1.so",
			"/usr/lib/libGLdispatch.so.0.0.0",
			"/usr/lib/libGL.so.1.7.0",
			"/usr/lib/libGLX_mesa.so.0.0.0",
			"/usr/lib/libGLX.so.0.0.0",
			"/usr/lib/libicudata.so.76.1",
			"/usr/lib/libicuuc.so.76.1",
			"/usr/lib/libLLVM.so.19.1",
			"/usr/lib/libncursesw.so.6.5",
			"/usr/lib/libpciaccess.so.0.11.1",
			"/usr/lib/libsensors.so.5.0.0",
			"/usr/lib/libSPIRV-Tools.so",
			"/usr/lib/libstdc++.so.6.0.33",
			"/usr/lib/libxcb-dri3.so.0.1.0",
			"/usr/lib/libxcb-glx.so.0.0.0",
			"/usr/lib/libxcb-present.so.0.0.0",
			"/usr/lib/libxcb-randr.so.0.1.0",
			"/usr/lib/libxcb-sync.so.1.0.0",
			"/usr/lib/libxcb-xfixes.so.0.0.0",
			"/usr/lib/libxml2.so.2.13.8",
			"/usr/lib/libxshmfence.so.1.0.0",
			"/usr/lib/libzstd.so.1.5.7",
			"/usr/lib/python3.13/lib-dynload/grp.cpython-313-x86_64-linux-gnu.so",
		},
	})
	programs.Add(structs.Program{
		PID:     11,
		Name:    "kitty",
		Command: "/usr/bin/kitty",
		Arguments: []string{
			"/usr/bin/kitty",
		},
		Folder: "/home/cookiengineer/Software",
		Environment: map[string]string{
			"USER":       "cookiengineer",
			"LD_PRELOAD": "/tmp/dropper.so",
		},
		Filesystem: []string{
			"/usr/bin/kitty",
			"/usr/lib/libdrm_amdgpu.so.1.124.0",
			"/usr/lib/libdrm_intel.so.1.124.0",
			"/usr/lib/libdrm.so.2.124.0",
			"/usr/lib/libedit.so.0.0.75",
			"/usr/lib/libelf-0.192.so",
			"/usr/lib/libgallium-25.0.4-arch1.1.so",
			"/usr/lib/libGLdispatch.so.0.0.0",
			"/usr/lib/libGL.so.1.7.0",
			"/usr/lib/libGLX_mesa.so.0.0.0",
			"/usr/lib/libGLX.so.0.0.0",
			"/usr/lib/libicudata.so.76.1",
			"/usr/lib/libicuuc.so.76.1",
			"/usr/lib/libLLVM.so.19.1",
			"/usr/lib/libncursesw.so.6.5",
			"/usr/lib/libpciaccess.so.0.11.1",
			"/usr/lib/libsensors.so.5.0.0",
			"/usr/lib/libSPIRV-Tools.so",
			"/usr/lib/libstdc++.so.6.0.33",
			"/usr/lib/libxcb-dri3.so.0.1.0",
			"/usr/lib/libxcb-glx.so.0.0.0",
			"/usr/lib/libxcb-present.so.0.0.0",
			"/usr/lib/libxcb-randr.so.0.1.0",
			"/usr/lib/libxcb-sync.so.1.0.0",
			"/usr/lib/libxcb-xfixes.so.0.0.0",
			"/usr/lib/libxml2.so.2.13.8",
			"/usr/lib/libxshmfence.so.1.0.0",
			"/usr/lib/libzstd.so.1.5.7",
			"/usr/lib/python3.13/lib-dynload/grp.cpython-313-x86_64-linux-gnu.so",
		},
	})
	programs.Add(structs.Program{
		PID:     13,
		Name:    "vim",
		Command: "/usr/bin/vim",
		Arguments: []string{
			"/usr/bin/vim",
		},
		Folder: "/home/cookiengineer/Software/project",
		Environment: map[string]string{
			"USER":       "cookiengineer",
			"LD_PRELOAD": "/tmp/exploit.so",
		},
		Filesystem: []string{
			"/usr/bin/vim",
			"/usr/lib/libacl.so.1.1.2302",
			"/usr/lib/libc.so.6",
			"/usr/lib/libgpm.so.2.1.0",
			"/usr/lib/libncursesw.so.6.5",
			"/usr/lib/libm.so.6",
			"/usr/lib/ld-linux-x86-64.so.2",
			"/usr/lib/ld-linux-x86-64.so.2",
		},
	})

	return programs

}

func TestPrograms(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		programs := mockupPrograms()

		program1 := programs.Get(1337)

		if program1 != nil {
			t.Errorf("Expected %d to be nil", program1.PID)
		}

		programs.Add(structs.Program{
			PID:     1337,
			Name:    "retrokit",
			Command: "/usr/bin/retrokit --sandbox=/tmp/sandbox",
			Arguments: []string{
				"/usr/bin/retrokit",
				"--sandbox=/tmp/sandbox",
			},
			Folder: "/tmp/sandbox",
			Environment: map[string]string{
				"USER": "cookiengineer",
			},
			Filesystem: []string{
				"/usr/bin/retrokit",
			},
		})

		program2 := programs.Get(1337)

		if program2 == nil {
			t.Errorf("Expected nil to be %d", program1.PID)
		} else if program2.PID != 1337 {
			t.Errorf("Expected %d to be %d", program1.PID, 1337)
		}

	})

	t.Run("Get()", func(t *testing.T) {

		programs := mockupPrograms()

		program1 := programs.Get(10)
		program2 := programs.Get(13)

		if program1 == nil {
			t.Errorf("Expected nil to be %s", "kitty")
		} else if program1.Name != "kitty" {
			t.Errorf("Expected %s to be %s", program1.Name, "kitty")
		}

		if program2 == nil {
			t.Errorf("Expected nil to be %s", "vim")
		} else if program2.Name != "vim" {
			t.Errorf("Expected %s to be %s", program1.Name, "vim")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		programs := mockupPrograms()

		found1 := programs.Query(matchers.Program{
			Name: "kitty",
			Command: "any",
		})

		found2 := programs.Query(matchers.Program{
			Name: "any",
			Command: "/usr/bin/vim",
		})

		if len(found1) == 2 {

			if found1[0].Command != "/usr/bin/kitty" {
				t.Errorf("Expected %s to be %s", found1[0].Command, "/usr/bin/kitty")
			}

			if found1[1].Command != "/usr/bin/kitty" {
				t.Errorf("Expected %s to be %s", found1[1].Command, "/usr/bin/kitty")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=kitty")
		}

		if len(found2) == 1 {

			if found2[0].Command != "/usr/bin/vim" {
				t.Errorf("Expected %s to be %s", found2[0].Command, "/usr/bin/vim")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Command=/usr/bin/vim")
		}

	})

	t.Run("QueryByEnvironmentVariable()", func(t *testing.T) {

		programs := mockupPrograms()

		found1 := programs.QueryByEnvironmentVariable("USER", "cookiengineer")
		found2 := programs.QueryByEnvironmentVariable("LD_PRELOAD", "any")
		found3 := programs.QueryByEnvironmentVariable("LD_PRELOAD", "/tmp/exploit.so")

		if len(found1) != 3 {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 3, "USER=cookiengineer")
		}

		if len(found2) != 2 {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "LD_PRELOAD=any")
		}

		if len(found3) != 1 {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "LD_PRELOAD=/tmp/exploit.so")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		programs := mockupPrograms()

		program1 := programs.Get(10)

		if program1 == nil {
			t.Errorf("Expected nil to be %s", "kitty")
		} else if program1.Name != "kitty" {
			t.Errorf("Expected %s to be %s", program1.Name, "kitty")
		}

		programs.Remove(10)

		program2 := programs.Get(10)

		if program2 != nil {
			t.Errorf("Expected %s to be nil", program2.Name)
		}

	})

}
